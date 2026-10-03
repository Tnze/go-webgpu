// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import (
	"runtime"
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu/sys"
)

type Adapter struct {
	inner   sys.Adapter
	cleanup runtime.Cleanup
}

func (a *Adapter) GetLimits(
	limits *Limits) Status {
	ret := sys.AdapterGetLimits(
		a.inner,
		limits,
	)
	return ret
}

func (a *Adapter) HasFeature(
	feature FeatureName) Bool {
	ret := sys.AdapterHasFeature(
		a.inner,
		feature,
	)
	return ret
}

func (a *Adapter) GetFeatures(
	features *SupportedFeatures) {
	var _features *sys.SupportedFeatures
	if features != nil {
		_features = new(features.unwrap())
	}

	sys.AdapterGetFeatures(
		a.inner,
		_features,
	)

	if features != nil {
		features.wrap(_features)
	}
}

func (a *Adapter) GetInfo(
	info *AdapterInfo) Status {
	var _info *sys.AdapterInfo
	if info != nil {
		_info = new(info.unwrap())
	}

	ret := sys.AdapterGetInfo(
		a.inner,
		_info,
	)

	if info != nil {
		info.wrap(_info)
	}
	return ret
}

func (a *Adapter) RequestDevice(
	descriptor *DeviceDescriptor) (RequestDeviceStatus, *Device, string) {
	var _descriptor *sys.DeviceDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	callback := requestDeviceCallback{
		C: make(chan requestDeviceCallbackResult, 1),
	}

	sys.AdapterRequestDevice(
		a.inner,
		_descriptor,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.device, ret.message
}

func (a *Adapter) owned() *Adapter {
	a.cleanup = runtime.AddCleanup(a, sys.AdapterRelease, a.inner)
	return a
}

func (a *Adapter) Release() {
	a.cleanup.Stop()
	sys.AdapterRelease(a.inner)
}

type BindGroup struct {
	inner   sys.BindGroup
	cleanup runtime.Cleanup
}

func (b *BindGroup) SetLabel(
	label string) {
	sys.BindGroupSetLabel(
		b.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (b *BindGroup) owned() *BindGroup {
	b.cleanup = runtime.AddCleanup(b, sys.BindGroupRelease, b.inner)
	return b
}

func (b *BindGroup) Release() {
	b.cleanup.Stop()
	sys.BindGroupRelease(b.inner)
}

type BindGroupLayout struct {
	inner   sys.BindGroupLayout
	cleanup runtime.Cleanup
}

func (b *BindGroupLayout) SetLabel(
	label string) {
	sys.BindGroupLayoutSetLabel(
		b.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (b *BindGroupLayout) owned() *BindGroupLayout {
	b.cleanup = runtime.AddCleanup(b, sys.BindGroupLayoutRelease, b.inner)
	return b
}

func (b *BindGroupLayout) Release() {
	b.cleanup.Stop()
	sys.BindGroupLayoutRelease(b.inner)
}

type Buffer struct {
	inner   sys.Buffer
	cleanup runtime.Cleanup
}

func (b *Buffer) MapAsync(
	mode MapMode,
	offset uintptr,
	size uintptr) (MapAsyncStatus, string) {
	callback := bufferMapCallback{
		C: make(chan bufferMapCallbackResult, 1),
	}

	sys.BufferMapAsync(
		b.inner,
		mode,
		offset,
		size,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.message
}

func (b *Buffer) GetMappedRange(
	offset uintptr,
	size uintptr) unsafe.Pointer {
	ret := sys.BufferGetMappedRange(
		b.inner,
		offset,
		size,
	)
	return ret
}

func (b *Buffer) GetConstMappedRange(
	offset uintptr,
	size uintptr) unsafe.Pointer {
	ret := sys.BufferGetConstMappedRange(
		b.inner,
		offset,
		size,
	)
	return ret
}

func (b *Buffer) ReadMappedRange(
	offset uintptr,
	data unsafe.Pointer,
	size uintptr) Status {
	ret := sys.BufferReadMappedRange(
		b.inner,
		offset,
		data,
		size,
	)
	return ret
}

func (b *Buffer) WriteMappedRange(
	offset uintptr,
	data unsafe.Pointer,
	size uintptr) Status {
	ret := sys.BufferWriteMappedRange(
		b.inner,
		offset,
		data,
		size,
	)
	return ret
}

func (b *Buffer) SetLabel(
	label string) {
	sys.BufferSetLabel(
		b.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (b *Buffer) GetUsage() BufferUsage {
	ret := sys.BufferGetUsage(
		b.inner,
	)
	return ret
}

func (b *Buffer) GetSize() uint64 {
	ret := sys.BufferGetSize(
		b.inner,
	)
	return ret
}

func (b *Buffer) GetMapState() BufferMapState {
	ret := sys.BufferGetMapState(
		b.inner,
	)
	return ret
}

func (b *Buffer) Unmap() {
	sys.BufferUnmap(
		b.inner,
	)
}

func (b *Buffer) Destroy() {
	sys.BufferDestroy(
		b.inner,
	)
}

func (b *Buffer) owned() *Buffer {
	b.cleanup = runtime.AddCleanup(b, sys.BufferRelease, b.inner)
	return b
}

func (b *Buffer) Release() {
	b.cleanup.Stop()
	sys.BufferRelease(b.inner)
}

type CommandBuffer struct {
	inner   sys.CommandBuffer
	cleanup runtime.Cleanup
}

func (c *CommandBuffer) SetLabel(
	label string) {
	sys.CommandBufferSetLabel(
		c.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (c *CommandBuffer) owned() *CommandBuffer {
	c.cleanup = runtime.AddCleanup(c, sys.CommandBufferRelease, c.inner)
	return c
}

func (c *CommandBuffer) Release() {
	c.cleanup.Stop()
	sys.CommandBufferRelease(c.inner)
}

type CommandEncoder struct {
	inner   sys.CommandEncoder
	cleanup runtime.Cleanup
}

func (c *CommandEncoder) Finish(
	descriptor *CommandBufferDescriptor) *CommandBuffer {
	var _descriptor *sys.CommandBufferDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.CommandEncoderFinish(
		c.inner,
		_descriptor,
	)
	return new(CommandBuffer{inner: ret}).owned()
}

func (c *CommandEncoder) BeginComputePass(
	descriptor *ComputePassDescriptor) *ComputePassEncoder {
	var _descriptor *sys.ComputePassDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.CommandEncoderBeginComputePass(
		c.inner,
		_descriptor,
	)
	return new(ComputePassEncoder{inner: ret}).owned()
}

func (c *CommandEncoder) BeginRenderPass(
	descriptor *RenderPassDescriptor) *RenderPassEncoder {
	var _descriptor *sys.RenderPassDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.CommandEncoderBeginRenderPass(
		c.inner,
		_descriptor,
	)
	return new(RenderPassEncoder{inner: ret}).owned()
}

func (c *CommandEncoder) CopyBufferToBuffer(
	source *Buffer,
	sourceOffset uint64,
	destination *Buffer,
	destinationOffset uint64,
	size uint64) {
	sys.CommandEncoderCopyBufferToBuffer(
		c.inner,
		source.inner,
		sourceOffset,
		destination.inner,
		destinationOffset,
		size,
	)
}

func (c *CommandEncoder) CopyBufferToTexture(
	source *TexelCopyBufferInfo,
	destination *TexelCopyTextureInfo,
	copySize *Extent3D) {
	var _source *sys.TexelCopyBufferInfo
	if source != nil {
		_source = new(source.unwrap())
	}

	var _destination *sys.TexelCopyTextureInfo
	if destination != nil {
		_destination = new(destination.unwrap())
	}

	sys.CommandEncoderCopyBufferToTexture(
		c.inner,
		_source,
		_destination,
		copySize,
	)
}

func (c *CommandEncoder) CopyTextureToBuffer(
	source *TexelCopyTextureInfo,
	destination *TexelCopyBufferInfo,
	copySize *Extent3D) {
	var _source *sys.TexelCopyTextureInfo
	if source != nil {
		_source = new(source.unwrap())
	}

	var _destination *sys.TexelCopyBufferInfo
	if destination != nil {
		_destination = new(destination.unwrap())
	}

	sys.CommandEncoderCopyTextureToBuffer(
		c.inner,
		_source,
		_destination,
		copySize,
	)
}

func (c *CommandEncoder) CopyTextureToTexture(
	source *TexelCopyTextureInfo,
	destination *TexelCopyTextureInfo,
	copySize *Extent3D) {
	var _source *sys.TexelCopyTextureInfo
	if source != nil {
		_source = new(source.unwrap())
	}

	var _destination *sys.TexelCopyTextureInfo
	if destination != nil {
		_destination = new(destination.unwrap())
	}

	sys.CommandEncoderCopyTextureToTexture(
		c.inner,
		_source,
		_destination,
		copySize,
	)
}

func (c *CommandEncoder) ClearBuffer(
	buffer *Buffer,
	offset uint64,
	size uint64) {
	sys.CommandEncoderClearBuffer(
		c.inner,
		buffer.inner,
		offset,
		size,
	)
}

func (c *CommandEncoder) InsertDebugMarker(
	markerLabel string) {
	sys.CommandEncoderInsertDebugMarker(
		c.inner,
		sys.StringView{
			Length: uint(len(markerLabel)),
			Data:   unsafe.StringData(markerLabel),
		},
	)
}

func (c *CommandEncoder) PopDebugGroup() {
	sys.CommandEncoderPopDebugGroup(
		c.inner,
	)
}

func (c *CommandEncoder) PushDebugGroup(
	groupLabel string) {
	sys.CommandEncoderPushDebugGroup(
		c.inner,
		sys.StringView{
			Length: uint(len(groupLabel)),
			Data:   unsafe.StringData(groupLabel),
		},
	)
}

func (c *CommandEncoder) ResolveQuerySet(
	querySet *QuerySet,
	firstQuery uint32,
	queryCount uint32,
	destination *Buffer,
	destinationOffset uint64) {
	sys.CommandEncoderResolveQuerySet(
		c.inner,
		querySet.inner,
		firstQuery,
		queryCount,
		destination.inner,
		destinationOffset,
	)
}

func (c *CommandEncoder) WriteTimestamp(
	querySet *QuerySet,
	queryIndex uint32) {
	sys.CommandEncoderWriteTimestamp(
		c.inner,
		querySet.inner,
		queryIndex,
	)
}

func (c *CommandEncoder) SetLabel(
	label string) {
	sys.CommandEncoderSetLabel(
		c.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (c *CommandEncoder) owned() *CommandEncoder {
	c.cleanup = runtime.AddCleanup(c, sys.CommandEncoderRelease, c.inner)
	return c
}

func (c *CommandEncoder) Release() {
	c.cleanup.Stop()
	sys.CommandEncoderRelease(c.inner)
}

type ComputePassEncoder struct {
	inner   sys.ComputePassEncoder
	cleanup runtime.Cleanup
}

func (c *ComputePassEncoder) InsertDebugMarker(
	markerLabel string) {
	sys.ComputePassEncoderInsertDebugMarker(
		c.inner,
		sys.StringView{
			Length: uint(len(markerLabel)),
			Data:   unsafe.StringData(markerLabel),
		},
	)
}

func (c *ComputePassEncoder) PopDebugGroup() {
	sys.ComputePassEncoderPopDebugGroup(
		c.inner,
	)
}

func (c *ComputePassEncoder) PushDebugGroup(
	groupLabel string) {
	sys.ComputePassEncoderPushDebugGroup(
		c.inner,
		sys.StringView{
			Length: uint(len(groupLabel)),
			Data:   unsafe.StringData(groupLabel),
		},
	)
}

func (c *ComputePassEncoder) SetPipeline(
	pipeline *ComputePipeline) {
	sys.ComputePassEncoderSetPipeline(
		c.inner,
		pipeline.inner,
	)
}

func (c *ComputePassEncoder) SetBindGroup(
	groupIndex uint32,
	group *BindGroup,
	dynamicOffsets []uint32) {
	sys.ComputePassEncoderSetBindGroup(
		c.inner,
		groupIndex,
		group.inner,
		dynamicOffsets,
	)
}

func (c *ComputePassEncoder) SetImmediates(
	offset uint32,
	data unsafe.Pointer,
	size uintptr) {
	sys.ComputePassEncoderSetImmediates(
		c.inner,
		offset,
		data,
		size,
	)
}

func (c *ComputePassEncoder) DispatchWorkgroups(
	workgroupCountX uint32,
	workgroupCountY uint32,
	workgroupCountZ uint32) {
	sys.ComputePassEncoderDispatchWorkgroups(
		c.inner,
		workgroupCountX,
		workgroupCountY,
		workgroupCountZ,
	)
}

func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(
	indirectBuffer *Buffer,
	indirectOffset uint64) {
	sys.ComputePassEncoderDispatchWorkgroupsIndirect(
		c.inner,
		indirectBuffer.inner,
		indirectOffset,
	)
}

func (c *ComputePassEncoder) End() {
	sys.ComputePassEncoderEnd(
		c.inner,
	)
}

func (c *ComputePassEncoder) SetLabel(
	label string) {
	sys.ComputePassEncoderSetLabel(
		c.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (c *ComputePassEncoder) owned() *ComputePassEncoder {
	c.cleanup = runtime.AddCleanup(c, sys.ComputePassEncoderRelease, c.inner)
	return c
}

func (c *ComputePassEncoder) Release() {
	c.cleanup.Stop()
	sys.ComputePassEncoderRelease(c.inner)
}

type ComputePipeline struct {
	inner   sys.ComputePipeline
	cleanup runtime.Cleanup
}

func (c *ComputePipeline) GetBindGroupLayout(
	groupIndex uint32) *BindGroupLayout {
	ret := sys.ComputePipelineGetBindGroupLayout(
		c.inner,
		groupIndex,
	)
	return new(BindGroupLayout{inner: ret}).owned()
}

func (c *ComputePipeline) SetLabel(
	label string) {
	sys.ComputePipelineSetLabel(
		c.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (c *ComputePipeline) owned() *ComputePipeline {
	c.cleanup = runtime.AddCleanup(c, sys.ComputePipelineRelease, c.inner)
	return c
}

func (c *ComputePipeline) Release() {
	c.cleanup.Stop()
	sys.ComputePipelineRelease(c.inner)
}

// TODO
//
// Releasing the last ref to a `WGPUDevice` also calls @ref wgpuDeviceDestroy.
// For more info, see @ref DeviceRelease.

type Device struct {
	inner   sys.Device
	cleanup runtime.Cleanup
}

func (d *Device) CreateBindGroup(
	descriptor *BindGroupDescriptor) *BindGroup {
	var _descriptor *sys.BindGroupDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateBindGroup(
		d.inner,
		_descriptor,
	)
	return new(BindGroup{inner: ret}).owned()
}

func (d *Device) CreateBindGroupLayout(
	descriptor *BindGroupLayoutDescriptor) *BindGroupLayout {
	var _descriptor *sys.BindGroupLayoutDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateBindGroupLayout(
		d.inner,
		_descriptor,
	)
	return new(BindGroupLayout{inner: ret}).owned()
}

func (d *Device) CreateBuffer(
	descriptor *BufferDescriptor) *Buffer {
	var _descriptor *sys.BufferDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateBuffer(
		d.inner,
		_descriptor,
	)
	return new(Buffer{inner: ret}).owned()
}

func (d *Device) CreateCommandEncoder(
	descriptor *CommandEncoderDescriptor) *CommandEncoder {
	var _descriptor *sys.CommandEncoderDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateCommandEncoder(
		d.inner,
		_descriptor,
	)
	return new(CommandEncoder{inner: ret}).owned()
}

func (d *Device) CreateComputePipeline(
	descriptor *ComputePipelineDescriptor) *ComputePipeline {
	var _descriptor *sys.ComputePipelineDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateComputePipeline(
		d.inner,
		_descriptor,
	)
	return new(ComputePipeline{inner: ret}).owned()
}

func (d *Device) CreateComputePipelineAsync(
	descriptor *ComputePipelineDescriptor) (CreatePipelineAsyncStatus, *ComputePipeline, string) {
	var _descriptor *sys.ComputePipelineDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	callback := createComputePipelineAsyncCallback{
		C: make(chan createComputePipelineAsyncCallbackResult, 1),
	}

	sys.DeviceCreateComputePipelineAsync(
		d.inner,
		_descriptor,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.pipeline, ret.message
}

func (d *Device) CreatePipelineLayout(
	descriptor *PipelineLayoutDescriptor) *PipelineLayout {
	var _descriptor *sys.PipelineLayoutDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreatePipelineLayout(
		d.inner,
		_descriptor,
	)
	return new(PipelineLayout{inner: ret}).owned()
}

func (d *Device) CreateQuerySet(
	descriptor *QuerySetDescriptor) *QuerySet {
	var _descriptor *sys.QuerySetDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateQuerySet(
		d.inner,
		_descriptor,
	)
	return new(QuerySet{inner: ret}).owned()
}

func (d *Device) CreateRenderPipelineAsync(
	descriptor *RenderPipelineDescriptor) (CreatePipelineAsyncStatus, *RenderPipeline, string) {
	var _descriptor *sys.RenderPipelineDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	callback := createRenderPipelineAsyncCallback{
		C: make(chan createRenderPipelineAsyncCallbackResult, 1),
	}

	sys.DeviceCreateRenderPipelineAsync(
		d.inner,
		_descriptor,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.pipeline, ret.message
}

func (d *Device) CreateRenderBundleEncoder(
	descriptor *RenderBundleEncoderDescriptor) *RenderBundleEncoder {
	var _descriptor *sys.RenderBundleEncoderDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateRenderBundleEncoder(
		d.inner,
		_descriptor,
	)
	return new(RenderBundleEncoder{inner: ret}).owned()
}

func (d *Device) CreateRenderPipeline(
	descriptor *RenderPipelineDescriptor) *RenderPipeline {
	var _descriptor *sys.RenderPipelineDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateRenderPipeline(
		d.inner,
		_descriptor,
	)
	return new(RenderPipeline{inner: ret}).owned()
}

func (d *Device) CreateSampler(
	descriptor *SamplerDescriptor) *Sampler {
	var _descriptor *sys.SamplerDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateSampler(
		d.inner,
		_descriptor,
	)
	return new(Sampler{inner: ret}).owned()
}

func (d *Device) CreateShaderModule(
	descriptor *ShaderModuleDescriptor) *ShaderModule {
	var _descriptor *sys.ShaderModuleDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateShaderModule(
		d.inner,
		_descriptor,
	)
	return new(ShaderModule{inner: ret}).owned()
}

func (d *Device) CreateTexture(
	descriptor *TextureDescriptor) *Texture {
	var _descriptor *sys.TextureDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.DeviceCreateTexture(
		d.inner,
		_descriptor,
	)
	return new(Texture{inner: ret}).owned()
}

func (d *Device) Destroy() {
	sys.DeviceDestroy(
		d.inner,
	)
}

func (d *Device) GetLostFuture() Future {
	ret := sys.DeviceGetLostFuture(
		d.inner,
	)
	return Future{
		Id: ret.Id,
	}
}

func (d *Device) GetLimits(
	limits *Limits) Status {
	ret := sys.DeviceGetLimits(
		d.inner,
		limits,
	)
	return ret
}

func (d *Device) HasFeature(
	feature FeatureName) Bool {
	ret := sys.DeviceHasFeature(
		d.inner,
		feature,
	)
	return ret
}

func (d *Device) GetFeatures(
	features *SupportedFeatures) {
	var _features *sys.SupportedFeatures
	if features != nil {
		_features = new(features.unwrap())
	}

	sys.DeviceGetFeatures(
		d.inner,
		_features,
	)

	if features != nil {
		features.wrap(_features)
	}
}

func (d *Device) GetAdapterInfo(
	adapterInfo *AdapterInfo) Status {
	var _adapterInfo *sys.AdapterInfo
	if adapterInfo != nil {
		_adapterInfo = new(adapterInfo.unwrap())
	}

	ret := sys.DeviceGetAdapterInfo(
		d.inner,
		_adapterInfo,
	)

	if adapterInfo != nil {
		adapterInfo.wrap(_adapterInfo)
	}
	return ret
}

func (d *Device) GetQueue() *Queue {
	ret := sys.DeviceGetQueue(
		d.inner,
	)
	return new(Queue{inner: ret}).owned()
}

func (d *Device) PushErrorScope(
	filter ErrorFilter) {
	sys.DevicePushErrorScope(
		d.inner,
		filter,
	)
}

func (d *Device) PopErrorScope() (PopErrorScopeStatus, ErrorType, string) {
	callback := popErrorScopeCallback{
		C: make(chan popErrorScopeCallbackResult, 1),
	}

	sys.DevicePopErrorScope(
		d.inner,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.errorType, ret.message
}

func (d *Device) SetLabel(
	label string) {
	sys.DeviceSetLabel(
		d.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (d *Device) owned() *Device {
	d.cleanup = runtime.AddCleanup(d, sys.DeviceRelease, d.inner)
	return d
}

func (d *Device) Release() {
	d.cleanup.Stop()
	sys.DeviceRelease(d.inner)
}

// A sampleable 2D texture that may perform 0-copy YUV sampling internally. Creation of @ref WGPUExternalTexture is extremely implementation-dependent and not defined in this header.

type ExternalTexture struct {
	inner   sys.ExternalTexture
	cleanup runtime.Cleanup
}

func (e *ExternalTexture) SetLabel(
	label string) {
	sys.ExternalTextureSetLabel(
		e.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (e *ExternalTexture) owned() *ExternalTexture {
	e.cleanup = runtime.AddCleanup(e, sys.ExternalTextureRelease, e.inner)
	return e
}

func (e *ExternalTexture) Release() {
	e.cleanup.Stop()
	sys.ExternalTextureRelease(e.inner)
}

type Instance struct {
	inner   sys.Instance
	cleanup runtime.Cleanup
}

func (i *Instance) CreateSurface(
	descriptor *SurfaceDescriptor) *Surface {
	var _descriptor *sys.SurfaceDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.InstanceCreateSurface(
		i.inner,
		_descriptor,
	)
	return new(Surface{inner: ret}).owned()
}

func (i *Instance) GetWGSLLanguageFeatures(
	features *SupportedWGSLLanguageFeatures) {
	var _features *sys.SupportedWGSLLanguageFeatures
	if features != nil {
		_features = new(features.unwrap())
	}

	sys.InstanceGetWGSLLanguageFeatures(
		i.inner,
		_features,
	)

	if features != nil {
		features.wrap(_features)
	}
}

func (i *Instance) HasWGSLLanguageFeature(
	feature WGSLLanguageFeatureName) Bool {
	ret := sys.InstanceHasWGSLLanguageFeature(
		i.inner,
		feature,
	)
	return ret
}

func (i *Instance) ProcessEvents() {
	sys.InstanceProcessEvents(
		i.inner,
	)
}

func (i *Instance) RequestAdapter(
	options *RequestAdapterOptions) (RequestAdapterStatus, *Adapter, string) {
	var _options *sys.RequestAdapterOptions
	if options != nil {
		_options = new(options.unwrap())
	}

	callback := requestAdapterCallback{
		C: make(chan requestAdapterCallbackResult, 1),
	}

	sys.InstanceRequestAdapter(
		i.inner,
		_options,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.adapter, ret.message
}

func (i *Instance) WaitAny(
	futureCount uintptr,
	futures *FutureWaitInfo,
	timeoutNS uint64) WaitStatus {
	ret := sys.InstanceWaitAny(
		i.inner,
		futureCount,
		futures,
		timeoutNS,
	)
	return ret
}

func (i *Instance) owned() *Instance {
	i.cleanup = runtime.AddCleanup(i, sys.InstanceRelease, i.inner)
	return i
}

func (i *Instance) Release() {
	i.cleanup.Stop()
	sys.InstanceRelease(i.inner)
}

type PipelineLayout struct {
	inner   sys.PipelineLayout
	cleanup runtime.Cleanup
}

func (p *PipelineLayout) SetLabel(
	label string) {
	sys.PipelineLayoutSetLabel(
		p.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (p *PipelineLayout) owned() *PipelineLayout {
	p.cleanup = runtime.AddCleanup(p, sys.PipelineLayoutRelease, p.inner)
	return p
}

func (p *PipelineLayout) Release() {
	p.cleanup.Stop()
	sys.PipelineLayoutRelease(p.inner)
}

type QuerySet struct {
	inner   sys.QuerySet
	cleanup runtime.Cleanup
}

func (q *QuerySet) SetLabel(
	label string) {
	sys.QuerySetSetLabel(
		q.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (q *QuerySet) GetType() QueryType {
	ret := sys.QuerySetGetType(
		q.inner,
	)
	return ret
}

func (q *QuerySet) GetCount() uint32 {
	ret := sys.QuerySetGetCount(
		q.inner,
	)
	return ret
}

func (q *QuerySet) Destroy() {
	sys.QuerySetDestroy(
		q.inner,
	)
}

func (q *QuerySet) owned() *QuerySet {
	q.cleanup = runtime.AddCleanup(q, sys.QuerySetRelease, q.inner)
	return q
}

func (q *QuerySet) Release() {
	q.cleanup.Stop()
	sys.QuerySetRelease(q.inner)
}

type Queue struct {
	inner   sys.Queue
	cleanup runtime.Cleanup
}

func (q *Queue) Submit(
	commands []*CommandBuffer) {
	_commands := make([]sys.CommandBuffer, len(commands))
	for i := range commands {
		_commands[i] = commands[i].inner
	}
	sys.QueueSubmit(
		q.inner,
		_commands,
	)
}

func (q *Queue) OnSubmittedWorkDone() (QueueWorkDoneStatus, string) {
	callback := queueWorkDoneCallback{
		C: make(chan queueWorkDoneCallbackResult, 1),
	}

	sys.QueueOnSubmittedWorkDone(
		q.inner,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.message
}

func (q *Queue) WriteBuffer(
	buffer *Buffer,
	bufferOffset uint64,
	data unsafe.Pointer,
	size uintptr) {
	sys.QueueWriteBuffer(
		q.inner,
		buffer.inner,
		bufferOffset,
		data,
		size,
	)
}

func (q *Queue) WriteTexture(
	destination *TexelCopyTextureInfo,
	data unsafe.Pointer,
	dataSize uintptr,
	dataLayout *TexelCopyBufferLayout,
	writeSize *Extent3D) {
	var _destination *sys.TexelCopyTextureInfo
	if destination != nil {
		_destination = new(destination.unwrap())
	}

	sys.QueueWriteTexture(
		q.inner,
		_destination,
		data,
		dataSize,
		dataLayout,
		writeSize,
	)
}

func (q *Queue) SetLabel(
	label string) {
	sys.QueueSetLabel(
		q.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (q *Queue) owned() *Queue {
	q.cleanup = runtime.AddCleanup(q, sys.QueueRelease, q.inner)
	return q
}

func (q *Queue) Release() {
	q.cleanup.Stop()
	sys.QueueRelease(q.inner)
}

type RenderBundle struct {
	inner   sys.RenderBundle
	cleanup runtime.Cleanup
}

func (r *RenderBundle) SetLabel(
	label string) {
	sys.RenderBundleSetLabel(
		r.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (r *RenderBundle) owned() *RenderBundle {
	r.cleanup = runtime.AddCleanup(r, sys.RenderBundleRelease, r.inner)
	return r
}

func (r *RenderBundle) Release() {
	r.cleanup.Stop()
	sys.RenderBundleRelease(r.inner)
}

type RenderBundleEncoder struct {
	inner   sys.RenderBundleEncoder
	cleanup runtime.Cleanup
}

func (r *RenderBundleEncoder) SetPipeline(
	pipeline *RenderPipeline) {
	sys.RenderBundleEncoderSetPipeline(
		r.inner,
		pipeline.inner,
	)
}

func (r *RenderBundleEncoder) SetBindGroup(
	groupIndex uint32,
	group *BindGroup,
	dynamicOffsets []uint32) {
	sys.RenderBundleEncoderSetBindGroup(
		r.inner,
		groupIndex,
		group.inner,
		dynamicOffsets,
	)
}

func (r *RenderBundleEncoder) SetImmediates(
	offset uint32,
	data unsafe.Pointer,
	size uintptr) {
	sys.RenderBundleEncoderSetImmediates(
		r.inner,
		offset,
		data,
		size,
	)
}

func (r *RenderBundleEncoder) Draw(
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32) {
	sys.RenderBundleEncoderDraw(
		r.inner,
		vertexCount,
		instanceCount,
		firstVertex,
		firstInstance,
	)
}

func (r *RenderBundleEncoder) DrawIndexed(
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32) {
	sys.RenderBundleEncoderDrawIndexed(
		r.inner,
		indexCount,
		instanceCount,
		firstIndex,
		baseVertex,
		firstInstance,
	)
}

func (r *RenderBundleEncoder) DrawIndirect(
	indirectBuffer *Buffer,
	indirectOffset uint64) {
	sys.RenderBundleEncoderDrawIndirect(
		r.inner,
		indirectBuffer.inner,
		indirectOffset,
	)
}

func (r *RenderBundleEncoder) DrawIndexedIndirect(
	indirectBuffer *Buffer,
	indirectOffset uint64) {
	sys.RenderBundleEncoderDrawIndexedIndirect(
		r.inner,
		indirectBuffer.inner,
		indirectOffset,
	)
}

func (r *RenderBundleEncoder) InsertDebugMarker(
	markerLabel string) {
	sys.RenderBundleEncoderInsertDebugMarker(
		r.inner,
		sys.StringView{
			Length: uint(len(markerLabel)),
			Data:   unsafe.StringData(markerLabel),
		},
	)
}

func (r *RenderBundleEncoder) PopDebugGroup() {
	sys.RenderBundleEncoderPopDebugGroup(
		r.inner,
	)
}

func (r *RenderBundleEncoder) PushDebugGroup(
	groupLabel string) {
	sys.RenderBundleEncoderPushDebugGroup(
		r.inner,
		sys.StringView{
			Length: uint(len(groupLabel)),
			Data:   unsafe.StringData(groupLabel),
		},
	)
}

func (r *RenderBundleEncoder) SetVertexBuffer(
	slot uint32,
	buffer *Buffer,
	offset uint64,
	size uint64) {
	sys.RenderBundleEncoderSetVertexBuffer(
		r.inner,
		slot,
		buffer.inner,
		offset,
		size,
	)
}

func (r *RenderBundleEncoder) SetIndexBuffer(
	buffer *Buffer,
	format IndexFormat,
	offset uint64,
	size uint64) {
	sys.RenderBundleEncoderSetIndexBuffer(
		r.inner,
		buffer.inner,
		format,
		offset,
		size,
	)
}

func (r *RenderBundleEncoder) Finish(
	descriptor *RenderBundleDescriptor) *RenderBundle {
	var _descriptor *sys.RenderBundleDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.RenderBundleEncoderFinish(
		r.inner,
		_descriptor,
	)
	return new(RenderBundle{inner: ret}).owned()
}

func (r *RenderBundleEncoder) SetLabel(
	label string) {
	sys.RenderBundleEncoderSetLabel(
		r.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (r *RenderBundleEncoder) owned() *RenderBundleEncoder {
	r.cleanup = runtime.AddCleanup(r, sys.RenderBundleEncoderRelease, r.inner)
	return r
}

func (r *RenderBundleEncoder) Release() {
	r.cleanup.Stop()
	sys.RenderBundleEncoderRelease(r.inner)
}

type RenderPassEncoder struct {
	inner   sys.RenderPassEncoder
	cleanup runtime.Cleanup
}

func (r *RenderPassEncoder) SetPipeline(
	pipeline *RenderPipeline) {
	sys.RenderPassEncoderSetPipeline(
		r.inner,
		pipeline.inner,
	)
}

func (r *RenderPassEncoder) SetBindGroup(
	groupIndex uint32,
	group *BindGroup,
	dynamicOffsets []uint32) {
	sys.RenderPassEncoderSetBindGroup(
		r.inner,
		groupIndex,
		group.inner,
		dynamicOffsets,
	)
}

func (r *RenderPassEncoder) SetImmediates(
	offset uint32,
	data unsafe.Pointer,
	size uintptr) {
	sys.RenderPassEncoderSetImmediates(
		r.inner,
		offset,
		data,
		size,
	)
}

func (r *RenderPassEncoder) Draw(
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32) {
	sys.RenderPassEncoderDraw(
		r.inner,
		vertexCount,
		instanceCount,
		firstVertex,
		firstInstance,
	)
}

func (r *RenderPassEncoder) DrawIndexed(
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	baseVertex int32,
	firstInstance uint32) {
	sys.RenderPassEncoderDrawIndexed(
		r.inner,
		indexCount,
		instanceCount,
		firstIndex,
		baseVertex,
		firstInstance,
	)
}

func (r *RenderPassEncoder) DrawIndirect(
	indirectBuffer *Buffer,
	indirectOffset uint64) {
	sys.RenderPassEncoderDrawIndirect(
		r.inner,
		indirectBuffer.inner,
		indirectOffset,
	)
}

func (r *RenderPassEncoder) DrawIndexedIndirect(
	indirectBuffer *Buffer,
	indirectOffset uint64) {
	sys.RenderPassEncoderDrawIndexedIndirect(
		r.inner,
		indirectBuffer.inner,
		indirectOffset,
	)
}

func (r *RenderPassEncoder) ExecuteBundles(
	bundles []*RenderBundle) {
	_bundles := make([]sys.RenderBundle, len(bundles))
	for i := range bundles {
		_bundles[i] = bundles[i].inner
	}
	sys.RenderPassEncoderExecuteBundles(
		r.inner,
		_bundles,
	)
}

func (r *RenderPassEncoder) InsertDebugMarker(
	markerLabel string) {
	sys.RenderPassEncoderInsertDebugMarker(
		r.inner,
		sys.StringView{
			Length: uint(len(markerLabel)),
			Data:   unsafe.StringData(markerLabel),
		},
	)
}

func (r *RenderPassEncoder) PopDebugGroup() {
	sys.RenderPassEncoderPopDebugGroup(
		r.inner,
	)
}

func (r *RenderPassEncoder) PushDebugGroup(
	groupLabel string) {
	sys.RenderPassEncoderPushDebugGroup(
		r.inner,
		sys.StringView{
			Length: uint(len(groupLabel)),
			Data:   unsafe.StringData(groupLabel),
		},
	)
}

func (r *RenderPassEncoder) SetStencilReference(
	reference uint32) {
	sys.RenderPassEncoderSetStencilReference(
		r.inner,
		reference,
	)
}

func (r *RenderPassEncoder) SetBlendConstant(
	color *Color) {
	sys.RenderPassEncoderSetBlendConstant(
		r.inner,
		color,
	)
}

func (r *RenderPassEncoder) SetViewport(
	x float32,
	y float32,
	width float32,
	height float32,
	minDepth float32,
	maxDepth float32) {
	sys.RenderPassEncoderSetViewport(
		r.inner,
		x,
		y,
		width,
		height,
		minDepth,
		maxDepth,
	)
}

func (r *RenderPassEncoder) SetScissorRect(
	x uint32,
	y uint32,
	width uint32,
	height uint32) {
	sys.RenderPassEncoderSetScissorRect(
		r.inner,
		x,
		y,
		width,
		height,
	)
}

func (r *RenderPassEncoder) SetVertexBuffer(
	slot uint32,
	buffer *Buffer,
	offset uint64,
	size uint64) {
	sys.RenderPassEncoderSetVertexBuffer(
		r.inner,
		slot,
		buffer.inner,
		offset,
		size,
	)
}

func (r *RenderPassEncoder) SetIndexBuffer(
	buffer *Buffer,
	format IndexFormat,
	offset uint64,
	size uint64) {
	sys.RenderPassEncoderSetIndexBuffer(
		r.inner,
		buffer.inner,
		format,
		offset,
		size,
	)
}

func (r *RenderPassEncoder) BeginOcclusionQuery(
	queryIndex uint32) {
	sys.RenderPassEncoderBeginOcclusionQuery(
		r.inner,
		queryIndex,
	)
}

func (r *RenderPassEncoder) EndOcclusionQuery() {
	sys.RenderPassEncoderEndOcclusionQuery(
		r.inner,
	)
}

func (r *RenderPassEncoder) End() {
	sys.RenderPassEncoderEnd(
		r.inner,
	)
}

func (r *RenderPassEncoder) SetLabel(
	label string) {
	sys.RenderPassEncoderSetLabel(
		r.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (r *RenderPassEncoder) owned() *RenderPassEncoder {
	r.cleanup = runtime.AddCleanup(r, sys.RenderPassEncoderRelease, r.inner)
	return r
}

func (r *RenderPassEncoder) Release() {
	r.cleanup.Stop()
	sys.RenderPassEncoderRelease(r.inner)
}

type RenderPipeline struct {
	inner   sys.RenderPipeline
	cleanup runtime.Cleanup
}

func (r *RenderPipeline) GetBindGroupLayout(
	groupIndex uint32) *BindGroupLayout {
	ret := sys.RenderPipelineGetBindGroupLayout(
		r.inner,
		groupIndex,
	)
	return new(BindGroupLayout{inner: ret}).owned()
}

func (r *RenderPipeline) SetLabel(
	label string) {
	sys.RenderPipelineSetLabel(
		r.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (r *RenderPipeline) owned() *RenderPipeline {
	r.cleanup = runtime.AddCleanup(r, sys.RenderPipelineRelease, r.inner)
	return r
}

func (r *RenderPipeline) Release() {
	r.cleanup.Stop()
	sys.RenderPipelineRelease(r.inner)
}

type Sampler struct {
	inner   sys.Sampler
	cleanup runtime.Cleanup
}

func (s *Sampler) SetLabel(
	label string) {
	sys.SamplerSetLabel(
		s.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (s *Sampler) owned() *Sampler {
	s.cleanup = runtime.AddCleanup(s, sys.SamplerRelease, s.inner)
	return s
}

func (s *Sampler) Release() {
	s.cleanup.Stop()
	sys.SamplerRelease(s.inner)
}

type ShaderModule struct {
	inner   sys.ShaderModule
	cleanup runtime.Cleanup
}

func (s *ShaderModule) GetCompilationInfo() (CompilationInfoRequestStatus, *CompilationInfo) {
	callback := compilationInfoCallback{
		C: make(chan compilationInfoCallbackResult, 1),
	}

	sys.ShaderModuleGetCompilationInfo(
		s.inner,
		callback.info(),
	)
	ret := <-callback.C
	return ret.status, ret.compilationInfo
}

func (s *ShaderModule) SetLabel(
	label string) {
	sys.ShaderModuleSetLabel(
		s.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (s *ShaderModule) owned() *ShaderModule {
	s.cleanup = runtime.AddCleanup(s, sys.ShaderModuleRelease, s.inner)
	return s
}

func (s *ShaderModule) Release() {
	s.cleanup.Stop()
	sys.ShaderModuleRelease(s.inner)
}

// An object used to continuously present image data to the user, see @ref Surfaces for more details.

type Surface struct {
	inner   sys.Surface
	cleanup runtime.Cleanup
}

func (s *Surface) Configure(
	config *SurfaceConfiguration) {
	var _config *sys.SurfaceConfiguration
	if config != nil {
		_config = new(config.unwrap())
	}

	sys.SurfaceConfigure(
		s.inner,
		_config,
	)
}

func (s *Surface) GetCapabilities(
	adapter *Adapter,
	capabilities *SurfaceCapabilities) Status {
	var _capabilities *sys.SurfaceCapabilities
	if capabilities != nil {
		_capabilities = new(capabilities.unwrap())
	}

	ret := sys.SurfaceGetCapabilities(
		s.inner,
		adapter.inner,
		_capabilities,
	)

	if capabilities != nil {
		capabilities.wrap(_capabilities)
	}
	return ret
}

func (s *Surface) GetCurrentTexture(
	surfaceTexture *SurfaceTexture) {
	var _surfaceTexture *sys.SurfaceTexture
	if surfaceTexture != nil {
		_surfaceTexture = new(surfaceTexture.unwrap())
	}

	sys.SurfaceGetCurrentTexture(
		s.inner,
		_surfaceTexture,
	)

	if surfaceTexture != nil {
		surfaceTexture.wrap(_surfaceTexture)
	}
}

func (s *Surface) Present() Status {
	ret := sys.SurfacePresent(
		s.inner,
	)
	return ret
}

func (s *Surface) Unconfigure() {
	sys.SurfaceUnconfigure(
		s.inner,
	)
}

func (s *Surface) SetLabel(
	label string) {
	sys.SurfaceSetLabel(
		s.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (s *Surface) owned() *Surface {
	s.cleanup = runtime.AddCleanup(s, sys.SurfaceRelease, s.inner)
	return s
}

func (s *Surface) Release() {
	s.cleanup.Stop()
	sys.SurfaceRelease(s.inner)
}

type Texture struct {
	inner   sys.Texture
	cleanup runtime.Cleanup
}

func (t *Texture) CreateView(
	descriptor *TextureViewDescriptor) *TextureView {
	var _descriptor *sys.TextureViewDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}

	ret := sys.TextureCreateView(
		t.inner,
		_descriptor,
	)
	return new(TextureView{inner: ret}).owned()
}

func (t *Texture) SetLabel(
	label string) {
	sys.TextureSetLabel(
		t.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (t *Texture) GetWidth() uint32 {
	ret := sys.TextureGetWidth(
		t.inner,
	)
	return ret
}

func (t *Texture) GetHeight() uint32 {
	ret := sys.TextureGetHeight(
		t.inner,
	)
	return ret
}

func (t *Texture) GetDepthOrArrayLayers() uint32 {
	ret := sys.TextureGetDepthOrArrayLayers(
		t.inner,
	)
	return ret
}

func (t *Texture) GetMipLevelCount() uint32 {
	ret := sys.TextureGetMipLevelCount(
		t.inner,
	)
	return ret
}

func (t *Texture) GetSampleCount() uint32 {
	ret := sys.TextureGetSampleCount(
		t.inner,
	)
	return ret
}

func (t *Texture) GetDimension() TextureDimension {
	ret := sys.TextureGetDimension(
		t.inner,
	)
	return ret
}

func (t *Texture) GetTextureBindingViewDimension() TextureViewDimension {
	ret := sys.TextureGetTextureBindingViewDimension(
		t.inner,
	)
	return ret
}

func (t *Texture) GetFormat() TextureFormat {
	ret := sys.TextureGetFormat(
		t.inner,
	)
	return ret
}

func (t *Texture) GetUsage() TextureUsage {
	ret := sys.TextureGetUsage(
		t.inner,
	)
	return ret
}

func (t *Texture) Destroy() {
	sys.TextureDestroy(
		t.inner,
	)
}

func (t *Texture) owned() *Texture {
	t.cleanup = runtime.AddCleanup(t, sys.TextureRelease, t.inner)
	return t
}

func (t *Texture) Release() {
	t.cleanup.Stop()
	sys.TextureRelease(t.inner)
}

type TextureView struct {
	inner   sys.TextureView
	cleanup runtime.Cleanup
}

func (t *TextureView) SetLabel(
	label string) {
	sys.TextureViewSetLabel(
		t.inner,
		sys.StringView{
			Length: uint(len(label)),
			Data:   unsafe.StringData(label),
		},
	)
}

func (t *TextureView) owned() *TextureView {
	t.cleanup = runtime.AddCleanup(t, sys.TextureViewRelease, t.inner)
	return t
}

func (t *TextureView) Release() {
	t.cleanup.Stop()
	sys.TextureViewRelease(t.inner)
}
