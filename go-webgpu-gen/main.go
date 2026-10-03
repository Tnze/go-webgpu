package main

import (
	"bytes"
	"embed"
	"flag"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/Tnze/go-webgpu/go-webgpu-gen/spec"
)

//go:embed templates/*
var temps embed.FS

var (
	src        = flag.String("src", "webgpu.yaml", "Input webgpu.yaml file path")
	out        = flag.String("out", "webgpu", "Codegen output path")
	importPath = flag.String("importPath", "github.com/Tnze/go-webgpu/webgpu", "Import path")
	formatCode = flag.Bool("format", true, "Format the generated source")
)

func main() {
	// Parse flags
	flag.Parse()

	var data spec.Spec
	// tmplFuncs["getDef"] = data.GetDef
	// tmplFuncs["isPureData"] = data.IsPureData

	// Parse templates
	tempSys := template.Must(template.
		New("root").Funcs(tmplFuncs).
		ParseFS(temps, "templates/sys/*.tmpl"))
	tempWrapper := template.Must(template.
		New("root").Funcs(tmplFuncs).
		ParseFS(temps, "templates/*.tmpl"))

	// Parse webgpu.yml
	srcFile, err := os.Open(*src)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", *src, err)
	}
	defer srcFile.Close()

	decoder := yaml.NewDecoder(srcFile)
	if err := decoder.Decode(&data); err != nil {
		log.Fatalf("Failed to decode webgpu.yaml: %v", err)
	}

	// Codegen
	data.ImportPath = *importPath
	data.PackageName = data.Name
	codegen(*out, tempWrapper, &data)
	data.PackageName = "sys"
	codegen(filepath.Join(*out, "sys"), tempSys, &data)
}

func codegen(out string, templates *template.Template, data *spec.Spec) {
	if err := os.MkdirAll(out, 0o755); err != nil {
		log.Fatalf("Failed to mkdir: %v", err)
	}
	var buffer bytes.Buffer
	for _, t := range templates.Templates() {
		filename := filepath.Join(out, strings.TrimSuffix(t.Name(), ".tmpl"))
		if !strings.HasSuffix(filename, ".go") {
			continue
		}
		log.Printf("Executing template %q", filename)

		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file %s: %v", filename, err)
		}

		if *formatCode {
			buffer.Reset()

			if err := t.Execute(&buffer, data); err != nil {
				log.Fatalf("Failed to execute template %s: %v", t.Name(), err)
			}

			code, err := format.Source(buffer.Bytes())
			if err != nil {
				log.Fatalf("Failed to formatting code: %v", err)
			}

			if _, err := f.Write(code); err != nil {
				log.Fatalf("Failed to written source file: %v", err)
			}
		} else {
			if err := t.Execute(f, data); err != nil {
				log.Fatalf("Failed to execute template %s: %v", t.Name(), err)
			}
		}

		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close file: %v", err)
		}
	}
}
