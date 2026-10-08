// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build !cgo

package sys

import (
	"syscall"
	"unsafe"
)

var (
	webgpuDLL                                        = syscall.NewLazyDLL("wgpu-native.dll")
	adapterGetLimitsProc                             = webgpuDLL.NewProc("AdapterGetLimits")
	adapterHasFeatureProc                            = webgpuDLL.NewProc("AdapterHasFeature")
	adapterGetFeaturesProc                           = webgpuDLL.NewProc("AdapterGetFeatures")
	adapterGetInfoProc                               = webgpuDLL.NewProc("AdapterGetInfo")
	adapterRequestDeviceProc                         = webgpuDLL.NewProc("AdapterRequestDevice")
	adapterAddRefProc                                = webgpuDLL.NewProc("AdapterAddRef")
	adapterReleaseProc                               = webgpuDLL.NewProc("AdapterRelease")
	bindGroupSetLabelProc                            = webgpuDLL.NewProc("BindGroupSetLabel")
	bindGroupAddRefProc                              = webgpuDLL.NewProc("BindGroupAddRef")
	bindGroupReleaseProc                             = webgpuDLL.NewProc("BindGroupRelease")
	bindGroupLayoutSetLabelProc                      = webgpuDLL.NewProc("BindGroupLayoutSetLabel")
	bindGroupLayoutAddRefProc                        = webgpuDLL.NewProc("BindGroupLayoutAddRef")
	bindGroupLayoutReleaseProc                       = webgpuDLL.NewProc("BindGroupLayoutRelease")
	bufferMapAsyncProc                               = webgpuDLL.NewProc("BufferMapAsync")
	bufferGetMappedRangeProc                         = webgpuDLL.NewProc("BufferGetMappedRange")
	bufferGetConstMappedRangeProc                    = webgpuDLL.NewProc("BufferGetConstMappedRange")
	bufferReadMappedRangeProc                        = webgpuDLL.NewProc("BufferReadMappedRange")
	bufferWriteMappedRangeProc                       = webgpuDLL.NewProc("BufferWriteMappedRange")
	bufferSetLabelProc                               = webgpuDLL.NewProc("BufferSetLabel")
	bufferGetUsageProc                               = webgpuDLL.NewProc("BufferGetUsage")
	bufferGetSizeProc                                = webgpuDLL.NewProc("BufferGetSize")
	bufferGetMapStateProc                            = webgpuDLL.NewProc("BufferGetMapState")
	bufferUnmapProc                                  = webgpuDLL.NewProc("BufferUnmap")
	bufferDestroyProc                                = webgpuDLL.NewProc("BufferDestroy")
	bufferAddRefProc                                 = webgpuDLL.NewProc("BufferAddRef")
	bufferReleaseProc                                = webgpuDLL.NewProc("BufferRelease")
	commandBufferSetLabelProc                        = webgpuDLL.NewProc("CommandBufferSetLabel")
	commandBufferAddRefProc                          = webgpuDLL.NewProc("CommandBufferAddRef")
	commandBufferReleaseProc                         = webgpuDLL.NewProc("CommandBufferRelease")
	commandEncoderFinishProc                         = webgpuDLL.NewProc("CommandEncoderFinish")
	commandEncoderBeginComputePassProc               = webgpuDLL.NewProc("CommandEncoderBeginComputePass")
	commandEncoderBeginRenderPassProc                = webgpuDLL.NewProc("CommandEncoderBeginRenderPass")
	commandEncoderCopyBufferToBufferProc             = webgpuDLL.NewProc("CommandEncoderCopyBufferToBuffer")
	commandEncoderCopyBufferToTextureProc            = webgpuDLL.NewProc("CommandEncoderCopyBufferToTexture")
	commandEncoderCopyTextureToBufferProc            = webgpuDLL.NewProc("CommandEncoderCopyTextureToBuffer")
	commandEncoderCopyTextureToTextureProc           = webgpuDLL.NewProc("CommandEncoderCopyTextureToTexture")
	commandEncoderClearBufferProc                    = webgpuDLL.NewProc("CommandEncoderClearBuffer")
	commandEncoderInsertDebugMarkerProc              = webgpuDLL.NewProc("CommandEncoderInsertDebugMarker")
	commandEncoderPopDebugGroupProc                  = webgpuDLL.NewProc("CommandEncoderPopDebugGroup")
	commandEncoderPushDebugGroupProc                 = webgpuDLL.NewProc("CommandEncoderPushDebugGroup")
	commandEncoderResolveQuerySetProc                = webgpuDLL.NewProc("CommandEncoderResolveQuerySet")
	commandEncoderWriteTimestampProc                 = webgpuDLL.NewProc("CommandEncoderWriteTimestamp")
	commandEncoderSetLabelProc                       = webgpuDLL.NewProc("CommandEncoderSetLabel")
	commandEncoderAddRefProc                         = webgpuDLL.NewProc("CommandEncoderAddRef")
	commandEncoderReleaseProc                        = webgpuDLL.NewProc("CommandEncoderRelease")
	computePassEncoderInsertDebugMarkerProc          = webgpuDLL.NewProc("ComputePassEncoderInsertDebugMarker")
	computePassEncoderPopDebugGroupProc              = webgpuDLL.NewProc("ComputePassEncoderPopDebugGroup")
	computePassEncoderPushDebugGroupProc             = webgpuDLL.NewProc("ComputePassEncoderPushDebugGroup")
	computePassEncoderSetPipelineProc                = webgpuDLL.NewProc("ComputePassEncoderSetPipeline")
	computePassEncoderSetBindGroupProc               = webgpuDLL.NewProc("ComputePassEncoderSetBindGroup")
	computePassEncoderSetImmediatesProc              = webgpuDLL.NewProc("ComputePassEncoderSetImmediates")
	computePassEncoderDispatchWorkgroupsProc         = webgpuDLL.NewProc("ComputePassEncoderDispatchWorkgroups")
	computePassEncoderDispatchWorkgroupsIndirectProc = webgpuDLL.NewProc("ComputePassEncoderDispatchWorkgroupsIndirect")
	computePassEncoderEndProc                        = webgpuDLL.NewProc("ComputePassEncoderEnd")
	computePassEncoderSetLabelProc                   = webgpuDLL.NewProc("ComputePassEncoderSetLabel")
	computePassEncoderAddRefProc                     = webgpuDLL.NewProc("ComputePassEncoderAddRef")
	computePassEncoderReleaseProc                    = webgpuDLL.NewProc("ComputePassEncoderRelease")
	computePipelineGetBindGroupLayoutProc            = webgpuDLL.NewProc("ComputePipelineGetBindGroupLayout")
	computePipelineSetLabelProc                      = webgpuDLL.NewProc("ComputePipelineSetLabel")
	computePipelineAddRefProc                        = webgpuDLL.NewProc("ComputePipelineAddRef")
	computePipelineReleaseProc                       = webgpuDLL.NewProc("ComputePipelineRelease")
	deviceCreateBindGroupProc                        = webgpuDLL.NewProc("DeviceCreateBindGroup")
	deviceCreateBindGroupLayoutProc                  = webgpuDLL.NewProc("DeviceCreateBindGroupLayout")
	deviceCreateBufferProc                           = webgpuDLL.NewProc("DeviceCreateBuffer")
	deviceCreateCommandEncoderProc                   = webgpuDLL.NewProc("DeviceCreateCommandEncoder")
	deviceCreateComputePipelineProc                  = webgpuDLL.NewProc("DeviceCreateComputePipeline")
	deviceCreateComputePipelineAsyncProc             = webgpuDLL.NewProc("DeviceCreateComputePipelineAsync")
	deviceCreatePipelineLayoutProc                   = webgpuDLL.NewProc("DeviceCreatePipelineLayout")
	deviceCreateQuerySetProc                         = webgpuDLL.NewProc("DeviceCreateQuerySet")
	deviceCreateRenderPipelineAsyncProc              = webgpuDLL.NewProc("DeviceCreateRenderPipelineAsync")
	deviceCreateRenderBundleEncoderProc              = webgpuDLL.NewProc("DeviceCreateRenderBundleEncoder")
	deviceCreateRenderPipelineProc                   = webgpuDLL.NewProc("DeviceCreateRenderPipeline")
	deviceCreateSamplerProc                          = webgpuDLL.NewProc("DeviceCreateSampler")
	deviceCreateShaderModuleProc                     = webgpuDLL.NewProc("DeviceCreateShaderModule")
	deviceCreateTextureProc                          = webgpuDLL.NewProc("DeviceCreateTexture")
	deviceDestroyProc                                = webgpuDLL.NewProc("DeviceDestroy")
	deviceGetLostFutureProc                          = webgpuDLL.NewProc("DeviceGetLostFuture")
	deviceGetLimitsProc                              = webgpuDLL.NewProc("DeviceGetLimits")
	deviceHasFeatureProc                             = webgpuDLL.NewProc("DeviceHasFeature")
	deviceGetFeaturesProc                            = webgpuDLL.NewProc("DeviceGetFeatures")
	deviceGetAdapterInfoProc                         = webgpuDLL.NewProc("DeviceGetAdapterInfo")
	deviceGetQueueProc                               = webgpuDLL.NewProc("DeviceGetQueue")
	devicePushErrorScopeProc                         = webgpuDLL.NewProc("DevicePushErrorScope")
	devicePopErrorScopeProc                          = webgpuDLL.NewProc("DevicePopErrorScope")
	deviceSetLabelProc                               = webgpuDLL.NewProc("DeviceSetLabel")
	deviceAddRefProc                                 = webgpuDLL.NewProc("DeviceAddRef")
	deviceReleaseProc                                = webgpuDLL.NewProc("DeviceRelease")
	externalTextureSetLabelProc                      = webgpuDLL.NewProc("ExternalTextureSetLabel")
	externalTextureAddRefProc                        = webgpuDLL.NewProc("ExternalTextureAddRef")
	externalTextureReleaseProc                       = webgpuDLL.NewProc("ExternalTextureRelease")
	instanceCreateSurfaceProc                        = webgpuDLL.NewProc("InstanceCreateSurface")
	instanceGetWGSLLanguageFeaturesProc              = webgpuDLL.NewProc("InstanceGetWGSLLanguageFeatures")
	instanceHasWGSLLanguageFeatureProc               = webgpuDLL.NewProc("InstanceHasWGSLLanguageFeature")
	instanceProcessEventsProc                        = webgpuDLL.NewProc("InstanceProcessEvents")
	instanceRequestAdapterProc                       = webgpuDLL.NewProc("InstanceRequestAdapter")
	instanceWaitAnyProc                              = webgpuDLL.NewProc("InstanceWaitAny")
	instanceAddRefProc                               = webgpuDLL.NewProc("InstanceAddRef")
	instanceReleaseProc                              = webgpuDLL.NewProc("InstanceRelease")
	pipelineLayoutSetLabelProc                       = webgpuDLL.NewProc("PipelineLayoutSetLabel")
	pipelineLayoutAddRefProc                         = webgpuDLL.NewProc("PipelineLayoutAddRef")
	pipelineLayoutReleaseProc                        = webgpuDLL.NewProc("PipelineLayoutRelease")
	querySetSetLabelProc                             = webgpuDLL.NewProc("QuerySetSetLabel")
	querySetGetTypeProc                              = webgpuDLL.NewProc("QuerySetGetType")
	querySetGetCountProc                             = webgpuDLL.NewProc("QuerySetGetCount")
	querySetDestroyProc                              = webgpuDLL.NewProc("QuerySetDestroy")
	querySetAddRefProc                               = webgpuDLL.NewProc("QuerySetAddRef")
	querySetReleaseProc                              = webgpuDLL.NewProc("QuerySetRelease")
	queueSubmitProc                                  = webgpuDLL.NewProc("QueueSubmit")
	queueOnSubmittedWorkDoneProc                     = webgpuDLL.NewProc("QueueOnSubmittedWorkDone")
	queueWriteBufferProc                             = webgpuDLL.NewProc("QueueWriteBuffer")
	queueWriteTextureProc                            = webgpuDLL.NewProc("QueueWriteTexture")
	queueSetLabelProc                                = webgpuDLL.NewProc("QueueSetLabel")
	queueAddRefProc                                  = webgpuDLL.NewProc("QueueAddRef")
	queueReleaseProc                                 = webgpuDLL.NewProc("QueueRelease")
	renderBundleSetLabelProc                         = webgpuDLL.NewProc("RenderBundleSetLabel")
	renderBundleAddRefProc                           = webgpuDLL.NewProc("RenderBundleAddRef")
	renderBundleReleaseProc                          = webgpuDLL.NewProc("RenderBundleRelease")
	renderBundleEncoderSetPipelineProc               = webgpuDLL.NewProc("RenderBundleEncoderSetPipeline")
	renderBundleEncoderSetBindGroupProc              = webgpuDLL.NewProc("RenderBundleEncoderSetBindGroup")
	renderBundleEncoderSetImmediatesProc             = webgpuDLL.NewProc("RenderBundleEncoderSetImmediates")
	renderBundleEncoderDrawProc                      = webgpuDLL.NewProc("RenderBundleEncoderDraw")
	renderBundleEncoderDrawIndexedProc               = webgpuDLL.NewProc("RenderBundleEncoderDrawIndexed")
	renderBundleEncoderDrawIndirectProc              = webgpuDLL.NewProc("RenderBundleEncoderDrawIndirect")
	renderBundleEncoderDrawIndexedIndirectProc       = webgpuDLL.NewProc("RenderBundleEncoderDrawIndexedIndirect")
	renderBundleEncoderInsertDebugMarkerProc         = webgpuDLL.NewProc("RenderBundleEncoderInsertDebugMarker")
	renderBundleEncoderPopDebugGroupProc             = webgpuDLL.NewProc("RenderBundleEncoderPopDebugGroup")
	renderBundleEncoderPushDebugGroupProc            = webgpuDLL.NewProc("RenderBundleEncoderPushDebugGroup")
	renderBundleEncoderSetVertexBufferProc           = webgpuDLL.NewProc("RenderBundleEncoderSetVertexBuffer")
	renderBundleEncoderSetIndexBufferProc            = webgpuDLL.NewProc("RenderBundleEncoderSetIndexBuffer")
	renderBundleEncoderFinishProc                    = webgpuDLL.NewProc("RenderBundleEncoderFinish")
	renderBundleEncoderSetLabelProc                  = webgpuDLL.NewProc("RenderBundleEncoderSetLabel")
	renderBundleEncoderAddRefProc                    = webgpuDLL.NewProc("RenderBundleEncoderAddRef")
	renderBundleEncoderReleaseProc                   = webgpuDLL.NewProc("RenderBundleEncoderRelease")
	renderPassEncoderSetPipelineProc                 = webgpuDLL.NewProc("RenderPassEncoderSetPipeline")
	renderPassEncoderSetBindGroupProc                = webgpuDLL.NewProc("RenderPassEncoderSetBindGroup")
	renderPassEncoderSetImmediatesProc               = webgpuDLL.NewProc("RenderPassEncoderSetImmediates")
	renderPassEncoderDrawProc                        = webgpuDLL.NewProc("RenderPassEncoderDraw")
	renderPassEncoderDrawIndexedProc                 = webgpuDLL.NewProc("RenderPassEncoderDrawIndexed")
	renderPassEncoderDrawIndirectProc                = webgpuDLL.NewProc("RenderPassEncoderDrawIndirect")
	renderPassEncoderDrawIndexedIndirectProc         = webgpuDLL.NewProc("RenderPassEncoderDrawIndexedIndirect")
	renderPassEncoderExecuteBundlesProc              = webgpuDLL.NewProc("RenderPassEncoderExecuteBundles")
	renderPassEncoderInsertDebugMarkerProc           = webgpuDLL.NewProc("RenderPassEncoderInsertDebugMarker")
	renderPassEncoderPopDebugGroupProc               = webgpuDLL.NewProc("RenderPassEncoderPopDebugGroup")
	renderPassEncoderPushDebugGroupProc              = webgpuDLL.NewProc("RenderPassEncoderPushDebugGroup")
	renderPassEncoderSetStencilReferenceProc         = webgpuDLL.NewProc("RenderPassEncoderSetStencilReference")
	renderPassEncoderSetBlendConstantProc            = webgpuDLL.NewProc("RenderPassEncoderSetBlendConstant")
	renderPassEncoderSetViewportProc                 = webgpuDLL.NewProc("RenderPassEncoderSetViewport")
	renderPassEncoderSetScissorRectProc              = webgpuDLL.NewProc("RenderPassEncoderSetScissorRect")
	renderPassEncoderSetVertexBufferProc             = webgpuDLL.NewProc("RenderPassEncoderSetVertexBuffer")
	renderPassEncoderSetIndexBufferProc              = webgpuDLL.NewProc("RenderPassEncoderSetIndexBuffer")
	renderPassEncoderBeginOcclusionQueryProc         = webgpuDLL.NewProc("RenderPassEncoderBeginOcclusionQuery")
	renderPassEncoderEndOcclusionQueryProc           = webgpuDLL.NewProc("RenderPassEncoderEndOcclusionQuery")
	renderPassEncoderEndProc                         = webgpuDLL.NewProc("RenderPassEncoderEnd")
	renderPassEncoderSetLabelProc                    = webgpuDLL.NewProc("RenderPassEncoderSetLabel")
	renderPassEncoderAddRefProc                      = webgpuDLL.NewProc("RenderPassEncoderAddRef")
	renderPassEncoderReleaseProc                     = webgpuDLL.NewProc("RenderPassEncoderRelease")
	renderPipelineGetBindGroupLayoutProc             = webgpuDLL.NewProc("RenderPipelineGetBindGroupLayout")
	renderPipelineSetLabelProc                       = webgpuDLL.NewProc("RenderPipelineSetLabel")
	renderPipelineAddRefProc                         = webgpuDLL.NewProc("RenderPipelineAddRef")
	renderPipelineReleaseProc                        = webgpuDLL.NewProc("RenderPipelineRelease")
	samplerSetLabelProc                              = webgpuDLL.NewProc("SamplerSetLabel")
	samplerAddRefProc                                = webgpuDLL.NewProc("SamplerAddRef")
	samplerReleaseProc                               = webgpuDLL.NewProc("SamplerRelease")
	shaderModuleGetCompilationInfoProc               = webgpuDLL.NewProc("ShaderModuleGetCompilationInfo")
	shaderModuleSetLabelProc                         = webgpuDLL.NewProc("ShaderModuleSetLabel")
	shaderModuleAddRefProc                           = webgpuDLL.NewProc("ShaderModuleAddRef")
	shaderModuleReleaseProc                          = webgpuDLL.NewProc("ShaderModuleRelease")
	surfaceConfigureProc                             = webgpuDLL.NewProc("SurfaceConfigure")
	surfaceGetCapabilitiesProc                       = webgpuDLL.NewProc("SurfaceGetCapabilities")
	surfaceGetCurrentTextureProc                     = webgpuDLL.NewProc("SurfaceGetCurrentTexture")
	surfacePresentProc                               = webgpuDLL.NewProc("SurfacePresent")
	surfaceUnconfigureProc                           = webgpuDLL.NewProc("SurfaceUnconfigure")
	surfaceSetLabelProc                              = webgpuDLL.NewProc("SurfaceSetLabel")
	surfaceAddRefProc                                = webgpuDLL.NewProc("SurfaceAddRef")
	surfaceReleaseProc                               = webgpuDLL.NewProc("SurfaceRelease")
	textureCreateViewProc                            = webgpuDLL.NewProc("TextureCreateView")
	textureSetLabelProc                              = webgpuDLL.NewProc("TextureSetLabel")
	textureGetWidthProc                              = webgpuDLL.NewProc("TextureGetWidth")
	textureGetHeightProc                             = webgpuDLL.NewProc("TextureGetHeight")
	textureGetDepthOrArrayLayersProc                 = webgpuDLL.NewProc("TextureGetDepthOrArrayLayers")
	textureGetMipLevelCountProc                      = webgpuDLL.NewProc("TextureGetMipLevelCount")
	textureGetSampleCountProc                        = webgpuDLL.NewProc("TextureGetSampleCount")
	textureGetDimensionProc                          = webgpuDLL.NewProc("TextureGetDimension")
	textureGetTextureBindingViewDimensionProc        = webgpuDLL.NewProc("TextureGetTextureBindingViewDimension")
	textureGetFormatProc                             = webgpuDLL.NewProc("TextureGetFormat")
	textureGetUsageProc                              = webgpuDLL.NewProc("TextureGetUsage")
	textureDestroyProc                               = webgpuDLL.NewProc("TextureDestroy")
	textureAddRefProc                                = webgpuDLL.NewProc("TextureAddRef")
	textureReleaseProc                               = webgpuDLL.NewProc("TextureRelease")
	textureViewSetLabelProc                          = webgpuDLL.NewProc("TextureViewSetLabel")
	textureViewAddRefProc                            = webgpuDLL.NewProc("TextureViewAddRef")
	textureViewReleaseProc                           = webgpuDLL.NewProc("TextureViewRelease")
)

