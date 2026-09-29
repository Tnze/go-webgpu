# go-webgpu

Go bindings for [WebGPU](https://www.w3.org/TR/webgpu/), generated from the
[webgpu-native/webgpu-headers](https://github.com/webgpu-native/webgpu-headers)
spec (`webgpu.yml`).

Two packages:

- **`webgpu`** — 1:1 mirror of the C API (types, enums, handles, raw functions).
- **`gpu`** — ergonomic wrapper on top of `webgpu` (Go methods, GC cleanup,
  blocking async). Most users want this one.

```go
import "github.com/Tnze/go-webgpu/gpu"

instance, err := gpu.CreateInstance(nil)
if err != nil {
	log.Fatal(err)
}
defer instance.Release()

adapter, status, msg := instance.RequestAdapter(&gpu.RequestAdapterOptions{})
if status != gpu.RequestAdapterStatusSuccess {
	log.Fatalf("request adapter: %d: %s", status, msg)
}
defer adapter.Release()
```

## Getting started

```bash
# 1. Fetch the headers and the wgpu-native library for your platform
./scripts/fetch-deps.sh

# 2. Run an example
go run ./examples/compute
go run ./examples/buffer
```

`scripts/fetch-deps.sh` populates `third_party/` with `webgpu.h` and the
wgpu-native binaries for your OS.

## Linking modes

Two independent choices: **how Go calls the C API**, and **how C symbols bind
to the underlying library**. Both are selected with build tags.

### 1. Go → C API

| Mode | Build tag | How it works |
|------|-----------|--------------|
| **Linker** (default) | *(none — requires cgo)* | The Go toolchain links `wgpu*` symbols at build time. |
| **Runtime load** | `webgpu_dynload` | Go loads the shared library and looks up symbols at startup. Windows uses `LoadLibrary`, macOS/Linux use `dlopen`. |

### 2. C symbols → native library (linker mode only)

| Mode | Build tag | Linker input | Runtime needs |
|------|-----------|--------------|---------------|
| **Dynamic** (default) | *(none)* | shared library (`wgpu_native.dll` / `libwgpu_native.so`) | the shared library |
| **DLL file** | `webgpu_dll` | the `.dll` / `.so` file path directly | the shared library |
| **Import library** | `webgpu_dllib` | `wgpu_native.dll.lib` (Windows import library) | `wgpu_native.dll` |
| **Static** | `webgpu_static` | `wgpu_native_static.lib` / `libwgpu_native.a` | nothing |

```bash
go build ./examples/compute                     # dynamic (default)
go build -tags webgpu_static ./examples/compute  # fully static
go build -tags webgpu_dynload ./examples/compute # runtime symbol lookup
go build -tags webgpu_dll ./examples/compute     # link the .dll file directly
go build -tags webgpu_dllib ./examples/compute   # link the import library
```

Notes:

- **`webgpu_static` on Windows** links `wgpu_native_static.lib`, which is an
  MSVC build. Use `clang` from a Visual Studio developer prompt
  (`vcvars64.bat` + `CC=clang`). MinGW's `gcc` cannot resolve the MSVC CRT
  symbols it needs.
- **`webgpu_dynload`** needs no link-time wgpu library at all. On Windows it
  works with `CGO_ENABLED=0`. On macOS/Linux it uses cgo for `dlopen`.
- **`webgpu_dllib`** is Windows-only in practice; on Unix it falls back to the
  shared library.

### Finding the shared library at runtime

The dynamic and runtime-load modes look for `wgpu_native.dll` /
`libwgpu_native.so` in this order:

1. `WGPU_NATIVE_DLL` environment variable (full path)
2. Next to the executable
3. `third_party/wgpu-native/<platform>/lib/` (development builds)
4. The system library search path

## Building with cgo on Windows

```bat
:: after scripts/fetch-deps.sh has populated third_party/
vcvars64.bat
set CC=clang
set CGO_ENABLED=1
go build ./webgpu/
```

With `CGO_ENABLED=0` on Windows the pure-Go syscall backend is used and no C
toolchain is required.

## Examples

| Example | What it shows |
|---------|---------------|
| `examples/compute` | Instance → adapter → device → storage buffer → compute dispatch → map-read |
| `examples/buffer` | Buffer write / read round-trip |
| `go-ecs-std/examples/sprite` | Window + surface + render pipeline + textured sprite |

## API notes

### Handles

Opaque WebGPU objects (`Device`, `Buffer`, `Texture`, …) are `unsafe.Pointer`
handles. In `gpu` they are wrapped structs with GC cleanup; call `Release()`
when done. In `webgpu` they are bare `unsafe.Pointer` types with package-level
functions (`DeviceCreateBuffer(d, …)`).

### Errors and blocking

Async C APIs are wrapped as blocking methods that wait on an internal channel
while pumping events:

```go
adapter, status, msg := instance.RequestAdapter(&gpu.RequestAdapterOptions{})
if status != gpu.RequestAdapterStatusSuccess {
	log.Fatalf("request adapter: %d: %s", status, msg)
}

status, msg = buf.Map(gpu.MapModeRead, 0, size) // blocking, not MapAsync
```

Create-style calls return `(T, error)`. Status-returning calls surface the raw
status enum alongside a message string.

## Regenerating the bindings

Only needed when `webgpu.yml` changes or you modify the generator.

```bash
# fetch the spec
curl -sL https://raw.githubusercontent.com/webgpu-native/webgpu-headers/main/webgpu.yml \
  -o webgpu.yml

# regenerate webgpu/ and gpu/
go run ./gen/ -spec webgpu.yml -out webgpu -out-wrap gpu
```

Flags:

- `-spec` — path to `webgpu.yml` (default `webgpu.yml`)
- `-out` — output directory for `package webgpu` (default `webgpu`)
- `-out-wrap` — output directory for `package gpu` (default `gpu`)

## Name conversion

| C name | Go name |
|---|---|
| `WGPUBlendFactor_Zero` | `BlendFactorZero` |
| `WGPUBufferUsage_MapRead` | `BufferUsageMapRead` |
| `WGPUBufferDescriptor` | `BufferDescriptor` |
| `wgpuDeviceCreateBuffer` | `DeviceCreateBuffer` |
| `wgpuCreateInstance` | `CreateInstance` |
| `array_layer_count_undefined` | `ArrayLayerCountUndefined` |

Go keywords get a `Val` suffix: `type` → `typeVal`.

## License

BSD-3-Clause (same as webgpu-native/webgpu-headers).
