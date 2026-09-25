// Command buffer demonstrates a storage-buffer round trip through the
// generated WebGPU bindings: instance → adapter → device → write → copy → map-read.
//
// Requires third_party deps (see scripts/fetch-deps.sh) and a working GPU backend.
package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu"
)

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
	const size = 256

	// MAP_WRITE pairs with COPY_SRC; MAP_READ pairs with COPY_DST.
	src := webgpu.DeviceCreateBuffer(device, webgpu.BufferDescriptor{
		Label: "src",
		Usage: webgpu.BufferUsageMapWrite | webgpu.BufferUsageCopySrc,
		Size:  size,
	})
	dst := webgpu.DeviceCreateBuffer(device, webgpu.BufferDescriptor{
		Label: "dst",
		Usage: webgpu.BufferUsageMapRead | webgpu.BufferUsageCopyDst,
		Size:  size,
	})
	if src == nil || dst == nil {
		fmt.Fprintln(os.Stderr, "CreateBuffer failed")
		os.Exit(1)
	}

	payload := make([]byte, size)
	for i := range payload {
		payload[i] = byte(i)
	}
	binary.LittleEndian.PutUint32(payload[0:], 0xC0FFEE00)

	// Map-write the source buffer and fill it.
	var (
		wStatus webgpu.MapAsyncStatus
		wDone   bool
	)
	webgpu.BufferMapAsync(src, webgpu.MapModeWrite, 0, size,
		func(status webgpu.MapAsyncStatus, message string) {
			wStatus = status
			wDone = true
		})
	wait(instance, &wDone)
	if wStatus != webgpu.MapAsyncStatusSuccess {
		fmt.Fprintln(os.Stderr, "MapAsync(write) failed:", wStatus)
		os.Exit(1)
	}
	wp := webgpu.BufferGetMappedRange(src, 0, size)
	if wp == nil {
		fmt.Fprintln(os.Stderr, "GetMappedRange(write) failed")
		os.Exit(1)
	}
	copy(unsafe.Slice((*byte)(wp), size), payload)
	webgpu.BufferUnmap(src)

	// Copy src → dst, then map-read dst.
	enc := webgpu.DeviceCreateCommandEncoder(device, webgpu.CommandEncoderDescriptor{})
	webgpu.CommandEncoderCopyBufferToBuffer(enc, src, 0, dst, 0, size)
	cmd := webgpu.CommandEncoderFinish(enc, webgpu.CommandBufferDescriptor{})
	webgpu.QueueSubmit(queue, []*webgpu.CommandBuffer{cmd})

	var (
		mapStatus webgpu.MapAsyncStatus
		mapMsg    string
		mapDone   bool
	)
	webgpu.BufferMapAsync(dst, webgpu.MapModeRead, 0, size,
		func(status webgpu.MapAsyncStatus, message string) {
			mapStatus, mapMsg = status, message
			mapDone = true
		})
	wait(instance, &mapDone)
	if mapStatus != webgpu.MapAsyncStatusSuccess {
		fmt.Fprintln(os.Stderr, "MapAsync failed:", mapStatus, mapMsg)
		os.Exit(1)
	}

	ptr := webgpu.BufferGetConstMappedRange(dst, 0, size)
	if ptr == nil {
		fmt.Fprintln(os.Stderr, "GetConstMappedRange failed")
		os.Exit(1)
	}
	got := unsafe.Slice((*byte)(ptr), size)
	magic := binary.LittleEndian.Uint32(got[0:])
	ok := true
	for i := range payload {
		if got[i] != payload[i] {
			fmt.Fprintf(os.Stderr, "mismatch at %d: got %#x want %#x\n", i, got[i], payload[i])
			fmt.Fprintf(os.Stderr, "got[0:16]  = % x\n", got[:16])
			fmt.Fprintf(os.Stderr, "want[0:16] = % x\n", payload[:16])
			ok = false
			break
		}
	}
	webgpu.BufferUnmap(dst)
	webgpu.BufferDestroy(src)
	webgpu.BufferDestroy(dst)

	if !ok {
		os.Exit(1)
	}
	fmt.Printf("buffer roundtrip ok (size=%d, magic=%#x)\n", size, magic)
}
