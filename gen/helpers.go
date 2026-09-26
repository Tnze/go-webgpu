package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/Tnze/go-webgpu/gen/parser"
)

var goKeywords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
}

// ---------------------------------------------------------------------------
// Template data structure
// ---------------------------------------------------------------------------

// TemplateData is the top-level data passed to every template.
type TemplateData struct {
	Spec          *parser.Spec
	Enums         []EnumData
	Bitflags      []BitflagData
	Structs       []StructData
	Funcs         []FuncData
	Handles       []HandleData
	Consts        []ConstData
	CallbackInfos []CallbackInfoData
	ErrorTypes    []ErrorTypeEntryData
}

// ErrorTypeEntryData describes one distinct Go error type generated
// from the WGPUErrorType enum.
type ErrorTypeEntryData struct {
	Name    string // Go type name, e.g. "ValidationError"
	Value   string // enum value, e.g. "ErrorTypeValidation"
	Display string // lowercase label for Error(), e.g. "validation"
}

// EnumData describes a Go enum type.
type EnumData struct {
	Name    string
	CName   string
	Entries []EnumEntryData
}

// EnumEntryData describes a single enum value.
type EnumEntryData struct {
	Name  string
	CName string
	Value string // explicit value or "" for iota
}

// BitflagData describes a Go bitflag type.
type BitflagData struct {
	Name    string
	CName   string
	Entries []BitflagEntryData
}

// BitflagEntryData is a single flag value.
type BitflagEntryData struct {
	Name  string
	CName string
	Value string
}

// StructData describes a Go struct.
type StructData struct {
	Name       string
	CName      string
	Type       string // "extensible", "standalone", "extension"
	Members    []StructMemberData
	HasDefault bool
}

// StructMemberData is a single struct field.
type StructMemberData struct {
	Name       string // Go PascalCase field name
	GoType     string // Go type string
	TypeRef    string // original YAML type reference (e.g. "object.device", "struct.color")
	CName      string // YAML field name
	Doc        string
	HasDefault bool
	DefaultVal string
	IsPtr      bool // true if YAML has pointer: immutable/mutable
	IsOptional bool // true if YAML has optional: true
	IsSlice    bool // true for array<T> data pointer field
	IsArray    bool // true for fixed arrays
	ArrayLen   int
}

// CallbackInfoData describes a callback type and the glue generated for it.
type CallbackInfoData struct {
	Name       string // e.g. DeviceLostCallbackInfo
	CName      string // e.g. WGPUDeviceLostCallbackInfo
	FnType     string // e.g. DeviceLostFn
	HasMode    bool   // true except uncaptured_error
	BaseName   string // e.g. RequestAdapter (pascalCase of callback name)
	ExportName string // e.g. goRequestAdapterCB; "" if no trampoline
	CFnType    string // e.g. WGPURequestAdapterCallback
	OpName     string // error op label, e.g. "Instance.RequestAdapter"
	// Args lists every C parameter in order (status/message included).
	Args []CallbackArgInfo
	// HasTrampoline is true for callbacks used as function arguments
	// (those get a //export trampoline + cgo.Handle). Embedded descriptor
	// callbacks (device_lost, uncaptured_error) are false.
	HasTrampoline bool
	// Precomputed conveniences for templates.
	StatusArg      *CallbackArgInfo
	MessageArg     *CallbackArgInfo
	ErrorTypeArg   *CallbackArgInfo // enum.error_type parameter (if any)
	ErrorMessageArg *CallbackArgInfo // message parameter feeding the typed error
	SuccessConst   string       // e.g. MapAsyncStatusSuccess
	CallArgs       []string     // Go value names passed to cb(...)
	ErrConvert     string       // Go expression producing the typed error ("" if none)
}

// CallbackArgInfo describes one C parameter of a callback trampoline.
type CallbackArgInfo struct {
	YAMLName   string // original YAML name (e.g. "type", "adapter")
	Name       string // Go parameter name for the C value (c_type, c_adapter)
	GoName     string // converted Go value name (typeVal, adapter)
	CGoType    string // cgo type in //export (C.WGPUAdapter, *C.WGPUCompilationInfo)
	ExternType string // C type in extern decl (WGPUAdapter, WGPUCompilationInfo*)
	GoType     string // Go type of the converted value (Adapter, ErrorType)
	Kind       string // status | message | object | enum | struct_ptr | string | value
	IsStatus   bool
	IsMessage  bool
}

