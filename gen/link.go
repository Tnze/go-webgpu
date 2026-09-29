package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Tnze/go-webgpu/gen/parser"
)

// ---------------------------------------------------------------------------
// Link-mode and dynload-shim generation.
//
// Two independent axes:
//
//	1. Go → C API: either the linker resolves wgpu* symbols (cgo), or Go
//	   loads the shared library at runtime and looks them up (dynload).
//	2. C symbols → native library (cgo only): the shared library, its import
//	   library, or the static library.
//
// Selected with build tags; see README. Generated files:
//
//	webgpu_link_dyn.go     cgo + dynamic link (default)
//	webgpu_link_dll.go     cgo + link wgpu_native.dll / libwgpu_native.so
//	webgpu_link_dllib.go   cgo + link wgpu_native.dll.lib
//	webgpu_link_static.go  cgo + link wgpu_native_static.lib / .a
//	webgpu_dynload_unix.go cgo + dlopen forwarders (macOS/Linux)
//	webgpu_shim.c          C forwarders used by the dynload backend
// ---------------------------------------------------------------------------

// LinkMode describes one cgo link configuration.
type LinkMode struct {
	// File is the generated Go file name (without .go).
	File string
	// Tag is the build-tag name that selects this mode; "" means default.
	Tag string
	// Description is used by the generator log.
	Description string
	// Flags are per-constraint LDFLAGS lines, e.g.
	//   {"windows,amd64": "${SRCDIR}/../third_party/.../wgpu_native.dll.lib"}.
	Flags []LinkFlag
}

// LinkFlag is one `#cgo <constraints> LDFLAGS: <flags>` line.
type LinkFlag struct {
	Constraints string // GOOS,GOARCH; "" for unprefixed
	Flags       string
}

// libDir returns the ${SRCDIR}-relative third_party lib directory for a target.
func libDir(goos, goarch string) string {
	var plat string
	switch goos {
	case "windows":
		if goarch == "arm64" {
			plat = "windows-aarch64-msvc"
		} else {
			plat = "windows-x86_64-msvc"
		}
	case "darwin":
		if goarch == "arm64" {
			plat = "macos-aarch64"
		} else {
			plat = "macos-x86_64"
		}
	default: // linux
		if goarch == "arm64" {
			plat = "linux-aarch64"
		} else {
			plat = "linux-x86_64"
		}
	}
	return "${SRCDIR}/../third_party/wgpu-native/" + plat + "/lib"
}

// targets lists the GOOS/GOARCH pairs we emit LDFLAGS for.
var targets = []struct{ goos, goarch string }{
	{"windows", "amd64"},
	{"windows", "arm64"},
	{"darwin", "arm64"},
	{"darwin", "amd64"},
	{"linux", "amd64"},
	{"linux", "arm64"},
}

func constraint(goos, goarch string) string {
	return goos + "," + goarch
}

// rpathFlag returns the -Wl,-rpath flag needed to find the shared library at
// runtime (no-op on Windows, which searches the executable directory / PATH).
func rpathFlag(dir string) string {
	return "-Wl,-rpath," + dir
}

// windowsStaticSyslibs are the Win32 import libraries wgpu_native_static.lib
// needs. They resolve to the usual .lib files under both cl/link and MinGW.
const windowsStaticSyslibs = " -lws2_32 -luserenv -lntdll -lbcrypt -ld3dcompiler" +
	" -lopengl32 -lgdi32 -luser32 -lshell32 -lole32 -ladvapi32 -lsynchronization" +
	" -lruntimeobject"

// linkModes returns every cgo link configuration, in documentation order.
func linkModes() []LinkMode {
	dyn := LinkMode{
		File:        "webgpu_link_dyn",
		Description: "cgo + dynamic link (shared library)",
	}
	dll := LinkMode{
		File:        "webgpu_link_dll",
		Tag:         "webgpu_dll",
		Description: "cgo + link the shared library file directly",
	}
	dllib := LinkMode{
		File:        "webgpu_link_dllib",
		Tag:         "webgpu_dllib",
		Description: "cgo + link the import library (wgpu_native.dll.lib)",
	}
	static := LinkMode{
		File:        "webgpu_link_static",
		Tag:         "webgpu_static",
		Description: "cgo + static link",
	}

	for _, t := range targets {
		dir := libDir(t.goos, t.goarch)
		c := constraint(t.goos, t.goarch)

		switch t.goos {
		case "windows":
			// cgo rejects bare .lib paths (see go.dev/s/invalidflag), so the
			// import and static libraries are named through -L/-l. -lfoo finds
			// foo.lib; -lwgpu_native.dll finds wgpu_native.dll.lib.
			dyn.Flags = append(dyn.Flags, LinkFlag{c, "-L" + dir + " -lwgpu_native"})
			dll.Flags = append(dll.Flags, LinkFlag{c, dir + "/wgpu_native.dll"})
			dllib.Flags = append(dllib.Flags, LinkFlag{c, "-L" + dir + " -lwgpu_native.dll"})
			// The static library is an MSVC build and pulls in the CRT and
			// Win32 networking/graphics stacks; link them explicitly.
			static.Flags = append(static.Flags, LinkFlag{c,
				"-L" + dir + " -lwgpu_native_static" + windowsStaticSyslibs})
		default:
			lib := "libwgpu_native.so"
			if t.goos == "darwin" {
				lib = "libwgpu_native.dylib"
			}
			// Default and "link the .so directly" are the same artifact on Unix.
			flags := "-L" + dir + " -lwgpu_native " + rpathFlag(dir)
			dyn.Flags = append(dyn.Flags, LinkFlag{c, flags})
			dll.Flags = append(dll.Flags, LinkFlag{c, dir + "/" + lib + " " + rpathFlag(dir)})
			// No import library on Unix; fall back to the shared library so the
			// tag still builds.
			dllib.Flags = append(dllib.Flags, LinkFlag{c, flags})
			static.Flags = append(static.Flags, LinkFlag{c, dir + "/libwgpu_native.a"})
		}
	}
	return []LinkMode{dyn, dll, dllib, static}
}

