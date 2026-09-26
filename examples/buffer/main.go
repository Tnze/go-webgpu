// Command buffer demonstrates a storage-buffer round trip through the
// generated WebGPU bindings: instance → adapter → device → write → copy → map-read.
//
// Requires third_party deps (see scripts/fetch-deps.sh) and a working GPU backend.
package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu"
)

func main() {
	instance := webgpu.CreateInstance(webgpu.InstanceDescriptor{})
	if instance == 0 {
		log.Fatal("create instance failed")
	}
	defer instance.Release()

	adapter, adapterStatus := instance.RequestAdapter(webgpu.RequestAdapterOptions{})
	if adapterStatus != webgpu.RequestAdapterStatusSuccess {
		log.Fatalf("request adapter: status %d", adapterStatus)
	}
	defer adapter.Release()

	device, deviceStatus := adapter.RequestDevice(webgpu.DeviceDescriptor{})
	if deviceStatus != webgpu.RequestDeviceStatusSuccess {
		log.Fatalf("request device: status %d", deviceStatus)
	}
	defer device.Release()

	queue := device.GetQueue()
	if queue == 0 {
		log.Fatal("get queue failed")
	}
	defer queue.Release()

	const size = 256

	// MAP_WRITE pairs with COPY_SRC; MAP_READ pairs with COPY_DST.
	src := device.CreateBuffer(webgpu.BufferDescriptor{
		Label: "src",
		Usage: webgpu.BufferUsageMapWrite | webgpu.BufferUsageCopySrc,
		Size:  size,
	})
	if src == 0 {
		log.Fatal("create src buffer failed")
	}
	defer src.Release()
	defer src.Destroy()

	dst := device.CreateBuffer(webgpu.BufferDescriptor{
		Label: "dst",
		Usage: webgpu.BufferUsageMapRead | webgpu.BufferUsageCopyDst,
		Size:  size,
	})
	if dst == 0 {
		log.Fatal("create dst buffer failed")
	}
	defer dst.Release()
	defer dst.Destroy()

	payload := make([]byte, size)
	for i := range payload {
		payload[i] = byte(i)
	}
	binary.LittleEndian.PutUint32(payload[0:], 0xC0FFEE00)

	// Map-write the source buffer and fill it.
	if status := src.Map(webgpu.MapModeWrite, 0, size); status != webgpu.MapAsyncStatusSuccess {
		log.Fatalf("map src: status %d", status)
	}
	wp := src.GetMappedRange(0, size)
	if wp == nil {
		log.Fatal("get src mapped range failed")
	}
	copy(unsafe.Slice((*byte)(wp), size), payload)
	src.Unmap()

	// Copy src → dst, then map-read dst.
	enc := device.CreateCommandEncoder(webgpu.CommandEncoderDescriptor{})
	if enc == 0 {
		log.Fatal("create command encoder failed")
	}
	defer enc.Release()

	enc.CopyBufferToBuffer(src, 0, dst, 0, size)
	cmd := enc.Finish(webgpu.CommandBufferDescriptor{})
	if cmd == 0 {
		log.Fatal("finish command buffer failed")
	}
	defer cmd.Release()

	queue.Submit([]webgpu.CommandBuffer{cmd})

	if status := dst.Map(webgpu.MapModeRead, 0, size); status != webgpu.MapAsyncStatusSuccess {
		log.Fatalf("map dst: status %d", status)
	}
	ptr := dst.GetConstMappedRange(0, size)
	if ptr == nil {
		log.Fatal("get dst mapped range failed")
	}
	got := unsafe.Slice((*byte)(ptr), size)
	magic := binary.LittleEndian.Uint32(got[0:])
	for i := range payload {
		if got[i] != payload[i] {
			dst.Unmap()
			log.Fatalf("mismatch at %d: got %#x want %#x", i, got[i], payload[i])
		}
	}
	dst.Unmap()

	fmt.Printf("buffer roundtrip ok (size=%d, magic=%#x)\n", size, magic)
}
