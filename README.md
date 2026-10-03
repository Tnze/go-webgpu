# go-webgpu

A WebGPU binding for Go programming language. Based on codegen.

## Re-generate codes

```
go run ./go-webgpu-gen \
    -src webgpu-headers/webgpu.yml \
    -out webgpu \
    -importPath "github.com/Tnze/go-webgpu/webgpu"
```
