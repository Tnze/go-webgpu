# go-webgpu

Go bindings for [WebGPU](https://www.w3.org/TR/webgpu/), auto-generated from the
[webgpu-native/webgpu-headers](https://github.com/webgpu-native/webgpu-headers)
specification (`webgpu.yml`).

## Architecture

```
go-webgpu/
├── gen/                        # code generator (not imported by users)
│   ├── main.go                 # entrypoint
│   ├── helpers.go              # type mapping, name conversion, template data
│   ├── parser/yaml.go          # webgpu.yml parser
│   └── templates/              # text/template files
│       ├── webgpu.go.tmpl              # constants, enums, bitflags, structs, handles
│       ├── webgpu_defaults.go.tmpl     # NewXxx() default-value constructors
│       ├── webgpu_cgo.go.tmpl          # CGO backend
│       └── webgpu_syscall.go.tmpl      # Windows syscall backend
├── webgpu/                     # generated Go package (package webgpu)
│   ├── webgpu.go               # common types — no build tags
│   ├── webgpu_defaults.go      # struct default constructors
│   ├── webgpu_cgo.go           # CGO function wrappers (//go:build cgo)
│   └── webgpu_syscall.go       # syscall wrappers (//go:build windows && !cgo)
├── third_party/                # fetched by scripts/fetch-deps.sh
│   └── wgpu-native/<plat>/lib/ # libwgpu_native.* / wgpu_native.dll
└── webgpu.yml                  # upstream spec (fetched by user)
```

### Dual-backend strategy

The library supports two mutually-exclusive calling conventions, selected by
build tags at compile time:

| Backend | Build constraint | When used | Requires |
|---------|-----------------|-----------|----------|
| **CGO** | `//go:build cgo` | Default on all platforms when CGO is enabled | `webgpu.h` header + native library |
| **Syscall** | `//go:build windows && !cgo` | Windows only, when `CGO_ENABLED=0` | `wgpu_native.dll` |

Both backends share the same public API defined in `webgpu.go` (types, constants,
handles). Only the function *implementations* differ.

CGO backend:
```go
func DeviceCreateBindGroup(d Device, descriptor *BindGroupDescriptor) BindGroup {
    c_d := (C.WGPUDevice)(unsafe.Pointer(d))
    c_descriptor := (*C.WGPUBindGroupDescriptor)(unsafe.Pointer(descriptor))
    c_result := C.wgpuDeviceCreateBindGroup(c_d, c_descriptor)
    return BindGroup(unsafe.Pointer(c_result))
}
```

Syscall backend:
```go
func DeviceCreateBindGroup(d Device, descriptor *BindGroupDescriptor) BindGroup {
    r1, _, _ := procDeviceCreateBindGroup.Call(uintptr(unsafe.Pointer(d)), uintptr(unsafe.Pointer(descriptor)))
    return BindGroup(unsafe.Pointer(r1))
}
```

### Handle representation

All WebGPU opaque objects (`WGPUDevice`, `WGPUBuffer`, …) are defined types
over `unsafe.Pointer`, matching the C layout (`WGPUNnnImpl *`). Zero is nil:

```go
type Device unsafe.Pointer // nil is null
```

Struct fields that hold objects use the same handle types, so they can be
assigned directly (`colorAtt.View = view`).

Go does not allow methods on types whose underlying type is `unsafe.Pointer`,
so the generated API is package-level functions rather than methods
(`DeviceCreateBuffer(d, …)` instead of `d.CreateBuffer(…)`).

### Object lifetime

WebGPU objects are reference-counted. Each returned Go handle owns one
reference. The binding does **not** install GC cleanup — the caller controls
lifetime and must call the matching `*Release` when done:

```go
buf := webgpu.DeviceCreateBuffer(d, &webgpu.BufferDescriptor{ /* ... */ })
defer webgpu.BufferRelease(buf)
```

`*Destroy` (Buffer, Device, Texture, QuerySet) destroys the GPU resource but
does not drop the reference — call `*Release` afterwards. Do not use a handle
after `*Release`.

## Type mapping reference

| webgpu.yml type | Go type | CGO type | Notes |
|---|---|---|---|
| `uint8/16/32/64` | `uint8/16/32/64` | `C.uint8_t` … `C.uint64_t` | |
| `int32/64` | `int32/64` | `C.int32_t` / `C.int64_t` | |
| `float32/64` | `float32/64` | `C.float` / `C.double` | |
| `bool` | `Bool` | `C.WGPUBool` | `type Bool uint32`; `True` / `False` |
| `usize` | `uintptr` | `C.size_t` | |
| `string_with_default_empty` | `string` | `*C.char` | CGO: `C.CString` + `defer C.free` |
| `nullable_float32` | `*float32` | — | |
| `enum.*` | `type X uint32` | `C.WGPUXxx` | Prefixed names: `BlendFactorZero` |
| `bitflag.*` | `type X uint64` | `C.WGPUXxx` | Auto `1<<N` values |
| `struct.*` | `*T` (pointer) | `*C.WGPUXxx` | Descriptors and out-params are passed as pointers |
| `object.*` | `X` (handle) | `C.WGPUXxx` | `type X unsafe.Pointer`; nil is null |
| `callback.*` | internal only | — | blocking wrappers hide C futures/callbacks |

### Errors and blocking

Operations surface raw C results: status enums stay enums, handles are
nil on failure, mapped pointers are nil on failure. Callback APIs
(`wgpuInstanceRequestAdapter`, `wgpuBufferMapAsync`, …) are wrapped in
blocking functions that wait on an internal channel while pumping events and
return the C message string alongside the status:

```go
adapter, status, msg := webgpu.InstanceRequestAdapter(instance, &webgpu.RequestAdapterOptions{})
if status != webgpu.RequestAdapterStatusSuccess {
	log.Fatalf("request adapter: %d: %s", status, msg)
}

status, msg = webgpu.BufferMap(buf, webgpu.MapModeRead, 0, size) // not MapAsync
```

`error_type` callbacks (`uncaptured_error`, `pop_error_scope`) fold the
message into a typed `error` (`*ValidationError`, `*OutOfMemoryError`, …).
| `c_void_*` | `unsafe.Pointer` | `unsafe.Pointer` | Platform-specific window handles |
| `array<T>` | `[]T` | ptr + count | Expanded to two C args |

### Default value mapping

Struct fields with `default:` in the YAML get explicit values in `NewXxx()`
constructors. The YAML defaults are the canonical WebGPU defaults — they are
**not** inferred from Go zero values.

| YAML default | Go expression | Example |
|---|---|---|
| `false` | `false` | `MappedAtCreation: false` |
| `0` | `0` | `MinBindingSize: 0` |
| `1` | `1` | `SampleCount: 1` |
| `0xFFFFFFFF` | `0xFFFFFFFF` | `StencilReadMask: 0xFFFFFFFF` |
| `constant.whole_size` | `WholeSize` | `Size: WholeSize` |
| `constant.limit_u32_undefined` | `LimitU32Undefined` | All limit fields |
| `constant.depth_clear_value_undefined` | `&DepthClearValueUndefined` | `*float32` pointer |
| `none` (bitflag) | `XxxNone` | `Usage: BufferUsageNone` |
| `all` (bitflag) | `XxxAll` | `WriteMask: ColorWriteMaskAll` |
| `zero` (struct) | `Xxx{}` | `Buffer: BufferBindingLayout{}` |
| bare enum value | `EnumName + Value` | `AlphaMode: CompositeAlphaModeAuto` |

## Generating bindings

### Prerequisites

- Go 1.21+
- The `webgpu.yml` spec file (from
  [webgpu-native/webgpu-headers](https://github.com/webgpu-native/webgpu-headers))

### Fetch the spec

```bash
curl -sL https://raw.githubusercontent.com/webgpu-native/webgpu-headers/main/webgpu.yml \
  -o webgpu.yml
```

### Run the generator

```bash
go run ./gen/ -spec webgpu.yml -out webgpu
```

Flags:
- `-spec` — path to `webgpu.yml` (default: `webgpu.yml`)
- `-out` — output directory (default: `webgpu`)

### Build

```bash
# With CGO (needs webgpu.h + native library installed)
go build ./webgpu/

# Windows, no CGO (needs wgpu_native.dll next to the binary or in PATH;
# dev builds via `go run` also find it under third_party/wgpu-native/)
CGO_ENABLED=0 GOOS=windows go build ./webgpu/
```

#### Windows + CGO

Build from a Visual Studio developer environment so clang can find the MSVC
headers/libs (`vcvars64.bat` or "x64 Native Tools Command Prompt"):

```bat
:: after scripts/fetch-deps.sh (Git Bash) has populated third_party/
vcvars64.bat
set CC=clang
set CGO_ENABLED=1
go build ./webgpu/
```

`third_party/wgpu-native/windows-x86_64-msvc/lib/wgpu_native.lib` must be the
DLL import library (`wgpu_native.dll.lib`); `fetch-deps.sh` copies it into
place. At runtime `wgpu_native.dll` must sit next to the executable or be on
PATH.

## Name conversion

The generator converts C names to idiomatic Go names:

| C name | Go name | Convention |
|---|---|---|
| `WGPUBlendFactor_Zero` | `BlendFactorZero` | Enum: prefix + PascalCase |
| `WGPUBufferUsage_MapRead` | `BufferUsageMapRead` | Bitflag: prefix + PascalCase |
| `WGPUBufferDescriptor` | `BufferDescriptor` | Struct: PascalCase |
| `wgpuDeviceCreateBuffer` | `DeviceCreateBuffer` | Object method: `Type + PascalCase` |
| `wgpuCreateInstance` | `CreateInstance` | Function: PascalCase |
| `array_layer_count_undefined` | `ArrayLayerCountUndefined` | Constant: PascalCase |

Object methods become package-level functions named `TypeMethod`
(`wgpuDeviceCreateBuffer` → `DeviceCreateBuffer`) with the handle as the first
parameter. Top-level functions such as `wgpuCreateInstance` keep their plain
PascalCase names.

Go keywords (`type`, `range`, etc.) get a `Val` suffix to avoid collisions:
`type` → `typeVal`.

## License

BSD-3-Clause (same as webgpu-native/webgpu-headers).
