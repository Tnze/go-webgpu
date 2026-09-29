package main

import (
	"fmt"
	"strings"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// Spec → TemplateData (low-level package).
// ---------------------------------------------------------------------------

func buildTemplateData(spec *parser.Spec) *TemplateData {
	td := &TemplateData{Spec: spec}

	for _, c := range spec.Constants {
		td.Consts = append(td.Consts, ConstData{
			Name:  pascalCase(c.Name),
			Value: constantGoValue(c),
			Doc:   cleanDoc(c.Doc),
		})
	}

	for _, obj := range spec.Objects {
		td.Handles = append(td.Handles, HandleData{
			Name:  pascalCase(obj.Name),
			CName: "WGPU" + pascalCase(obj.Name),
			Doc:   cleanDoc(obj.Doc),
		})
	}

	td.Enums = buildEnums(spec)
	td.ErrorTypes = buildErrorTypes(spec)
	td.Bitflags = buildBitflags(spec)
	td.CallbackInfos = buildCallbackInfos(spec)
	td.Structs = buildStructs(spec)
	td.Funcs = buildFuncs(spec)

	// Bodies are backend-specific; render them once here so the templates
	// only stitch lines (see callbody.go).
	for i := range td.Funcs {
		if td.Funcs[i].IsRelease {
			continue
		}
		td.Funcs[i].CgoBody = tabBody(callBody(td.Funcs[i], BackendCGO))
		td.Funcs[i].SyscallBody = tabBody(callBody(td.Funcs[i], BackendSyscall))
	}

	return td
}

func buildEnums(spec *parser.Spec) []EnumData {
	var out []EnumData
	for _, e := range spec.Enums {
		enumName := pascalCase(e.Name)
		ed := EnumData{
			Name:  enumName,
			CName: "WGPU" + enumName,
			Doc:   cleanDoc(e.Doc),
		}
		// YAML `null` entries occupy a numeric slot (usually 0 = undefined)
		// without emitting a Go constant, so values match the C ABI.
		val := 0
		for _, entry := range e.Entries {
			if entry.IsNull {
				val++
				continue
			}
			entryName := pascalCase(entry.Name)
			ed.Entries = append(ed.Entries, EnumEntryData{
				Name:  enumName + entryName,
				CName: "WGPU" + enumName + "_" + entryName,
				Value: fmt.Sprintf("%d", val),
				Doc:   cleanDoc(entry.Doc),
			})
			val++
		}
		out = append(out, ed)
	}
	return out
}

func buildErrorTypes(spec *parser.Spec) []ErrorTypeEntryData {
	var out []ErrorTypeEntryData
	for _, e := range spec.Enums {
		if e.Name != "error_type" {
			continue
		}
		for _, entry := range e.Entries {
			if entry.IsNull || entry.Name == "no_error" {
				continue
			}
			out = append(out, ErrorTypeEntryData{
				Name:    pascalCase(entry.Name) + "Error",
				Value:   "ErrorType" + pascalCase(entry.Name),
				Display: strings.ReplaceAll(entry.Name, "_", " "),
			})
		}
	}
	return out
}

func buildBitflags(spec *parser.Spec) []BitflagData {
	var out []BitflagData
	for _, bf := range spec.Bitflags {
		bfName := pascalCase(bf.Name)
		bd := BitflagData{
			Name:  bfName,
			CName: "WGPU" + bfName,
			Doc:   cleanDoc(bf.Doc),
		}
		// Auto-assign power-of-2 values starting from 1 (index 0 = "none" = 0).
		autoVal := uint(1)
		for _, entry := range bf.Entries {
			entryName := pascalCase(entry.Name)
			be := BitflagEntryData{
				Name:  bfName + entryName,
				CName: "WGPU" + bfName + "_" + entryName,
				Doc:   cleanDoc(entry.Doc),
			}
			switch {
			case entry.Value != nil:
				be.Value = fmt.Sprintf("%d", *entry.Value)
			case len(entry.ValueCombination) > 0:
				parts := make([]string, len(entry.ValueCombination))
				for i, v := range entry.ValueCombination {
					parts[i] = bfName + pascalCase(v)
				}
				be.Value = strings.Join(parts, " | ")
			case strings.EqualFold(entry.Name, "none"):
				be.Value = "0"
			default:
				be.Value = fmt.Sprintf("1 << %d", autoVal-1)
				autoVal++
			}
			bd.Entries = append(bd.Entries, be)
		}
		out = append(out, bd)
	}
	return out
}

func buildCallbackInfos(spec *parser.Spec) []CallbackInfoData {
	var out []CallbackInfoData
	seen := map[string]bool{}
	for _, cb := range spec.Callbacks {
		infoName := pascalCase(cb.Name) + "CallbackInfo"
		if seen[infoName] {
			continue
		}
		used, hasFn, opName := callbackUsage(spec, cb.Name)
		if !used {
			continue
		}
		seen[infoName] = true

		ci := CallbackInfoData{
			Name:          infoName,
			CName:         "WGPU" + infoName,
			Doc:           cleanDoc(cb.Doc),
			FnType:        pascalCase(cb.Name) + "Fn",
			HasMode:       cb.Name != "uncaptured_error",
			BaseName:      pascalCase(cb.Name),
			ExportName:    "go" + pascalCase(cb.Name) + "CB",
			CFnType:       "WGPU" + pascalCase(cb.Name) + "Callback",
			OpName:        opName,
			HasTrampoline: true, // every used callback needs a Go-reachable trampoline
			OneShot:       cb.Name != "uncaptured_error",
			FnParams:      pureCBParams(cb),
		}
		if hasFn {
			ci.ExportName = "go" + pascalCase(cb.Name) + "CB"
			ci.CFnType = "WGPU" + pascalCase(cb.Name) + "Callback"
		}

		// Classify C parameters:
		//   status → raw status enum; error_type → folded with following message;
		//   message → trailing string (or feeds the typed error).
		sawErrorType := false
		for _, a := range cb.Args {
			ai := callbackArgInfo(a)
			if ai.Kind == "error_type" {
				sawErrorType = true
			}
			if sawErrorType && !ai.IsStatus && ai.Kind != "error_type" && a.Name == "message" && isStringType(a.Type) {
				ai.Kind = "error_msg"
				ai.IsMessage = true
				sawErrorType = false
			}
			ci.Args = append(ci.Args, ai)
		}
		for i := range ci.Args {
			switch {
			case ci.Args[i].IsStatus:
				ci.StatusArg = &ci.Args[i]
				ci.SuccessConst = pascalCase(strings.TrimPrefix(cb.Args[i].Type, "enum.")) + "Success"
			case ci.Args[i].Kind == "error_type":
				ci.ErrorTypeArg = &ci.Args[i]
			case ci.Args[i].Kind == "error_msg":
				ci.ErrorMessageArg = &ci.Args[i]
			case ci.Args[i].IsMessage:
				ci.MessageArg = &ci.Args[i]
			default:
				ci.CallArgs = append(ci.CallArgs, ci.Args[i].GoName)
			}
		}
		if ci.ErrorTypeArg != nil {
			ci.CallArgs = append(ci.CallArgs, "err")
			msgExpr := `""`
			if ci.ErrorMessageArg != nil {
				msgExpr = "cgoGoString(" + ci.ErrorMessageArg.Name + ")"
			}
			ci.ErrConvert = "errorFromErrorType(ErrorType(" + ci.ErrorTypeArg.Name + "), " + msgExpr + ")"
		}
		if ci.StatusArg != nil {
			ci.CallArgs = append(ci.CallArgs, ci.StatusArg.GoName)
		}
		if ci.MessageArg != nil {
			ci.CallArgs = append(ci.CallArgs, ci.MessageArg.GoName)
		}
		out = append(out, ci)
	}
	return out
}

// callbackUsage reports whether a callback is referenced, whether it is used
// as a function argument (needs a trampoline), and the owning operation name.
func callbackUsage(spec *parser.Spec, cbName string) (used, hasFn bool, opName string) {
	kind := "callback." + cbName
	for _, s := range spec.Structs {
		for _, m := range s.Members {
			if m.Type == kind {
				used = true
			}
		}
	}
	for _, f := range spec.Functions {
		if f.Callback == kind {
			used, hasFn = true, true
			opName = pascalCase(f.Name)
		}
	}
	for _, obj := range spec.Objects {
		for _, m := range obj.Methods {
			if m.Callback == kind {
				used, hasFn = true, true
				methodName := pascalCase(m.Name)
				if m.Name == "map_async" {
					methodName = "Map"
				}
				opName = pascalCase(obj.Name) + methodName
			}
		}
	}
	return
}

func buildStructs(spec *parser.Spec) []StructData {
	var out []StructData
	for _, s := range spec.Structs {
		sd := StructData{
			Name:        pascalCase(s.Name),
			CName:       "WGPU" + pascalCase(s.Name),
			Type:        s.Type,
			Doc:         cleanDoc(s.Doc),
			FreeMembers: s.FreeMembers,
		}
		// C chain header depends on the struct kind:
		//   extensible*: WGPUChainedStruct const * nextInChain
		//   extension:   WGPUChainedStruct chain { next, sType }
		switch s.Type {
		case "extensible", "extensible_callback_arg":
			sd.Members = append(sd.Members, StructMemberData{
				Name: "NextInChain", GoType: "unsafe.Pointer",
				TypeRef: "c_void_const*", CName: "nextInChain",
			})
		case "extension":
			sd.Members = append(sd.Members, StructMemberData{
				Name: "NextInChain", GoType: "unsafe.Pointer",
				TypeRef: "c_void_const*", CName: "next",
			})
			sd.Members = append(sd.Members, StructMemberData{
				Name: "SType", GoType: "SType", TypeRef: "enum.s_type",
				CName: "sType", HasDefault: true, DefaultVal: "SType" + pascalCase(s.Name),
			})
			sd.HasDefault = true
		}
		for _, m := range s.Members {
			if isArrayType(m.Type) {
				// C layout: size_t xxxCount; T const * xxx;
				inner := extractArrayInner(m.Type)
				sd.Members = append(sd.Members,
					StructMemberData{
						Name: pascalCase(m.Name) + "Count", GoType: "uintptr",
						TypeRef: "usize", CName: m.Name + "_count",
					},
					StructMemberData{
						Name: pascalCase(m.Name), GoType: "*" + goTypeForRef(inner, spec),
						TypeRef: inner, CName: m.Name, Doc: cleanDoc(m.Doc),
						IsPtr: true, IsOptional: m.Optional, IsSlice: true,
					})
				continue
			}
			md := StructMemberData{
				Name: pascalCase(m.Name), GoType: goTypeForMember(m),
				TypeRef: m.Type, CName: m.Name, Doc: cleanDoc(m.Doc),
				IsPtr:      m.Pointer == "immutable" || m.Pointer == "mutable",
				IsOptional: m.Optional,
			}
			if m.Default != nil {
				md.HasDefault = true
				md.DefaultVal = memberDefaultValue(m, spec)
			}
			sd.Members = append(sd.Members, md)
			if md.HasDefault {
				sd.HasDefault = true
			}
		}
		out = append(out, sd)
	}
	return out
}

// buildFuncs emits package-level functions. Handles are unsafe.Pointer-based
// and Go forbids methods on them, so object methods become TypeMethod name.
func buildFuncs(spec *parser.Spec) []FuncData {
	var out []FuncData
	for _, f := range spec.Functions {
		fd := buildFuncData(f, spec)
		fd.Ident = fd.Name
		out = append(out, fd)
	}
	for _, obj := range spec.Objects {
		objName := pascalCase(obj.Name)
		for _, m := range obj.Methods {
			fd := buildFuncData(m, spec)
			// Low-level keeps the C name: wgpuBufferMapAsync → BufferMapAsync.
			fd.CName = "wgpu" + objName + pascalCase(m.Name)
			fd.GoArgs = append([]FuncArgData{{
				Name: receiverName(objName), GoType: objName,
				TypeRef: "object." + obj.Name, CName: obj.Name,
			}}, fd.GoArgs...)
			fd.Name = objName + pascalCase(m.Name)
			fd.Ident = fd.Name
			fd.OpName = fd.Name
			out = append(out, fd)
		}
		// Release comes last among the type's functions.
		out = append(out, FuncData{
			Name: objName + "Release", CName: "wgpu" + objName + "Release",
			Ident: objName + "Release", IsRelease: true, ObjName: objName,
			OpName: objName + "Release", RetKind: "none",
			GoArgs: []FuncArgData{{
				Name: receiverName(objName), GoType: objName,
				TypeRef: "object." + obj.Name, CName: obj.Name,
			}},
		})
	}
	return out
}

func buildFuncData(f parser.Function, spec *parser.Spec) FuncData {
	fd := FuncData{
		Name:  pascalCase(f.Name),
		CName: "wgpu" + pascalCase(f.Name),
		Doc:   cleanDoc(f.Doc),
	}
	// Async APIs are pure primitives: they take the raw callback-info record
	// by value and return a Future. Blocking / error folding lives in the gpu
	// wrapper layer.
	if f.Callback != "" {
		cbName := strings.TrimPrefix(f.Callback, "callback.")
		fd.HasCallback = true
		fd.CallbackName = cbName
		fd.CallbackFn = pascalCase(cbName) + "Fn"
		fd.ReturnRef = "struct.future"
		fd.RetKind = "future"
		fd.GoReturn = "Future"
	} else if f.Returns != nil && f.Returns.Type != "" && f.Returns.Type != "void" {
		fd.ReturnRef = f.Returns.Type
		fd.GoReturn = goTypeForRef(f.Returns.Type, spec)
		switch {
		case isHandleType(f.Returns.Type):
			fd.RetKind = "handle"
		case f.Returns.Type == "enum.status":
			fd.RetKind, fd.GoReturn = "status", "Status"
		case f.Returns.Type == "enum.wait_status":
			fd.RetKind, fd.GoReturn = "wait", "WaitStatus"
		case f.Returns.Type == "c_void_mapped_range_ptr" || strings.HasPrefix(f.Returns.Type, "c_void"):
			fd.RetKind = "pointer"
		case f.Returns.Type == "bool":
			fd.RetKind = "bool"
		case f.Returns.Type == "struct.future":
			fd.RetKind = "future"
		default:
			fd.RetKind = "value"
		}
	} else {
		fd.RetKind = "none"
	}
	fd.OpName = fd.Name

	for _, arg := range f.Args {
		goType := goTypeForRef(arg.Type, spec)
		isSlice := isArrayType(arg.Type)
		var sliceCPtr string
		if isSlice {
			inner := extractArrayInner(arg.Type)
			innerGo := goTypeForRef(inner, spec)
			goType = "[]" + innerGo
			sliceCPtr = "*" + cgoTypeName(innerGo)
		}
		isPtr := arg.Pointer == "immutable" || arg.Pointer == "mutable"
		if isStructType(arg.Type) && isPtr && !isSlice {
			goType = "*" + goType
		}
		fad := FuncArgData{
			Name: camelCase(arg.Name), GoType: goType, TypeRef: arg.Type,
			CName: arg.Name, IsPtr: isPtr, IsOptional: arg.Optional,
			IsSlice: isSlice, IsStr: isStringType(arg.Type), SliceCPtr: sliceCPtr,
		}
		if isStructType(arg.Type) {
			argName := camelCase(arg.Name)
			structName := pascalCase(strings.TrimPrefix(arg.Type, "struct."))
			for _, s := range spec.Structs {
				if pascalCase(s.Name) == structName {
					fad.HasChain = s.Type == "extensible" || s.Type == "extensible_callback_arg" || s.Type == "extension"
					collectStructPins(spec, arg.Type, argName, &fad.StrFields, &fad.PtrFields, &fad.SliceFields, 0)
					break
				}
			}
		}
		fd.GoArgs = append(fd.GoArgs, fad)
	}
	// The raw callback-info record is a by-value C parameter; mirror it as a
	// Go by-value struct so callers cannot pass a nil that the C API rejects.
	if fd.HasCallback {
		fd.GoArgs = append(fd.GoArgs, FuncArgData{
			Name:     "callbackInfo",
			GoType:   pascalCase(fd.CallbackName) + "CallbackInfo",
			TypeRef:  "callbackinfo." + fd.CallbackName,
			CName:    "callbackInfo",
			IsCBInfo: true,
		})
	}
	return fd
}

// pureCBParams returns the Go callback signature in C argument order, using
// pure type mapping only (no error folding, no reordering): strings become
// string, handles become their named types, enums their named types.
func pureCBParams(cb parser.Callback) []string {
	var out []string
	for _, a := range cb.Args {
		name := camelCase(a.Name)
		var typ string
		switch {
		case isStringType(a.Type):
			typ = "string"
		case isHandleType(a.Type):
			typ = pascalCase(strings.TrimPrefix(a.Type, "object."))
		case strings.HasPrefix(a.Type, "enum."):
			typ = pascalCase(strings.TrimPrefix(a.Type, "enum."))
		case isStructType(a.Type) && (a.Pointer == "immutable" || a.Pointer == "mutable"):
			typ = "*" + pascalCase(strings.TrimPrefix(a.Type, "struct."))
		default:
			typ = goTypeForRef(a.Type, nil)
		}
		out = append(out, name+" "+typ)
	}
	return out
}

// callbackArgInfo classifies one callback C parameter for trampoline generation.
func callbackArgInfo(a parser.CallbackArg) CallbackArgInfo {
	ai := CallbackArgInfo{
		YAMLName: a.Name,
		Name:     "c_" + camelCase(a.Name),
		GoName:   camelCase(a.Name),
		GoType:   goTypeForRef(a.Type, nil),
		PureType: pureCBType(a),
	}
	switch {
	case isStringType(a.Type):
		ai.ExternType, ai.CGoType = "WGPUStringView", "C.WGPUStringView"
	case strings.HasPrefix(a.Type, "enum."):
		n := "WGPU" + pascalCase(strings.TrimPrefix(a.Type, "enum."))
		ai.ExternType, ai.CGoType = n, "C."+n
	case strings.HasPrefix(a.Type, "object."):
		n := "WGPU" + pascalCase(strings.TrimPrefix(a.Type, "object."))
		ai.ExternType, ai.CGoType = n, "C."+n
	case strings.HasPrefix(a.Type, "struct."):
		n := "WGPU" + pascalCase(strings.TrimPrefix(a.Type, "struct."))
		if a.Pointer == "immutable" || a.Pointer == "mutable" {
			ai.ExternType, ai.CGoType = n+"*", "*C."+n
		} else {
			ai.ExternType, ai.CGoType = n, "C."+n
		}
	default:
		ai.ExternType, ai.CGoType = "void*", "unsafe.Pointer"
	}
	switch {
	case a.Name == "status" && strings.HasPrefix(a.Type, "enum."):
		ai.Kind, ai.IsStatus = "status", true
		ai.GoType = pascalCase(strings.TrimPrefix(a.Type, "enum."))
	case a.Type == "enum.error_type":
		ai.Kind, ai.GoName, ai.GoType = "error_type", "err", "error"
	case a.Name == "message" && isStringType(a.Type):
		ai.Kind, ai.IsMessage = "message", true
	case strings.HasPrefix(a.Type, "object."):
		ai.Kind = "object"
	case strings.HasPrefix(a.Type, "struct.") && (a.Pointer == "immutable" || a.Pointer == "mutable"):
		ai.Kind = "struct_ptr"
		ai.GoType = pascalCase(strings.TrimPrefix(a.Type, "struct."))
	case strings.HasPrefix(a.Type, "enum."):
		ai.Kind = "enum"
		ai.GoType = pascalCase(strings.TrimPrefix(a.Type, "enum."))
	case isStringType(a.Type):
		ai.Kind, ai.GoType = "string", "string"
	default:
		ai.Kind = "value"
	}
	return ai
}

// goCallbackArgs returns the Go parameters of a callback func type. Status
// enums pass through raw; error_type folds with its message into an error.
func goCallbackArgs(cb parser.Callback) []CallbackArgData {
	var args []CallbackArgData
	var statusGoType string
	var messageArg *CallbackArgData
	skipMessage := false
	for _, a := range cb.Args {
		if skipMessage {
			skipMessage = false
			continue
		}
		if a.Name == "status" && strings.HasPrefix(a.Type, "enum.") {
			statusGoType = pascalCase(strings.TrimPrefix(a.Type, "enum."))
			continue
		}
		if a.Type == "enum.error_type" {
			args = append(args, CallbackArgData{Name: "err", GoType: "error"})
			skipMessage = true
			continue
		}
		if a.Name == "message" && isStringType(a.Type) {
			messageArg = &CallbackArgData{Name: "message", GoType: "string"}
			continue
		}
		args = append(args, CallbackArgData{
			Name:   camelCase(a.Name),
			GoType: goTypeForRef(a.Type, nil),
		})
	}
	if statusGoType != "" {
		args = append(args, CallbackArgData{Name: "status", GoType: statusGoType})
	}
	if messageArg != nil {
		args = append(args, *messageArg)
	}
	return args
}
