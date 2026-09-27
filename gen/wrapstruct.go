package main

import (
	"fmt"
	"strings"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// Wrapper struct mirrors and field conversions (package gpu).
// ---------------------------------------------------------------------------

func buildWrapStruct(s parser.Struct, spec *parser.Spec) WrapStruct {
	name := pascalCase(s.Name)
	ws := WrapStruct{Name: name, Doc: cleanDoc(s.Doc), HasFree: s.FreeMembers}
	for _, m := range s.Members {
		if m.Default != nil {
			ws.HasDefault = true
		}
	}

	// Extension-chain conveniences that don't map 1:1 to YAML members.
	switch name {
	case "ShaderModuleDescriptor":
		ws.ExtraDoc = "// WGSL is the WGSL shader source, chained as ShaderSourceWGSL."
		ws.ExtraField = "WGSL *ShaderSourceWGSL"
		ws.ExtraToRaw = []string{
			"if s.WGSL != nil {",
			"\tv := webgpu.NewShaderSourceWGSL()",
			"\tv.Code = s.WGSL.Code",
			"\traw.NextInChain = unsafe.Pointer(&v)",
			"}",
		}
	case "SurfaceDescriptor":
		ws.ExtraDoc = "// Source is the platform window/layer to bind, e.g. *SurfaceSourceWindowsHWND."
		ws.ExtraField = "Source SurfaceSource"
		ws.ExtraToRaw = []string{
			"switch src := s.Source.(type) {",
			"case *SurfaceSourceWindowsHWND:",
			"\tv := webgpu.NewSurfaceSourceWindowsHWND()",
			"\tv.Hinstance = src.Hinstance",
			"\tv.Hwnd = src.Hwnd",
			"\traw.NextInChain = unsafe.Pointer(&v)",
			"case *SurfaceSourceMetalLayer:",
			"\tv := webgpu.NewSurfaceSourceMetalLayer()",
			"\tv.Layer = src.Layer",
			"\traw.NextInChain = unsafe.Pointer(&v)",
			"}",
		}
	}

	for _, m := range s.Members {
		if wf, ok := buildWrapField(m, spec); ok {
			ws.Fields = append(ws.Fields, wf)
		}
	}
	return ws
}

// buildWrapField mirrors one struct member. Returns ok=false for fields that
// are hidden in the wrapper (chain headers, callback infos).
func buildWrapField(m parser.StructMember, spec *parser.Spec) (WrapField, bool) {
	name := pascalCase(m.Name)
	if name == "NextInChain" || name == "SType" {
		return WrapField{}, false
	}
	if strings.HasPrefix(m.Type, "callback.") {
		return WrapField{}, false
	}
	doc := cleanDoc(m.Doc)
	src, dst := "s."+name, "raw."+name

	if isArrayType(m.Type) {
		return buildWrapSliceField(name, doc, src, dst, extractArrayInner(m.Type), spec), true
	}

	wf := WrapField{Name: name, Doc: doc}
	switch {
	case m.Type == "bool":
		wf.Type = "bool"
		wf.ToRaw = []string{fmt.Sprintf("if %s { %s = webgpu.True }", src, dst)}
		wf.FromRaw = []string{fmt.Sprintf("out.%s = %s == webgpu.True", name, dst)}

	case isStringType(m.Type):
		wf.Type = "string"
		wf.ToRaw = []string{fmt.Sprintf("%s = %s", dst, src)}
		wf.FromRaw = []string{fmt.Sprintf("out.%s = %s", name, dst)}

	case isHandleType(m.Type):
		hn := pascalCase(strings.TrimPrefix(m.Type, "object."))
		wf.Type = "*" + hn
		wf.ToRaw = []string{fmt.Sprintf("if %s != nil { %s = %s.raw() }", src, dst, src)}
		wf.FromRaw = []string{fmt.Sprintf("out.%s = new%s(%s)", name, hn, dst)}

	case isStructType(m.Type):
		sn := pascalCase(strings.TrimPrefix(m.Type, "struct."))
		if m.Optional || m.Pointer == "immutable" || m.Pointer == "mutable" {
			wf.Type = "*" + sn
			wf.ToRaw = []string{
				fmt.Sprintf("if %s != nil {", src),
				fmt.Sprintf("\tt := %s.toRaw()", src),
				fmt.Sprintf("\t%s = &t", dst),
				"}",
			}
			wf.FromRaw = []string{fmt.Sprintf("if %s != nil { out.%s = &%s }", dst, name, "*"+sn+"{}")}
		} else {
			wf.Type = sn
			wf.ToRaw = []string{fmt.Sprintf("%s = %s.toRaw()", dst, src)}
			wf.FromRaw = []string{fmt.Sprintf("out.%s = %s{}; _ = %s", name, sn, dst)}
		}

	case m.Type == "nullable_float32", strings.HasPrefix(m.Type, "c_void"):
		wf.Type = "unsafe.Pointer"
		if m.Type == "nullable_float32" {
			wf.Type = "*float32"
		}
		wf.ToRaw = []string{fmt.Sprintf("%s = %s", dst, src)}
		wf.FromRaw = []string{fmt.Sprintf("out.%s = %s", name, dst)}

	default:
		gt := goTypeForRef(m.Type, spec)
		if gt == "Bool" {
			wf.Type = "bool"
			wf.ToRaw = []string{fmt.Sprintf("if %s { %s = webgpu.True }", src, dst)}
			wf.FromRaw = []string{fmt.Sprintf("out.%s = %s == webgpu.True", name, dst)}
		} else {
			wf.Type = gt
			if hasNonZeroDefault(m) {
				wf.SkipZero = true
				wf.ToRaw = []string{fmt.Sprintf("if %s != 0 { %s = %s }", src, dst, src)}
			} else {
				wf.ToRaw = []string{fmt.Sprintf("%s = %s", dst, src)}
			}
			wf.FromRaw = []string{fmt.Sprintf("out.%s = %s", name, dst)}
		}
	}
	return wf, true
}

// buildWrapSliceField folds a Count+Pointer pair into a single slice field.
func buildWrapSliceField(name, doc, src, dst, inner string, spec *parser.Spec) WrapField {
	wf := WrapField{Name: name, Doc: doc}
	toRaw := func(elemType, convert string) []string {
		return []string{
			fmt.Sprintf("if len(%s) > 0 {", src),
			fmt.Sprintf("\tvs := make([]%s, len(%s))", elemType, src),
			fmt.Sprintf("\tfor i := range %s { vs[i] = %s[i]%s }", src, src, convert),
			fmt.Sprintf("\t%sCount = uintptr(len(vs))", dst),
			fmt.Sprintf("\t%s = &vs[0]", dst),
			"}",
		}
	}
	fromLoop := func(conv string) []string {
		return []string{
			fmt.Sprintf("for _, v := range unsafe.Slice(%s, %sCount) {", dst, dst),
			fmt.Sprintf("\tout.%s = append(out.%s, %s)", name, name, conv),
			"}",
		}
	}

	switch {
	case isHandleType(inner):
		hn := pascalCase(strings.TrimPrefix(inner, "object."))
		wf.Type = "[]*" + hn
		wf.ToRaw = toRaw("webgpu."+hn, ".raw()")
		wf.FromRaw = fromLoop("new" + hn + "(v)")
	case isStructType(inner):
		sn := pascalCase(strings.TrimPrefix(inner, "struct."))
		wf.Type = "[]" + sn
		wf.ToRaw = toRaw("webgpu."+sn, ".toRaw()")
		wf.FromRaw = fromLoop("v")
	case inner == "bool":
		wf.Type = "[]bool"
		wf.ToRaw = []string{
			fmt.Sprintf("if len(%s) > 0 {", src),
			fmt.Sprintf("\tvs := make([]webgpu.Bool, len(%s))", src),
			fmt.Sprintf("\tfor i := range %s { if %s[i] { vs[i] = webgpu.True } }", src, src),
			fmt.Sprintf("\t%sCount = uintptr(len(vs))", dst),
			fmt.Sprintf("\t%s = &vs[0]", dst),
			"}",
		}
		wf.FromRaw = fromLoop("v == webgpu.True")
	default:
		gt := goTypeForRef(inner, spec)
		wf.Type = "[]" + gt
		wf.ToRaw = []string{
			fmt.Sprintf("if len(%s) > 0 {", src),
			fmt.Sprintf("\t%sCount = uintptr(len(%s))", dst, src),
			fmt.Sprintf("\t%s = &%s[0]", dst, src),
			"}",
		}
		wf.FromRaw = []string{
			fmt.Sprintf("out.%s = append(out.%s[:0:0], unsafe.Slice(%s, %sCount)...)", name, name, dst, dst),
		}
	}
	return wf
}

// hasNonZeroDefault reports whether the member's YAML default is a non-zero
// sentinel (so toRaw should not overwrite it with the Go zero value).
func hasNonZeroDefault(m parser.StructMember) bool {
	if m.Default == nil {
		return false
	}
	switch v := m.Default.(type) {
	case int:
		return v != 0
	case float64:
		return v != 0
	case bool:
		return v
	case string:
		return strings.HasPrefix(v, "constant.")
	}
	return false
}
