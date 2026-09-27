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

	"github.com/Tnze/go-webgpu/gpu"
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
	instance, err := gpu.CreateInstance(nil)
	if err != nil {
		log.Fatal("create instance failed: ", err)
	}
	defer instance.Release()

	adapter, adapterStatus, adapterMsg := instance.RequestAdapter(&gpu.RequestAdapterOptions{})
	if adapterStatus != gpu.RequestAdapterStatusSuccess {
		log.Fatalf("request adapter: status %d: %s", adapterStatus, adapterMsg)
	}
	defer adapter.Release()

	device, deviceStatus, deviceMsg := adapter.RequestDevice(&gpu.DeviceDescriptor{})
	if deviceStatus != gpu.RequestDeviceStatusSuccess {
		log.Fatalf("request device: status %d: %s", deviceStatus, deviceMsg)
	}
	defer device.Release()

	queue := device.GetQueue()
	if queue == nil {
		log.Fatal("get queue failed")
	}
	defer queue.Release()

	// Storage buffer: written from the queue, read/written by the shader, copied out.
	storage, err := device.CreateBuffer(&gpu.BufferDescriptor{
		Label: "storage",
		Usage: gpu.BufferUsageStorage | gpu.BufferUsageCopySrc | gpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	if err != nil {
		log.Fatal("create storage buffer failed: ", err)
	}
	defer storage.Release()
	defer storage.Destroy()

	// MAP_READ pairs with COPY_DST.
	readback, err := device.CreateBuffer(&gpu.BufferDescriptor{
		Label: "readback",
		Usage: gpu.BufferUsageMapRead | gpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	if err != nil {
		log.Fatal("create readback buffer failed: ", err)
	}
	defer readback.Release()
	defer readback.Destroy()

	input := make([]uint32, numElements)
	for i := range input {
		input[i] = uint32(i)
	}
	data := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(input))), storageBytes)
	queue.WriteBuffer(storage, 0, data)
	runtime.KeepAlive(input)

	module, err := device.CreateShaderModule(&gpu.ShaderModuleDescriptor{
		Label: "compute",
		WGSL:  &gpu.ShaderSourceWGSL{Code: wgsl},
	})
	if err != nil {
		log.Fatal("create shader module failed: ", err)
	}
	defer module.Release()

	pipeline, err := device.CreateComputePipeline(&gpu.ComputePipelineDescriptor{
		Label: "compute",
		Compute: gpu.ComputeState{
			Module:     module,
			EntryPoint: "main",
		},
	})
	if err != nil {
		log.Fatal("create compute pipeline failed: ", err)
	}
	defer pipeline.Release()

	bgLayout := pipeline.GetBindGroupLayout(0)
	if bgLayout == nil {
		log.Fatal("get bind group layout failed")
	}
	defer bgLayout.Release()

	bindGroup, err := device.CreateBindGroup(&gpu.BindGroupDescriptor{
		Label:  "bind-group",
		Layout: bgLayout,
		Entries: []gpu.BindGroupEntry{{
			Binding: 0,
			Buffer:  storage,
			Size:    gpu.WholeSize,
		}},
	})
	if err != nil {
		log.Fatal("create bind group failed: ", err)
	}
	defer bindGroup.Release()

	enc, err := device.CreateCommandEncoder(&gpu.CommandEncoderDescriptor{})
	if err != nil {
		log.Fatal("create command encoder failed: ", err)
	}
	defer enc.Release()

	pass := enc.BeginComputePass(&gpu.ComputePassDescriptor{Label: "compute"})
	if pass == nil {
		log.Fatal("begin compute pass failed")
	}
	pass.SetPipeline(pipeline)
	pass.SetBindGroup(0, bindGroup, nil)
	pass.DispatchWorkgroups((numElements+workgroupX-1)/workgroupX, 1, 1)
	pass.End()
	pass.Release()

	enc.CopyBufferToBuffer(storage, 0, readback, 0, storageBytes)
	cmd, err := enc.Finish(&gpu.CommandBufferDescriptor{})
	if err != nil {
		log.Fatal("finish command buffer failed: ", err)
	}
	defer cmd.Release()

	queue.Submit([]*gpu.CommandBuffer{cmd})

	if status, msg := readback.Map(gpu.MapModeRead, 0, storageBytes); status != gpu.MapAsyncStatusSuccess {
		log.Fatalf("map readback: status %d: %s", status, msg)
	}
	ptr := readback.GetConstMappedRange(0, storageBytes)
	if ptr == nil {
		log.Fatal("get readback mapped range failed")
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
