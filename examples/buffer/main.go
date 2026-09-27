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

	"github.com/Tnze/go-webgpu/gpu"
)

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

	const size = 256

	// MAP_WRITE pairs with COPY_SRC; MAP_READ pairs with COPY_DST.
	src, err := device.CreateBuffer(&gpu.BufferDescriptor{
		Label: "src",
		Usage: gpu.BufferUsageMapWrite | gpu.BufferUsageCopySrc,
		Size:  size,
	})
	if err != nil {
		log.Fatal("create src buffer failed: ", err)
	}
	defer src.Release()
	defer src.Destroy()

	dst, err := device.CreateBuffer(&gpu.BufferDescriptor{
		Label: "dst",
		Usage: gpu.BufferUsageMapRead | gpu.BufferUsageCopyDst,
		Size:  size,
	})
	if err != nil {
		log.Fatal("create dst buffer failed: ", err)
	}
	defer dst.Release()
	defer dst.Destroy()

	payload := make([]byte, size)
	for i := range payload {
		payload[i] = byte(i)
	}
	binary.LittleEndian.PutUint32(payload[0:], 0xC0FFEE00)

	// Map-write the source buffer and fill it.
	if status, msg := src.Map(gpu.MapModeWrite, 0, size); status != gpu.MapAsyncStatusSuccess {
		log.Fatalf("map src: status %d: %s", status, msg)
	}
	wp := src.GetMappedRange(0, size)
	if wp == nil {
		log.Fatal("get src mapped range failed")
	}
	copy(unsafe.Slice((*byte)(wp), size), payload)
	src.Unmap()

	// Copy src → dst, then map-read dst.
	enc, err := device.CreateCommandEncoder(&gpu.CommandEncoderDescriptor{})
	if err != nil {
		log.Fatal("create command encoder failed: ", err)
	}
	defer enc.Release()

	enc.CopyBufferToBuffer(src, 0, dst, 0, size)
	cmd, err := enc.Finish(&gpu.CommandBufferDescriptor{})
	if err != nil {
		log.Fatal("finish command buffer failed: ", err)
	}
	defer cmd.Release()

	queue.Submit([]*gpu.CommandBuffer{cmd})

	if status, msg := dst.Map(gpu.MapModeRead, 0, size); status != gpu.MapAsyncStatusSuccess {
		log.Fatalf("map dst: status %d: %s", status, msg)
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
