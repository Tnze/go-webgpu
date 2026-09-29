package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------------
// Function-body emission.
//
// The pinner + marshaling + call + return logic is identical between the cgo
// and syscall backends except for a handful of type/call spellings. Rather
// than duplicating a hundred lines of text/template, each function's body is
// precomputed here as a line slice per backend — the same approach the gpu
// wrapper layer uses (see wrapmethod.go).
// ---------------------------------------------------------------------------

// Backend selects the call convention used to render a function body.
type Backend int

const (
	BackendCGO Backend = iota
	BackendSyscall
)

// callBody renders the complete body of fd (without the surrounding braces)
// for the given backend. Returned lines are already tab-indented one level.
func callBody(fd FuncData, be Backend) []string {
	var lines []string
	if needsPinner(fd) {
		lines = append(lines, "var pinner runtime.Pinner", "defer pinner.Unpin()")
		lines = append(lines, pinLines(fd)...)
	}
	lines = append(lines, convLines(fd, be)...)
	lines = append(lines, callLines(fd, be)...)
	return lines
}

// needsPinner reports whether the wrapper must keep Go pointers alive across
// the C call.
func needsPinner(fd FuncData) bool {
	for _, a := range fd.GoArgs {
		if isStructType(a.TypeRef) || a.IsSlice || a.IsStr || a.IsCBInfo {
			return true
		}
	}
	return false
}

// pinLines emits runtime.Pinner.Pin calls for every argument that may hold a
// Go pointer, including nested strings/slices/pointers of by-value structs.
func pinLines(fd FuncData) []string {
	var lines []string
	for _, a := range fd.GoArgs {
		switch {
		case a.IsCBInfo:
			lines = append(lines, fmt.Sprintf("pinner.Pin(&%s)", a.Name))
		case isStructType(a.TypeRef) && a.IsPtr:
			lines = append(lines, fmt.Sprintf("if %s != nil {", a.Name))
			lines = append(lines, "\tpinner.Pin("+a.Name+")")
			lines = append(lines, structPins(a, "\t")...)
			lines = append(lines, "}")
		case isStructType(a.TypeRef):
			lines = append(lines, "pinner.Pin(&"+a.Name+")")
			lines = append(lines, structPins(a, "")...)
		case a.IsSlice:
			lines = append(lines,
				fmt.Sprintf("if len(%s) > 0 {", a.Name),
				fmt.Sprintf("\tpinner.Pin(&%s[0])", a.Name),
				"}")
		}
	}
	return lines
}

// structPins emits pins for the Go pointers reachable from a by-value struct
// argument. indent is the leading whitespace of the enclosing block.
func structPins(a FuncArgData, indent string) []string {
	var lines []string
	if a.HasChain {
		lines = append(lines,
			fmt.Sprintf("%sif %s.NextInChain != nil {", indent, a.Name),
			fmt.Sprintf("%s\tpinner.Pin(%s.NextInChain)", indent, a.Name),
			fmt.Sprintf("%s}", indent))
	}
	for _, f := range a.StrFields {
		lines = append(lines,
			fmt.Sprintf("%sif len(%s) > 0 {", indent, f),
			fmt.Sprintf("%s\tpinner.Pin(unsafe.StringData(%s))", indent, f),
			fmt.Sprintf("%s}", indent))
	}
	for _, f := range a.PtrFields {
		lines = append(lines,
			fmt.Sprintf("%sif %s != nil {", indent, f),
			fmt.Sprintf("%s\tpinner.Pin(%s)", indent, f),
			fmt.Sprintf("%s}", indent))
	}
	for _, f := range a.SliceFields {
		lines = append(lines,
			fmt.Sprintf("%sif %s != nil {", indent, f),
			fmt.Sprintf("%s\tpinner.Pin(%s)", indent, f),
			fmt.Sprintf("%s}", indent))
	}
	return lines
}

// convLines emits per-argument conversions that differ between backends
// (string views, and the C-typed locals cgo needs).
func convLines(fd FuncData, be Backend) []string {
	var lines []string
	for _, a := range fd.GoArgs {
		if !a.IsStr {
			continue
		}
		if be == BackendCGO {
			lines = append(lines, fmt.Sprintf("c_%s := cgoStringView(%s, &pinner)", a.Name, a.Name))
		} else {
			lines = append(lines, fmt.Sprintf("%s_sv := syscallStringView(%s, &pinner)", a.Name, a.Name))
		}
	}
	if be == BackendCGO {
		for _, a := range fd.GoArgs {
			if a.IsStr {
				// Already emitted as cgoStringView above.
				continue
			}
			lines = append(lines, cgoLocal(a)...)
		}
	}
	return lines
}

