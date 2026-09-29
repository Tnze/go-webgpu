package main

import (
	"html/template"
	"strconv"
	"strings"
	"unicode"

	"github.com/Tnze/go-webgpu/go-webgpu-gen/spec"
)

var tmplFuncs = template.FuncMap{
	"goDoc": func(doc, indent string) string {
		doc = strings.TrimSpace(doc)

		if doc == "TODO" {
			return ""
		}

		var out strings.Builder
		for line := range strings.Lines(doc) {
			out.WriteString("//")
			if line := strings.TrimSpace(line); line != "" {
				out.WriteByte(' ')
				out.WriteString(line)
			}
			out.WriteString("\n")
			out.WriteString(indent)
		}
		return out.String()
	},
	"pascalCase": pascalCase,
	"camelCase":  camelCase,
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
	"structDefaultValue": func(input spec.ParameterType) string {
		if input.Default == nil {
			return ""
		}
		if val, found := strings.CutPrefix(*input.Default, "constant."); found {
			return pascalCase(val)
		}
		if *input.Default == "zero" || *input.Default == "false" || *input.Default == "0" || *input.Default == "0.0" {
			return ""
		}
		if input.Type == "float32" || input.Type == "float64" {
			return *input.Default
		}
		if typ, found := strings.CutPrefix(input.Type, "bitflag."); found {
			return pascalCase(typ) + pascalCase(*input.Default)
		}
		if typ, found := strings.CutPrefix(input.Type, "enum."); found {
			return pascalCase(typ) + pascalCase(*input.Default)
		}

		return *input.Default
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
	"isZeroValue": func(input string) bool {
		switch input {
		case "zero", "false", "0", "0.0":
			return true
		}
		return false
	},
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
	if strings.HasPrefix(t, "callback.") {
		return pascalCase(t[9:]) + "CallbackInfo"
	}
	if dot := strings.IndexByte(t, '.'); dot > 0 {
		return pascalCase(t[dot+1:])
	}
	switch t {
	case "out_string", "string_with_default_empty", "nullable_string":
		return "StringView"
	case "uint8", "uint16", "uint32", "uint64",
		"int8", "int16", "int32", "int64":
		return t
	case "float32", "nullable_float32":
		return "float32"
	case "float64", "float64_supertype":
		return "float64"
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
	if strings.HasPrefix(t, "callback.") {
		return "WGPU" + pascalCase(t[9:]) + "CallbackInfo"
	}
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