// FuncData describes a function to generate.
// When IsMethod is set, GoArgs[0] is the receiver and Name omits the
// object prefix (e.g. Device.CreateBuffer, not DeviceCreateBuffer).
// IsRelease marks the per-type Release method, which is emitted after the
// type's other methods and has a special body (Cleanup.Stop + wgpu*Release).
//
// Return convention: raw C results are surfaced directly — status enums
// stay enums, handles stay handles (zero = failure), pointer returns stay
// pointers (nil = failure). Callback APIs become blocking multi-return
// functions whose last values are the status enum and the C message string
// (the message is folded into a typed error for error_type callbacks).
type FuncData struct {
	Name          string
	CName         string
	Ident         string // unique Go identifier for internal symbols (e.g. proc vars)
	GoArgs        []FuncArgData
	GoReturn      string
	ReturnRef     string // original YAML type reference for the return type
	RetKind       string // handle | status | pointer | wait | future | bool | value | none
	ReturnsError  bool   // Go signature includes error
	SuccessConst  string // e.g. StatusSuccess for status/wait kinds
	OpName        string // error op label, e.g. "Device.CreateBuffer"
	Doc           string
	IsMethod      bool
	IsRelease     bool
	ObjName       string // receiver type name without '*', e.g. "Device"
	HasCallback   bool
	CallbackFn    string // Go callback type, e.g. RequestAdapterFn
	CallbackName  string // original callback name, e.g. request_adapter
	ReturnsFuture bool
	// BlockResults are the non-error results of a blocking wrapper for
	// callback APIs (e.g. [{Name:adapter, GoType:Adapter}]).
	BlockResults []CallbackArgData
}

// FuncArgData is a function argument.
type FuncArgData struct {
	Name        string // Go parameter name
	GoType      string // Go type string
	TypeRef     string // original YAML type reference
	CName       string // C argument name
	IsPtr       bool   // pointer: immutable/mutable
	IsOptional  bool
	IsSlice     bool
	IsStr       bool
	SliceFields []string // Go field names of slice members (for struct args needing pin)
	SliceCPtr   string   // cgo pointer type for slice data, e.g. *C.uint32_t
}

// HandleData describes a WebGPU handle (opaque object type).
type HandleData struct {
	Name  string
	CName string
}

// ConstData describes a named constant.
type ConstData struct {
	Name  string
	Value string
	Doc   string
}

// ---------------------------------------------------------------------------
// buildTemplateData transforms the parsed spec into template-ready data.
// ---------------------------------------------------------------------------

