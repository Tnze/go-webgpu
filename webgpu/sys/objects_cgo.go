// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo

package sys

// #cgo CFLAGS: -I${SRCDIR}/../../webgpu-headers
// #include <webgpu.h>
import "C"

import "unsafe"

func AdapterGetLimits(adapter Adapter, limits *Limits) Status {
	return Status(C.wgpuAdapterGetLimits(
		C.WGPUAdapter(adapter),
		(*C.WGPULimits)(unsafe.Pointer(limits)),
	))
}

func AdapterHasFeature(adapter Adapter, feature FeatureName) Bool {
	return Bool(C.wgpuAdapterHasFeature(
		C.WGPUAdapter(adapter),
		C.WGPUFeatureName(feature),
	))
}

// Get the list of @ref WGPUFeatureName values supported by the adapter.
func AdapterGetFeatures(adapter Adapter, features *SupportedFeatures) {
	C.wgpuAdapterGetFeatures(
		C.WGPUAdapter(adapter),
		(*C.WGPUSupportedFeatures)(unsafe.Pointer(features)),
	)
}

func AdapterGetInfo(adapter Adapter, info *AdapterInfo) Status {
	return Status(C.wgpuAdapterGetInfo(
		C.WGPUAdapter(adapter),
		(*C.WGPUAdapterInfo)(unsafe.Pointer(info)),
	))
}

