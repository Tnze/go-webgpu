package main

import (
	"strings"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// YAML type-reference → Go / CGO type mapping, and type predicates.
// ---------------------------------------------------------------------------

// goTypeForRef maps a YAML type reference to its Go type name.
// spec is optional; it is currently unused but kept for future struct-aware mapping.
func goTypeForRef(ref string, spec *parser.Spec) string {
	if ref == "" {
		return "unsafe.Pointer"
	}
	switch ref {
	case "bool":
		return "Bool"
	case "uint8", "uint16", "uint32", "uint64",
		"int32", "int64", "float32", "float64", "float64_supertype",
		"usize", "string_with_default_empty", "out_string", "nullable_string":
		if ref == "float64_supertype" {
			return "float64"
		}
		if isStringType(ref) {
			return "string"
		}
		if ref == "usize" {
			return "uintptr"
		}
		return ref
	}
	switch {
	case strings.HasPrefix(ref, "enum."):
		return pascalCase(strings.TrimPrefix(ref, "enum."))
	case strings.HasPrefix(ref, "bitflag."):
		return pascalCase(strings.TrimPrefix(ref, "bitflag."))
	case strings.HasPrefix(ref, "struct."):
		return pascalCase(strings.TrimPrefix(ref, "struct."))
	case strings.HasPrefix(ref, "object."):
		return pascalCase(strings.TrimPrefix(ref, "object."))
	case strings.HasPrefix(ref, "callback."):
		return pascalCase(strings.TrimPrefix(ref, "callback.")) + "Fn"
	case strings.HasPrefix(ref, "constant."):
		return "uint32"
	case strings.HasPrefix(ref, "c_void"):
		return "unsafe.Pointer"
	case ref == "nullable_float32":
		return "*float32"
	}
	return pascalCase(ref)
}

// goTypeForMember maps a struct member to its Go field type.
func goTypeForMember(m parser.StructMember) string {
	// WGPUStringView (ptr+len) matches a Go string header on 64-bit.
	if isStringType(m.Type) {
		return "string"
	}
	if strings.HasPrefix(m.Type, "callback.") {
		return pascalCase(strings.TrimPrefix(m.Type, "callback.")) + "CallbackInfo"
	}
	// Handles inside structs are the named object types (type X unsafe.Pointer).
	if isHandleType(m.Type) {
		return pascalCase(strings.TrimPrefix(m.Type, "object."))
	}
	// Optional / pointed-to structs become Go pointers.
	if isStructType(m.Type) && (m.Optional || m.Pointer == "immutable" || m.Pointer == "mutable") {
		return "*" + pascalCase(strings.TrimPrefix(m.Type, "struct."))
	}
	if isArrayType(m.Type) {
		return "[]" + goTypeForRef(extractArrayInner(m.Type), nil)
	}
	return goTypeForRef(m.Type, nil)
}

// cgoTypeName maps a Go type name to the corresponding C.WGPU… type.
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

// --- type predicates on YAML refs ---

func isHandleType(ref string) bool   { return strings.HasPrefix(ref, "object.") }
func isFuncPtrType(ref string) bool  { return strings.HasPrefix(ref, "callback.") }
func isArrayType(ref string) bool    { return strings.HasPrefix(ref, "array<") }
func isStructType(ref string) bool   { return strings.HasPrefix(ref, "struct.") }
func isNullableType(ref string) bool { return isHandleType(ref) || ref == "nullable_string" }

func isStringType(ref string) bool {
	return ref == "string_with_default_empty" || ref == "out_string" || ref == "nullable_string"
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
		if isArrayType(m.Type) && name == m.Name+"_count" {
			return true
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

// collectStructPins walks a struct type and records expressions of members
// that hold Go pointers (strings, arrays, optional/pointer fields) so the
// generated wrappers can pin them before passing the struct to C.
// By-value nested structs are walked recursively; pointer fields are pinned
// at one level (their pointees are assumed free of Go pointers).
func collectStructPins(spec *parser.Spec, typeRef, expr string, strFields, ptrFields, sliceFields *[]string, depth int) {
	if spec == nil || depth > 8 {
		return
	}
	structName := pascalCase(strings.TrimPrefix(typeRef, "struct."))
	for _, s := range spec.Structs {
		if pascalCase(s.Name) != structName {
			continue
		}
		for _, m := range s.Members {
			field := expr + "." + pascalCase(m.Name)
			switch {
			case isStringType(m.Type):
				*strFields = append(*strFields, field)
			case isArrayType(m.Type):
				*sliceFields = append(*sliceFields, field)
			case isStructType(m.Type):
				if m.Pointer == "immutable" || m.Pointer == "mutable" || m.Optional {
					*ptrFields = append(*ptrFields, field)
				} else {
					collectStructPins(spec, m.Type, field, strFields, ptrFields, sliceFields, depth+1)
				}
			}
		}
		return
	}
}
