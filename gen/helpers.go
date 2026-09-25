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

// CallbackInfoData describes an embedded C callback-info struct.
type CallbackInfoData struct {
	Name    string // e.g. DeviceLostCallbackInfo
	CName   string // e.g. WGPUDeviceLostCallbackInfo
	FnType  string // e.g. DeviceLostFn
	HasMode bool   // true except uncaptured_error
}

// FuncData describes a function to generate.
type FuncData struct {
	Name          string
	CName         string
	GoArgs        []FuncArgData
	GoReturn      string
	ReturnRef     string // original YAML type reference for the return type
	Doc           string
	IsMethod      bool
	ObjName       string
	HasCallback   bool
	CallbackFn    string // Go callback type, e.g. RequestAdapterFn
	CallbackName  string // original callback name, e.g. request_adapter
	ReturnsFuture bool
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
			}
		}
		for _, obj := range spec.Objects {
			for _, m := range obj.Methods {
				if m.Callback == "callback."+cb.Name {
					used = true
				}
			}
		}
		if !used {
			continue
		}
		seenCB[infoName] = true
		td.CallbackInfos = append(td.CallbackInfos, CallbackInfoData{
			Name:    infoName,
			CName:   "WGPU" + infoName,
			FnType:  pascalCase(cb.Name) + "Fn",
			HasMode: cb.Name != "uncaptured_error",
		})
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

	// Functions
	for _, f := range spec.Functions {
		td.Funcs = append(td.Funcs, buildFuncData(f, spec))
	}
	for _, obj := range spec.Objects {
		for _, m := range obj.Methods {
			fd := buildFuncData(m, spec)
			fd.IsMethod = true
			fd.ObjName = pascalCase(obj.Name)
			// Prefix method name with object name to avoid collisions.
			fd.Name = fd.ObjName + fd.Name
			// C API always prefixes methods with the object name.
			fd.CName = "wgpu" + fd.ObjName + pascalCase(m.Name)
			// Prepend the object handle as the first argument.
			handleArg := FuncArgData{
				Name:    camelCase(obj.Name),
				GoType:  "*" + pascalCase(obj.Name),
				TypeRef: "object." + obj.Name,
				CName:   obj.Name,
			}
			fd.GoArgs = append([]FuncArgData{handleArg}, fd.GoArgs...)
			td.Funcs = append(td.Funcs, fd)
		}
	}

	return td
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
		// All callback-taking APIs return a Future in the C ABI.
		fd.GoReturn = "Future"
		fd.ReturnRef = "struct.future"
		fd.ReturnsFuture = true
	} else if f.Returns != nil && f.Returns.Type != "" && f.Returns.Type != "void" {
		fd.GoReturn = goTypeForRef(f.Returns.Type, spec)
		fd.ReturnRef = f.Returns.Type
	}
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
		fad := FuncArgData{
			Name:       camelCase(arg.Name),
			GoType:     goType,
			TypeRef:    arg.Type,
			CName:      arg.Name,
			IsPtr:      arg.Pointer == "immutable" || arg.Pointer == "mutable",
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
		return "bool"
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
		return "*" + pascalCase(strings.TrimPrefix(ref, "object."))
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
	// C ABI: WGPUBool is uint32_t.
	if m.Type == "bool" {
		return "uint32"
	}
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
	case "bool":
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
	case "bool", "uint8", "uint16", "uint32", "uint64",
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
			// C WGPUBool is uint32.
			if v {
				return "1"
			}
			return "0"
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
		return "true"
	case "false":
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