func buildTemplateData(spec *parser.Spec) *TemplateData {
	td := &TemplateData{Spec: spec}

	// Constants
	for _, c := range spec.Constants {
		td.Consts = append(td.Consts, ConstData{
			Name:  pascalCase(c.Name),
			Value: constantGoValue(c),
			Doc:   c.Doc,
		})
	}

	// Handles (objects)
	for _, obj := range spec.Objects {
		td.Handles = append(td.Handles, HandleData{
			Name:  pascalCase(obj.Name),
			CName: "WGPU" + pascalCase(obj.Name),
		})
	}

	// Enums
	for _, e := range spec.Enums {
		enumName := pascalCase(e.Name)
		ed := EnumData{
			Name:  enumName,
			CName: "WGPU" + enumName,
		}
		// YAML `null` entries occupy a numeric slot (usually 0 = undefined)
		// without emitting a Go constant. Keep the slot so values match the C ABI.
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
			})
			val++
		}
		td.Enums = append(td.Enums, ed)
	}

	// Distinct Go error types for WGPUErrorType entries.
	// NoError is skipped — it maps to nil, not an error value.
	for _, e := range spec.Enums {
		if e.Name != "error_type" {
			continue
		}
		for _, entry := range e.Entries {
			if entry.IsNull || entry.Name == "no_error" {
				continue
			}
			td.ErrorTypes = append(td.ErrorTypes, ErrorTypeEntryData{
				Name:    pascalCase(entry.Name) + "Error",
				Value:   "ErrorType" + pascalCase(entry.Name),
				Display: strings.ReplaceAll(entry.Name, "_", " "),
			})
		}
	}

	// Bitflags
	for _, bf := range spec.Bitflags {
		bfName := pascalCase(bf.Name)
		bd := BitflagData{
			Name:  bfName,
			CName: "WGPU" + bfName,
		}
		// Auto-assign power-of-2 values starting from 1 (skip index 0 = "none" = 0).
		autoVal := uint(1)
		for _, entry := range bf.Entries {
			entryName := pascalCase(entry.Name)
			be := BitflagEntryData{
				Name:  bfName + entryName,
				CName: "WGPU" + bfName + "_" + entryName,
			}
			if entry.Value != nil {
				be.Value = fmt.Sprintf("%d", *entry.Value)
			} else if len(entry.ValueCombination) > 0 {
				parts := make([]string, len(entry.ValueCombination))
				for i, v := range entry.ValueCombination {
					parts[i] = bfName + pascalCase(v)
				}
				be.Value = strings.Join(parts, " | ")
			} else if strings.EqualFold(entry.Name, "none") {
				be.Value = "0"
			} else {
				be.Value = fmt.Sprintf("1 << %d", autoVal-1)
				autoVal++
			}
			bd.Entries = append(bd.Entries, be)
		}
		td.Bitflags = append(td.Bitflags, bd)
	}

	// Callback info structs (embedded in descriptors such as DeviceDescriptor).
	seenCB := map[string]bool{}
	for _, cb := range spec.Callbacks {
		infoName := pascalCase(cb.Name) + "CallbackInfo"
		if seenCB[infoName] {
			continue
		}
		// Only emit info types that appear as struct members or as async-arg params.
		used := false
		hasFn := false
		var opName string
		for _, s := range spec.Structs {
			for _, m := range s.Members {
				if m.Type == "callback."+cb.Name {
					used = true
				}
			}
		}
		for _, f := range spec.Functions {
			if f.Callback == "callback."+cb.Name {
				used = true
				hasFn = true
				opName = pascalCase(f.Name)
			}
		}
		for _, obj := range spec.Objects {
			for _, m := range obj.Methods {
				if m.Callback == "callback."+cb.Name {
					used = true
					hasFn = true
					methodName := pascalCase(m.Name)
					if m.Name == "map_async" {
						methodName = "Map"
					}
					opName = pascalCase(obj.Name) + "." + methodName
				}
			}
		}
		if !used {
			continue
		}
		seenCB[infoName] = true

		ci := CallbackInfoData{
			Name:         infoName,
			CName:        "WGPU" + infoName,
			FnType:       pascalCase(cb.Name) + "Fn",
			HasMode:      cb.Name != "uncaptured_error",
			BaseName:     pascalCase(cb.Name),
			OpName:       opName,
			HasTrampoline: hasFn,
		}
		if hasFn {
			ci.ExportName = "go" + pascalCase(cb.Name) + "CB"
			ci.CFnType = "WGPU" + pascalCase(cb.Name) + "Callback"
		}

		// Classify each C parameter.
		//   status         → raw status enum (before trailing message)
		//   error_type     → folded with following message into typed error
		//   message        → trailing string (or feeds the typed error)
		//   everything else → passed through raw
		sawErrorType := false
		for _, a := range cb.Args {
			ai := callbackArgInfo(a)
			if ai.Kind == "error_type" {
				sawErrorType = true
			}
			// The message immediately after error_type feeds the typed error.
			if sawErrorType && !ai.IsStatus && ai.Kind != "error_type" && a.Name == "message" && isStringType(a.Type) {
				ai.Kind = "error_msg"
				ai.IsMessage = true
				sawErrorType = false
			}
			ci.Args = append(ci.Args, ai)
		}
		// Point into the slice so the pointers stay valid.
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
			ci.ErrConvert = "errorFromErrorType(" + "ErrorType(" + ci.ErrorTypeArg.Name + "), " + msgExpr + ")"
		}
		if ci.StatusArg != nil {
			ci.CallArgs = append(ci.CallArgs, ci.StatusArg.GoName)
		}
		// Standalone messages (not folded into a typed error) are passed through.
		if ci.MessageArg != nil {
			ci.CallArgs = append(ci.CallArgs, ci.MessageArg.GoName)
		}

		td.CallbackInfos = append(td.CallbackInfos, ci)
	}

	// Structs
	for _, s := range spec.Structs {
		sd := StructData{
			Name:  pascalCase(s.Name),
			CName: "WGPU" + pascalCase(s.Name),
			Type:  s.Type,
		}
		// C layout of the chain header depends on the struct kind:
		//   extensible*: WGPUChainedStruct const * nextInChain
		//   extension:   WGPUChainedStruct chain { next, sType }
		switch s.Type {
		case "extensible", "extensible_callback_arg":
			sd.Members = append(sd.Members, StructMemberData{
				Name:    "NextInChain",
				GoType:  "unsafe.Pointer",
				TypeRef: "c_void_const*",
				CName:   "nextInChain",
			})
		case "extension":
			sd.Members = append(sd.Members, StructMemberData{
				Name:    "NextInChain",
				GoType:  "unsafe.Pointer",
				TypeRef: "c_void_const*",
				CName:   "next",
			})
			sd.Members = append(sd.Members, StructMemberData{
				Name:       "SType",
				GoType:     "SType",
				TypeRef:    "enum.s_type",
				CName:      "sType",
				HasDefault: true,
				DefaultVal: "SType" + pascalCase(s.Name),
			})
			sd.HasDefault = true
		}
		for _, m := range s.Members {
			if isArrayType(m.Type) {
				// C layout: size_t xxxCount; T const * xxx;
				inner := extractArrayInner(m.Type)
				countName := pascalCase(m.Name) + "Count"
				sd.Members = append(sd.Members, StructMemberData{
					Name:    countName,
					GoType:  "uintptr",
					TypeRef: "usize",
					CName:   m.Name + "_count",
				})
				sd.Members = append(sd.Members, StructMemberData{
					Name:       pascalCase(m.Name),
					GoType:     "*" + goTypeForRef(inner, spec),
					TypeRef:    inner,
					CName:      m.Name,
					Doc:        m.Doc,
					IsPtr:      true,
					IsOptional: m.Optional,
					IsSlice:    true,
				})
				continue
			}
			md := StructMemberData{
				Name:       pascalCase(m.Name),
				GoType:     goTypeForMember(m),
				TypeRef:    m.Type,
				CName:      m.Name,
				Doc:        m.Doc,
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
		td.Structs = append(td.Structs, sd)
	}

	// Functions. Prefer a Go method whenever the call has an object handle
	// to use as a receiver; otherwise emit a package-level function.
	for _, f := range spec.Functions {
		fd := buildFuncData(f, spec)
		if tryAsMethod(&fd, f.Name) {
			fd.Ident = fd.ObjName + fd.Name
		} else {
			fd.Ident = fd.Name
		}
		td.Funcs = append(td.Funcs, fd)
	}
	for _, obj := range spec.Objects {
		objName := pascalCase(obj.Name)
		// Reserved method names on every handle type.
		used := map[string]bool{"Handle": true, "Release": true}
		for _, m := range obj.Methods {
			fd := buildFuncData(m, spec)
			methodName := pascalCase(m.Name)
			// MapAsync is blocking in this API; expose it as Map.
			// Keep the C name based on the spec name.
			fd.CName = "wgpu" + objName + pascalCase(m.Name)
			if m.Name == "map_async" {
				methodName = "Map"
			}
			// Prepend the object handle as the first argument (the receiver).
			handleArg := FuncArgData{
				Name:    receiverName(objName),
				GoType:  objName,
				TypeRef: "object." + obj.Name,
				CName:   obj.Name,
			}
			fd.GoArgs = append([]FuncArgData{handleArg}, fd.GoArgs...)

			if used[methodName] {
				// Name collision on the receiver type: fall back to a function.
				fd.Name = objName + methodName
				fd.Ident = fd.Name
			} else {
				used[methodName] = true
				fd.IsMethod = true
				fd.ObjName = objName
				fd.Name = methodName
				fd.Ident = objName + methodName
				fd.OpName = objName + "." + methodName
			}
			td.Funcs = append(td.Funcs, fd)
		}

		// Release always comes last among the type's methods.
		fd := FuncData{
			Name:      "Release",
			CName:     "wgpu" + objName + "Release",
			Ident:     objName + "Release",
			IsMethod:  true,
			IsRelease: true,
			ObjName:   objName,
			OpName:    objName + ".Release",
			RetKind:   "none",
			Doc:       "drops this handle's reference to the native object",
			GoArgs: []FuncArgData{{
				Name:    receiverName(objName),
				GoType:  objName,
				TypeRef: "object." + obj.Name,
				CName:   obj.Name,
			}},
		}
		td.Funcs = append(td.Funcs, fd)
	}

	return td
}

// tryAsMethod converts a function into a method when its first argument is an
// object handle. The handle argument becomes the receiver. Returns false when
// no receiver is available (or the name is already taken on that type).
func tryAsMethod(fd *FuncData, rawName string) bool {
	if len(fd.GoArgs) == 0 || !isHandleType(fd.GoArgs[0].TypeRef) {
		return false
	}
	recv := fd.GoArgs[0]
	objName := pascalCase(strings.TrimPrefix(recv.TypeRef, "object."))
	methodName := pascalCase(rawName)
	// Avoid colliding with generated Handle() / Release() methods.
	if methodName == "Handle" || methodName == "Release" {
		return false
	}
	fd.IsMethod = true
	fd.ObjName = objName
	fd.Name = methodName
	fd.GoArgs[0].Name = receiverName(objName)
	fd.OpName = objName + "." + methodName
	return true
}

// receiverName returns the conventional single-letter Go receiver name
// for a type (Device → d, CommandEncoder → c).
func receiverName(typeName string) string {
	if typeName == "" {
		return "x"
	}
	return strings.ToLower(typeName[:1])
}

func buildFuncData(f parser.Function, spec *parser.Spec) FuncData {
	fd := FuncData{
		Name:  pascalCase(f.Name),
		CName: "wgpu" + pascalCase(f.Name),
		Doc:   f.Doc,
	}
	if f.Callback != "" {
		fd.HasCallback = true
		fd.CallbackName = strings.TrimPrefix(f.Callback, "callback.")
		fd.CallbackFn = pascalCase(fd.CallbackName) + "Fn"
		// Blocking Go API: the C Future/callback is an implementation detail.
		// Results (including the status enum) are returned via an internal channel.
		fd.RetKind = "blocking"
		fd.ReturnsError = false
		fd.GoReturn = ""
		fd.ReturnRef = "struct.future" // C still returns a Future we ignore
		if spec != nil {
			if cb := spec.GetCallback(fd.CallbackName); cb != nil {
				fd.BlockResults = goCallbackArgs(*cb)
			}
		}
		// MapAsync is blocking in this API; expose it as Map.
		if f.Name == "map_async" {
			fd.Name = "Map"
		}
	} else if f.Returns != nil && f.Returns.Type != "" && f.Returns.Type != "void" {
		fd.ReturnRef = f.Returns.Type
		fd.GoReturn = goTypeForRef(f.Returns.Type, spec)
		switch {
		case isHandleType(f.Returns.Type):
			fd.RetKind = "handle"
		case f.Returns.Type == "enum.status":
			fd.RetKind = "status"
			fd.GoReturn = "Status"
		case f.Returns.Type == "enum.wait_status":
			fd.RetKind = "wait"
			fd.GoReturn = "WaitStatus"
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
	// Op label for errors: "CreateInstance", … (method path sets "Device.CreateBuffer")
	fd.OpName = fd.Name
	for _, arg := range f.Args {
		goType := goTypeForRef(arg.Type, spec)
		isSlice := isArrayType(arg.Type)
		var fadSliceCPtr string
		if isSlice {
			inner := extractArrayInner(arg.Type)
			innerGo := goTypeForRef(inner, spec)
			goType = "[]" + innerGo
			fadSliceCPtr = "*" + cgoTypeName(innerGo)
		}
		isPtr := arg.Pointer == "immutable" || arg.Pointer == "mutable"
		// Struct args are C pointers (descriptors / out-params); mirror that in Go.
		if isStructType(arg.Type) && isPtr && !isSlice {
			goType = "*" + goType
		}
		fad := FuncArgData{
			Name:       camelCase(arg.Name),
			GoType:     goType,
			TypeRef:    arg.Type,
			CName:      arg.Name,
			IsPtr:      isPtr,
			IsOptional: arg.Optional,
			IsSlice:    isSlice,
			IsStr:      isStringType(arg.Type),
			SliceCPtr:  fadSliceCPtr,
		}
		// For struct args, precompute slice fields that need pinning.
		// Store full expressions like "descriptor.Entries" for template use.
		if isStructType(arg.Type) && spec != nil {
			argName := camelCase(arg.Name)
			structName := pascalCase(strings.TrimPrefix(arg.Type, "struct."))
			for _, s := range spec.Structs {
				if pascalCase(s.Name) == structName {
					for _, m := range s.Members {
						if isArrayType(m.Type) {
							fad.SliceFields = append(fad.SliceFields, argName+"."+pascalCase(m.Name))
						}
					}
					break
				}
			}
		}
		fd.GoArgs = append(fd.GoArgs, fad)
	}
	if fd.HasCallback {
		fd.GoArgs = append(fd.GoArgs, FuncArgData{
			Name:    "callback",
			GoType:  fd.CallbackFn,
			TypeRef: "callback." + fd.CallbackName,
			CName:   "callback",
		})
	}
	return fd
}

// callbackArgInfo classifies a single callback C parameter and computes
// the names/types needed for trampoline generation.
func callbackArgInfo(a parser.CallbackArg) CallbackArgInfo {
	ai := CallbackArgInfo{
		YAMLName: a.Name,
		Name:     "c_" + camelCase(a.Name),
		GoName:   camelCase(a.Name),
		GoType:   goTypeForRef(a.Type, nil),
	}
	// C types for the extern declaration and the cgo //export signature.
	switch {
	case isStringType(a.Type):
		ai.ExternType = "WGPUStringView"
		ai.CGoType = "C.WGPUStringView"
	case strings.HasPrefix(a.Type, "enum."):
		n := "WGPU" + pascalCase(strings.TrimPrefix(a.Type, "enum."))
		ai.ExternType = n
		ai.CGoType = "C." + n
	case strings.HasPrefix(a.Type, "object."):
		n := "WGPU" + pascalCase(strings.TrimPrefix(a.Type, "object."))
		ai.ExternType = n
		ai.CGoType = "C." + n
	case strings.HasPrefix(a.Type, "struct."):
		n := "WGPU" + pascalCase(strings.TrimPrefix(a.Type, "struct."))
		if a.Pointer == "immutable" || a.Pointer == "mutable" {
			ai.ExternType = n + "*"
			ai.CGoType = "*C." + n
		} else {
			ai.ExternType = n
			ai.CGoType = "C." + n
		}
	default:
		ai.ExternType = "void*"
		ai.CGoType = "unsafe.Pointer"
	}
	// Kind classification.
	switch {
	case a.Name == "status" && strings.HasPrefix(a.Type, "enum."):
		ai.Kind = "status"
		ai.IsStatus = true
		ai.GoType = pascalCase(strings.TrimPrefix(a.Type, "enum."))
	case a.Type == "enum.error_type":
		ai.Kind = "error_type"
		ai.GoName = "err"
		ai.GoType = "error"
	case a.Name == "message" && isStringType(a.Type):
		ai.Kind = "message"
		ai.IsMessage = true
	case strings.HasPrefix(a.Type, "object."):
		ai.Kind = "object"
	case strings.HasPrefix(a.Type, "struct.") && (a.Pointer == "immutable" || a.Pointer == "mutable"):
		ai.Kind = "struct_ptr"
		ai.GoType = pascalCase(strings.TrimPrefix(a.Type, "struct."))
	case strings.HasPrefix(a.Type, "enum."):
		ai.Kind = "enum"
		ai.GoType = pascalCase(strings.TrimPrefix(a.Type, "enum."))
	case isStringType(a.Type):
		ai.Kind = "string"
		ai.GoType = "string"
	default:
		ai.Kind = "value"
	}
	return ai
}

// ---------------------------------------------------------------------------
// Name mapping
// ---------------------------------------------------------------------------

func pascalCase(s string) string {
	if s == "" {
		return s
	}
	s = strings.TrimPrefix(s, "WGPU")

	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		switch strings.ToLower(p) {
		case "gpu":
			b.WriteString("GPU")
		case "d3d11", "d3d12":
			b.WriteString(strings.ToUpper(p))
		case "rgb", "rgba", "bgra", "rg":
			b.WriteString(strings.ToUpper(p))
		default:
			runes := []rune(p)
			runes[0] = unicode.ToUpper(runes[0])
			b.WriteString(string(runes))
		}
	}
	return b.String()
}

func camelCase(s string) string {
	if s == "" {
		return s
	}
	ps := pascalCase(s)
	runes := []rune(ps)
	runes[0] = unicode.ToLower(runes[0])
	result := string(runes)
	// Avoid Go keywords.
	if goKeywords[result] {
		result = result + "Val"
	}
	return result
}

func enumValue(enumName, entryName string) string {
	return pascalCase(enumName) + "_" + pascalCase(entryName)
}

// ---------------------------------------------------------------------------
// Type mapping
// ---------------------------------------------------------------------------

func goTypeForRef(ref string, spec *parser.Spec) string {
	if ref == "" {
		return "unsafe.Pointer"
	}
	switch ref {
	case "bool":
		return "Bool"
	case "uint8":
		return "uint8"
	case "uint16":
		return "uint16"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "float32":
		return "float32"
	case "float64":
		return "float64"
	case "float64_supertype":
		return "float64"
	case "usize":
		return "uintptr"
	case "string_with_default_empty", "out_string", "nullable_string":
		return "string"
	}
	if strings.HasPrefix(ref, "enum.") {
		return pascalCase(strings.TrimPrefix(ref, "enum."))
	}
	if strings.HasPrefix(ref, "bitflag.") {
		return pascalCase(strings.TrimPrefix(ref, "bitflag."))
	}
	if strings.HasPrefix(ref, "struct.") {
		return pascalCase(strings.TrimPrefix(ref, "struct."))
	}
	if strings.HasPrefix(ref, "object.") {
		// Handles are thin typed uintptr values (zero is null).
		return pascalCase(strings.TrimPrefix(ref, "object."))
	}
	if strings.HasPrefix(ref, "callback.") {
		return pascalCase(strings.TrimPrefix(ref, "callback.")) + "Fn"
	}
	if strings.HasPrefix(ref, "constant.") {
		return "uint32"
	}
	// Platform-specific void pointer types → unsafe.Pointer
	if strings.HasPrefix(ref, "c_void") {
		return "unsafe.Pointer"
	}
	// Nullable float32 → *float32
	if ref == "nullable_float32" {
		return "*float32"
	}
	return pascalCase(ref)
}

func goTypeForMember(m parser.StructMember) string {
	// Nullable / default-empty strings are WGPUStringView (pointer + length),
	// which matches a Go string header on 64-bit.
	if m.Type == "nullable_string" || isStringType(m.Type) {
		return "string"
	}
	// Embedded callback-info records (WGPUNnnCallbackInfo).
	if strings.HasPrefix(m.Type, "callback.") {
		return pascalCase(strings.TrimPrefix(m.Type, "callback.")) + "CallbackInfo"
	}
	// Object handles inside structs are raw C pointers (WGPUNnnImpl*).
	if strings.HasPrefix(m.Type, "object.") {
		return "uintptr"
	}
	// Optional / pointed-to structs are pointers in the C layout.
	if isStructType(m.Type) && (m.Optional || m.Pointer == "immutable" || m.Pointer == "mutable") {
		return "*" + pascalCase(strings.TrimPrefix(m.Type, "struct."))
	}
	if isArrayType(m.Type) {
		inner := extractArrayInner(m.Type)
		return "[]" + goTypeForRef(inner, nil)
	}
	return goTypeForRef(m.Type, nil)
}

// ---------------------------------------------------------------------------
// CGO type name
// ---------------------------------------------------------------------------

func cgoTypeName(goType string) string {
	if strings.HasPrefix(goType, "*") {
		return "C.WGPU" + strings.TrimPrefix(goType, "*")
	}
	switch goType {
	case "bool", "Bool":
		return "C.WGPUBool"
	case "uint8":
		return "C.uint8_t"
	case "uint16":
		return "C.uint16_t"
	case "uint32":
		return "C.uint32_t"
	case "uint64":
		return "C.uint64_t"
	case "int32":
		return "C.int32_t"
	case "int64":
		return "C.int64_t"
	case "float32":
		return "C.float"
	case "float64":
		return "C.double"
	case "uintptr":
		return "C.size_t"
	case "string":
		return "C.WGPUStringView"
	case "unsafe.Pointer":
		return "unsafe.Pointer"
	}
	return "C.WGPU" + goType
}

// ---------------------------------------------------------------------------
// Type predicates
// ---------------------------------------------------------------------------

func isHandleType(ref string) bool {
	return strings.HasPrefix(ref, "object.")
}

func isFuncPtrType(ref string) bool {
	return strings.HasPrefix(ref, "callback.")
}

func isNullableType(ref string) bool {
	return strings.HasPrefix(ref, "object.") || ref == "nullable_string"
}

func isArrayType(ref string) bool {
	return strings.HasPrefix(ref, "array<")
}

func isStringType(ref string) bool {
	return ref == "string_with_default_empty" || ref == "out_string" || ref == "nullable_string"
}

func isStructType(ref string) bool {
	return strings.HasPrefix(ref, "struct.")
}

func isPrimitiveGoType(t string) bool {
	switch t {
	case "bool", "Bool", "uint8", "uint16", "uint32", "uint64",
		"int32", "int64", "float32", "float64", "uintptr":
		return true
	}
	return false
}

func isArrayLenField(name string, members []parser.StructMember) bool {
	for _, m := range members {
		if isArrayType(m.Type) {
			if name == m.Name+"_count" {
				return true
			}
		}
	}
	return false
}

func extractArrayInner(ref string) string {
	if strings.HasPrefix(ref, "array<") && strings.HasSuffix(ref, ">") {
		return ref[len("array<") : len(ref)-1]
	}
	return ref
}

// ---------------------------------------------------------------------------
// Default value expressions
// ---------------------------------------------------------------------------

func memberDefaultValue(m parser.StructMember, spec *parser.Spec) string {
	if m.Default == nil {
		return goZeroValue(goTypeForRef(m.Type, spec))
	}
	goType := goTypeForRef(m.Type, spec)
	switch v := m.Default.(type) {
	case bool:
		if m.Type == "bool" {
			if v {
				return "True"
			}
			return "False"
		}
		if v {
			return "true"
		}
		return "false"
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%g", v)
	case string:
		result := resolveDefaultString(v, m.Type, spec)
		// If the field is a pointer type and the default is a value, take its address.
		if strings.HasPrefix(goType, "*") && !strings.HasPrefix(result, "&") && result != "nil" {
			result = "&" + result
		}
		return result
	}
	return goZeroValue(goType)
}

func resolveDefaultString(val, typeRef string, spec *parser.Spec) string {
	switch val {
	case "zero":
		return goZeroValue(goTypeForRef(typeRef, spec))
	case "nan":
		return "&DepthClearValueUndefined"
	case "constant.whole_size":
		return "WholeSize"
	case "constant.whole_map_size":
		return "WholeMapSize"
	case "constant.undefined":
		return "Undefined"
	case "true":
		if typeRef == "bool" {
			return "True"
		}
		return "true"
	case "false":
		if typeRef == "bool" {
			return "False"
		}
		return "false"
	}

	// Bare enum default (e.g. "auto" for enum.composite_alpha_mode).
	if strings.HasPrefix(typeRef, "enum.") {
		enumName := pascalCase(strings.TrimPrefix(typeRef, "enum."))
		return enumName + pascalCase(val)
	}
	// Bare bitflag default (e.g. "render_attachment" for bitflag.texture_usage).
	if strings.HasPrefix(typeRef, "bitflag.") {
		bfName := pascalCase(strings.TrimPrefix(typeRef, "bitflag."))
		return bfName + pascalCase(val)
	}

	if strings.HasPrefix(val, "constant.") {
		return pascalCase(strings.TrimPrefix(val, "constant."))
	}
	return val
}

func goZeroValue(goType string) string {
	switch goType {
	case "bool":
		return "false"
	case "Bool":
		return "False"
	case "uint8", "uint16", "uint32", "uint64",
		"int32", "int64", "float32", "float64", "uintptr":
		return "0"
	case "string":
		return `""`
	case "unsafe.Pointer":
		return "nil"
	}
	if strings.HasPrefix(goType, "*") {
		return "nil"
	}
	return goType + "{}"
}

func constantGoValue(c parser.Constant) string {
	switch c.Value {
	case "uint32_max":
		return "0xFFFFFFFF"
	case "uint64_max":
		return "0xFFFFFFFFFFFFFFFF"
	case "usize_max":
		return "0xFFFFFFFFFFFFFFFF"
	case "nan":
		return "float32(math.NaN())"
	default:
		return c.Value
	}
}

// ---------------------------------------------------------------------------
// Misc helpers
// ---------------------------------------------------------------------------

var commentRe = regexp.MustCompile(`(?m)^`)

func commentLine(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	return commentRe.ReplaceAllString(s, "// ")
}

// commentIndent is like commentLine but prefixes each line with a tab
// for use inside struct/const blocks.
func commentIndent(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	return commentRe.ReplaceAllString(s, "\t// ")
}

func goTypeName(ref string) string {
	return goTypeForRef(ref, nil)
}

// CallbackArgData is one parameter of a Go callback type.
type CallbackArgData struct {
	Name   string
	GoType string
}

// goCallbackArgs returns the Go parameters for a callback. Status enums
// are passed as raw values (never wrapped in error). The message string is
// dropped. An error_type argument is folded with its following message into
// a single typed error parameter.
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
			skipMessage = true // following message feeds the error
			continue
		}
		if a.Name == "message" && isStringType(a.Type) {
			// Surfaced as a trailing string, unless folded into a typed error.
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

// goCallbackSig renders a callback func type.
func goCallbackSig(cb parser.Callback) string {
	var b strings.Builder
	b.WriteString("func(")
	args := goCallbackArgs(cb)
	for i, a := range args {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%s %s", a.Name, a.GoType)
	}
	b.WriteString(")")
	return b.String()
}

// goSignature renders a complete Go func/method signature (without the body).
func goSignature(fd FuncData) string {
	var b strings.Builder
	args := fd.GoArgs
	if fd.IsMethod && len(args) > 0 {
		recv := args[0]
		args = args[1:]
		fmt.Fprintf(&b, "func (%s %s) %s(", recv.Name, recv.GoType, fd.Name)
	} else {
		fmt.Fprintf(&b, "func %s(", fd.Name)
	}
	first := true
	for _, a := range args {
		if strings.HasPrefix(a.TypeRef, "callback.") {
			continue // blocking wrappers own the callback
		}
		if !first {
			b.WriteString(", ")
		}
		first = false
		fmt.Fprintf(&b, "%s %s", a.Name, a.GoType)
	}
	b.WriteString(")")
	switch {
	case fd.RetKind == "blocking":
		if len(fd.BlockResults) > 0 {
			b.WriteString(" (")
			for i, r := range fd.BlockResults {
				if i > 0 {
					b.WriteString(", ")
				}
				b.WriteString(r.GoType)
			}
			b.WriteString(")")
		}
	case fd.GoReturn != "":
		fmt.Fprintf(&b, " %s", fd.GoReturn)
	}
	return b.String()
}

// structSliceFields returns the Go field names of a struct that are slices
// (need pinning when passed to C). Used by the CGO template.
func structSliceFields(ref string, td *TemplateData) []string {
	name := pascalCase(strings.TrimPrefix(ref, "struct."))
	for _, s := range td.Structs {
		if s.Name == name {
			var fields []string
			for _, m := range s.Members {
				if m.IsSlice {
					fields = append(fields, m.Name)
				}
			}
			return fields
		}
	}
	return nil
}
