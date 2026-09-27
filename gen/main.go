// Command gen reads the webgpu.yml spec from webgpu-native/webgpu-headers
// and generates Go bindings for the webgpu package.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Tnze/go-webgpu/gen/parser"
)

var (
	specPath  = flag.String("spec", "webgpu.yml", "path to webgpu.yml spec file")
	outputDir = flag.String("out", "webgpu", "output directory for generated Go files")
)

func main() {
	flag.Parse()

	spec, err := parser.ParseFile(*specPath)
	if err != nil {
		log.Fatalf("failed to parse spec: %v", err)
	}

	if err := os.MkdirAll(*outputDir, 0o755); err != nil {
		log.Fatalf("failed to create output dir: %v", err)
	}

	files, err := generate(spec)
	if err != nil {
		log.Fatalf("generation failed: %v", err)
	}

	for name, data := range files {
		p := filepath.Join(*outputDir, name)
		if err := os.WriteFile(p, data, 0o644); err != nil {
			log.Fatalf("write %s: %v", p, err)
		}
		fmt.Printf("generated %s (%d bytes)\n", p, len(data))
	}
}

// generate returns a map of filename -> file content for the webgpu package.
func generate(spec *parser.Spec) (map[string][]byte, error) {
	tmpl, err := loadTemplates()
	if err != nil {
		return nil, fmt.Errorf("load templates: %w", err)
	}

	data := buildTemplateData(spec)
	files := make(map[string][]byte)

	for name, t := range tmpl {
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return nil, fmt.Errorf("execute template %q: %w", name, err)
		}
		// Strip .tmpl suffix from output filename.
		outName := strings.TrimSuffix(name, ".tmpl")
		// Run gofmt on .go files (skip if format.Source fails, e.g. cgo preamble).
		if strings.HasSuffix(outName, ".go") {
			if formatted, err := format.Source(buf.Bytes()); err == nil {
				files[outName] = formatted
				continue
			}
		}
		files[outName] = buf.Bytes()
	}

	return files, nil
}

func loadTemplates() (map[string]*template.Template, error) {
	funcMap := template.FuncMap{
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

	tmplDir := "gen/templates"
	entries, err := os.ReadDir(tmplDir)
	if err != nil {
		return nil, fmt.Errorf("read template dir: %w", err)
	}

	tmpls := make(map[string]*template.Template)
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".tmpl") {
			continue
		}
		t, err := template.New(e.Name()).Funcs(funcMap).ParseFiles(filepath.Join(tmplDir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("parse template %s: %w", e.Name(), err)
		}
		tmpls[e.Name()] = t
	}

	return tmpls, nil
}

// defaultGoValue returns the zero value for a Go type (used in templates).
func defaultGoValue(goType string) string {
	return goZeroValue(goType)
}
