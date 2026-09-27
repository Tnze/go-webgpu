package main

import (
	"fmt"
	"strings"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// Wrapper method generation (package gpu).
//
// Each method precomputes: parameter list, low-level call arguments,
// conversion pre/post statements, and the return-shape body.
// ---------------------------------------------------------------------------

func buildWrapHandle(obj parser.Object, spec *parser.Spec) WrapHandle {
	name := pascalCase(obj.Name)
	wh := WrapHandle{Name: name, Doc: cleanDoc(obj.Doc)}
	for _, m := range obj.Methods {
		wh.Methods = append(wh.Methods, buildWrapMethod(name, m, spec))
	}
	return wh
}

func buildCreateInstance() WrapMethod {
	return WrapMethod{
		Name: "CreateInstance",
		Sig:  "func CreateInstance(desc *InstanceDescriptor) (*Instance, error)",
		Body: []string{
			"raw := desc.toRaw()",
			"h := webgpu.CreateInstance(&raw)",
			"runtime.KeepAlive(desc)",
			"if h == nil {",
			"\treturn nil, ErrCreate",
			"}",
			"return newInstance(h), nil",
		},
	}
}

// wrapParam is one wrapper method parameter plus its low-level conversion.
type wrapParam struct {
	name, typ string
}

// wrapResult is one wrapper method result.
type wrapResult struct {
	name, typ, wrap string
}

func buildWrapMethod(recv string, m parser.Function, spec *parser.Spec) WrapMethod {
	methodName := pascalCase(m.Name)
	llName := recv + pascalCase(m.Name)
	if m.Name == "map_async" {
		methodName, llName = "Map", recv+"Map"
	}
	if methodName == "Release" {
		return WrapMethod{Name: "Release"} // emitted by the handle block
	}

	// Special-cased out-param APIs.
	if recv == "Surface" && methodName == "GetCurrentTexture" {
		return WrapMethod{
			Name: methodName,
			Sig:  "func (h *Surface) GetCurrentTexture() (*Texture, webgpu.SurfaceGetCurrentTextureStatus)",
			Body: []string{
				"var st webgpu.SurfaceTexture",
				"webgpu.SurfaceGetCurrentTexture(h.h, &st)",
				"return newTexture(st.Texture), st.Status",
			},
		}
	}
	if recv == "Surface" && methodName == "GetCapabilities" {
		return WrapMethod{
			Name: methodName,
			Sig:  "func (h *Surface) GetCapabilities(adapter *Adapter) (*SurfaceCapabilities, webgpu.Status)",
			Body: []string{
				"var caps webgpu.SurfaceCapabilities",
				"status := webgpu.SurfaceGetCapabilities(h.h, adapter.raw(), &caps)",
				"out := surfaceCapabilitiesFromRaw(&caps)",
				"runtime.KeepAlive(adapter)",
				"return out, status",
			},
		}
	}

	// --- parameters ---
	var params []wrapParam
	llArgs := []string{"h.h"}
	var pre, post []string
	var skipNextSize bool
	var prevDataName string

	for _, a := range m.Args {
		if strings.HasPrefix(a.Type, "callback.") {
			continue
		}
		n := camelCase(a.Name)

		switch {
		case isHandleType(a.Type):
			hn := pascalCase(strings.TrimPrefix(a.Type, "object."))
			params = append(params, wrapParam{n, "*" + hn})
			llArgs = append(llArgs, n+".raw()")

		case isStructType(a.Type):
			sn := pascalCase(strings.TrimPrefix(a.Type, "struct."))
			raw := n + "Raw"
			if a.Pointer == "immutable" || a.Pointer == "mutable" {
				params = append(params, wrapParam{n, "*" + sn})
				pre = append(pre,
					fmt.Sprintf("var %s *webgpu.%s", raw, sn),
					fmt.Sprintf("if %s != nil {", n),
					fmt.Sprintf("\tt := %s.toRaw()", n),
					fmt.Sprintf("\t%s = &t", raw),
					"}",
				)
				llArgs = append(llArgs, raw)
			} else {
				params = append(params, wrapParam{n, sn})
				pre = append(pre, fmt.Sprintf("%s := %s.toRaw()", raw, n))
				llArgs = append(llArgs, "&"+raw)
			}
			post = append(post, "runtime.KeepAlive("+n+")")

		case isArrayType(a.Type):
			typ, convPre, convArg, convPost := wrapSliceArg(n, extractArrayInner(a.Type), spec)
			params = append(params, wrapParam{n, typ})
			pre = append(pre, convPre...)
			llArgs = append(llArgs, convArg)
			post = append(post, convPost...)

		case isStringType(a.Type):
			params = append(params, wrapParam{n, "string"})
			llArgs = append(llArgs, n)

		case strings.HasPrefix(a.Type, "c_void") && n == "data":
			// (data, dataSize) pair becomes a single []byte parameter.
			params = append(params, wrapParam{n, "[]byte"})
			pre = append(pre,
				fmt.Sprintf("var %sPtr unsafe.Pointer", n),
				fmt.Sprintf("if len(%s) > 0 { %sPtr = unsafe.Pointer(&%s[0]) }", n, n, n),
			)
			llArgs = append(llArgs, n+"Ptr")
			post = append(post, "runtime.KeepAlive("+n+")")
			skipNextSize, prevDataName = true, n

		case skipNextSize && (a.Type == "usize" || a.Type == "uint64"):
			llArgs = append(llArgs, "uintptr(len("+prevDataName+"))")
			skipNextSize = false // size is derived from the []byte param

		case strings.HasPrefix(a.Type, "c_void"):
			params = append(params, wrapParam{n, "unsafe.Pointer"})
			llArgs = append(llArgs, n)

		default:
			gt := goTypeForRef(a.Type, spec)
			if gt == "Bool" || a.Type == "bool" {
				params = append(params, wrapParam{n, "bool"})
				tmp := n + "Raw"
				pre = append(pre,
					fmt.Sprintf("var %s webgpu.Bool", tmp),
					fmt.Sprintf("if %s { %s = webgpu.True }", n, tmp),
				)
				llArgs = append(llArgs, tmp)
			} else {
				params = append(params, wrapParam{n, gt})
				llArgs = append(llArgs, n)
			}
		}
	}

	var ps []string
	for _, p := range params {
		ps = append(ps, p.name+" "+p.typ)
	}
	sig := fmt.Sprintf("func (h *%s) %s(%s)", recv, methodName, strings.Join(ps, ", "))
	llCall := fmt.Sprintf("webgpu.%s(%s)", llName, strings.Join(llArgs, ", "))

	// --- results ---
	var results []wrapResult
	style := "void"

	switch {
	case m.Callback != "":
		if cb := spec.GetCallback(strings.TrimPrefix(m.Callback, "callback.")); cb != nil {
			for _, ra := range goCallbackArgs(*cb) {
				t := ra.GoType
				switch {
				case isKnownHandle(spec, t):
					results = append(results, wrapResult{lowerFirst(t), "*" + t, ""})
				case t == "error" || t == "string":
					results = append(results, wrapResult{ra.Name, t, ""})
				default:
					results = append(results, wrapResult{ra.Name, "webgpu." + t, ""})
				}
			}
		}
		style = "status_tuple"

	case m.Returns != nil && m.Returns.Type != "" && m.Returns.Type != "void":
		r := m.Returns.Type
		switch {
		case isHandleType(r):
			hn := pascalCase(strings.TrimPrefix(r, "object."))
			results = append(results, wrapResult{lowerFirst(hn), "*" + hn, "new" + hn + "(rv)"})
			if strings.HasPrefix(methodName, "Create") || methodName == "Finish" {
				results = append(results, wrapResult{"err", "error", ""})
				style = "handle_error"
			} else {
				style = "handle"
			}
		case r == "enum.status":
			results, style = append(results, wrapResult{"status", "webgpu.Status", "rv"}), "status"
		case r == "enum.wait_status":
			results, style = append(results, wrapResult{"status", "webgpu.WaitStatus", "rv"}), "status"
		case r == "bool":
			results, style = append(results, wrapResult{"ok", "bool", "rv == webgpu.True"}), "bool"
		case r == "struct.compilation_info":
			results, style = append(results, wrapResult{"info", "*CompilationInfo", "compilationInfoFromRaw(rv)"}), "compilation_info"
		default:
			gt := goTypeForRef(r, spec)
			switch {
			case gt == "Bool":
				results, style = append(results, wrapResult{"ok", "bool", "rv == webgpu.True"}), "bool"
			case isStructType(r):
				results, style = append(results, wrapResult{"v", "webgpu." + gt, "rv"}), "values"
			default:
				results, style = append(results, wrapResult{"v", gt, "rv"}), "values"
			}
		}
	}

	switch len(results) {
	case 0:
	case 1:
		sig += " " + results[0].typ
	default:
		var rs []string
		for _, r := range results {
			rs = append(rs, r.name+" "+r.typ)
		}
		sig += " (" + strings.Join(rs, ", ") + ")"
	}

	// --- body ---
	body := append([]string{}, pre...)
	body = append(body, wrapBody(style, llCall, results, post, spec)...)
	return WrapMethod{Name: methodName, Doc: cleanDoc(m.Doc), Sig: sig, Body: body}
}

// wrapSliceArg returns the wrapper type and conversion code for a []T argument.
func wrapSliceArg(name, inner string, spec *parser.Spec) (typ string, pre []string, arg string, post []string) {
	switch {
	case isHandleType(inner):
		hn := pascalCase(strings.TrimPrefix(inner, "object."))
		tmp := name + "Raw"
		return "[]*" + hn,
			[]string{
				fmt.Sprintf("%s := make([]webgpu.%s, len(%s))", tmp, hn, name),
				fmt.Sprintf("for i := range %s { %s[i] = %s[i].raw() }", tmp, tmp, name),
			},
			tmp,
			[]string{"runtime.KeepAlive(" + name + ")"}
	case isStructType(inner):
		sn := pascalCase(strings.TrimPrefix(inner, "struct."))
		tmp := name + "Raw"
		return "[]" + sn,
			[]string{
				fmt.Sprintf("%s := make([]webgpu.%s, len(%s))", tmp, sn, name),
				fmt.Sprintf("for i := range %s { %s[i] = %s[i].toRaw() }", tmp, tmp, name),
			},
			tmp,
			[]string{"runtime.KeepAlive(" + name + ")"}
	case inner == "bool" || goTypeForRef(inner, spec) == "Bool":
		tmp := name + "Raw"
		return "[]bool",
			[]string{
				fmt.Sprintf("%s := make([]webgpu.Bool, len(%s))", tmp, name),
				fmt.Sprintf("for i := range %s { if %s[i] { %s[i] = webgpu.True } }", tmp, name, tmp),
			},
			tmp,
			[]string{"runtime.KeepAlive(" + name + ")"}
	default:
		return "[]" + goTypeForRef(inner, spec), nil, name, nil
	}
}

// wrapBody emits the return-shape body after the conversion pre-statements.
func wrapBody(style, llCall string, results []wrapResult, post []string, spec *parser.Spec) []string {
	switch style {
	case "void":
		return append([]string{llCall}, post...)

	case "handle_error":
		return append([]string{
			"rv := " + llCall,
		}, append(post,
			"if rv == nil {",
			"\treturn nil, ErrCreate",
			"}",
			"return "+results[0].wrap+", nil",
		)...)

	case "handle":
		return append([]string{"rv := " + llCall}, append(post, "return "+results[0].wrap)...)

	case "status", "values":
		return append([]string{"rv := " + llCall}, append(post, "return rv")...)

	case "bool":
		return append([]string{"rv := " + llCall}, append(post, "return rv == webgpu.True")...)

	case "compilation_info":
		return append([]string{"info, status := " + llCall}, append(post, "return compilationInfoFromRaw(info), status")...)

	case "status_tuple":
		var names, rets []string
		for i, r := range results {
			local := fmt.Sprintf("r%d", i)
			names = append(names, local)
			switch {
			case strings.HasPrefix(r.typ, "*") && isKnownHandle(spec, strings.TrimPrefix(r.typ, "*")):
				rets = append(rets, "new"+strings.TrimPrefix(r.typ, "*")+"("+local+")")
			case r.typ == "bool":
				rets = append(rets, local+" == webgpu.True")
			default:
				rets = append(rets, local)
			}
		}
		out := append([]string{strings.Join(names, ", ") + " := " + llCall}, post...)
		if len(rets) > 0 {
			out = append(out, "return "+strings.Join(rets, ", "))
		}
		return out
	}
	return nil
}