// cgoLocal emits the C-typed local for one cgo argument.
func cgoLocal(a FuncArgData) []string {
	n := a.Name
	switch {
	case a.IsCBInfo:
		// By-value C struct: cast the pointer and dereference.
		return []string{fmt.Sprintf("c_%s := *(*C.WGPU%s)(unsafe.Pointer(&%s))", n, a.GoType, n)}
	case isHandleType(a.TypeRef):
		return []string{fmt.Sprintf("c_%s := (%s)(unsafe.Pointer(%s))", n, cgoTypeName(a.GoType), n)}
	case a.IsSlice:
		return []string{
			fmt.Sprintf("var c_%s_ptr %s", n, a.SliceCPtr),
			fmt.Sprintf("if len(%s) > 0 {", n),
			fmt.Sprintf("\tc_%s_ptr = (%s)(unsafe.Pointer(&%s[0]))", n, a.SliceCPtr, n),
			"}",
			fmt.Sprintf("c_%s_count := (C.size_t)(len(%s))", n, n),
		}
	case isStructType(a.TypeRef):
		if a.IsPtr {
			return []string{fmt.Sprintf("c_%s := (*%s)(unsafe.Pointer(%s))", n, cgoTypeName(a.GoType), n)}
		}
		return []string{fmt.Sprintf("c_%s := (*%s)(unsafe.Pointer(&%s))", n, cgoTypeName(a.GoType), n)}
	default:
		return []string{fmt.Sprintf("c_%s := (%s)(%s)", n, cgoTypeName(a.GoType), n)}
	}
}

// callLines emits the invocation and the return statement.
func callLines(fd FuncData, be Backend) []string {
	var args []string
	for _, a := range fd.GoArgs {
		args = append(args, callArg(a, be))
	}
	joined := strings.Join(args, ", ")

	if be == BackendCGO {
		call := fmt.Sprintf("C.%s(%s)", fd.CName, joined)
		if fd.ReturnRef == "" {
			return []string{call}
		}
		return []string{"c_result := " + call, returnStmt(fd, "c_result")}
	}

	call := fmt.Sprintf("proc%s.Call(%s)", fd.Ident, joined)
	if fd.ReturnRef == "" {
		return []string{call}
	}
	return []string{"r1, _, _ := " + call, returnStmt(fd, "r1")}
}

// callArg renders one argument at the call site.
func callArg(a FuncArgData, be Backend) string {
	cgo := be == BackendCGO
	switch {
	case a.IsSlice:
		if cgo {
			return fmt.Sprintf("c_%s_count, c_%s_ptr", a.Name, a.Name)
		}
		return fmt.Sprintf("uintptr(len(%s)), uintptr(slicePtr(%s))", a.Name, a.Name)
	case a.IsCBInfo:
		if cgo {
			return "c_" + a.Name
		}
		return fmt.Sprintf("uintptr(unsafe.Pointer(&%s))", a.Name)
	case isHandleType(a.TypeRef):
		if cgo {
			return "c_" + a.Name
		}
		return fmt.Sprintf("uintptr(unsafe.Pointer(%s))", a.Name)
	case a.IsStr:
		if cgo {
			return "c_" + a.Name
		}
		return fmt.Sprintf("uintptr(unsafe.Pointer(&%s_sv))", a.Name)
	case isStructType(a.TypeRef):
		if cgo {
			return "c_" + a.Name
		}
		if a.IsPtr {
			return fmt.Sprintf("uintptr(unsafe.Pointer(%s))", a.Name)
		}
		return fmt.Sprintf("uintptr(unsafe.Pointer(&%s))", a.Name)
	default:
		if cgo {
			return "c_" + a.Name
		}
		return fmt.Sprintf("uintptr(%s)", a.Name)
	}
}

// returnStmt converts the raw call result to fd's Go return type.
func returnStmt(fd FuncData, src string) string {
	switch fd.RetKind {
	case "handle":
		return fmt.Sprintf("return %s(unsafe.Pointer(%s))", fd.GoReturn, src)
	case "status":
		return fmt.Sprintf("return Status(%s)", src)
	case "wait":
		return fmt.Sprintf("return WaitStatus(%s)", src)
	case "pointer":
		return fmt.Sprintf("return unsafe.Pointer(%s)", src)
	case "bool":
		return fmt.Sprintf("return Bool(%s)", src)
	case "future":
		// cgo yields the WGPUFuture struct; syscall yields the id register.
		if src == "c_result" {
			return fmt.Sprintf("return Future{Id: uint64(%s.id)}", src)
		}
		return fmt.Sprintf("return Future{Id: uint64(%s)}", src)
	default:
		return fmt.Sprintf("return (%s)(%s)", fd.GoReturn, src)
	}
}

// tabBody indents body lines one level for embedding inside a func.
func tabBody(body []string) []string {
	out := make([]string, len(body))
	for i, l := range body {
		if l == "" {
			out[i] = l
			continue
		}
		out[i] = "\t" + l
	}
	return out
}
