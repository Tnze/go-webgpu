// Command compute runs a WGSL compute shader that doubles every u32 in a
// storage buffer: instance → adapter → device → dispatch → copy → map-read.
//
// Requires third_party deps (see scripts/fetch-deps.sh) and a working GPU backend.
package main

import (
	"fmt"
	"log"
	"runtime"
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu"
)

const (
	numElements  = 256
	workgroupX   = 64
	storageBytes = numElements * 4
)

// wgsl doubles each element of the bound storage buffer.
const wgsl = `
@group(0) @binding(0) var<storage, read_write> data: array<u32>;

@compute @workgroup_size(64)
fn main(@builtin(global_invocation_id) id: vec3<u32>) {
	if (id.x < arrayLength(&data)) {
		data[id.x] = data[id.x] * 2u;
	}
}
`

func main() {
	instance, err := webgpu.CreateInstance(webgpu.InstanceDescriptor{})
	if err != nil {
		log.Fatal(err)
	}
	defer instance.Release()

	adapter, err := instance.RequestAdapter(webgpu.RequestAdapterOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer adapter.Release()

	device, err := adapter.RequestDevice(webgpu.DeviceDescriptor{})
	if err != nil {
		log.Fatal(err)
	}
	defer device.Release()

	queue, err := device.GetQueue()
	if err != nil {
		log.Fatal(err)
	}
	defer queue.Release()

	// Storage buffer: written from the queue, read/written by the shader, copied out.
	storage, err := device.CreateBuffer(webgpu.BufferDescriptor{
		Label: "storage",
		Usage: webgpu.BufferUsageStorage | webgpu.BufferUsageCopySrc | webgpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Release()
	defer storage.Destroy()

	// MAP_READ pairs with COPY_DST.
	readback, err := device.CreateBuffer(webgpu.BufferDescriptor{
		Label: "readback",
		Usage: webgpu.BufferUsageMapRead | webgpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer readback.Release()
	defer readback.Destroy()

	input := make([]uint32, numElements)
	for i := range input {
		input[i] = uint32(i)
	}
	queue.WriteBuffer(storage, 0, unsafe.Pointer(unsafe.SliceData(input)), storageBytes)
	runtime.KeepAlive(input)

	// ShaderSourceWGSL is an extension struct: NewShaderSourceWGSL sets the
	// chain sType, and the struct is chained via ShaderModuleDescriptor.NextInChain.
	// The chained object (and its string payload) must stay pinned while C reads it.
	var pinner runtime.Pinner
	defer pinner.Unpin()
	src := webgpu.NewShaderSourceWGSL()
	src.Code = wgsl
	pinner.Pin(&src)
	pinner.Pin(unsafe.StringData(src.Code))
	module, err := device.CreateShaderModule(webgpu.ShaderModuleDescriptor{
		Label:       "compute",
		NextInChain: unsafe.Pointer(&src),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer module.Release()

	// Layout 0 = auto: the pipeline derives its bind group layout from the shader.
	pipeline, err := device.CreateComputePipeline(webgpu.ComputePipelineDescriptor{
		Label: "compute",
		Compute: webgpu.ComputeState{
			Module:     module.Handle(),
			EntryPoint: "main",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer pipeline.Release()

	bgLayout, err := pipeline.GetBindGroupLayout(0)
	if err != nil {
		log.Fatal(err)
	}
	defer bgLayout.Release()

	entries := []webgpu.BindGroupEntry{{
		Binding: 0,
		Buffer:  storage.Handle(),
		Size:    webgpu.WholeSize,
	}}
	bindGroup, err := device.CreateBindGroup(webgpu.BindGroupDescriptor{
		Label:        "bind-group",
		Layout:       bgLayout.Handle(),
		EntriesCount: 1,
		Entries:      &entries[0],
	})
	runtime.KeepAlive(entries)
	if err != nil {
		log.Fatal(err)
	}
	defer bindGroup.Release()

	enc, err := device.CreateCommandEncoder(webgpu.CommandEncoderDescriptor{})
	if err != nil {
		log.Fatal(err)
	}
	defer enc.Release()

	pass, err := enc.BeginComputePass(webgpu.ComputePassDescriptor{Label: "compute"})
	if err != nil {
		log.Fatal(err)
	}
	pass.SetPipeline(pipeline)
	pass.SetBindGroup(0, bindGroup, nil)
	pass.DispatchWorkgroups((numElements+workgroupX-1)/workgroupX, 1, 1)
	pass.End()
	pass.Release()

	enc.CopyBufferToBuffer(storage, 0, readback, 0, storageBytes)
	cmd, err := enc.Finish(webgpu.CommandBufferDescriptor{})
	if err != nil {
		log.Fatal(err)
	}
	defer cmd.Release()

	queue.Submit([]webgpu.CommandBuffer{cmd})

	if err := readback.Map(webgpu.MapModeRead, 0, storageBytes); err != nil {
		log.Fatal(err)
	}
	ptr, err := readback.GetConstMappedRange(0, storageBytes)
	if err != nil {
		log.Fatal(err)
	}
	got := unsafe.Slice((*uint32)(ptr), numElements)
	for i := range input {
		want := input[i] * 2
		if got[i] != want {
			readback.Unmap()
			log.Fatalf("mismatch at %d: got %d want %d", i, got[i], want)
		}
	}
	head := [4]uint32{}
	copy(head[:], got)
	readback.Unmap()

	fmt.Printf("compute ok (elements=%d, workgroup=%d, got[0..3]=%v)\n",
		numElements, workgroupX, head)
}
