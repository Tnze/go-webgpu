// Command compute runs a WGSL compute shader that doubles every u32 in a
// storage buffer: instance → adapter → device → dispatch → copy → map-read.
//
// Requires third_party deps (see scripts/fetch-deps.sh) and a working GPU backend.
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
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

// wait polls process-events until done is set by the async callback.
// (wgpu-native's wgpuInstanceWaitAny is not implemented for all builds.)
func wait(instance *webgpu.Instance, done *bool) {
	deadline := time.Now().Add(10 * time.Second)
	for !*done {
		if time.Now().After(deadline) {
			fmt.Fprintln(os.Stderr, "timeout waiting for async callback")
			os.Exit(1)
		}
		webgpu.InstanceProcessEvents(instance)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	instance := webgpu.CreateInstance(webgpu.InstanceDescriptor{})
	if instance == nil {
		fmt.Fprintln(os.Stderr, "CreateInstance failed")
		os.Exit(1)
	}

	var (
		adapter *webgpu.Adapter
		aErr    string
		aDone   bool
	)
	webgpu.InstanceRequestAdapter(instance, webgpu.RequestAdapterOptions{},
		func(status webgpu.RequestAdapterStatus, a *webgpu.Adapter, message string) {
			adapter, aErr = a, message
			if status != webgpu.RequestAdapterStatusSuccess {
				aErr = fmt.Sprintf("request adapter status=%v %s", status, message)
			}
			aDone = true
		})
	wait(instance, &aDone)
	if adapter == nil {
		fmt.Fprintln(os.Stderr, "RequestAdapter failed:", aErr)
		os.Exit(1)
	}

	var (
		device *webgpu.Device
		dErr   string
		dDone  bool
	)
	webgpu.AdapterRequestDevice(adapter, webgpu.DeviceDescriptor{},
		func(status webgpu.RequestDeviceStatus, d *webgpu.Device, message string) {
			device, dErr = d, message
			if status != webgpu.RequestDeviceStatusSuccess {
				dErr = fmt.Sprintf("request device status=%v %s", status, message)
			}
			dDone = true
		})
	wait(instance, &dDone)
	if device == nil {
		fmt.Fprintln(os.Stderr, "RequestDevice failed:", dErr)
		os.Exit(1)
	}

	queue := webgpu.DeviceGetQueue(device)

	// Storage buffer: written from the queue, read/written by the shader, copied out.
	storage := webgpu.DeviceCreateBuffer(device, webgpu.BufferDescriptor{
		Label: "storage",
		Usage: webgpu.BufferUsageStorage | webgpu.BufferUsageCopySrc | webgpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	// MAP_READ pairs with COPY_DST.
	readback := webgpu.DeviceCreateBuffer(device, webgpu.BufferDescriptor{
		Label: "readback",
		Usage: webgpu.BufferUsageMapRead | webgpu.BufferUsageCopyDst,
		Size:  storageBytes,
	})
	if storage == nil || readback == nil {
		fmt.Fprintln(os.Stderr, "CreateBuffer failed")
		os.Exit(1)
	}

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
	module := webgpu.DeviceCreateShaderModule(device, webgpu.ShaderModuleDescriptor{
		Label:       "compute",
		NextInChain: unsafe.Pointer(&src),
	})
	if module == nil {
		fmt.Fprintln(os.Stderr, "CreateShaderModule failed")
		os.Exit(1)
	}

	// Layout 0 = auto: the pipeline derives its bind group layout from the shader.
	pipeline := webgpu.DeviceCreateComputePipeline(device, webgpu.ComputePipelineDescriptor{
		Label: "compute",
		Compute: webgpu.ComputeState{
			Module:     module.Handle(),
			EntryPoint: "main",
		},
	})
	if pipeline == nil {
		fmt.Fprintln(os.Stderr, "CreateComputePipeline failed")
		os.Exit(1)
	}

	bgLayout := webgpu.ComputePipelineGetBindGroupLayout(pipeline, 0)
	if bgLayout == nil {
		fmt.Fprintln(os.Stderr, "GetBindGroupLayout failed")
		os.Exit(1)
	}
	entries := []webgpu.BindGroupEntry{{
		Binding: 0,
		Buffer:  storage.Handle(),
		Size:    webgpu.WholeSize,
	}}
	bindGroup := webgpu.DeviceCreateBindGroup(device, webgpu.BindGroupDescriptor{
		Label:        "bind-group",
		Layout:       bgLayout.Handle(),
		EntriesCount: 1,
		Entries:      &entries[0],
	})
	runtime.KeepAlive(entries)
	if bindGroup == nil {
		fmt.Fprintln(os.Stderr, "CreateBindGroup failed")
		os.Exit(1)
	}

	enc := webgpu.DeviceCreateCommandEncoder(device, webgpu.CommandEncoderDescriptor{})
	pass := webgpu.CommandEncoderBeginComputePass(enc, webgpu.ComputePassDescriptor{Label: "compute"})
	webgpu.ComputePassEncoderSetPipeline(pass, pipeline)
	webgpu.ComputePassEncoderSetBindGroup(pass, 0, bindGroup, nil)
	webgpu.ComputePassEncoderDispatchWorkgroups(pass, (numElements+workgroupX-1)/workgroupX, 1, 1)
	webgpu.ComputePassEncoderEnd(pass)
	webgpu.CommandEncoderCopyBufferToBuffer(enc, storage, 0, readback, 0, storageBytes)
	cmd := webgpu.CommandEncoderFinish(enc, webgpu.CommandBufferDescriptor{})
	webgpu.QueueSubmit(queue, []*webgpu.CommandBuffer{cmd})

	var (
		mapStatus webgpu.MapAsyncStatus
		mapMsg    string
		mapDone   bool
	)
	webgpu.BufferMapAsync(readback, webgpu.MapModeRead, 0, storageBytes,
		func(status webgpu.MapAsyncStatus, message string) {
			mapStatus, mapMsg = status, message
			mapDone = true
		})
	wait(instance, &mapDone)
	if mapStatus != webgpu.MapAsyncStatusSuccess {
		fmt.Fprintln(os.Stderr, "MapAsync failed:", mapStatus, mapMsg)
		os.Exit(1)
	}

	ptr := webgpu.BufferGetConstMappedRange(readback, 0, storageBytes)
	if ptr == nil {
		fmt.Fprintln(os.Stderr, "GetConstMappedRange failed")
		os.Exit(1)
	}
	got := unsafe.Slice((*uint32)(ptr), numElements)
	ok := true
	for i := range input {
		want := input[i] * 2
		if got[i] != want {
			fmt.Fprintf(os.Stderr, "mismatch at %d: got %d want %d\n", i, got[i], want)
			ok = false
			break
		}
	}
	head := [4]uint32{}
	copy(head[:], got)
	webgpu.BufferUnmap(readback)
	webgpu.BufferDestroy(storage)
	webgpu.BufferDestroy(readback)

	if !ok {
		os.Exit(1)
	}
	fmt.Printf("compute ok (elements=%d, workgroup=%d, got[0..3]=%v)\n",
		numElements, workgroupX, head)
}