// generateLinkFiles emits one tiny Go file per link mode. Each file carries
// only the #cgo LDFLAGS directives; the call wrappers live in webgpu_cgo.go.
func generateLinkFiles(spec *parser.Spec) (map[string][]byte, error) {
	// Every cgo file needs the include path for webgpu.h.
	cflags := "#cgo CFLAGS: -I${SRCDIR}/../third_party/webgpu-headers"

	files := map[string][]byte{}
	for _, mode := range linkModes() {
		tags := linkModeTags(mode)
		var b bytes.Buffer
		fmt.Fprintf(&b, "// Code generated by gen; DO NOT EDIT.\n")
		fmt.Fprintf(&b, "//go:build %s\n\n", tags)
		fmt.Fprintf(&b, "package webgpu\n\n/*\n%s\n", cflags)
		for _, f := range mode.Flags {
			if f.Constraints != "" {
				fmt.Fprintf(&b, "#cgo %s LDFLAGS: %s\n", f.Constraints, f.Flags)
			} else {
				fmt.Fprintf(&b, "#cgo LDFLAGS: %s\n", f.Flags)
			}
		}
		b.WriteString("*/\nimport \"C\"\n")
		files[mode.File+".go"] = formatSource(mode.File+".go", b.Bytes())
	}
	return files, nil
}

// linkModeTags renders the build constraint selecting one link mode.
//
// All modes are mutually exclusive and all are suppressed by webgpu_dynload,
// which does no link-time binding at all.
func linkModeTags(mode LinkMode) string {
	// cgo && !webgpu_dynload && (this mode's tag, or none of the others).
	var parts []string
	parts = append(parts, "cgo", "!webgpu_dynload")
	for _, other := range linkModes() {
		if other.Tag == "" {
			continue
		}
		if other.Tag == mode.Tag {
			parts = append(parts, other.Tag)
		} else {
			parts = append(parts, "!"+other.Tag)
		}
	}
	return strings.Join(parts, " && ")
}

// ---------------------------------------------------------------------------
// Dynload shim (macOS/Linux, cgo + dlopen).
//
// webgpu_shim.c defines every wgpu* entry point as a forwarder through a
// dlsym'd function pointer, so the Go call wrappers are identical to the
// linker-based backend. The file is always compiled when cgo is on, but its
// body is empty unless -DWGPU_DYNLOAD is set by webgpu_dynload_unix.go.
// ---------------------------------------------------------------------------

// shimFuncs collects every C entry point the Go wrappers call.
func shimFuncs(spec *parser.Spec) []CSignature {
	var out []CSignature
	for _, f := range spec.Functions {
		out = append(out, buildCSignature(f))
	}
	for _, obj := range spec.Objects {
		for _, m := range obj.Methods {
			out = append(out, buildMethodCSignature(obj, m))
		}
		out = append(out, releaseCSignature(pascalCase(obj.Name)))
	}
	for _, s := range spec.Structs {
		if s.FreeMembers {
			out = append(out, freeMembersCSignature(pascalCase(s.Name)))
		}
	}
	return out
}

