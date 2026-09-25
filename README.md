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
func DeviceCreateBindGroup(device *Device, descriptor BindGroupDescriptor) *BindGroup {
    c_device := (C.WGPUDevice)(unsafe.Pointer(device.inner))
    c_descriptor := (*C.WGPUBindGroupDescriptor)(unsafe.Pointer(&descriptor))
    c_result := C.wgpuCreateBindGroup(c_device, c_descriptor)
    return &BindGroup{inner: uintptr(unsafe.Pointer(c_result))}
}
```

Syscall backend:
```go
func DeviceCreateBindGroup(device *Device, descriptor BindGroupDescriptor) *BindGroup {
    device_v := uintptr(device.inner)
    descriptor_v := uintptr(unsafe.Pointer(&descriptor))
    r1, _, _ := procDeviceCreateBindGroup.Call(device_v, descriptor_v)
    return &BindGroup{inner: r1}
}
```

### Handle representation

All WebGPU opaque objects (`WGPUDevice`, `WGPUBuffer`, etc.) are represented as:

```go
type Device struct {
    inner uintptr // stores the native C pointer as a uintptr
}
```

- **CGO**: cast `uintptr` → `unsafe.Pointer` → `C.WGPUXxx` when calling C functions.
- **Syscall**: pass `uintptr` directly to `LazyProc.Call()`.

This avoids importing `unsafe` in the common types file while keeping both
backends zero-allocation.

## Type mapping reference

| webgpu.yml type | Go type | CGO type | Notes |
|---|---|---|---|
| `uint8/16/32/64` | `uint8/16/32/64` | `C.uint8_t` … `C.uint64_t` | |
| `int32/64` | `int32/64` | `C.int32_t` / `C.int64_t` | |
| `float32/64` | `float32/64` | `C.float` / `C.double` | |
| `bool` | `bool` | `C.int` | |
| `usize` | `uintptr` | `C.size_t` | |
| `string_with_default_empty` | `string` | `*C.char` | CGO: `C.CString` + `defer C.free` |
| `nullable_float32` | `*float32` | — | |
| `enum.*` | `type X uint32` | `C.WGPUXxx` | Prefixed names: `BlendFactorZero` |
| `bitflag.*` | `type X uint32` | `C.WGPUXxx` | Auto `1<<N` values |
| `struct.*` | Go struct | `C.WGPUXxx` | Passed by pointer to C |
| `object.*` | `*X` (handle) | `C.WGPUXxx` | Stored as `inner uintptr` |
| `callback.*` | `type XFn func(…)` | — | `Fn` suffix avoids name collisions |
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
| `wgpuDeviceCreateBuffer` | `DeviceCreateBuffer` | Method: Object + PascalCase |
| `wgpuCreateInstance` | `CreateInstance` | Function: PascalCase |
| `array_layer_count_undefined` | `ArrayLayerCountUndefined` | Constant: PascalCase |

Go keywords (`type`, `range`, etc.) get a `Val` suffix to avoid collisions:
`type` → `typeVal`.

## License

BSD-3-Clause (same as webgpu-native/webgpu-headers).
