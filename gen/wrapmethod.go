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

func wrapCBResults(cb parser.Callback, spec *parser.Spec) []wrapResult {
	var out []wrapResult
	for _, ra := range goCallbackArgs(cb) {
		t := ra.GoType
		switch {
		case isKnownHandle(spec, t):
			out = append(out, wrapResult{lowerFirst(t), "*" + t, ""})
		case t == "error":
			out = append(out, wrapResult{"err", "error", ""})
		case t == "string":
			out = append(out, wrapResult{ra.Name, "string", ""})
		case t == "CompilationInfo":
			out = append(out, wrapResult{ra.Name, "*CompilationInfo", ""})
		default:
			out = append(out, wrapResult{ra.Name, "webgpu." + t, ""})
		}
	}
	return out
}

// pureCBType is the webgpu-level Go type of one C callback argument.
func pureCBType(a parser.CallbackArg) string {
	switch {
	case isStringType(a.Type):
		return "string"
	case isHandleType(a.Type):
		return pascalCase(strings.TrimPrefix(a.Type, "object."))
	case strings.HasPrefix(a.Type, "enum."):
		return pascalCase(strings.TrimPrefix(a.Type, "enum."))
	case isStructType(a.Type) && (a.Pointer == "immutable" || a.Pointer == "mutable"):
		return "*" + pascalCase(strings.TrimPrefix(a.Type, "struct."))
	default:
		return goTypeForRef(a.Type, nil)
	}
}

// webgpuQualify prefixes a named webgpu type with its package name so it can
// be referenced from the gpu package. Builtins and pointers are handled.
func webgpuQualify(t string) string {
	if t == "string" || t == "unsafe.Pointer" || isPrimitiveGoType(t) {
		return t
	}
	if strings.HasPrefix(t, "*") {
		return "*" + webgpuQualify(t[1:])
	}
	return "webgpu." + t
}

// modeArg is the leading mode argument of NewXxxCallbackInfo (none for
// callbacks without a configurable mode).
func modeArg(base string) string {
	if base == "UncapturedError" {
		return ""
	}
	return "webgpu.CallbackModeAllowProcessEvents, "
}

// wrapResultExprs returns the Go expressions that turn the pure callback
// parameters (C order) into the gpu-layer results (wrapCBResults order).
func wrapResultExprs(cb parser.Callback, spec *parser.Spec) []string {
	// Name of each pure parameter (C order).
	pure := map[string]string{} // yaml name → Go param name
	for _, a := range cb.Args {
		pure[a.Name] = camelCase(a.Name)
	}
	var errorMsg string
	for _, a := range cb.Args {
		if isStringType(a.Type) && a.Name == "message" {
			errorMsg = pure[a.Name]
		}
	}

	var exprs []string
	skipNext := false
	// Walk the same shape as goCallbackArgs to build expressions.
	var statusExpr, messageExpr string
	for _, a := range cb.Args {
		if skipNext {
			skipNext = false
			continue
		}
		switch {
		case a.Name == "status" && strings.HasPrefix(a.Type, "enum."):
			statusExpr = pure[a.Name]
		case a.Type == "enum.error_type":
			exprs = append(exprs, fmt.Sprintf("errorFromErrorType(webgpu.ErrorType(%s), %s)", pure[a.Name], errorMsg))
			skipNext = true
		case a.Name == "message" && isStringType(a.Type):
			messageExpr = pure[a.Name]
		default:
			t := goTypeForRef(a.Type, nil)
			switch {
			case isKnownHandle(spec, t):
				exprs = append(exprs, fmt.Sprintf("new%s(%s)", t, pure[a.Name]))
			case isStructType(a.Type):
				exprs = append(exprs, fmt.Sprintf("compilationInfoFromRawPtr(unsafe.Pointer(%s))", pure[a.Name]))
			default:
				exprs = append(exprs, pure[a.Name])
			}
		}
	}
	if statusExpr != "" {
		exprs = append(exprs, statusExpr)
	}
	if messageExpr != "" {
		exprs = append(exprs, messageExpr)
	}
	return exprs
}

func buildWrapMethod(recv string, m parser.Function, spec *parser.Spec) WrapMethod {
	methodName := pascalCase(m.Name)
	llName := recv + pascalCase(m.Name)
	if m.Name == "map_async" {
		// Low-level keeps the C name (BufferMapAsync); only the gpu method is Map.
		methodName = "Map"
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
		// Blocking sync-over-async: install a Go closure via
		// webgpu.NewXxxCallbackInfo (backend-uniform), call the raw primitive,
		// then pump events until the callback fires.
		cbName := strings.TrimPrefix(m.Callback, "callback.")
		cb := spec.GetCallback(cbName)
		if cb == nil {
			return WrapMethod{Name: methodName, Sig: sig}
		}
		base := pascalCase(cbName)
		results = append(results, wrapCBResults(*cb, spec)...)
		sig = attachResults(sig, results)

		// Closure parameters follow the C signature (pure webgpu types,
		// qualified with the webgpu package name).
		var pureParams []string
		for _, a := range cb.Args {
			pureParams = append(pureParams, camelCase(a.Name)+" "+webgpuQualify(pureCBType(a)))
		}
		exprs := wrapResultExprs(*cb, spec)

		var fields, assigns []string
		for _, r := range results {
			fields = append(fields, r.name+" "+r.typ)
			assigns = append(assigns, r.name)
		}
		body := append([]string{}, pre...)
		body = append(body, "type _res struct {")
		for _, f := range fields {
			body = append(body, "\t"+f)
		}
		body = append(body,
			"}",
			"_ch := make(chan _res, 1)",
			fmt.Sprintf("_cbInfo := webgpu.New%s(%swebgpu.%s(func(%s) {",
				base+"CallbackInfo", modeArg(base), base+"Fn", strings.Join(pureParams, ", ")),
			fmt.Sprintf("\t_ch <- _res{%s}", strings.Join(exprs, ", ")),
			"}))",
		)
		ll := fmt.Sprintf("webgpu.%s(%s)", llName, strings.Join(append(llArgs, "_cbInfo"), ", "))
		body = append(body, ll)
		body = append(body, post...)
		body = append(body,
			"_out := waitRecv(_ch)",
			"runtime.KeepAlive(_cbInfo)",
		)
		var rets []string
		for _, r := range results {
			rets = append(rets, "_out."+r.name)
		}
		if len(rets) > 0 {
			body = append(body, "return "+strings.Join(rets, ", "))
		}
		return WrapMethod{Name: methodName, Doc: cleanDoc(m.Doc), Sig: sig, Body: body}

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

	sig = attachResults(sig, results)

	// --- body ---
	body := append([]string{}, pre...)
	body = append(body, wrapBody(style, llCall, results, post, spec)...)
	return WrapMethod{Name: methodName, Doc: cleanDoc(m.Doc), Sig: sig, Body: body}
}

// attachResults appends the result clause to a signature that has none yet.
func attachResults(sig string, results []wrapResult) string {
	switch len(results) {
	case 0:
		return sig
	case 1:
		return sig + " " + results[0].typ
	default:
		var rs []string
		for _, r := range results {
			rs = append(rs, r.name+" "+r.typ)
		}
		return sig + " (" + strings.Join(rs, ", ") + ")"
	}
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
	}
	return nil
}