// generateDynloadShim renders webgpu_shim.c.
func generateDynloadShim(spec *parser.Spec) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("// Code generated by gen; DO NOT EDIT.\n")
	b.WriteString("//\n")
	b.WriteString("// Thin forwarders that call through dlsym'd function pointers.\n")
	b.WriteString("// Compiled empty unless WGPU_DYNLOAD is defined; see webgpu_dynload_unix.go.\n\n")
	b.WriteString("//go:build cgo && webgpu_dynload && !windows\n\n")
	b.WriteString("#ifdef WGPU_DYNLOAD\n\n")
	b.WriteString("#include \"webgpu.h\"\n")
	b.WriteString("#include <dlfcn.h>\n")
	b.WriteString("#include <stdio.h>\n")
	b.WriteString("#include <stdlib.h>\n\n")

	b.WriteString("static void *wgpu_shim_handle;\n\n")

	// Function-pointer declarations.
	for _, s := range shimFuncs(spec) {
		fmt.Fprintf(&b, "%s;\n", s.PointerDecl())
	}
	b.WriteString("\n")

	// Loader.
	b.WriteString("static void wgpu_shim_load_sym(void **dst, const char *name) {\n")
	b.WriteString("\tvoid *p = dlsym(wgpu_shim_handle, name);\n")
	b.WriteString("\tif (!p) {\n")
	b.WriteString("\t\tfprintf(stderr, \"go-webgpu: missing symbol %s: %s\\n\", name, dlerror());\n")
	b.WriteString("\t\tabort();\n")
	b.WriteString("\t}\n")
	b.WriteString("\t*dst = p;\n")
	b.WriteString("}\n\n")

	b.WriteString("void wgpuLoadLibrary(const char *path) {\n")
	b.WriteString("\tif (wgpu_shim_handle) {\n")
	b.WriteString("\t\treturn;\n")
	b.WriteString("\t}\n")
	b.WriteString("\twgpu_shim_handle = dlopen(path, RTLD_NOW | RTLD_LOCAL);\n")
	b.WriteString("\tif (!wgpu_shim_handle) {\n")
	b.WriteString("\t\tfprintf(stderr, \"go-webgpu: dlopen %s: %s\\n\", path, dlerror());\n")
	b.WriteString("\t\tabort();\n")
	b.WriteString("\t}\n")
	for _, s := range shimFuncs(spec) {
		fmt.Fprintf(&b, "\twgpu_shim_load_sym((void **)&p_%s, \"%s\");\n", s.Name, s.Name)
	}
	b.WriteString("}\n\n")

	// Forwarders.
	for _, s := range shimFuncs(spec) {
		fmt.Fprintf(&b, "%s {\n", s.Prototype())
		if s.Return == "void" {
			fmt.Fprintf(&b, "\tp_%s(%s);\n", s.Name, s.CallArgs())
		} else {
			fmt.Fprintf(&b, "\treturn p_%s(%s);\n", s.Name, s.CallArgs())
		}
		b.WriteString("}\n\n")
	}

	b.WriteString("#endif /* WGPU_DYNLOAD */\n")
	return b.Bytes(), nil
}

// generateDynloadUnix renders webgpu_dynload_unix.go: the -DWGPU_DYNLOAD
// switch, the -ldl link flag, library search, and the init hook.
func generateDynloadUnix(spec *parser.Spec) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("// Code generated by gen; DO NOT EDIT.\n")
	b.WriteString("//go:build cgo && webgpu_dynload && !windows\n\n")
	b.WriteString(`// Package webgpu runtime-loading backend for macOS and Linux.
//
// Instead of linking wgpu_native at build time, this file turns on the C
// forwarders in webgpu_shim.c and resolves every wgpu* symbol through dlopen
// at startup. The Go call wrappers are shared with the linker-based backend.
package webgpu

/*
#cgo CFLAGS: -I${SRCDIR}/../third_party/webgpu-headers -DWGPU_DYNLOAD
#cgo linux LDFLAGS: -ldl
#include <stdlib.h>
void wgpuLoadLibrary(const char *path);
*/
import "C"

import (
	"os"
	"path/filepath"
	"runtime"
	"unsafe"
)

// loadLibrary resolves the shared-library path and hands it to the C shim.
func loadLibrary() {
	path := resolveLibraryPath()
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.wgpuLoadLibrary(cpath)
}

// resolveLibraryPath picks the wgpu_native shared library to dlopen.
// Search order: WGPU_NATIVE_DLL, next to the executable, third_party, bare name.
func resolveLibraryPath() string {
	name := "libwgpu_native.so"
	if runtime.GOOS == "darwin" {
		name = "libwgpu_native.dylib"
	}
	if p := os.Getenv("WGPU_NATIVE_DLL"); p != "" {
		return p
	}
	if exe, err := os.Executable(); err == nil {
		if p := filepath.Join(filepath.Dir(exe), name); fileExists(p) {
			return p
		}
	}
	if _, file, _, ok := runtime.Caller(0); ok {
		pattern := filepath.Join(filepath.Dir(file), "..", "third_party", "wgpu-native", runtime.GOOS+"-*", "lib", name)
		if matches, err := filepath.Glob(pattern); err == nil && len(matches) > 0 {
			return matches[0]
		}
	}
	return name
}

func fileExists(p string) bool {
	st, err := os.Stat(p)
	return err == nil && !st.IsDir()
}

func init() {
	loadLibrary()
}
`)
	return formatSource("webgpu_dynload_unix.go", b.Bytes()), nil
}
