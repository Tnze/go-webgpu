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

	const size = 256

	// MAP_WRITE pairs with COPY_SRC; MAP_READ pairs with COPY_DST.
	src := webgpu.DeviceCreateBuffer(device, &webgpu.BufferDescriptor{
		Label: "src",
		Usage: webgpu.BufferUsageMapWrite | webgpu.BufferUsageCopySrc,
		Size:  size,
	})
	if src == nil {
		log.Fatal("create src buffer failed")
	}
	defer webgpu.BufferRelease(src)
	defer webgpu.BufferDestroy(src)

	dst := webgpu.DeviceCreateBuffer(device, &webgpu.BufferDescriptor{
		Label: "dst",
		Usage: webgpu.BufferUsageMapRead | webgpu.BufferUsageCopyDst,
		Size:  size,
	})
	if dst == nil {
		log.Fatal("create dst buffer failed")
	}
	defer webgpu.BufferRelease(dst)
	defer webgpu.BufferDestroy(dst)

	payload := make([]byte, size)
	for i := range payload {
		payload[i] = byte(i)
	}
	binary.LittleEndian.PutUint32(payload[0:], 0xC0FFEE00)

	// Map-write the source buffer and fill it.
	if status, msg := webgpu.BufferMap(src, webgpu.MapModeWrite, 0, size); status != webgpu.MapAsyncStatusSuccess {
		log.Fatalf("map src: status %d: %s", status, msg)
	}
	wp := webgpu.BufferGetMappedRange(src, 0, size)
	if wp == nil {
		log.Fatal("get src mapped range failed")
	}
	copy(unsafe.Slice((*byte)(wp), size), payload)
	webgpu.BufferUnmap(src)

	// Copy src → dst, then map-read dst.
	enc := webgpu.DeviceCreateCommandEncoder(device, &webgpu.CommandEncoderDescriptor{})
	if enc == nil {
		log.Fatal("create command encoder failed")
	}
	defer webgpu.CommandEncoderRelease(enc)

	webgpu.CommandEncoderCopyBufferToBuffer(enc, src, 0, dst, 0, size)
	cmd := webgpu.CommandEncoderFinish(enc, &webgpu.CommandBufferDescriptor{})
	if cmd == nil {
		log.Fatal("finish command buffer failed")
	}
	defer webgpu.CommandBufferRelease(cmd)

	webgpu.QueueSubmit(queue, []webgpu.CommandBuffer{cmd})

	if status, msg := webgpu.BufferMap(dst, webgpu.MapModeRead, 0, size); status != webgpu.MapAsyncStatusSuccess {
		log.Fatalf("map dst: status %d: %s", status, msg)
	}
	ptr := webgpu.BufferGetConstMappedRange(dst, 0, size)
	if ptr == nil {
		log.Fatal("get dst mapped range failed")
	}
	got := unsafe.Slice((*byte)(ptr), size)
	magic := binary.LittleEndian.Uint32(got[0:])
	for i := range payload {
		if got[i] != payload[i] {
			webgpu.BufferUnmap(dst)
			log.Fatalf("mismatch at %d: got %#x want %#x", i, got[i], payload[i])
		}
	}
	webgpu.BufferUnmap(dst)

	fmt.Printf("buffer roundtrip ok (size=%d, magic=%#x)\n", size, magic)
}