func AdapterRequestDevice(adapter Adapter, descriptor *DeviceDescriptor, callback RequestDeviceCallbackInfo) {
	C.wgpuAdapterRequestDevice(
		C.WGPUAdapter(adapter),
		(*C.WGPUDeviceDescriptor)(unsafe.Pointer(descriptor)),
		C.WGPURequestDeviceCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPURequestDeviceCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

func AdapterAddRef(adapter Adapter) {
	C.wgpuAdapterAddRef(C.WGPUAdapter(adapter))
}

func AdapterRelease(adapter Adapter) {
	C.wgpuAdapterRelease(C.WGPUAdapter(adapter))
}

func BindGroupSetLabel(bindGroup BindGroup, label StringView) {
	C.wgpuBindGroupSetLabel(
		C.WGPUBindGroup(bindGroup),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func BindGroupAddRef(bindGroup BindGroup) {
	C.wgpuBindGroupAddRef(C.WGPUBindGroup(bindGroup))
}

func BindGroupRelease(bindGroup BindGroup) {
	C.wgpuBindGroupRelease(C.WGPUBindGroup(bindGroup))
}

func BindGroupLayoutSetLabel(bindGroupLayout BindGroupLayout, label StringView) {
	C.wgpuBindGroupLayoutSetLabel(
		C.WGPUBindGroupLayout(bindGroupLayout),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func BindGroupLayoutAddRef(bindGroupLayout BindGroupLayout) {
	C.wgpuBindGroupLayoutAddRef(C.WGPUBindGroupLayout(bindGroupLayout))
}

func BindGroupLayoutRelease(bindGroupLayout BindGroupLayout) {
	C.wgpuBindGroupLayoutRelease(C.WGPUBindGroupLayout(bindGroupLayout))
}

func BufferMapAsync(buffer Buffer, mode MapMode, offset uintptr, size uintptr, callback BufferMapCallbackInfo) {
	C.wgpuBufferMapAsync(
		C.WGPUBuffer(buffer),
		C.WGPUMapMode(mode),
		C.size_t(offset),
		C.size_t(size),
		C.WGPUBufferMapCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPUBufferMapCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

// Returns a mutable pointer to beginning of the mapped range.
// See @ref MappedRangeBehavior for error conditions and guarantees.
// This function is safe to call inside spontaneous callbacks (see @ref CallbackReentrancy).
//
// In Wasm, if `memcpy`ing into this range, prefer using @ref wgpuBufferWriteMappedRange
// instead for better performance.
func BufferGetMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.wgpuBufferGetMappedRange(
		C.WGPUBuffer(buffer),
		C.size_t(offset),
		C.size_t(size),
	))
}

// Returns a const pointer to beginning of the mapped range.
// It must not be written; writing to this range causes undefined behavior.
// See @ref MappedRangeBehavior for error conditions and guarantees.
// This function is safe to call inside spontaneous callbacks (see @ref CallbackReentrancy).
//
// In Wasm, if `memcpy`ing from this range, prefer using @ref wgpuBufferReadMappedRange
// instead for better performance.
func BufferGetConstMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.wgpuBufferGetConstMappedRange(
		C.WGPUBuffer(buffer),
		C.size_t(offset),
		C.size_t(size),
	))
}

// Copies a range of data from the buffer mapping into the provided destination pointer.
// See @ref MappedRangeBehavior for error conditions and guarantees.
// This function is safe to call inside spontaneous callbacks (see @ref CallbackReentrancy).
//
// In Wasm, this is more efficient than copying from a mapped range into a `malloc`'d range.
func BufferReadMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {
	return Status(C.wgpuBufferReadMappedRange(
		C.WGPUBuffer(buffer),
		C.size_t(offset),
		unsafe.Pointer(data),
		C.size_t(size),
	))
}

// Copies a range of data from the provided source pointer into the buffer mapping.
// See @ref MappedRangeBehavior for error conditions and guarantees.
// This function is safe to call inside spontaneous callbacks (see @ref CallbackReentrancy).
//
// In Wasm, this is more efficient than copying from a `malloc`'d range into a mapped range.
func BufferWriteMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {
	return Status(C.wgpuBufferWriteMappedRange(
		C.WGPUBuffer(buffer),
		C.size_t(offset),
		unsafe.Pointer(data),
		C.size_t(size),
	))
}

func BufferSetLabel(buffer Buffer, label StringView) {
	C.wgpuBufferSetLabel(
		C.WGPUBuffer(buffer),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func BufferGetUsage(buffer Buffer) BufferUsage {
	return BufferUsage(C.wgpuBufferGetUsage(
		C.WGPUBuffer(buffer),
	))
}

func BufferGetSize(buffer Buffer) uint64 {
	return uint64(C.wgpuBufferGetSize(
		C.WGPUBuffer(buffer),
	))
}

func BufferGetMapState(buffer Buffer) BufferMapState {
	return BufferMapState(C.wgpuBufferGetMapState(
		C.WGPUBuffer(buffer),
	))
}

func BufferUnmap(buffer Buffer) {
	C.wgpuBufferUnmap(
		C.WGPUBuffer(buffer),
	)
}

func BufferDestroy(buffer Buffer) {
	C.wgpuBufferDestroy(
		C.WGPUBuffer(buffer),
	)
}

func BufferAddRef(buffer Buffer) {
	C.wgpuBufferAddRef(C.WGPUBuffer(buffer))
}

func BufferRelease(buffer Buffer) {
	C.wgpuBufferRelease(C.WGPUBuffer(buffer))
}

func CommandBufferSetLabel(commandBuffer CommandBuffer, label StringView) {
	C.wgpuCommandBufferSetLabel(
		C.WGPUCommandBuffer(commandBuffer),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func CommandBufferAddRef(commandBuffer CommandBuffer) {
	C.wgpuCommandBufferAddRef(C.WGPUCommandBuffer(commandBuffer))
}

func CommandBufferRelease(commandBuffer CommandBuffer) {
	C.wgpuCommandBufferRelease(C.WGPUCommandBuffer(commandBuffer))
}

func CommandEncoderFinish(commandEncoder CommandEncoder, descriptor *CommandBufferDescriptor) CommandBuffer {
	return CommandBuffer(C.wgpuCommandEncoderFinish(
		C.WGPUCommandEncoder(commandEncoder),
		(*C.WGPUCommandBufferDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func CommandEncoderBeginComputePass(commandEncoder CommandEncoder, descriptor *ComputePassDescriptor) ComputePassEncoder {
	return ComputePassEncoder(C.wgpuCommandEncoderBeginComputePass(
		C.WGPUCommandEncoder(commandEncoder),
		(*C.WGPUComputePassDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func CommandEncoderBeginRenderPass(commandEncoder CommandEncoder, descriptor *RenderPassDescriptor) RenderPassEncoder {
	return RenderPassEncoder(C.wgpuCommandEncoderBeginRenderPass(
		C.WGPUCommandEncoder(commandEncoder),
		(*C.WGPURenderPassDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func CommandEncoderCopyBufferToBuffer(commandEncoder CommandEncoder, source Buffer, source_offset uint64, destination Buffer, destination_offset uint64, size uint64) {
	C.wgpuCommandEncoderCopyBufferToBuffer(
		C.WGPUCommandEncoder(commandEncoder),
		C.WGPUBuffer(source),
		C.uint64_t(source_offset),
		C.WGPUBuffer(destination),
		C.uint64_t(destination_offset),
		C.uint64_t(size),
	)
}

func CommandEncoderCopyBufferToTexture(commandEncoder CommandEncoder, source *TexelCopyBufferInfo, destination *TexelCopyTextureInfo, copy_size *Extent3D) {
	C.wgpuCommandEncoderCopyBufferToTexture(
		C.WGPUCommandEncoder(commandEncoder),
		(*C.WGPUTexelCopyBufferInfo)(unsafe.Pointer(source)),
		(*C.WGPUTexelCopyTextureInfo)(unsafe.Pointer(destination)),
		(*C.WGPUExtent3D)(unsafe.Pointer(copy_size)),
	)
}

func CommandEncoderCopyTextureToBuffer(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyBufferInfo, copy_size *Extent3D) {
	C.wgpuCommandEncoderCopyTextureToBuffer(
		C.WGPUCommandEncoder(commandEncoder),
		(*C.WGPUTexelCopyTextureInfo)(unsafe.Pointer(source)),
		(*C.WGPUTexelCopyBufferInfo)(unsafe.Pointer(destination)),
		(*C.WGPUExtent3D)(unsafe.Pointer(copy_size)),
	)
}

func CommandEncoderCopyTextureToTexture(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyTextureInfo, copy_size *Extent3D) {
	C.wgpuCommandEncoderCopyTextureToTexture(
		C.WGPUCommandEncoder(commandEncoder),
		(*C.WGPUTexelCopyTextureInfo)(unsafe.Pointer(source)),
		(*C.WGPUTexelCopyTextureInfo)(unsafe.Pointer(destination)),
		(*C.WGPUExtent3D)(unsafe.Pointer(copy_size)),
	)
}

func CommandEncoderClearBuffer(commandEncoder CommandEncoder, buffer Buffer, offset uint64, size uint64) {
	C.wgpuCommandEncoderClearBuffer(
		C.WGPUCommandEncoder(commandEncoder),
		C.WGPUBuffer(buffer),
		C.uint64_t(offset),
		C.uint64_t(size),
	)
}

func CommandEncoderInsertDebugMarker(commandEncoder CommandEncoder, marker_label StringView) {
	C.wgpuCommandEncoderInsertDebugMarker(
		C.WGPUCommandEncoder(commandEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(marker_label.Data)),
			length: C.size_t(marker_label.Length),
		},
	)
}

func CommandEncoderPopDebugGroup(commandEncoder CommandEncoder) {
	C.wgpuCommandEncoderPopDebugGroup(
		C.WGPUCommandEncoder(commandEncoder),
	)
}

func CommandEncoderPushDebugGroup(commandEncoder CommandEncoder, group_label StringView) {
	C.wgpuCommandEncoderPushDebugGroup(
		C.WGPUCommandEncoder(commandEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(group_label.Data)),
			length: C.size_t(group_label.Length),
		},
	)
}

func CommandEncoderResolveQuerySet(commandEncoder CommandEncoder, query_set QuerySet, first_query uint32, query_count uint32, destination Buffer, destination_offset uint64) {
	C.wgpuCommandEncoderResolveQuerySet(
		C.WGPUCommandEncoder(commandEncoder),
		C.WGPUQuerySet(query_set),
		C.uint32_t(first_query),
		C.uint32_t(query_count),
		C.WGPUBuffer(destination),
		C.uint64_t(destination_offset),
	)
}

func CommandEncoderWriteTimestamp(commandEncoder CommandEncoder, query_set QuerySet, query_index uint32) {
	C.wgpuCommandEncoderWriteTimestamp(
		C.WGPUCommandEncoder(commandEncoder),
		C.WGPUQuerySet(query_set),
		C.uint32_t(query_index),
	)
}

func CommandEncoderSetLabel(commandEncoder CommandEncoder, label StringView) {
	C.wgpuCommandEncoderSetLabel(
		C.WGPUCommandEncoder(commandEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func CommandEncoderAddRef(commandEncoder CommandEncoder) {
	C.wgpuCommandEncoderAddRef(C.WGPUCommandEncoder(commandEncoder))
}

func CommandEncoderRelease(commandEncoder CommandEncoder) {
	C.wgpuCommandEncoderRelease(C.WGPUCommandEncoder(commandEncoder))
}

func ComputePassEncoderInsertDebugMarker(computePassEncoder ComputePassEncoder, marker_label StringView) {
	C.wgpuComputePassEncoderInsertDebugMarker(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(marker_label.Data)),
			length: C.size_t(marker_label.Length),
		},
	)
}

func ComputePassEncoderPopDebugGroup(computePassEncoder ComputePassEncoder) {
	C.wgpuComputePassEncoderPopDebugGroup(
		C.WGPUComputePassEncoder(computePassEncoder),
	)
}

func ComputePassEncoderPushDebugGroup(computePassEncoder ComputePassEncoder, group_label StringView) {
	C.wgpuComputePassEncoderPushDebugGroup(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(group_label.Data)),
			length: C.size_t(group_label.Length),
		},
	)
}

func ComputePassEncoderSetPipeline(computePassEncoder ComputePassEncoder, pipeline ComputePipeline) {
	C.wgpuComputePassEncoderSetPipeline(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.WGPUComputePipeline(pipeline),
	)
}

func ComputePassEncoderSetBindGroup(computePassEncoder ComputePassEncoder, group_index uint32, group BindGroup, dynamic_offsets []uint32) {
	C.wgpuComputePassEncoderSetBindGroup(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.uint32_t(group_index),
		C.WGPUBindGroup(group),
		C.size_t(len(dynamic_offsets)),
		(*C.uint32_t)(unsafe.Pointer(&dynamic_offsets[0])),
	)
}

func ComputePassEncoderSetImmediates(computePassEncoder ComputePassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	C.wgpuComputePassEncoderSetImmediates(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.uint32_t(offset),
		unsafe.Pointer(data),
		C.size_t(size),
	)
}

func ComputePassEncoderDispatchWorkgroups(computePassEncoder ComputePassEncoder, workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	C.wgpuComputePassEncoderDispatchWorkgroups(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.uint32_t(workgroupCountX),
		C.uint32_t(workgroupCountY),
		C.uint32_t(workgroupCountZ),
	)
}

func ComputePassEncoderDispatchWorkgroupsIndirect(computePassEncoder ComputePassEncoder, indirect_buffer Buffer, indirect_offset uint64) {
	C.wgpuComputePassEncoderDispatchWorkgroupsIndirect(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.WGPUBuffer(indirect_buffer),
		C.uint64_t(indirect_offset),
	)
}

func ComputePassEncoderEnd(computePassEncoder ComputePassEncoder) {
	C.wgpuComputePassEncoderEnd(
		C.WGPUComputePassEncoder(computePassEncoder),
	)
}

func ComputePassEncoderSetLabel(computePassEncoder ComputePassEncoder, label StringView) {
	C.wgpuComputePassEncoderSetLabel(
		C.WGPUComputePassEncoder(computePassEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func ComputePassEncoderAddRef(computePassEncoder ComputePassEncoder) {
	C.wgpuComputePassEncoderAddRef(C.WGPUComputePassEncoder(computePassEncoder))
}

func ComputePassEncoderRelease(computePassEncoder ComputePassEncoder) {
	C.wgpuComputePassEncoderRelease(C.WGPUComputePassEncoder(computePassEncoder))
}

func ComputePipelineGetBindGroupLayout(computePipeline ComputePipeline, group_index uint32) BindGroupLayout {
	return BindGroupLayout(C.wgpuComputePipelineGetBindGroupLayout(
		C.WGPUComputePipeline(computePipeline),
		C.uint32_t(group_index),
	))
}

func ComputePipelineSetLabel(computePipeline ComputePipeline, label StringView) {
	C.wgpuComputePipelineSetLabel(
		C.WGPUComputePipeline(computePipeline),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func ComputePipelineAddRef(computePipeline ComputePipeline) {
	C.wgpuComputePipelineAddRef(C.WGPUComputePipeline(computePipeline))
}

func ComputePipelineRelease(computePipeline ComputePipeline) {
	C.wgpuComputePipelineRelease(C.WGPUComputePipeline(computePipeline))
}

func DeviceCreateBindGroup(device Device, descriptor *BindGroupDescriptor) BindGroup {
	return BindGroup(C.wgpuDeviceCreateBindGroup(
		C.WGPUDevice(device),
		(*C.WGPUBindGroupDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateBindGroupLayout(device Device, descriptor *BindGroupLayoutDescriptor) BindGroupLayout {
	return BindGroupLayout(C.wgpuDeviceCreateBindGroupLayout(
		C.WGPUDevice(device),
		(*C.WGPUBindGroupLayoutDescriptor)(unsafe.Pointer(descriptor)),
	))
}

// TODO
//
// If @ref WGPUBufferDescriptor::mappedAtCreation is `true` and the mapping allocation fails,
// returns `NULL`.
func DeviceCreateBuffer(device Device, descriptor *BufferDescriptor) Buffer {
	return Buffer(C.wgpuDeviceCreateBuffer(
		C.WGPUDevice(device),
		(*C.WGPUBufferDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateCommandEncoder(device Device, descriptor *CommandEncoderDescriptor) CommandEncoder {
	return CommandEncoder(C.wgpuDeviceCreateCommandEncoder(
		C.WGPUDevice(device),
		(*C.WGPUCommandEncoderDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateComputePipeline(device Device, descriptor *ComputePipelineDescriptor) ComputePipeline {
	return ComputePipeline(C.wgpuDeviceCreateComputePipeline(
		C.WGPUDevice(device),
		(*C.WGPUComputePipelineDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateComputePipelineAsync(device Device, descriptor *ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallbackInfo) {
	C.wgpuDeviceCreateComputePipelineAsync(
		C.WGPUDevice(device),
		(*C.WGPUComputePipelineDescriptor)(unsafe.Pointer(descriptor)),
		C.WGPUCreateComputePipelineAsyncCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPUCreateComputePipelineAsyncCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

func DeviceCreatePipelineLayout(device Device, descriptor *PipelineLayoutDescriptor) PipelineLayout {
	return PipelineLayout(C.wgpuDeviceCreatePipelineLayout(
		C.WGPUDevice(device),
		(*C.WGPUPipelineLayoutDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateQuerySet(device Device, descriptor *QuerySetDescriptor) QuerySet {
	return QuerySet(C.wgpuDeviceCreateQuerySet(
		C.WGPUDevice(device),
		(*C.WGPUQuerySetDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateRenderPipelineAsync(device Device, descriptor *RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallbackInfo) {
	C.wgpuDeviceCreateRenderPipelineAsync(
		C.WGPUDevice(device),
		(*C.WGPURenderPipelineDescriptor)(unsafe.Pointer(descriptor)),
		C.WGPUCreateRenderPipelineAsyncCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPUCreateRenderPipelineAsyncCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

func DeviceCreateRenderBundleEncoder(device Device, descriptor *RenderBundleEncoderDescriptor) RenderBundleEncoder {
	return RenderBundleEncoder(C.wgpuDeviceCreateRenderBundleEncoder(
		C.WGPUDevice(device),
		(*C.WGPURenderBundleEncoderDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateRenderPipeline(device Device, descriptor *RenderPipelineDescriptor) RenderPipeline {
	return RenderPipeline(C.wgpuDeviceCreateRenderPipeline(
		C.WGPUDevice(device),
		(*C.WGPURenderPipelineDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateSampler(device Device, descriptor *SamplerDescriptor) Sampler {
	return Sampler(C.wgpuDeviceCreateSampler(
		C.WGPUDevice(device),
		(*C.WGPUSamplerDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateShaderModule(device Device, descriptor *ShaderModuleDescriptor) ShaderModule {
	return ShaderModule(C.wgpuDeviceCreateShaderModule(
		C.WGPUDevice(device),
		(*C.WGPUShaderModuleDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceCreateTexture(device Device, descriptor *TextureDescriptor) Texture {
	return Texture(C.wgpuDeviceCreateTexture(
		C.WGPUDevice(device),
		(*C.WGPUTextureDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func DeviceDestroy(device Device) {
	C.wgpuDeviceDestroy(
		C.WGPUDevice(device),
	)
}

func DeviceGetLostFuture(device Device) Future {
	return Future{Id: uint64(C.wgpuDeviceGetLostFuture(
		C.WGPUDevice(device),
	).id)}
}

func DeviceGetLimits(device Device, limits *Limits) Status {
	return Status(C.wgpuDeviceGetLimits(
		C.WGPUDevice(device),
		(*C.WGPULimits)(unsafe.Pointer(limits)),
	))
}

func DeviceHasFeature(device Device, feature FeatureName) Bool {
	return Bool(C.wgpuDeviceHasFeature(
		C.WGPUDevice(device),
		C.WGPUFeatureName(feature),
	))
}

// Get the list of @ref WGPUFeatureName values supported by the device.
func DeviceGetFeatures(device Device, features *SupportedFeatures) {
	C.wgpuDeviceGetFeatures(
		C.WGPUDevice(device),
		(*C.WGPUSupportedFeatures)(unsafe.Pointer(features)),
	)
}

func DeviceGetAdapterInfo(device Device, adapter_info *AdapterInfo) Status {
	return Status(C.wgpuDeviceGetAdapterInfo(
		C.WGPUDevice(device),
		(*C.WGPUAdapterInfo)(unsafe.Pointer(adapter_info)),
	))
}

func DeviceGetQueue(device Device) Queue {
	return Queue(C.wgpuDeviceGetQueue(
		C.WGPUDevice(device),
	))
}

// Pushes an error scope to the current thread's error scope stack.
// See @ref ErrorScopes.
func DevicePushErrorScope(device Device, filter ErrorFilter) {
	C.wgpuDevicePushErrorScope(
		C.WGPUDevice(device),
		C.WGPUErrorFilter(filter),
	)
}

// Pops an error scope to the current thread's error scope stack,
// asynchronously returning the result. See @ref ErrorScopes.
func DevicePopErrorScope(device Device, callback PopErrorScopeCallbackInfo) {
	C.wgpuDevicePopErrorScope(
		C.WGPUDevice(device),
		C.WGPUPopErrorScopeCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPUPopErrorScopeCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

func DeviceSetLabel(device Device, label StringView) {
	C.wgpuDeviceSetLabel(
		C.WGPUDevice(device),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func DeviceAddRef(device Device) {
	C.wgpuDeviceAddRef(C.WGPUDevice(device))
}

func DeviceRelease(device Device) {
	C.wgpuDeviceRelease(C.WGPUDevice(device))
}

func ExternalTextureSetLabel(externalTexture ExternalTexture, label StringView) {
	C.wgpuExternalTextureSetLabel(
		C.WGPUExternalTexture(externalTexture),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func ExternalTextureAddRef(externalTexture ExternalTexture) {
	C.wgpuExternalTextureAddRef(C.WGPUExternalTexture(externalTexture))
}

func ExternalTextureRelease(externalTexture ExternalTexture) {
	C.wgpuExternalTextureRelease(C.WGPUExternalTexture(externalTexture))
}

// Creates a @ref WGPUSurface, see @ref Surface-Creation for more details.
func InstanceCreateSurface(instance Instance, descriptor *SurfaceDescriptor) Surface {
	return Surface(C.wgpuInstanceCreateSurface(
		C.WGPUInstance(instance),
		(*C.WGPUSurfaceDescriptor)(unsafe.Pointer(descriptor)),
	))
}

// Get the list of @ref WGPUWGSLLanguageFeatureName values supported by the instance.
func InstanceGetWGSLLanguageFeatures(instance Instance, features *SupportedWGSLLanguageFeatures) {
	C.wgpuInstanceGetWGSLLanguageFeatures(
		C.WGPUInstance(instance),
		(*C.WGPUSupportedWGSLLanguageFeatures)(unsafe.Pointer(features)),
	)
}

func InstanceHasWGSLLanguageFeature(instance Instance, feature WGSLLanguageFeatureName) Bool {
	return Bool(C.wgpuInstanceHasWGSLLanguageFeature(
		C.WGPUInstance(instance),
		C.WGPUWGSLLanguageFeatureName(feature),
	))
}

// Processes asynchronous events on this `WGPUInstance`, calling any callbacks for asynchronous operations created with @ref WGPUCallbackMode_AllowProcessEvents.
//
// See @ref Process-Events for more information.
func InstanceProcessEvents(instance Instance) {
	C.wgpuInstanceProcessEvents(
		C.WGPUInstance(instance),
	)
}

func InstanceRequestAdapter(instance Instance, options *RequestAdapterOptions, callback RequestAdapterCallbackInfo) {
	C.wgpuInstanceRequestAdapter(
		C.WGPUInstance(instance),
		(*C.WGPURequestAdapterOptions)(unsafe.Pointer(options)),
		C.WGPURequestAdapterCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPURequestAdapterCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

// Wait for at least one WGPUFuture in `futures` to complete, and call callbacks of the respective completed asynchronous operations.
//
// See @ref Wait-Any for more information.
func InstanceWaitAny(instance Instance, future_count uintptr, futures *FutureWaitInfo, timeout_NS uint64) WaitStatus {
	return WaitStatus(C.wgpuInstanceWaitAny(
		C.WGPUInstance(instance),
		C.size_t(future_count),
		(*C.WGPUFutureWaitInfo)(unsafe.Pointer(futures)),
		C.uint64_t(timeout_NS),
	))
}

func InstanceAddRef(instance Instance) {
	C.wgpuInstanceAddRef(C.WGPUInstance(instance))
}

func InstanceRelease(instance Instance) {
	C.wgpuInstanceRelease(C.WGPUInstance(instance))
}

func PipelineLayoutSetLabel(pipelineLayout PipelineLayout, label StringView) {
	C.wgpuPipelineLayoutSetLabel(
		C.WGPUPipelineLayout(pipelineLayout),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func PipelineLayoutAddRef(pipelineLayout PipelineLayout) {
	C.wgpuPipelineLayoutAddRef(C.WGPUPipelineLayout(pipelineLayout))
}

func PipelineLayoutRelease(pipelineLayout PipelineLayout) {
	C.wgpuPipelineLayoutRelease(C.WGPUPipelineLayout(pipelineLayout))
}

func QuerySetSetLabel(querySet QuerySet, label StringView) {
	C.wgpuQuerySetSetLabel(
		C.WGPUQuerySet(querySet),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func QuerySetGetType(querySet QuerySet) QueryType {
	return QueryType(C.wgpuQuerySetGetType(
		C.WGPUQuerySet(querySet),
	))
}

func QuerySetGetCount(querySet QuerySet) uint32 {
	return uint32(C.wgpuQuerySetGetCount(
		C.WGPUQuerySet(querySet),
	))
}

func QuerySetDestroy(querySet QuerySet) {
	C.wgpuQuerySetDestroy(
		C.WGPUQuerySet(querySet),
	)
}

func QuerySetAddRef(querySet QuerySet) {
	C.wgpuQuerySetAddRef(C.WGPUQuerySet(querySet))
}

func QuerySetRelease(querySet QuerySet) {
	C.wgpuQuerySetRelease(C.WGPUQuerySet(querySet))
}

func QueueSubmit(queue Queue, commands []CommandBuffer) {
	C.wgpuQueueSubmit(
		C.WGPUQueue(queue),
		C.size_t(len(commands)),
		(*C.WGPUCommandBuffer)(unsafe.Pointer(&commands[0])),
	)
}

func QueueOnSubmittedWorkDone(queue Queue, callback QueueWorkDoneCallbackInfo) {
	C.wgpuQueueOnSubmittedWorkDone(
		C.WGPUQueue(queue),
		C.WGPUQueueWorkDoneCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPUQueueWorkDoneCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

// Produces a @ref DeviceError both content-timeline (`size` alignment) and device-timeline
// errors defined by the WebGPU specification.
func QueueWriteBuffer(queue Queue, buffer Buffer, buffer_offset uint64, data unsafe.Pointer, size uintptr) {
	C.wgpuQueueWriteBuffer(
		C.WGPUQueue(queue),
		C.WGPUBuffer(buffer),
		C.uint64_t(buffer_offset),
		unsafe.Pointer(data),
		C.size_t(size),
	)
}

func QueueWriteTexture(queue Queue, destination *TexelCopyTextureInfo, data unsafe.Pointer, data_size uintptr, data_layout *TexelCopyBufferLayout, write_size *Extent3D) {
	C.wgpuQueueWriteTexture(
		C.WGPUQueue(queue),
		(*C.WGPUTexelCopyTextureInfo)(unsafe.Pointer(destination)),
		unsafe.Pointer(data),
		C.size_t(data_size),
		(*C.WGPUTexelCopyBufferLayout)(unsafe.Pointer(data_layout)),
		(*C.WGPUExtent3D)(unsafe.Pointer(write_size)),
	)
}

func QueueSetLabel(queue Queue, label StringView) {
	C.wgpuQueueSetLabel(
		C.WGPUQueue(queue),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func QueueAddRef(queue Queue) {
	C.wgpuQueueAddRef(C.WGPUQueue(queue))
}

func QueueRelease(queue Queue) {
	C.wgpuQueueRelease(C.WGPUQueue(queue))
}

func RenderBundleSetLabel(renderBundle RenderBundle, label StringView) {
	C.wgpuRenderBundleSetLabel(
		C.WGPURenderBundle(renderBundle),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func RenderBundleAddRef(renderBundle RenderBundle) {
	C.wgpuRenderBundleAddRef(C.WGPURenderBundle(renderBundle))
}

func RenderBundleRelease(renderBundle RenderBundle) {
	C.wgpuRenderBundleRelease(C.WGPURenderBundle(renderBundle))
}

func RenderBundleEncoderSetPipeline(renderBundleEncoder RenderBundleEncoder, pipeline RenderPipeline) {
	C.wgpuRenderBundleEncoderSetPipeline(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.WGPURenderPipeline(pipeline),
	)
}

func RenderBundleEncoderSetBindGroup(renderBundleEncoder RenderBundleEncoder, group_index uint32, group BindGroup, dynamic_offsets []uint32) {
	C.wgpuRenderBundleEncoderSetBindGroup(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.uint32_t(group_index),
		C.WGPUBindGroup(group),
		C.size_t(len(dynamic_offsets)),
		(*C.uint32_t)(unsafe.Pointer(&dynamic_offsets[0])),
	)
}

func RenderBundleEncoderSetImmediates(renderBundleEncoder RenderBundleEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	C.wgpuRenderBundleEncoderSetImmediates(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.uint32_t(offset),
		unsafe.Pointer(data),
		C.size_t(size),
	)
}

func RenderBundleEncoderDraw(renderBundleEncoder RenderBundleEncoder, vertex_count uint32, instance_count uint32, first_vertex uint32, first_instance uint32) {
	C.wgpuRenderBundleEncoderDraw(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.uint32_t(vertex_count),
		C.uint32_t(instance_count),
		C.uint32_t(first_vertex),
		C.uint32_t(first_instance),
	)
}

func RenderBundleEncoderDrawIndexed(renderBundleEncoder RenderBundleEncoder, index_count uint32, instance_count uint32, first_index uint32, base_vertex int32, first_instance uint32) {
	C.wgpuRenderBundleEncoderDrawIndexed(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.uint32_t(index_count),
		C.uint32_t(instance_count),
		C.uint32_t(first_index),
		C.int32_t(base_vertex),
		C.uint32_t(first_instance),
	)
}

func RenderBundleEncoderDrawIndirect(renderBundleEncoder RenderBundleEncoder, indirect_buffer Buffer, indirect_offset uint64) {
	C.wgpuRenderBundleEncoderDrawIndirect(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.WGPUBuffer(indirect_buffer),
		C.uint64_t(indirect_offset),
	)
}

func RenderBundleEncoderDrawIndexedIndirect(renderBundleEncoder RenderBundleEncoder, indirect_buffer Buffer, indirect_offset uint64) {
	C.wgpuRenderBundleEncoderDrawIndexedIndirect(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.WGPUBuffer(indirect_buffer),
		C.uint64_t(indirect_offset),
	)
}

func RenderBundleEncoderInsertDebugMarker(renderBundleEncoder RenderBundleEncoder, marker_label StringView) {
	C.wgpuRenderBundleEncoderInsertDebugMarker(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(marker_label.Data)),
			length: C.size_t(marker_label.Length),
		},
	)
}

func RenderBundleEncoderPopDebugGroup(renderBundleEncoder RenderBundleEncoder) {
	C.wgpuRenderBundleEncoderPopDebugGroup(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
	)
}

func RenderBundleEncoderPushDebugGroup(renderBundleEncoder RenderBundleEncoder, group_label StringView) {
	C.wgpuRenderBundleEncoderPushDebugGroup(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(group_label.Data)),
			length: C.size_t(group_label.Length),
		},
	)
}

func RenderBundleEncoderSetVertexBuffer(renderBundleEncoder RenderBundleEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {
	C.wgpuRenderBundleEncoderSetVertexBuffer(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.uint32_t(slot),
		C.WGPUBuffer(buffer),
		C.uint64_t(offset),
		C.uint64_t(size),
	)
}

func RenderBundleEncoderSetIndexBuffer(renderBundleEncoder RenderBundleEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	C.wgpuRenderBundleEncoderSetIndexBuffer(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.WGPUBuffer(buffer),
		C.WGPUIndexFormat(format),
		C.uint64_t(offset),
		C.uint64_t(size),
	)
}

func RenderBundleEncoderFinish(renderBundleEncoder RenderBundleEncoder, descriptor *RenderBundleDescriptor) RenderBundle {
	return RenderBundle(C.wgpuRenderBundleEncoderFinish(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		(*C.WGPURenderBundleDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func RenderBundleEncoderSetLabel(renderBundleEncoder RenderBundleEncoder, label StringView) {
	C.wgpuRenderBundleEncoderSetLabel(
		C.WGPURenderBundleEncoder(renderBundleEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func RenderBundleEncoderAddRef(renderBundleEncoder RenderBundleEncoder) {
	C.wgpuRenderBundleEncoderAddRef(C.WGPURenderBundleEncoder(renderBundleEncoder))
}

func RenderBundleEncoderRelease(renderBundleEncoder RenderBundleEncoder) {
	C.wgpuRenderBundleEncoderRelease(C.WGPURenderBundleEncoder(renderBundleEncoder))
}

func RenderPassEncoderSetPipeline(renderPassEncoder RenderPassEncoder, pipeline RenderPipeline) {
	C.wgpuRenderPassEncoderSetPipeline(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.WGPURenderPipeline(pipeline),
	)
}

func RenderPassEncoderSetBindGroup(renderPassEncoder RenderPassEncoder, group_index uint32, group BindGroup, dynamic_offsets []uint32) {
	C.wgpuRenderPassEncoderSetBindGroup(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(group_index),
		C.WGPUBindGroup(group),
		C.size_t(len(dynamic_offsets)),
		(*C.uint32_t)(unsafe.Pointer(&dynamic_offsets[0])),
	)
}

func RenderPassEncoderSetImmediates(renderPassEncoder RenderPassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	C.wgpuRenderPassEncoderSetImmediates(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(offset),
		unsafe.Pointer(data),
		C.size_t(size),
	)
}

func RenderPassEncoderDraw(renderPassEncoder RenderPassEncoder, vertex_count uint32, instance_count uint32, first_vertex uint32, first_instance uint32) {
	C.wgpuRenderPassEncoderDraw(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(vertex_count),
		C.uint32_t(instance_count),
		C.uint32_t(first_vertex),
		C.uint32_t(first_instance),
	)
}

func RenderPassEncoderDrawIndexed(renderPassEncoder RenderPassEncoder, index_count uint32, instance_count uint32, first_index uint32, base_vertex int32, first_instance uint32) {
	C.wgpuRenderPassEncoderDrawIndexed(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(index_count),
		C.uint32_t(instance_count),
		C.uint32_t(first_index),
		C.int32_t(base_vertex),
		C.uint32_t(first_instance),
	)
}

func RenderPassEncoderDrawIndirect(renderPassEncoder RenderPassEncoder, indirect_buffer Buffer, indirect_offset uint64) {
	C.wgpuRenderPassEncoderDrawIndirect(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.WGPUBuffer(indirect_buffer),
		C.uint64_t(indirect_offset),
	)
}

func RenderPassEncoderDrawIndexedIndirect(renderPassEncoder RenderPassEncoder, indirect_buffer Buffer, indirect_offset uint64) {
	C.wgpuRenderPassEncoderDrawIndexedIndirect(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.WGPUBuffer(indirect_buffer),
		C.uint64_t(indirect_offset),
	)
}

func RenderPassEncoderExecuteBundles(renderPassEncoder RenderPassEncoder, bundles []RenderBundle) {
	C.wgpuRenderPassEncoderExecuteBundles(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.size_t(len(bundles)),
		(*C.WGPURenderBundle)(unsafe.Pointer(&bundles[0])),
	)
}

func RenderPassEncoderInsertDebugMarker(renderPassEncoder RenderPassEncoder, marker_label StringView) {
	C.wgpuRenderPassEncoderInsertDebugMarker(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(marker_label.Data)),
			length: C.size_t(marker_label.Length),
		},
	)
}

func RenderPassEncoderPopDebugGroup(renderPassEncoder RenderPassEncoder) {
	C.wgpuRenderPassEncoderPopDebugGroup(
		C.WGPURenderPassEncoder(renderPassEncoder),
	)
}

func RenderPassEncoderPushDebugGroup(renderPassEncoder RenderPassEncoder, group_label StringView) {
	C.wgpuRenderPassEncoderPushDebugGroup(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(group_label.Data)),
			length: C.size_t(group_label.Length),
		},
	)
}

func RenderPassEncoderSetStencilReference(renderPassEncoder RenderPassEncoder, reference uint32) {
	C.wgpuRenderPassEncoderSetStencilReference(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(reference),
	)
}

func RenderPassEncoderSetBlendConstant(renderPassEncoder RenderPassEncoder, color *Color) {
	C.wgpuRenderPassEncoderSetBlendConstant(
		C.WGPURenderPassEncoder(renderPassEncoder),
		(*C.WGPUColor)(unsafe.Pointer(color)),
	)
}

// TODO
//
// If any argument is non-finite, produces a @ref NonFiniteFloatValueError.
func RenderPassEncoderSetViewport(renderPassEncoder RenderPassEncoder, x float32, y float32, width float32, height float32, min_depth float32, max_depth float32) {
	C.wgpuRenderPassEncoderSetViewport(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.float(x),
		C.float(y),
		C.float(width),
		C.float(height),
		C.float(min_depth),
		C.float(max_depth),
	)
}

func RenderPassEncoderSetScissorRect(renderPassEncoder RenderPassEncoder, x uint32, y uint32, width uint32, height uint32) {
	C.wgpuRenderPassEncoderSetScissorRect(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(x),
		C.uint32_t(y),
		C.uint32_t(width),
		C.uint32_t(height),
	)
}

func RenderPassEncoderSetVertexBuffer(renderPassEncoder RenderPassEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {
	C.wgpuRenderPassEncoderSetVertexBuffer(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(slot),
		C.WGPUBuffer(buffer),
		C.uint64_t(offset),
		C.uint64_t(size),
	)
}

func RenderPassEncoderSetIndexBuffer(renderPassEncoder RenderPassEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	C.wgpuRenderPassEncoderSetIndexBuffer(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.WGPUBuffer(buffer),
		C.WGPUIndexFormat(format),
		C.uint64_t(offset),
		C.uint64_t(size),
	)
}

func RenderPassEncoderBeginOcclusionQuery(renderPassEncoder RenderPassEncoder, query_index uint32) {
	C.wgpuRenderPassEncoderBeginOcclusionQuery(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.uint32_t(query_index),
	)
}

func RenderPassEncoderEndOcclusionQuery(renderPassEncoder RenderPassEncoder) {
	C.wgpuRenderPassEncoderEndOcclusionQuery(
		C.WGPURenderPassEncoder(renderPassEncoder),
	)
}

func RenderPassEncoderEnd(renderPassEncoder RenderPassEncoder) {
	C.wgpuRenderPassEncoderEnd(
		C.WGPURenderPassEncoder(renderPassEncoder),
	)
}

func RenderPassEncoderSetLabel(renderPassEncoder RenderPassEncoder, label StringView) {
	C.wgpuRenderPassEncoderSetLabel(
		C.WGPURenderPassEncoder(renderPassEncoder),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func RenderPassEncoderAddRef(renderPassEncoder RenderPassEncoder) {
	C.wgpuRenderPassEncoderAddRef(C.WGPURenderPassEncoder(renderPassEncoder))
}

func RenderPassEncoderRelease(renderPassEncoder RenderPassEncoder) {
	C.wgpuRenderPassEncoderRelease(C.WGPURenderPassEncoder(renderPassEncoder))
}

func RenderPipelineGetBindGroupLayout(renderPipeline RenderPipeline, group_index uint32) BindGroupLayout {
	return BindGroupLayout(C.wgpuRenderPipelineGetBindGroupLayout(
		C.WGPURenderPipeline(renderPipeline),
		C.uint32_t(group_index),
	))
}

func RenderPipelineSetLabel(renderPipeline RenderPipeline, label StringView) {
	C.wgpuRenderPipelineSetLabel(
		C.WGPURenderPipeline(renderPipeline),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func RenderPipelineAddRef(renderPipeline RenderPipeline) {
	C.wgpuRenderPipelineAddRef(C.WGPURenderPipeline(renderPipeline))
}

func RenderPipelineRelease(renderPipeline RenderPipeline) {
	C.wgpuRenderPipelineRelease(C.WGPURenderPipeline(renderPipeline))
}

func SamplerSetLabel(sampler Sampler, label StringView) {
	C.wgpuSamplerSetLabel(
		C.WGPUSampler(sampler),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func SamplerAddRef(sampler Sampler) {
	C.wgpuSamplerAddRef(C.WGPUSampler(sampler))
}

func SamplerRelease(sampler Sampler) {
	C.wgpuSamplerRelease(C.WGPUSampler(sampler))
}

func ShaderModuleGetCompilationInfo(shaderModule ShaderModule, callback CompilationInfoCallbackInfo) {
	C.wgpuShaderModuleGetCompilationInfo(
		C.WGPUShaderModule(shaderModule),
		C.WGPUCompilationInfoCallbackInfo{
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(callback.NextInChain)),
			mode:        C.WGPUCallbackMode(callback.Mode),
			callback:    C.WGPUCompilationInfoCallback(callback.Callback),
			userdata1:   callback.Userdata1,
			userdata2:   callback.Userdata2,
		},
	)
}

func ShaderModuleSetLabel(shaderModule ShaderModule, label StringView) {
	C.wgpuShaderModuleSetLabel(
		C.WGPUShaderModule(shaderModule),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func ShaderModuleAddRef(shaderModule ShaderModule) {
	C.wgpuShaderModuleAddRef(C.WGPUShaderModule(shaderModule))
}

func ShaderModuleRelease(shaderModule ShaderModule) {
	C.wgpuShaderModuleRelease(C.WGPUShaderModule(shaderModule))
}

// Configures parameters for rendering to `surface`.
// Produces a @ref DeviceError for all content-timeline errors defined by the WebGPU specification.
//
// See @ref Surface-Configuration for more details.
func SurfaceConfigure(surface Surface, config *SurfaceConfiguration) {
	C.wgpuSurfaceConfigure(
		C.WGPUSurface(surface),
		(*C.WGPUSurfaceConfiguration)(unsafe.Pointer(config)),
	)
}

// Provides information on how `adapter` is able to use `surface`.
// See @ref Surface-Capabilities for more details.
func SurfaceGetCapabilities(surface Surface, adapter Adapter, capabilities *SurfaceCapabilities) Status {
	return Status(C.wgpuSurfaceGetCapabilities(
		C.WGPUSurface(surface),
		C.WGPUAdapter(adapter),
		(*C.WGPUSurfaceCapabilities)(unsafe.Pointer(capabilities)),
	))
}

// Returns the @ref WGPUTexture to render to `surface` this frame along with metadata on the frame.
// Returns `NULL` and @ref WGPUSurfaceGetCurrentTextureStatus_Error if the surface is not configured.
//
// See @ref Surface-Presenting for more details.
func SurfaceGetCurrentTexture(surface Surface, surface_texture *SurfaceTexture) {
	C.wgpuSurfaceGetCurrentTexture(
		C.WGPUSurface(surface),
		(*C.WGPUSurfaceTexture)(unsafe.Pointer(surface_texture)),
	)
}

// Shows `surface`'s current texture to the user.
// See @ref Surface-Presenting for more details.
func SurfacePresent(surface Surface) Status {
	return Status(C.wgpuSurfacePresent(
		C.WGPUSurface(surface),
	))
}

// Removes the configuration for `surface`.
// See @ref Surface-Configuration for more details.
func SurfaceUnconfigure(surface Surface) {
	C.wgpuSurfaceUnconfigure(
		C.WGPUSurface(surface),
	)
}

// Modifies the label used to refer to `surface`.
func SurfaceSetLabel(surface Surface, label StringView) {
	C.wgpuSurfaceSetLabel(
		C.WGPUSurface(surface),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func SurfaceAddRef(surface Surface) {
	C.wgpuSurfaceAddRef(C.WGPUSurface(surface))
}

func SurfaceRelease(surface Surface) {
	C.wgpuSurfaceRelease(C.WGPUSurface(surface))
}

func TextureCreateView(texture Texture, descriptor *TextureViewDescriptor) TextureView {
	return TextureView(C.wgpuTextureCreateView(
		C.WGPUTexture(texture),
		(*C.WGPUTextureViewDescriptor)(unsafe.Pointer(descriptor)),
	))
}

func TextureSetLabel(texture Texture, label StringView) {
	C.wgpuTextureSetLabel(
		C.WGPUTexture(texture),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func TextureGetWidth(texture Texture) uint32 {
	return uint32(C.wgpuTextureGetWidth(
		C.WGPUTexture(texture),
	))
}

func TextureGetHeight(texture Texture) uint32 {
	return uint32(C.wgpuTextureGetHeight(
		C.WGPUTexture(texture),
	))
}

func TextureGetDepthOrArrayLayers(texture Texture) uint32 {
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(
		C.WGPUTexture(texture),
	))
}

func TextureGetMipLevelCount(texture Texture) uint32 {
	return uint32(C.wgpuTextureGetMipLevelCount(
		C.WGPUTexture(texture),
	))
}

func TextureGetSampleCount(texture Texture) uint32 {
	return uint32(C.wgpuTextureGetSampleCount(
		C.WGPUTexture(texture),
	))
}

func TextureGetDimension(texture Texture) TextureDimension {
	return TextureDimension(C.wgpuTextureGetDimension(
		C.WGPUTexture(texture),
	))
}

func TextureGetTextureBindingViewDimension(texture Texture) TextureViewDimension {
	return TextureViewDimension(C.wgpuTextureGetTextureBindingViewDimension(
		C.WGPUTexture(texture),
	))
}

func TextureGetFormat(texture Texture) TextureFormat {
	return TextureFormat(C.wgpuTextureGetFormat(
		C.WGPUTexture(texture),
	))
}

func TextureGetUsage(texture Texture) TextureUsage {
	return TextureUsage(C.wgpuTextureGetUsage(
		C.WGPUTexture(texture),
	))
}

func TextureDestroy(texture Texture) {
	C.wgpuTextureDestroy(
		C.WGPUTexture(texture),
	)
}

func TextureAddRef(texture Texture) {
	C.wgpuTextureAddRef(C.WGPUTexture(texture))
}

func TextureRelease(texture Texture) {
	C.wgpuTextureRelease(C.WGPUTexture(texture))
}

func TextureViewSetLabel(textureView TextureView, label StringView) {
	C.wgpuTextureViewSetLabel(
		C.WGPUTextureView(textureView),
		C.WGPUStringView{
			data:   (*C.char)(unsafe.Pointer(label.Data)),
			length: C.size_t(label.Length),
		},
	)
}

func TextureViewAddRef(textureView TextureView) {
	C.wgpuTextureViewAddRef(C.WGPUTextureView(textureView))
}

func TextureViewRelease(textureView TextureView) {
	C.wgpuTextureViewRelease(C.WGPUTextureView(textureView))
}
