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
└── webgpu.yml                  # upstream spec (fetched by user)
```

### Dual-backend strategy

The library supports two mutually-exclusive calling conventions, selected by
build tags at compile time:

| Backend | Build constraint | When used | Requires |
|---------|-----------------|-----------|----------|
| **CGO** | `//go:build cgo` | Default on all platforms when CGO is enabled | `webgpu.h` header + native library |
| **Syscall** | `//go:build windows && !cgo` | Windows only, when `CGO_ENABLED=0` | `webgpu.dll` in PATH or next to binary |

Both backends share the same public API defined in `webgpu.go` (types, constants,
handles). Only the function *implementations* differ.

CGO backend:
```go
func (d Device) CreateBindGroup(descriptor BindGroupDescriptor) BindGroup {
    c_d := (C.WGPUDevice)(unsafe.Pointer(d.Handle()))
    c_descriptor := (*C.WGPUBindGroupDescriptor)(unsafe.Pointer(&descriptor))
    c_result := C.wgpuDeviceCreateBindGroup(c_d, c_descriptor)
    return BindGroup(uintptr(unsafe.Pointer(c_result)))
}
```

Syscall backend:
```go
func (d Device) CreateBindGroup(descriptor BindGroupDescriptor) BindGroup {
    d_v := d.Handle()
    descriptor_v := uintptr(unsafe.Pointer(&descriptor))
    r1, _, _ := procDeviceCreateBindGroup.Call(d_v, descriptor_v)
    return BindGroup(r1)
}
```

### Handle representation

All WebGPU opaque objects (`WGPUDevice`, `WGPUBuffer`, etc.) are thin typed
uintptr values. Wrapping is zero-allocation:

```go
type Device uintptr // zero is null
```

- **CGO**: `unsafe.Pointer(d.Handle())` → `C.WGPUXxx` when calling C functions.
- **Syscall**: `d.Handle()` passed to `LazyProc.Call()`.

### Object lifetime

WebGPU objects are reference-counted. Each returned Go handle owns one
reference. The binding does **not** install GC cleanup — the caller controls
lifetime and must call `Release()` when done:

```go
buf := d.CreateBuffer(webgpu.BufferDescriptor{ /* ... */ })
defer buf.Release()
```

`Destroy()` (Buffer, Device, Texture, QuerySet) destroys the GPU resource but
does not drop the reference — call `Release()` afterwards. `Handle()` returns
the raw pointer without taking ownership. Do not use a handle after `Release()`.

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
| `bitflag.*` | `type X uint32` | `C.WGPUXxx` | Auto `1<<N` values |
| `struct.*` | `*T` (pointer) | `*C.WGPUXxx` | Descriptors and out-params are passed as pointers |
| `object.*` | `X` (handle) | `C.WGPUXxx` | `type X uintptr`; zero is null |
| `callback.*` | internal only | — | blocking wrappers hide C futures/callbacks |

### Errors and blocking

Operations surface raw C results: status enums stay enums, handles are
zero on failure, mapped pointers are nil on failure. Callback APIs
(`wgpuInstanceRequestAdapter`, `wgpuBufferMapAsync`, …) are wrapped in
blocking methods that wait on an internal channel while pumping events and
return the C message string alongside the status:

```go
adapter, status, msg := instance.RequestAdapter(webgpu.RequestAdapterOptions{})
if status != webgpu.RequestAdapterStatusSuccess {
	log.Fatalf("request adapter: %d: %s", status, msg)
}

status, msg = buf.Map(webgpu.MapModeRead, 0, size) // not MapAsync
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

# Windows, no CGO (needs webgpu.dll)
CGO_ENABLED=0 GOOS=windows go build ./webgpu/
```

## Name conversion

The generator converts C names to idiomatic Go names:

| C name | Go name | Convention |
|---|---|---|
| `WGPUBlendFactor_Zero` | `BlendFactorZero` | Enum: prefix + PascalCase |
| `WGPUBufferUsage_MapRead` | `BufferUsageMapRead` | Bitflag: prefix + PascalCase |
| `WGPUBufferDescriptor` | `BufferDescriptor` | Struct: PascalCase |
| `wgpuDeviceCreateBuffer` | `Device.CreateBuffer` | Method: receiver + PascalCase |
| `wgpuCreateInstance` | `CreateInstance` | Function: PascalCase |
| `array_layer_count_undefined` | `ArrayLayerCountUndefined` | Constant: PascalCase |

Object methods become Go methods whenever a receiver is available (the object
handle is the first argument). Top-level functions without a handle — for
example `wgpuCreateInstance` — stay package-level functions.

Go keywords (`type`, `range`, etc.) get a `Val` suffix to avoid collisions:
`type` → `typeVal`.

## License

BSD-3-Clause (same as webgpu-native/webgpu-headers).
