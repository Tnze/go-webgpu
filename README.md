# go-webgpu

A WebGPU binding for Go programming language. Based on codegen.

## Re-generate codes

```
go run ./go-webgpu-gen \
    -src webgpu-headers/webgpu.yml \
    -out webgpu \
    -importPath "github.com/Tnze/go-webgpu/webgpu"
```

## Project Configuration

To use this library in your project. Considering how do you want to link to the WebGPU implementation.

There is mainly two ways to call WebGPU APIs supported by this library: By CGO or `syscall.LazyDLL/LazyProc` (Windows only).

### By CGO

By using CGO, you can link WebGPU implementation dynamically or statically.
To doing so, you should set the `CGO_LDFLAGS` before building your Go project, specific
the WebGPU implementation library linkage options.

Further more, you can have more flexible controls by using generic `LDFLAGS` options,
e.g., static linking.

For example, on Linux/macOS:

```bash
export CGO_ENABLE=1
export CGO_LDFLAGS="-L/path/to/wgpu-native-lib -lwgpu-native"
```

On Windows:

```powershell
$env:CC="C:\Program Files\Microsoft Visual Studio\18\Community\VC\Tools\Llvm\x64\bin\clang.exe"
$env:CGO_ENABLE="1"
$env:CGO_LDFLAGS="-L/path/to/wgpu-native-lib -lwgpu-native"
```

### By `syscall`

On windows, we can call DLLs without enabling CGO, by using `syscall` std package.
In this case, make sure the `wgpu-native.dll` in the searching path is only thing we need to do.

TODO: make the DLL name configurable.
