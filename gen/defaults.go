package main

import (
	"fmt"
	"strings"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// YAML default → Go expression mapping.
// ---------------------------------------------------------------------------

// memberDefaultValue renders the Go expression for a struct member's default.
func memberDefaultValue(m parser.StructMember, spec *parser.Spec) string {
	if m.Default == nil {
		return goZeroValueRef(m.Type, goTypeForRef(m.Type, spec))
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
		if strings.HasPrefix(goType, "*") && !strings.HasPrefix(result, "&") && result != "nil" {
			result = "&" + result
		}
		return result
	}
	return goZeroValue(goType)
}

// resolveDefaultString expands symbolic defaults (zero, nan, constant.*, …).
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
	if strings.HasPrefix(typeRef, "enum.") {
		return pascalCase(strings.TrimPrefix(typeRef, "enum.")) + pascalCase(val)
	}
	if strings.HasPrefix(typeRef, "bitflag.") {
		return pascalCase(strings.TrimPrefix(typeRef, "bitflag.")) + pascalCase(val)
	}
	if strings.HasPrefix(val, "constant.") {
		return pascalCase(strings.TrimPrefix(val, "constant."))
	}
	return val
}

// goZeroValue is the zero value for a Go type name.
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

// goZeroValueRef is goZeroValue with the original type reference, so object
// handles (type X unsafe.Pointer) get nil.
func goZeroValueRef(typeRef, goType string) string {
	if isHandleType(typeRef) {
		return "nil"
	}
	return goZeroValue(goType)
}

// constantGoValue maps YAML constant values to Go expressions.
func constantGoValue(c parser.Constant) string {
	switch c.Value {
	case "uint32_max":
		return "0xFFFFFFFF"
	case "uint64_max", "usize_max":
		return "0xFFFFFFFFFFFFFFFF"
	case "nan":
		return "float32(math.NaN())"
	default:
		return c.Value
	}
}
