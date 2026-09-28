package main

import (
	"embed"
	"flag"
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
	src = flag.String("src", "webgpu.yaml", "Input webgpu.yaml file path")
	out = flag.String("out", ".", "Output path")
)

func main() {
	// Parse flags
	flag.Parse()

	// Parse templates
	temp := template.Must(template.New("root").Funcs(tmplFuncs).ParseFS(temps, "templates/*.tmpl"))

	// Parse webgpu.yml
	srcFile, err := os.Open(*src)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", *src, err)
	}
	defer srcFile.Close()

	var data spec.Spec
	decoder := yaml.NewDecoder(srcFile)
	if err := decoder.Decode(&data); err != nil {
		log.Fatalf("Failed to decode webgpu.yaml: %v", err)
	}

	// Codegen
	if err := os.MkdirAll(*out, 0o755); err != nil {
		log.Fatalf("Failed to mkdir: %v", err)
	}

	for _, t := range temp.Templates() {
		log.Printf("Executing template %q", t.Name())

		filename := strings.TrimSuffix(t.Name(), ".tmpl")
		f, err := os.Create(filepath.Join(*out, filename))
		if err != nil {
			log.Fatalf("Failed to create file %s: %v", filename, err)
		}

		if err := t.Execute(f, data); err != nil {
			log.Fatalf("Failed to execute template %s: %v", t.Name(), err)
		}

		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close file: %v", err)
		}
	}
}
