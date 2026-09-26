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

	const size = 256

	// MAP_WRITE pairs with COPY_SRC; MAP_READ pairs with COPY_DST.
	src, err := device.CreateBuffer(webgpu.BufferDescriptor{
		Label: "src",
		Usage: webgpu.BufferUsageMapWrite | webgpu.BufferUsageCopySrc,
		Size:  size,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer src.Release()
	defer src.Destroy()

	dst, err := device.CreateBuffer(webgpu.BufferDescriptor{
		Label: "dst",
		Usage: webgpu.BufferUsageMapRead | webgpu.BufferUsageCopyDst,
		Size:  size,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Release()
	defer dst.Destroy()

	payload := make([]byte, size)
	for i := range payload {
		payload[i] = byte(i)
	}
	binary.LittleEndian.PutUint32(payload[0:], 0xC0FFEE00)

	// Map-write the source buffer and fill it.
	if err := src.Map(webgpu.MapModeWrite, 0, size); err != nil {
		log.Fatal(err)
	}
	wp, err := src.GetMappedRange(0, size)
	if err != nil {
		log.Fatal(err)
	}
	copy(unsafe.Slice((*byte)(wp), size), payload)
	src.Unmap()

	// Copy src → dst, then map-read dst.
	enc, err := device.CreateCommandEncoder(webgpu.CommandEncoderDescriptor{})
	if err != nil {
		log.Fatal(err)
	}
	defer enc.Release()

	enc.CopyBufferToBuffer(src, 0, dst, 0, size)
	cmd, err := enc.Finish(webgpu.CommandBufferDescriptor{})
	if err != nil {
		log.Fatal(err)
	}
	defer cmd.Release()

	queue.Submit([]webgpu.CommandBuffer{cmd})

	if err := dst.Map(webgpu.MapModeRead, 0, size); err != nil {
		log.Fatal(err)
	}
	ptr, err := dst.GetConstMappedRange(0, size)
	if err != nil {
		log.Fatal(err)
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