func AdapterGetLimits(adapter Adapter, limits *Limits) Status {

	_, _, err := adapterGetLimitsProc.Call()
	if err != nil {
		panic(err)
	}
}

func AdapterHasFeature(adapter Adapter, feature FeatureName) Bool {

	_, _, err := adapterHasFeatureProc.Call()
	if err != nil {
		panic(err)
	}
}

func AdapterGetFeatures(adapter Adapter, features *SupportedFeatures) {

	_, _, err := adapterGetFeaturesProc.Call()
	if err != nil {
		panic(err)
	}
}

func AdapterGetInfo(adapter Adapter, info *AdapterInfo) Status {

	_, _, err := adapterGetInfoProc.Call()
	if err != nil {
		panic(err)
	}
}

func AdapterRequestDevice(adapter Adapter, descriptor *DeviceDescriptor, callback RequestDeviceCallbackInfo) {

	_, _, err := adapterRequestDeviceProc.Call()
	if err != nil {
		panic(err)
	}
}

func AdapterAddRef(adapter Adapter) {
	adapterAddRefProc.Call(uintptr(adapter))
}

func AdapterRelease(adapter Adapter) {
	adapterReleaseProc.Call(uintptr(adapter))
}

func BindGroupSetLabel(bindGroup BindGroup, label StringView) {

	_, _, err := bindGroupSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func BindGroupAddRef(bindGroup BindGroup) {
	bindGroupAddRefProc.Call(uintptr(bindGroup))
}

func BindGroupRelease(bindGroup BindGroup) {
	bindGroupReleaseProc.Call(uintptr(bindGroup))
}

func BindGroupLayoutSetLabel(bindGroupLayout BindGroupLayout, label StringView) {

	_, _, err := bindGroupLayoutSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func BindGroupLayoutAddRef(bindGroupLayout BindGroupLayout) {
	bindGroupLayoutAddRefProc.Call(uintptr(bindGroupLayout))
}

func BindGroupLayoutRelease(bindGroupLayout BindGroupLayout) {
	bindGroupLayoutReleaseProc.Call(uintptr(bindGroupLayout))
}

func BufferMapAsync(buffer Buffer, mode MapMode, offset uintptr, size uintptr, callback BufferMapCallbackInfo) {

	_, _, err := bufferMapAsyncProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferGetMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {

	_, _, err := bufferGetMappedRangeProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferGetConstMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {

	_, _, err := bufferGetConstMappedRangeProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferReadMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {

	_, _, err := bufferReadMappedRangeProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferWriteMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {

	_, _, err := bufferWriteMappedRangeProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferSetLabel(buffer Buffer, label StringView) {

	_, _, err := bufferSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferGetUsage(buffer Buffer) BufferUsage {

	_, _, err := bufferGetUsageProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferGetSize(buffer Buffer) uint64 {

	_, _, err := bufferGetSizeProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferGetMapState(buffer Buffer) BufferMapState {

	_, _, err := bufferGetMapStateProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferUnmap(buffer Buffer) {

	_, _, err := bufferUnmapProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferDestroy(buffer Buffer) {

	_, _, err := bufferDestroyProc.Call()
	if err != nil {
		panic(err)
	}
}

func BufferAddRef(buffer Buffer) {
	bufferAddRefProc.Call(uintptr(buffer))
}

func BufferRelease(buffer Buffer) {
	bufferReleaseProc.Call(uintptr(buffer))
}

func CommandBufferSetLabel(commandBuffer CommandBuffer, label StringView) {

	_, _, err := commandBufferSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandBufferAddRef(commandBuffer CommandBuffer) {
	commandBufferAddRefProc.Call(uintptr(commandBuffer))
}

func CommandBufferRelease(commandBuffer CommandBuffer) {
	commandBufferReleaseProc.Call(uintptr(commandBuffer))
}

func CommandEncoderFinish(commandEncoder CommandEncoder, descriptor *CommandBufferDescriptor) CommandBuffer {

	_, _, err := commandEncoderFinishProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderBeginComputePass(commandEncoder CommandEncoder, descriptor *ComputePassDescriptor) ComputePassEncoder {

	_, _, err := commandEncoderBeginComputePassProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderBeginRenderPass(commandEncoder CommandEncoder, descriptor *RenderPassDescriptor) RenderPassEncoder {

	_, _, err := commandEncoderBeginRenderPassProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderCopyBufferToBuffer(commandEncoder CommandEncoder, source Buffer, sourceOffset uint64, destination Buffer, destinationOffset uint64, size uint64) {

	_, _, err := commandEncoderCopyBufferToBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderCopyBufferToTexture(commandEncoder CommandEncoder, source *TexelCopyBufferInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {

	_, _, err := commandEncoderCopyBufferToTextureProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderCopyTextureToBuffer(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyBufferInfo, copySize *Extent3D) {

	_, _, err := commandEncoderCopyTextureToBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderCopyTextureToTexture(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {

	_, _, err := commandEncoderCopyTextureToTextureProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderClearBuffer(commandEncoder CommandEncoder, buffer Buffer, offset uint64, size uint64) {

	_, _, err := commandEncoderClearBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderInsertDebugMarker(commandEncoder CommandEncoder, markerLabel StringView) {

	_, _, err := commandEncoderInsertDebugMarkerProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderPopDebugGroup(commandEncoder CommandEncoder) {

	_, _, err := commandEncoderPopDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderPushDebugGroup(commandEncoder CommandEncoder, groupLabel StringView) {

	_, _, err := commandEncoderPushDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderResolveQuerySet(commandEncoder CommandEncoder, querySet QuerySet, firstQuery uint32, queryCount uint32, destination Buffer, destinationOffset uint64) {

	_, _, err := commandEncoderResolveQuerySetProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderWriteTimestamp(commandEncoder CommandEncoder, querySet QuerySet, queryIndex uint32) {

	_, _, err := commandEncoderWriteTimestampProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderSetLabel(commandEncoder CommandEncoder, label StringView) {

	_, _, err := commandEncoderSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func CommandEncoderAddRef(commandEncoder CommandEncoder) {
	commandEncoderAddRefProc.Call(uintptr(commandEncoder))
}

func CommandEncoderRelease(commandEncoder CommandEncoder) {
	commandEncoderReleaseProc.Call(uintptr(commandEncoder))
}

func ComputePassEncoderInsertDebugMarker(computePassEncoder ComputePassEncoder, markerLabel StringView) {

	_, _, err := computePassEncoderInsertDebugMarkerProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderPopDebugGroup(computePassEncoder ComputePassEncoder) {

	_, _, err := computePassEncoderPopDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderPushDebugGroup(computePassEncoder ComputePassEncoder, groupLabel StringView) {

	_, _, err := computePassEncoderPushDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetPipeline(computePassEncoder ComputePassEncoder, pipeline ComputePipeline) {

	_, _, err := computePassEncoderSetPipelineProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetBindGroup(computePassEncoder ComputePassEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {

	_, _, err := computePassEncoderSetBindGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetImmediates(computePassEncoder ComputePassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {

	_, _, err := computePassEncoderSetImmediatesProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderDispatchWorkgroups(computePassEncoder ComputePassEncoder, workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {

	_, _, err := computePassEncoderDispatchWorkgroupsProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderDispatchWorkgroupsIndirect(computePassEncoder ComputePassEncoder, indirectBuffer Buffer, indirectOffset uint64) {

	_, _, err := computePassEncoderDispatchWorkgroupsIndirectProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderEnd(computePassEncoder ComputePassEncoder) {

	_, _, err := computePassEncoderEndProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetLabel(computePassEncoder ComputePassEncoder, label StringView) {

	_, _, err := computePassEncoderSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderAddRef(computePassEncoder ComputePassEncoder) {
	computePassEncoderAddRefProc.Call(uintptr(computePassEncoder))
}

func ComputePassEncoderRelease(computePassEncoder ComputePassEncoder) {
	computePassEncoderReleaseProc.Call(uintptr(computePassEncoder))
}

func ComputePipelineGetBindGroupLayout(computePipeline ComputePipeline, groupIndex uint32) BindGroupLayout {

	_, _, err := computePipelineGetBindGroupLayoutProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePipelineSetLabel(computePipeline ComputePipeline, label StringView) {

	_, _, err := computePipelineSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func ComputePipelineAddRef(computePipeline ComputePipeline) {
	computePipelineAddRefProc.Call(uintptr(computePipeline))
}

func ComputePipelineRelease(computePipeline ComputePipeline) {
	computePipelineReleaseProc.Call(uintptr(computePipeline))
}

func DeviceCreateBindGroup(device Device, descriptor *BindGroupDescriptor) BindGroup {

	_, _, err := deviceCreateBindGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateBindGroupLayout(device Device, descriptor *BindGroupLayoutDescriptor) BindGroupLayout {

	_, _, err := deviceCreateBindGroupLayoutProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateBuffer(device Device, descriptor *BufferDescriptor) Buffer {

	_, _, err := deviceCreateBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateCommandEncoder(device Device, descriptor *CommandEncoderDescriptor) CommandEncoder {

	_, _, err := deviceCreateCommandEncoderProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateComputePipeline(device Device, descriptor *ComputePipelineDescriptor) ComputePipeline {

	_, _, err := deviceCreateComputePipelineProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateComputePipelineAsync(device Device, descriptor *ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallbackInfo) {

	_, _, err := deviceCreateComputePipelineAsyncProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreatePipelineLayout(device Device, descriptor *PipelineLayoutDescriptor) PipelineLayout {

	_, _, err := deviceCreatePipelineLayoutProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateQuerySet(device Device, descriptor *QuerySetDescriptor) QuerySet {

	_, _, err := deviceCreateQuerySetProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateRenderPipelineAsync(device Device, descriptor *RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallbackInfo) {

	_, _, err := deviceCreateRenderPipelineAsyncProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateRenderBundleEncoder(device Device, descriptor *RenderBundleEncoderDescriptor) RenderBundleEncoder {

	_, _, err := deviceCreateRenderBundleEncoderProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateRenderPipeline(device Device, descriptor *RenderPipelineDescriptor) RenderPipeline {

	_, _, err := deviceCreateRenderPipelineProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateSampler(device Device, descriptor *SamplerDescriptor) Sampler {

	_, _, err := deviceCreateSamplerProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateShaderModule(device Device, descriptor *ShaderModuleDescriptor) ShaderModule {

	_, _, err := deviceCreateShaderModuleProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceCreateTexture(device Device, descriptor *TextureDescriptor) Texture {

	_, _, err := deviceCreateTextureProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceDestroy(device Device) {

	_, _, err := deviceDestroyProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceGetLostFuture(device Device) Future {

	_, _, err := deviceGetLostFutureProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceGetLimits(device Device, limits *Limits) Status {

	_, _, err := deviceGetLimitsProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceHasFeature(device Device, feature FeatureName) Bool {

	_, _, err := deviceHasFeatureProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceGetFeatures(device Device, features *SupportedFeatures) {

	_, _, err := deviceGetFeaturesProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceGetAdapterInfo(device Device, adapterInfo *AdapterInfo) Status {

	_, _, err := deviceGetAdapterInfoProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceGetQueue(device Device) Queue {

	_, _, err := deviceGetQueueProc.Call()
	if err != nil {
		panic(err)
	}
}

func DevicePushErrorScope(device Device, filter ErrorFilter) {

	_, _, err := devicePushErrorScopeProc.Call()
	if err != nil {
		panic(err)
	}
}

func DevicePopErrorScope(device Device, callback PopErrorScopeCallbackInfo) {

	_, _, err := devicePopErrorScopeProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceSetLabel(device Device, label StringView) {

	_, _, err := deviceSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func DeviceAddRef(device Device) {
	deviceAddRefProc.Call(uintptr(device))
}

func DeviceRelease(device Device) {
	deviceReleaseProc.Call(uintptr(device))
}

func ExternalTextureSetLabel(externalTexture ExternalTexture, label StringView) {

	_, _, err := externalTextureSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func ExternalTextureAddRef(externalTexture ExternalTexture) {
	externalTextureAddRefProc.Call(uintptr(externalTexture))
}

func ExternalTextureRelease(externalTexture ExternalTexture) {
	externalTextureReleaseProc.Call(uintptr(externalTexture))
}

func InstanceCreateSurface(instance Instance, descriptor *SurfaceDescriptor) Surface {

	_, _, err := instanceCreateSurfaceProc.Call()
	if err != nil {
		panic(err)
	}
}

func InstanceGetWGSLLanguageFeatures(instance Instance, features *SupportedWGSLLanguageFeatures) {

	_, _, err := instanceGetWGSLLanguageFeaturesProc.Call()
	if err != nil {
		panic(err)
	}
}

func InstanceHasWGSLLanguageFeature(instance Instance, feature WGSLLanguageFeatureName) Bool {

	_, _, err := instanceHasWGSLLanguageFeatureProc.Call()
	if err != nil {
		panic(err)
	}
}

func InstanceProcessEvents(instance Instance) {

	_, _, err := instanceProcessEventsProc.Call()
	if err != nil {
		panic(err)
	}
}

func InstanceRequestAdapter(instance Instance, options *RequestAdapterOptions, callback RequestAdapterCallbackInfo) {

	_, _, err := instanceRequestAdapterProc.Call()
	if err != nil {
		panic(err)
	}
}

func InstanceWaitAny(instance Instance, futureCount uintptr, futures *FutureWaitInfo, timeoutNS uint64) WaitStatus {

	_, _, err := instanceWaitAnyProc.Call()
	if err != nil {
		panic(err)
	}
}

func InstanceAddRef(instance Instance) {
	instanceAddRefProc.Call(uintptr(instance))
}

func InstanceRelease(instance Instance) {
	instanceReleaseProc.Call(uintptr(instance))
}

func PipelineLayoutSetLabel(pipelineLayout PipelineLayout, label StringView) {

	_, _, err := pipelineLayoutSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func PipelineLayoutAddRef(pipelineLayout PipelineLayout) {
	pipelineLayoutAddRefProc.Call(uintptr(pipelineLayout))
}

func PipelineLayoutRelease(pipelineLayout PipelineLayout) {
	pipelineLayoutReleaseProc.Call(uintptr(pipelineLayout))
}

func QuerySetSetLabel(querySet QuerySet, label StringView) {

	_, _, err := querySetSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func QuerySetGetType(querySet QuerySet) QueryType {

	_, _, err := querySetGetTypeProc.Call()
	if err != nil {
		panic(err)
	}
}

func QuerySetGetCount(querySet QuerySet) uint32 {

	_, _, err := querySetGetCountProc.Call()
	if err != nil {
		panic(err)
	}
}

func QuerySetDestroy(querySet QuerySet) {

	_, _, err := querySetDestroyProc.Call()
	if err != nil {
		panic(err)
	}
}

func QuerySetAddRef(querySet QuerySet) {
	querySetAddRefProc.Call(uintptr(querySet))
}

func QuerySetRelease(querySet QuerySet) {
	querySetReleaseProc.Call(uintptr(querySet))
}

func QueueSubmit(queue Queue, commands []CommandBuffer) {

	_, _, err := queueSubmitProc.Call()
	if err != nil {
		panic(err)
	}
}

func QueueOnSubmittedWorkDone(queue Queue, callback QueueWorkDoneCallbackInfo) {

	_, _, err := queueOnSubmittedWorkDoneProc.Call()
	if err != nil {
		panic(err)
	}
}

func QueueWriteBuffer(queue Queue, buffer Buffer, bufferOffset uint64, data unsafe.Pointer, size uintptr) {

	_, _, err := queueWriteBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func QueueWriteTexture(queue Queue, destination *TexelCopyTextureInfo, data unsafe.Pointer, dataSize uintptr, dataLayout *TexelCopyBufferLayout, writeSize *Extent3D) {

	_, _, err := queueWriteTextureProc.Call()
	if err != nil {
		panic(err)
	}
}

func QueueSetLabel(queue Queue, label StringView) {

	_, _, err := queueSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func QueueAddRef(queue Queue) {
	queueAddRefProc.Call(uintptr(queue))
}

func QueueRelease(queue Queue) {
	queueReleaseProc.Call(uintptr(queue))
}

func RenderBundleSetLabel(renderBundle RenderBundle, label StringView) {

	_, _, err := renderBundleSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleAddRef(renderBundle RenderBundle) {
	renderBundleAddRefProc.Call(uintptr(renderBundle))
}

func RenderBundleRelease(renderBundle RenderBundle) {
	renderBundleReleaseProc.Call(uintptr(renderBundle))
}

func RenderBundleEncoderSetPipeline(renderBundleEncoder RenderBundleEncoder, pipeline RenderPipeline) {

	_, _, err := renderBundleEncoderSetPipelineProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetBindGroup(renderBundleEncoder RenderBundleEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {

	_, _, err := renderBundleEncoderSetBindGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetImmediates(renderBundleEncoder RenderBundleEncoder, offset uint32, data unsafe.Pointer, size uintptr) {

	_, _, err := renderBundleEncoderSetImmediatesProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDraw(renderBundleEncoder RenderBundleEncoder, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {

	_, _, err := renderBundleEncoderDrawProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDrawIndexed(renderBundleEncoder RenderBundleEncoder, indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {

	_, _, err := renderBundleEncoderDrawIndexedProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDrawIndirect(renderBundleEncoder RenderBundleEncoder, indirectBuffer Buffer, indirectOffset uint64) {

	_, _, err := renderBundleEncoderDrawIndirectProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDrawIndexedIndirect(renderBundleEncoder RenderBundleEncoder, indirectBuffer Buffer, indirectOffset uint64) {

	_, _, err := renderBundleEncoderDrawIndexedIndirectProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderInsertDebugMarker(renderBundleEncoder RenderBundleEncoder, markerLabel StringView) {

	_, _, err := renderBundleEncoderInsertDebugMarkerProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderPopDebugGroup(renderBundleEncoder RenderBundleEncoder) {

	_, _, err := renderBundleEncoderPopDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderPushDebugGroup(renderBundleEncoder RenderBundleEncoder, groupLabel StringView) {

	_, _, err := renderBundleEncoderPushDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetVertexBuffer(renderBundleEncoder RenderBundleEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {

	_, _, err := renderBundleEncoderSetVertexBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetIndexBuffer(renderBundleEncoder RenderBundleEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {

	_, _, err := renderBundleEncoderSetIndexBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderFinish(renderBundleEncoder RenderBundleEncoder, descriptor *RenderBundleDescriptor) RenderBundle {

	_, _, err := renderBundleEncoderFinishProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetLabel(renderBundleEncoder RenderBundleEncoder, label StringView) {

	_, _, err := renderBundleEncoderSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderAddRef(renderBundleEncoder RenderBundleEncoder) {
	renderBundleEncoderAddRefProc.Call(uintptr(renderBundleEncoder))
}

func RenderBundleEncoderRelease(renderBundleEncoder RenderBundleEncoder) {
	renderBundleEncoderReleaseProc.Call(uintptr(renderBundleEncoder))
}

func RenderPassEncoderSetPipeline(renderPassEncoder RenderPassEncoder, pipeline RenderPipeline) {

	_, _, err := renderPassEncoderSetPipelineProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetBindGroup(renderPassEncoder RenderPassEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {

	_, _, err := renderPassEncoderSetBindGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetImmediates(renderPassEncoder RenderPassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {

	_, _, err := renderPassEncoderSetImmediatesProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDraw(renderPassEncoder RenderPassEncoder, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {

	_, _, err := renderPassEncoderDrawProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDrawIndexed(renderPassEncoder RenderPassEncoder, indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {

	_, _, err := renderPassEncoderDrawIndexedProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDrawIndirect(renderPassEncoder RenderPassEncoder, indirectBuffer Buffer, indirectOffset uint64) {

	_, _, err := renderPassEncoderDrawIndirectProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDrawIndexedIndirect(renderPassEncoder RenderPassEncoder, indirectBuffer Buffer, indirectOffset uint64) {

	_, _, err := renderPassEncoderDrawIndexedIndirectProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderExecuteBundles(renderPassEncoder RenderPassEncoder, bundles []RenderBundle) {

	_, _, err := renderPassEncoderExecuteBundlesProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderInsertDebugMarker(renderPassEncoder RenderPassEncoder, markerLabel StringView) {

	_, _, err := renderPassEncoderInsertDebugMarkerProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderPopDebugGroup(renderPassEncoder RenderPassEncoder) {

	_, _, err := renderPassEncoderPopDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderPushDebugGroup(renderPassEncoder RenderPassEncoder, groupLabel StringView) {

	_, _, err := renderPassEncoderPushDebugGroupProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetStencilReference(renderPassEncoder RenderPassEncoder, reference uint32) {

	_, _, err := renderPassEncoderSetStencilReferenceProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetBlendConstant(renderPassEncoder RenderPassEncoder, color *Color) {

	_, _, err := renderPassEncoderSetBlendConstantProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetViewport(renderPassEncoder RenderPassEncoder, x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) {

	_, _, err := renderPassEncoderSetViewportProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetScissorRect(renderPassEncoder RenderPassEncoder, x uint32, y uint32, width uint32, height uint32) {

	_, _, err := renderPassEncoderSetScissorRectProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetVertexBuffer(renderPassEncoder RenderPassEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {

	_, _, err := renderPassEncoderSetVertexBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetIndexBuffer(renderPassEncoder RenderPassEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {

	_, _, err := renderPassEncoderSetIndexBufferProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderBeginOcclusionQuery(renderPassEncoder RenderPassEncoder, queryIndex uint32) {

	_, _, err := renderPassEncoderBeginOcclusionQueryProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderEndOcclusionQuery(renderPassEncoder RenderPassEncoder) {

	_, _, err := renderPassEncoderEndOcclusionQueryProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderEnd(renderPassEncoder RenderPassEncoder) {

	_, _, err := renderPassEncoderEndProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetLabel(renderPassEncoder RenderPassEncoder, label StringView) {

	_, _, err := renderPassEncoderSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderAddRef(renderPassEncoder RenderPassEncoder) {
	renderPassEncoderAddRefProc.Call(uintptr(renderPassEncoder))
}

func RenderPassEncoderRelease(renderPassEncoder RenderPassEncoder) {
	renderPassEncoderReleaseProc.Call(uintptr(renderPassEncoder))
}

func RenderPipelineGetBindGroupLayout(renderPipeline RenderPipeline, groupIndex uint32) BindGroupLayout {

	_, _, err := renderPipelineGetBindGroupLayoutProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPipelineSetLabel(renderPipeline RenderPipeline, label StringView) {

	_, _, err := renderPipelineSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func RenderPipelineAddRef(renderPipeline RenderPipeline) {
	renderPipelineAddRefProc.Call(uintptr(renderPipeline))
}

func RenderPipelineRelease(renderPipeline RenderPipeline) {
	renderPipelineReleaseProc.Call(uintptr(renderPipeline))
}

func SamplerSetLabel(sampler Sampler, label StringView) {

	_, _, err := samplerSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func SamplerAddRef(sampler Sampler) {
	samplerAddRefProc.Call(uintptr(sampler))
}

func SamplerRelease(sampler Sampler) {
	samplerReleaseProc.Call(uintptr(sampler))
}

func ShaderModuleGetCompilationInfo(shaderModule ShaderModule, callback CompilationInfoCallbackInfo) {

	_, _, err := shaderModuleGetCompilationInfoProc.Call()
	if err != nil {
		panic(err)
	}
}

func ShaderModuleSetLabel(shaderModule ShaderModule, label StringView) {

	_, _, err := shaderModuleSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func ShaderModuleAddRef(shaderModule ShaderModule) {
	shaderModuleAddRefProc.Call(uintptr(shaderModule))
}

func ShaderModuleRelease(shaderModule ShaderModule) {
	shaderModuleReleaseProc.Call(uintptr(shaderModule))
}

func SurfaceConfigure(surface Surface, config *SurfaceConfiguration) {

	_, _, err := surfaceConfigureProc.Call()
	if err != nil {
		panic(err)
	}
}

func SurfaceGetCapabilities(surface Surface, adapter Adapter, capabilities *SurfaceCapabilities) Status {

	_, _, err := surfaceGetCapabilitiesProc.Call()
	if err != nil {
		panic(err)
	}
}

func SurfaceGetCurrentTexture(surface Surface, surfaceTexture *SurfaceTexture) {

	_, _, err := surfaceGetCurrentTextureProc.Call()
	if err != nil {
		panic(err)
	}
}

func SurfacePresent(surface Surface) Status {

	_, _, err := surfacePresentProc.Call()
	if err != nil {
		panic(err)
	}
}

func SurfaceUnconfigure(surface Surface) {

	_, _, err := surfaceUnconfigureProc.Call()
	if err != nil {
		panic(err)
	}
}

func SurfaceSetLabel(surface Surface, label StringView) {

	_, _, err := surfaceSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func SurfaceAddRef(surface Surface) {
	surfaceAddRefProc.Call(uintptr(surface))
}

func SurfaceRelease(surface Surface) {
	surfaceReleaseProc.Call(uintptr(surface))
}

func TextureCreateView(texture Texture, descriptor *TextureViewDescriptor) TextureView {

	_, _, err := textureCreateViewProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureSetLabel(texture Texture, label StringView) {

	_, _, err := textureSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetWidth(texture Texture) uint32 {

	_, _, err := textureGetWidthProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetHeight(texture Texture) uint32 {

	_, _, err := textureGetHeightProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetDepthOrArrayLayers(texture Texture) uint32 {

	_, _, err := textureGetDepthOrArrayLayersProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetMipLevelCount(texture Texture) uint32 {

	_, _, err := textureGetMipLevelCountProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetSampleCount(texture Texture) uint32 {

	_, _, err := textureGetSampleCountProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetDimension(texture Texture) TextureDimension {

	_, _, err := textureGetDimensionProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetTextureBindingViewDimension(texture Texture) TextureViewDimension {

	_, _, err := textureGetTextureBindingViewDimensionProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetFormat(texture Texture) TextureFormat {

	_, _, err := textureGetFormatProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureGetUsage(texture Texture) TextureUsage {

	_, _, err := textureGetUsageProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureDestroy(texture Texture) {

	_, _, err := textureDestroyProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureAddRef(texture Texture) {
	textureAddRefProc.Call(uintptr(texture))
}

func TextureRelease(texture Texture) {
	textureReleaseProc.Call(uintptr(texture))
}

func TextureViewSetLabel(textureView TextureView, label StringView) {

	_, _, err := textureViewSetLabelProc.Call()
	if err != nil {
		panic(err)
	}
}

func TextureViewAddRef(textureView TextureView) {
	textureViewAddRefProc.Call(uintptr(textureView))
}

func TextureViewRelease(textureView TextureView) {
	textureViewReleaseProc.Call(uintptr(textureView))
}
