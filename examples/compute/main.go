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
	instance := webgpu.CreateInstance(&webgpu.InstanceDescriptor{})
	if instance == nil {
		log.Fatal("create instance failed")
	}
	defer webgpu.InstanceRelease(instance)

	adapter, adapterStatus, adapterMsg := webgpu.InstanceRequestAdapter(instance, &webgpu.RequestAdapterOptions{})
	if adapterStatus != webgpu.RequestAdapterStatusSuccess {
		log.Fatalf("request adapter: status %d: %s", adapterStatus, adapterMsg)
	}
	defer webgpu.AdapterRelease(adapter)

	device, deviceStatus, deviceMsg := webgpu.AdapterRequestDevice(adapter, &webgpu.DeviceDescriptor{})
	if deviceStatus != webgpu.RequestDeviceStatusSuccess {
		log.Fatalf("request device: status %d: %s", deviceStatus, deviceMsg)
	}
	defer webgpu.DeviceRelease(device)

	queue := webgpu.DeviceGetQueue(device)
	if queue == nil {
		log.Fatal("get queue failed")
	}
	defer webgpu.QueueRelease(queue)

	// Storage buffer: written from the queue, read/written by the shader, copied out.
	storage := webgpu.DeviceCreateBuffer(device, &webgpu.BufferDescriptor{
		Label: "storage",
		Usage: webgpu.BufferUsageStorage | webgpu.BufferUsageCopySrc | webgpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	if storage == nil {
		log.Fatal("create storage buffer failed")
	}
	defer webgpu.BufferRelease(storage)
	defer webgpu.BufferDestroy(storage)

	// MAP_READ pairs with COPY_DST.
	readback := webgpu.DeviceCreateBuffer(device, &webgpu.BufferDescriptor{
		Label: "readback",
		Usage: webgpu.BufferUsageMapRead | webgpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	if readback == nil {
		log.Fatal("create readback buffer failed")
	}
	defer webgpu.BufferRelease(readback)
	defer webgpu.BufferDestroy(readback)

	input := make([]uint32, numElements)
	for i := range input {
		input[i] = uint32(i)
	}
	webgpu.QueueWriteBuffer(queue, storage, 0, unsafe.Pointer(unsafe.SliceData(input)), storageBytes)
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
	module := webgpu.DeviceCreateShaderModule(device, &webgpu.ShaderModuleDescriptor{
		Label:       "compute",
		NextInChain: unsafe.Pointer(&src),
	})
	if module == nil {
		log.Fatal("create shader module failed")
	}
	defer webgpu.ShaderModuleRelease(module)

	// Layout 0 = auto: the pipeline derives its bind group layout from the shader.
	pipeline := webgpu.DeviceCreateComputePipeline(device, &webgpu.ComputePipelineDescriptor{
		Label: "compute",
		Compute: webgpu.ComputeState{
			Module:     module,
			EntryPoint: "main",
		},
	})
	if pipeline == nil {
		log.Fatal("create compute pipeline failed")
	}
	defer webgpu.ComputePipelineRelease(pipeline)

	bgLayout := webgpu.ComputePipelineGetBindGroupLayout(pipeline, 0)
	if bgLayout == nil {
		log.Fatal("get bind group layout failed")
	}
	defer webgpu.BindGroupLayoutRelease(bgLayout)

	entries := []webgpu.BindGroupEntry{{
		Binding: 0,
		Buffer:  storage,
		Size:    webgpu.WholeSize,
	}}
	bindGroup := webgpu.DeviceCreateBindGroup(device, &webgpu.BindGroupDescriptor{
		Label:        "bind-group",
		Layout:       bgLayout,
		EntriesCount: 1,
		Entries:      &entries[0],
	})
	runtime.KeepAlive(entries)
	if bindGroup == nil {
		log.Fatal("create bind group failed")
	}
	defer webgpu.BindGroupRelease(bindGroup)

	enc := webgpu.DeviceCreateCommandEncoder(device, &webgpu.CommandEncoderDescriptor{})
	if enc == nil {
		log.Fatal("create command encoder failed")
	}
	defer webgpu.CommandEncoderRelease(enc)

	pass := webgpu.CommandEncoderBeginComputePass(enc, &webgpu.ComputePassDescriptor{Label: "compute"})
	if pass == nil {
		log.Fatal("begin compute pass failed")
	}
	webgpu.ComputePassEncoderSetPipeline(pass, pipeline)
	webgpu.ComputePassEncoderSetBindGroup(pass, 0, bindGroup, nil)
	webgpu.ComputePassEncoderDispatchWorkgroups(pass, (numElements+workgroupX-1)/workgroupX, 1, 1)
	webgpu.ComputePassEncoderEnd(pass)
	webgpu.ComputePassEncoderRelease(pass)

	webgpu.CommandEncoderCopyBufferToBuffer(enc, storage, 0, readback, 0, storageBytes)
	cmd := webgpu.CommandEncoderFinish(enc, &webgpu.CommandBufferDescriptor{})
	if cmd == nil {
		log.Fatal("finish command buffer failed")
	}
	defer webgpu.CommandBufferRelease(cmd)

	webgpu.QueueSubmit(queue, []webgpu.CommandBuffer{cmd})

	if status, msg := webgpu.BufferMap(readback, webgpu.MapModeRead, 0, storageBytes); status != webgpu.MapAsyncStatusSuccess {
		log.Fatalf("map readback: status %d: %s", status, msg)
	}
	ptr := webgpu.BufferGetConstMappedRange(readback, 0, storageBytes)
	if ptr == nil {
		log.Fatal("get readback mapped range failed")
	}
	got := unsafe.Slice((*uint32)(ptr), numElements)
	for i := range input {
		want := input[i] * 2
		if got[i] != want {
			webgpu.BufferUnmap(readback)
			log.Fatalf("mismatch at %d: got %d want %d", i, got[i], want)
		}
	}
	head := [4]uint32{}
	copy(head[:], got)
	webgpu.BufferUnmap(readback)

	fmt.Printf("compute ok (elements=%d, workgroup=%d, got[0..3]=%v)\n",
		numElements, workgroupX, head)
}
