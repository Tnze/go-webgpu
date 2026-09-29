package main

import (
	"fmt"
	"strings"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// C API signature emission for the dynload shim.
//
// The shim is a C file that defines every wgpu* entry point as a thin forwarder
// through a dlsym'd function pointer. It must reproduce the exact C prototype
// from webgpu.h, which we derive from the YAML type model.
// ---------------------------------------------------------------------------

// cTypeRef maps a YAML type reference to the C type spelling used by webgpu.h.
// pointer is one of "", "immutable", "mutable"; optional adds WGPU_NULLABLE.
func cTypeRef(ref, pointer string, optional bool) string {
	nullable := ""
	if optional {
		nullable = "WGPU_NULLABLE "
	}
	switch {
	case ref == "void":
		return "void"
	case isStringType(ref):
		return "WGPUStringView"
	case ref == "bool":
		return "WGPUBool"
	case ref == "usize":
		return "size_t"
	case ref == "nullable_float32":
		return nullable + "float *"
	case strings.HasPrefix(ref, "c_void"):
		if pointer == "immutable" {
			return nullable + "void const *"
		}
		return nullable + "void *"
	case ref == "float64_supertype":
		return "double"
	case ref == "uint8", ref == "uint16", ref == "uint32", ref == "uint64",
		ref == "int32", ref == "int64", ref == "float32", ref == "float64":
		return cPrimType(ref)
	}

	switch {
	case strings.HasPrefix(ref, "enum."):
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "enum."))
	case strings.HasPrefix(ref, "bitflag."):
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "bitflag."))
	case strings.HasPrefix(ref, "object."):
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "object."))
	case strings.HasPrefix(ref, "struct."):
		name := "WGPU" + pascalCase(strings.TrimPrefix(ref, "struct."))
		switch pointer {
		case "immutable":
			return nullable + name + " const *"
		case "mutable":
			return nullable + name + " *"
		default:
			return name
		}
	case strings.HasPrefix(ref, "callback."):
		// By-value callback-info records (WGPUNnnCallbackInfo).
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "callback.")) + "CallbackInfo"
	}
	return "void *"
}

func cPrimType(ref string) string {
	switch ref {
	case "uint8":
		return "uint8_t"
	case "uint16":
		return "uint16_t"
	case "uint32":
		return "uint32_t"
	case "uint64":
		return "uint64_t"
	case "int32":
		return "int32_t"
	case "int64":
		return "int64_t"
	case "float32":
		return "float"
	case "float64", "float64_supertype":
		return "double"
	}
	return "void *"
}

// cReturnFrom maps a YAML return type to its C spelling.
func cReturnFrom(returns *parser.ReturnType, hasCallback bool) string {
	if hasCallback {
		return "WGPUFuture"
	}
	if returns == nil || returns.Type == "" || returns.Type == "void" {
		return "void"
	}
	ref := returns.Type
	switch {
	case ref == "bool":
		return "WGPUBool"
	case ref == "usize":
		return "size_t"
	case isStringType(ref):
		return "WGPUStringView"
	case strings.HasPrefix(ref, "c_void"):
		if returns.Pointer == "immutable" {
			return "void const *"
		}
		return "void *"
	case strings.HasPrefix(ref, "enum."):
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "enum."))
	case strings.HasPrefix(ref, "bitflag."):
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "bitflag."))
	case strings.HasPrefix(ref, "object."):
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "object."))
	case strings.HasPrefix(ref, "struct."):
		return "WGPU" + pascalCase(strings.TrimPrefix(ref, "struct."))
	case ref == "nullable_float32":
		return "float *"
	default:
		return cPrimType(ref)
	}
}

// CArg is one parameter of a generated C prototype. Arrays expand to a
// count/pointer pair, matching the C header.
type CArg struct {
	Name string
	Type string
}

// CSignature is the C prototype of one wgpu* entry point.
type CSignature struct {
	Name   string
	Return string
	Args   []CArg
}

// Prototype renders "RET name(T0 a0, T1 a1)".
func (s CSignature) Prototype() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s %s(", s.Return, s.Name)
	for i, a := range s.Args {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%s %s", a.Type, a.Name)
	}
	b.WriteString(")")
	return b.String()
}

// CallArgs renders the argument list for a forwarder call.
func (s CSignature) CallArgs() string {
	names := make([]string, len(s.Args))
	for i, a := range s.Args {
		names[i] = a.Name
	}
	return strings.Join(names, ", ")
}

// PointerDecl renders the function-pointer variable declaration.
func (s CSignature) PointerDecl() string {
	var b strings.Builder
	fmt.Fprintf(&b, "static %s (*p_%s)(", s.Return, s.Name)
	for i, a := range s.Args {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(a.Type)
	}
	b.WriteString(")")
	return b.String()
}

// appendArrayArgs expands array<T> into count + pointer parameters.
func appendArrayArgs(dst []CArg, a parser.FunctionArg) []CArg {
	if isArrayType(a.Type) {
		inner := extractArrayInner(a.Type)
		return append(dst,
			CArg{Name: a.Name + "_count", Type: "size_t"},
			CArg{Name: a.Name, Type: cTypeRef(inner, "immutable", a.Optional)},
		)
	}
	return append(dst, CArg{
		Name: a.Name,
		Type: cTypeRef(a.Type, a.Pointer, a.Optional),
	})
}

// buildCSignature derives the C prototype of a free function.
func buildCSignature(f parser.Function) CSignature {
	sig := CSignature{
		Name:   "wgpu" + pascalCase(f.Name),
		Return: cReturnFrom(f.Returns, f.Callback != ""),
	}
	for _, a := range f.Args {
		sig.Args = appendArrayArgs(sig.Args, a)
	}
	return sig
}

// buildMethodCSignature derives the C prototype of an object method: the
// receiver is prepended and the symbol name gains the object prefix.
func buildMethodCSignature(obj parser.Object, m parser.Function) CSignature {
	sig := CSignature{
		Name:   "wgpu" + pascalCase(obj.Name) + pascalCase(m.Name),
		Return: cReturnFrom(m.Returns, m.Callback != ""),
		Args:   []CArg{{Name: obj.Name, Type: "WGPU" + pascalCase(obj.Name)}},
	}
	for _, a := range m.Args {
		sig.Args = appendArrayArgs(sig.Args, a)
	}
	return sig
}

// releaseCSignature builds the synthetic Release prototype for an object.
func releaseCSignature(objName string) CSignature {
	return CSignature{
		Name:   "wgpu" + objName + "Release",
		Return: "void",
		Args:   []CArg{{Name: "value", Type: "WGPU" + objName}},
	}
}

// freeMembersCSignature builds the WGPUNnnFreeMembers prototype.
func freeMembersCSignature(structName string) CSignature {
	return CSignature{
		Name:   "wgpu" + structName + "FreeMembers",
		Return: "void",
		Args:   []CArg{{Name: "value", Type: "WGPU" + structName}},
	}
}
