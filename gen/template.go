package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// Template loading and template-facing helper functions.
// ---------------------------------------------------------------------------

const (
	lowTmplDir  = "gen/templates"
	wrapTmplDir = "gen/templates/wrap"
)

// baseFuncMap is shared by both template sets.
func baseFuncMap() template.FuncMap {
	return template.FuncMap{
		"camel":         camelCase,
		"pascal":        pascalCase,
		"enumVal":       enumValue,
		"hasPrefix":     strings.HasPrefix,
		"hasSuffix":     strings.HasSuffix,
		"trimPrefix":    strings.TrimPrefix,
		"trimSuffix":    strings.TrimSuffix,
		"upper":         strings.ToUpper,
		"lower":         strings.ToLower,
		"title":         strings.Title,
		"contains":      strings.Contains,
		"replace":       strings.Replace,
		"replaceAll":    strings.ReplaceAll,
		"join":          strings.Join,
		"defaultType":   defaultGoValue,
		"goType":        goTypeName,
		"cgoType":       cgoTypeName,
		"isHandle":      isHandleType,
		"isFuncPtr":     isFuncPtrType,
		"isNullable":    isNullableType,
		"isArray":       isArrayType,
		"isString":      isStringType,
		"isStruct":      isStructType,
		"isArrayLen":    isArrayLenField,
		"comment":       commentLine,
		"commentLine":   commentLine,
		"commentIndent": commentIndent,
		"sig":           goSignature,
		"recv":          receiverName,
		"cbSig":         goCallbackSig,
		"lowerFirst":    lowerFirst,
		"handleType": func(goReturn string) string {
			return strings.TrimPrefix(goReturn, "*")
		},
		"sub": func(a, b int) int { return a - b },
		"add": func(a, b int) int { return a + b },
		"deref": func(s *string) string {
			if s == nil {
				return ""
			}
			return *s
		},
	}
}

// loadTemplates parses the low-level package templates (gen/templates/*.tmpl).
func loadTemplates() (map[string]*template.Template, error) {
	return parseTmplDir(lowTmplDir, "webgpu", baseFuncMap())
}

// loadWrapTemplates parses the wrapper package templates (gen/templates/wrap/*.tmpl).
func loadWrapTemplates() (map[string]*template.Template, error) {
	return parseTmplDir(wrapTmplDir, "wrap", baseFuncMap())
}

// parseTmplDir parses every *.tmpl in dir and keys them by the name with the
// .tmpl suffix stripped (the output filename).
func parseTmplDir(dir, name string, funcMap template.FuncMap) (map[string]*template.Template, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read template dir %s: %w", dir, err)
	}
	tmpls := map[string]*template.Template{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".tmpl") {
			continue
		}
		t, err := template.New(e.Name()).Funcs(funcMap).ParseFiles(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("parse template %s: %w", e.Name(), err)
		}
		tmpls[strings.TrimSuffix(e.Name(), ".tmpl")] = t
	}
	_ = name
	return tmpls, nil
}

// --- template signature helpers ---

// defaultGoValue is the zero value for a Go type (template helper).
func defaultGoValue(goType string) string {
	return goZeroValue(goType)
}

// goTypeName maps a YAML type ref to a Go type (template helper).
func goTypeName(ref string) string {
	return goTypeForRef(ref, nil)
}

// goCallbackSig renders a callback func type.
func goCallbackSig(cb parser.Callback) string {
	var b strings.Builder
	b.WriteString("func(")
	for i, a := range goCallbackArgs(cb) {
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
		if !first {
			b.WriteString(", ")
		}
		first = false
		fmt.Fprintf(&b, "%s %s", a.Name, a.GoType)
	}
	b.WriteString(")")
	if fd.GoReturn != "" {
		fmt.Fprintf(&b, " %s", fd.GoReturn)
	}
	return b.String()
}

// structSliceFields returns the slice member names of a struct (CGO template helper).
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
