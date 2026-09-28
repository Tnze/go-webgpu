package main

import (
	"html/template"
	"strconv"
	"strings"
	"unicode"
)

var tmplFuncs = template.FuncMap{
	"singleLineComments": func(input string) string {
		var out strings.Builder
		for line := range strings.Lines(input) {
			out.WriteString("// ")
			out.WriteString(strings.TrimSpace(line))
			out.WriteString("\n")
		}
		return out.String()
	},
	"singleLineCommentsIndent": func(input, indent string) string {
		var out strings.Builder
		for line := range strings.Lines(input) {
			out.WriteString(indent)
			out.WriteString("// ")
			out.WriteString(strings.TrimSpace(line))
			out.WriteString("\n")
		}
		return out.String()
	},
	"constantCase": strings.ToUpper,
	"pascalCase":   pascalCase,
	"camelCase":    camelCase,
	"toGoLiteral": func(s string) string {
		switch s {
		case "usize_max":
			return "math.MaxUint"
		case "uint32_max":
			return "math.MaxUint32"
		case "uint64_max":
			return "math.MaxUint64"
		case "nan":
			return "math.NaN()"
		default:
			return s
		}
	},
	"bitFlagValue": func(v int) string {
		switch v {
		case 0:
			return "0"
		default:
			return "1 << " + strconv.Itoa(v-1)
		}
	},
	"arrayElem": func(typeName string) string {
		if strings.HasPrefix(typeName, "array<") && strings.HasSuffix(typeName, ">") {
			return typeName[6 : len(typeName)-1]
		}
		return ""
	},
	"goType":    goType,
	"cType":     cType,
	"trimSpace": strings.TrimSpace,
	"hasPrefix": strings.HasPrefix,
	"hasSuffix": strings.HasSuffix,
}

func pascalCase(s string) string {
	var out strings.Builder
	out.Grow(len(s))
	nextUpper := true
	for _, c := range s {
		if nextUpper {
			out.WriteRune(unicode.ToUpper(c))
			nextUpper = false
		} else {
			if c == '_' {
				nextUpper = true
			} else {
				out.WriteRune(c)
			}
		}
	}
	return out.String()
}

func camelCase(s string) string {
	var out strings.Builder
	out.Grow(len(s))
	nextUpper := false
	for _, c := range s {
		if nextUpper {
			out.WriteRune(unicode.ToUpper(c))
			nextUpper = false
		} else {
			if c == '_' {
				nextUpper = true
			} else {
				out.WriteRune(c)
			}
		}
	}
	return out.String()
}

func goType(t string) string {
	if dot := strings.IndexByte(t, '.'); dot > 0 {
		return pascalCase(t[dot+1:])
	}
	switch t {
	case "out_string", "string_with_default_empty", "nullable_string":
		return "StringView"
	case "uint8", "uint16", "uint32", "uint64",
		"int8", "int16", "int32", "int64",
		"float32", "float64":
		return t
	case "usize":
		return "uintptr"
	case "c_void",
		"c_void_data_ptr",
		"c_void_mapped_range_ptr",
		"c_void_a_native_window",
		"c_void_ca_metal_layer",
		"c_void_h_instance",
		"c_void_h_wnd",
		"c_void_wl_display",
		"c_void_wl_surface",
		"c_void_x11_display",
		"c_void_xcb_connection":
		return "unsafe.Pointer"
	}
	return pascalCase(t)
}

func cType(t string) (ctype string) {
	if dot := strings.IndexByte(t, '.'); dot > 0 {
		return "WGPU" + pascalCase(t[dot+1:])
	}
	switch t {
	case "bool":
		ctype = "WGPUBool"
	case "nullable_string", "string_with_default_empty", "out_string":
		ctype = "WGPUStringView"
	case "uint16":
		ctype = "uint16_t"
	case "uint32":
		ctype = "uint32_t"
	case "uint64":
		ctype = "uint64_t"
	case "usize":
		ctype = "size_t"
	case "int16":
		ctype = "int16_t"
	case "int32":
		ctype = "int32_t"
	case "float32", "nullable_float32":
		ctype = "float"
	case "float64", "float64_supertype":
		ctype = "double"
	case "c_void",
		"c_void_data_ptr",
		"c_void_mapped_range_ptr",
		"c_void_a_native_window",
		"c_void_ca_metal_layer",
		"c_void_h_instance",
		"c_void_h_wnd",
		"c_void_wl_display",
		"c_void_wl_surface",
		"c_void_x11_display",
		"c_void_xcb_connection":
		ctype = "void*"
	}
	return
}
