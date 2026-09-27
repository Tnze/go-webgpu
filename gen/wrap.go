package main

import (
	"bytes"
	"fmt"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// High-level wrapper package (package gpu) — model, entry, data assembly.
//
// All conversion / call-shape logic is precomputed into string slices; the
// text/template files in gen/templates/wrap/ only stitch those lines together.
// ---------------------------------------------------------------------------

// WrapData is the root passed to the wrap templates.
type WrapData struct {
	Package  string
	Handles  []WrapHandle
	Structs  []WrapStruct
	Enums    []EnumData
	Bitflags []BitflagData
	Consts   []ConstData
	Aliases  []WrapAlias // enum/bitflag constants re-exported from webgpu
	Helpers  string      // extra source appended to types.go
}

// WrapAlias is a re-exported enum or bitflag constant.
type WrapAlias struct {
	Name  string
	Value string // webgpu.XXX
}

// WrapHandle is an opaque object wrapped with runtime.Cleanup.
type WrapHandle struct {
	Name    string
	Doc     string
	Methods []WrapMethod
}

// WrapMethod is one generated method (or package-level function).
type WrapMethod struct {
	Name string
	Doc  string
	Sig  string   // full signature
	Body []string // body lines (one leading tab added by the template)
}

// WrapStruct is a Go-idiomatic mirror of a transparent C struct.
type WrapStruct struct {
	Name       string
	Doc        string
	Fields     []WrapField
	HasFree    bool
	HasDefault bool // low-level NewXxx() applies YAML defaults
	ExtraDoc   string
	ExtraField string
	ExtraToRaw []string
}

// WrapField is one field of a wrapper struct with precomputed conversions.
type WrapField struct {
	Name     string
	Type     string
	Doc      string
	SkipZero bool     // don't overwrite non-zero YAML default with zero
	ToRaw    []string // lines filling raw.Field from s.Field
	FromRaw  []string // lines filling out.Field from raw.Field
}

// generateWrapPackage builds the wrapper package source files.
func generateWrapPackage(spec *parser.Spec) (map[string][]byte, error) {
	tmpls, err := loadWrapTemplates()
	if err != nil {
		return nil, err
	}
	data := buildWrapData(spec)
	files := map[string][]byte{}
	for name, t := range tmpls {
		var buf bytes.Buffer
		if err := t.Execute(&buf, data); err != nil {
			return nil, fmt.Errorf("execute wrap template %q: %w", name, err)
		}
		files[name] = formatSource(name, buf.Bytes())
	}
	return files, nil
}

func buildWrapData(spec *parser.Spec) *WrapData {
	wd := &WrapData{Package: "gpu", Helpers: wrapHelpersSrc}

	for _, c := range spec.Constants {
		wd.Consts = append(wd.Consts, ConstData{
			Name:  pascalCase(c.Name),
			Value: constantGoValue(c),
			Doc:   cleanDoc(c.Doc),
		})
	}
	for _, e := range spec.Enums {
		wd.Enums = append(wd.Enums, EnumData{Name: pascalCase(e.Name), Doc: cleanDoc(e.Doc)})
		for _, en := range e.Entries {
			if en.IsNull {
				continue
			}
			n := pascalCase(e.Name) + pascalCase(en.Name)
			wd.Aliases = append(wd.Aliases, WrapAlias{Name: n, Value: "webgpu." + n})
		}
	}
	for _, bf := range spec.Bitflags {
		wd.Bitflags = append(wd.Bitflags, BitflagData{Name: pascalCase(bf.Name), Doc: cleanDoc(bf.Doc)})
		for _, en := range bf.Entries {
			n := pascalCase(bf.Name) + pascalCase(en.Name)
			wd.Aliases = append(wd.Aliases, WrapAlias{Name: n, Value: "webgpu." + n})
		}
	}
	wd.Aliases = append(wd.Aliases,
		WrapAlias{Name: "True", Value: "webgpu.True"},
		WrapAlias{Name: "False", Value: "webgpu.False"},
	)

	for _, s := range spec.Structs {
		// Extension-chain structs are provided by wrapHelpersSrc.
		if s.Type == "extension" {
			continue
		}
		wd.Structs = append(wd.Structs, buildWrapStruct(s, spec))
	}
	for _, obj := range spec.Objects {
		wh := buildWrapHandle(obj, spec)
		if wh.Name == "Instance" {
			wh.Methods = append([]WrapMethod{buildCreateInstance()}, wh.Methods...)
		}
		wd.Handles = append(wd.Handles, wh)
	}
	return wd
}

// isKnownHandle reports whether name is an object type in the spec.
func isKnownHandle(spec *parser.Spec, name string) bool {
	for _, o := range spec.Objects {
		if pascalCase(o.Name) == name {
			return true
		}
	}
	return false
}

// wrapHelpersSrc is appended to types.go: hand-written extension-chain types
// and out-param converters that are awkward to express as generated mirrors.
const wrapHelpersSrc = `
func compilationInfoFromRaw(raw webgpu.CompilationInfo) *CompilationInfo {
	msgs := make([]CompilationMessage, 0, raw.MessagesCount)
	for _, m := range unsafe.Slice(raw.Messages, raw.MessagesCount) {
		msgs = append(msgs, CompilationMessage{
			Message: m.Message,
			Type:    m.Type,
			LineNum: m.LineNum,
			LinePos: m.LinePos,
			Offset:  m.Offset,
			Length:  m.Length,
		})
	}
	return &CompilationInfo{Messages: msgs}
}

// ShaderSourceWGSL is the WGSL shader source extension chain.
type ShaderSourceWGSL struct {
	Code string
}

// SurfaceSource is a platform-specific surface binding source.
type SurfaceSource interface{ surfaceSource() }

type SurfaceSourceWindowsHWND struct {
	Hinstance unsafe.Pointer
	Hwnd      unsafe.Pointer
}

func (*SurfaceSourceWindowsHWND) surfaceSource() {}

type SurfaceSourceMetalLayer struct {
	Layer unsafe.Pointer
}

func (*SurfaceSourceMetalLayer) surfaceSource() {}
`
