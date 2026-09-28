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
	"pascalCase": func(s string) string {
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
	},
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
	"trimSpace": strings.TrimSpace,
}
