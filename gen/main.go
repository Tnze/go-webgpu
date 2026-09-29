// Command gen reads the webgpu.yml spec from webgpu-native/webgpu-headers
// and generates the low-level webgpu bindings plus the ergonomic gpu wrapper.
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

	"github.com/Tnze/go-webgpu/gen/parser"
)

var (
	specPath   = flag.String("spec", "webgpu.yml", "path to webgpu.yml spec file")
	outputDir  = flag.String("out", "webgpu", "output directory for the low-level package")
	wrapOutDir = flag.String("out-wrap", "gpu", "output directory for the ergonomic wrapper package")
)

func main() {
	flag.Parse()

	spec, err := parser.ParseFile(*specPath)
	if err != nil {
		log.Fatalf("parse spec: %v", err)
	}

	files, err := generate(spec)
	if err != nil {
		log.Fatalf("generate webgpu: %v", err)
	}
	if err := writeFiles(*outputDir, files); err != nil {
		log.Fatal(err)
	}

	wrapFiles, err := generateWrapPackage(spec)
	if err != nil {
		log.Fatalf("generate gpu: %v", err)
	}
	if err := writeFiles(*wrapOutDir, wrapFiles); err != nil {
		log.Fatal(err)
	}
}

// generate renders the low-level webgpu package templates plus the
// link-mode / dynload files that select how C symbols bind to the native
// library.
func generate(spec *parser.Spec) (map[string][]byte, error) {
	tmpls, err := loadTemplates()
	if err != nil {
		return nil, err
	}
	data := buildTemplateData(spec)
	files := make(map[string][]byte, len(tmpls)+8)
	for name, t := range tmpls {
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return nil, fmt.Errorf("execute %q: %w", name, err)
		}
		files[name] = formatSource(name, buf.Bytes())
	}

	linkFiles, err := generateLinkFiles(spec)
	if err != nil {
		return nil, fmt.Errorf("generate link files: %w", err)
	}
	for name, data := range linkFiles {
		files[name] = data
	}

	shim, err := generateDynloadShim(spec)
	if err != nil {
		return nil, fmt.Errorf("generate dynload shim: %w", err)
	}
	files["webgpu_shim.c"] = shim

	dynload, err := generateDynloadUnix(spec)
	if err != nil {
		return nil, fmt.Errorf("generate dynload unix: %w", err)
	}
	files["webgpu_dynload_unix.go"] = dynload

	return files, nil
}

// writeFiles creates dir, replaces any previously generated sources, and
// writes the new ones.
func writeFiles(dir string, files map[string][]byte) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}
	// Drop stale generated files so renamed outputs do not linger.
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read %s: %w", dir, err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, ".c") {
			continue
		}
		if _, keep := files[name]; !keep {
			if err := os.Remove(filepath.Join(dir, name)); err != nil {
				return fmt.Errorf("remove stale %s: %w", filepath.Join(dir, name), err)
			}
		}
	}
	for name, data := range files {
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, data, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", p, err)
		}
		fmt.Printf("generated %s (%d bytes)\n", p, len(data))
	}
	return nil
}

// formatSource runs gofmt on Go sources. The cgo preamble can make
// format.Source fail, in which case the raw bytes are kept.
func formatSource(name string, data []byte) []byte {
	if !strings.HasSuffix(name, ".go") {
		return data
	}
	if formatted, err := format.Source(data); err == nil {
		return formatted
	}
	return data
}
