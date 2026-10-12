// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build !cgo && amd64

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
	ret, _, err := adapterGetLimitsProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(limits)),
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func AdapterHasFeature(adapter Adapter, feature FeatureName) Bool {
	ret, _, err := adapterHasFeatureProc.Call(
		uintptr(adapter),
		uintptr(feature), // enum.feature_name
	)
	if err != nil {
		panic(err)
	}
	return Bool(ret)
}

func AdapterGetFeatures(adapter Adapter, features *SupportedFeatures) {
	_, _, err := adapterGetFeaturesProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(features)),
	)
	if err != nil {
		panic(err)
	}
}

func AdapterGetInfo(adapter Adapter, info *AdapterInfo) Status {
	ret, _, err := adapterGetInfoProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(info)),
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func AdapterRequestDevice(adapter Adapter, descriptor *DeviceDescriptor, callback RequestDeviceCallbackInfo) {
	_, _, err := adapterRequestDeviceProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(descriptor)),
	)
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
	_, _, err := bindGroupSetLabelProc.Call(
		uintptr(bindGroup),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := bindGroupLayoutSetLabelProc.Call(
		uintptr(bindGroupLayout),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := bufferMapAsyncProc.Call(
		uintptr(buffer),
		uintptr(mode),   // bitflag.map_mode
		uintptr(offset), // usize
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
}

func BufferGetMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {
	ret, _, err := bufferGetMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(ret)
}

func BufferGetConstMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {
	ret, _, err := bufferGetConstMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(ret)
}

func BufferReadMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {
	ret, _, err := bufferReadMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func BufferWriteMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {
	ret, _, err := bufferWriteMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func BufferSetLabel(buffer Buffer, label StringView) {
	_, _, err := bufferSetLabelProc.Call(
		uintptr(buffer),
		uintptr(unsafe.Pointer(&label)),
	)
	if err != nil {
		panic(err)
	}
}

func BufferGetUsage(buffer Buffer) BufferUsage {
	ret, _, err := bufferGetUsageProc.Call(
		uintptr(buffer),
	)
	if err != nil {
		panic(err)
	}
	return BufferUsage(ret)
}

func BufferGetSize(buffer Buffer) uint64 {
	ret, _, err := bufferGetSizeProc.Call(
		uintptr(buffer),
	)
	if err != nil {
		panic(err)
	}
	return uint64(ret)
}

func BufferGetMapState(buffer Buffer) BufferMapState {
	ret, _, err := bufferGetMapStateProc.Call(
		uintptr(buffer),
	)
	if err != nil {
		panic(err)
	}
	return BufferMapState(ret)
}

func BufferUnmap(buffer Buffer) {
	_, _, err := bufferUnmapProc.Call(
		uintptr(buffer),
	)
	if err != nil {
		panic(err)
	}
}

func BufferDestroy(buffer Buffer) {
	_, _, err := bufferDestroyProc.Call(
		uintptr(buffer),
	)
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
	_, _, err := commandBufferSetLabelProc.Call(
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(&label)),
	)
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
	ret, _, err := commandEncoderFinishProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return CommandBuffer(ret)
}

func CommandEncoderBeginComputePass(commandEncoder CommandEncoder, descriptor *ComputePassDescriptor) ComputePassEncoder {
	ret, _, err := commandEncoderBeginComputePassProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return ComputePassEncoder(ret)
}

func CommandEncoderBeginRenderPass(commandEncoder CommandEncoder, descriptor *RenderPassDescriptor) RenderPassEncoder {
	ret, _, err := commandEncoderBeginRenderPassProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return RenderPassEncoder(ret)
}

func CommandEncoderCopyBufferToBuffer(commandEncoder CommandEncoder, source Buffer, sourceOffset uint64, destination Buffer, destinationOffset uint64, size uint64) {
	_, _, err := commandEncoderCopyBufferToBufferProc.Call(
		uintptr(commandEncoder),
		uintptr(source),            // object.buffer
		uintptr(sourceOffset),      // uint64
		uintptr(destination),       // object.buffer
		uintptr(destinationOffset), // uint64
		uintptr(size),              // uint64
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderCopyBufferToTexture(commandEncoder CommandEncoder, source *TexelCopyBufferInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {
	_, _, err := commandEncoderCopyBufferToTextureProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(copySize)),
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderCopyTextureToBuffer(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyBufferInfo, copySize *Extent3D) {
	_, _, err := commandEncoderCopyTextureToBufferProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(copySize)),
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderCopyTextureToTexture(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {
	_, _, err := commandEncoderCopyTextureToTextureProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(copySize)),
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderClearBuffer(commandEncoder CommandEncoder, buffer Buffer, offset uint64, size uint64) {
	_, _, err := commandEncoderClearBufferProc.Call(
		uintptr(commandEncoder),
		uintptr(buffer), // object.buffer
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderInsertDebugMarker(commandEncoder CommandEncoder, markerLabel StringView) {
	_, _, err := commandEncoderInsertDebugMarkerProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderPopDebugGroup(commandEncoder CommandEncoder) {
	_, _, err := commandEncoderPopDebugGroupProc.Call(
		uintptr(commandEncoder),
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderPushDebugGroup(commandEncoder CommandEncoder, groupLabel StringView) {
	_, _, err := commandEncoderPushDebugGroupProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderResolveQuerySet(commandEncoder CommandEncoder, querySet QuerySet, firstQuery uint32, queryCount uint32, destination Buffer, destinationOffset uint64) {
	_, _, err := commandEncoderResolveQuerySetProc.Call(
		uintptr(commandEncoder),
		uintptr(querySet),          // object.query_set
		uintptr(firstQuery),        // uint32
		uintptr(queryCount),        // uint32
		uintptr(destination),       // object.buffer
		uintptr(destinationOffset), // uint64
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderWriteTimestamp(commandEncoder CommandEncoder, querySet QuerySet, queryIndex uint32) {
	_, _, err := commandEncoderWriteTimestampProc.Call(
		uintptr(commandEncoder),
		uintptr(querySet),   // object.query_set
		uintptr(queryIndex), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func CommandEncoderSetLabel(commandEncoder CommandEncoder, label StringView) {
	_, _, err := commandEncoderSetLabelProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := computePassEncoderInsertDebugMarkerProc.Call(
		uintptr(computePassEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderPopDebugGroup(computePassEncoder ComputePassEncoder) {
	_, _, err := computePassEncoderPopDebugGroupProc.Call(
		uintptr(computePassEncoder),
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderPushDebugGroup(computePassEncoder ComputePassEncoder, groupLabel StringView) {
	_, _, err := computePassEncoderPushDebugGroupProc.Call(
		uintptr(computePassEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetPipeline(computePassEncoder ComputePassEncoder, pipeline ComputePipeline) {
	_, _, err := computePassEncoderSetPipelineProc.Call(
		uintptr(computePassEncoder),
		uintptr(pipeline), // object.compute_pipeline
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetBindGroup(computePassEncoder ComputePassEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {
	_, _, err := computePassEncoderSetBindGroupProc.Call(
		uintptr(computePassEncoder),
		uintptr(groupIndex), // uint32
		uintptr(group),      // object.bind_group
		uintptr(len(dynamicOffsets)),
		uintptr(unsafe.Pointer(unsafe.SliceData(dynamicOffsets))),
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetImmediates(computePassEncoder ComputePassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	_, _, err := computePassEncoderSetImmediatesProc.Call(
		uintptr(computePassEncoder),
		uintptr(offset), // uint32
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderDispatchWorkgroups(computePassEncoder ComputePassEncoder, workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	_, _, err := computePassEncoderDispatchWorkgroupsProc.Call(
		uintptr(computePassEncoder),
		uintptr(workgroupCountX), // uint32
		uintptr(workgroupCountY), // uint32
		uintptr(workgroupCountZ), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderDispatchWorkgroupsIndirect(computePassEncoder ComputePassEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, err := computePassEncoderDispatchWorkgroupsIndirectProc.Call(
		uintptr(computePassEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderEnd(computePassEncoder ComputePassEncoder) {
	_, _, err := computePassEncoderEndProc.Call(
		uintptr(computePassEncoder),
	)
	if err != nil {
		panic(err)
	}
}

func ComputePassEncoderSetLabel(computePassEncoder ComputePassEncoder, label StringView) {
	_, _, err := computePassEncoderSetLabelProc.Call(
		uintptr(computePassEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
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
	ret, _, err := computePipelineGetBindGroupLayoutProc.Call(
		uintptr(computePipeline),
		uintptr(groupIndex), // uint32
	)
	if err != nil {
		panic(err)
	}
	return BindGroupLayout(ret)
}

func ComputePipelineSetLabel(computePipeline ComputePipeline, label StringView) {
	_, _, err := computePipelineSetLabelProc.Call(
		uintptr(computePipeline),
		uintptr(unsafe.Pointer(&label)),
	)
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
	ret, _, err := deviceCreateBindGroupProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return BindGroup(ret)
}

func DeviceCreateBindGroupLayout(device Device, descriptor *BindGroupLayoutDescriptor) BindGroupLayout {
	ret, _, err := deviceCreateBindGroupLayoutProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return BindGroupLayout(ret)
}

func DeviceCreateBuffer(device Device, descriptor *BufferDescriptor) Buffer {
	ret, _, err := deviceCreateBufferProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return Buffer(ret)
}

func DeviceCreateCommandEncoder(device Device, descriptor *CommandEncoderDescriptor) CommandEncoder {
	ret, _, err := deviceCreateCommandEncoderProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return CommandEncoder(ret)
}

func DeviceCreateComputePipeline(device Device, descriptor *ComputePipelineDescriptor) ComputePipeline {
	ret, _, err := deviceCreateComputePipelineProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return ComputePipeline(ret)
}

func DeviceCreateComputePipelineAsync(device Device, descriptor *ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallbackInfo) {
	_, _, err := deviceCreateComputePipelineAsyncProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
}

func DeviceCreatePipelineLayout(device Device, descriptor *PipelineLayoutDescriptor) PipelineLayout {
	ret, _, err := deviceCreatePipelineLayoutProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return PipelineLayout(ret)
}

func DeviceCreateQuerySet(device Device, descriptor *QuerySetDescriptor) QuerySet {
	ret, _, err := deviceCreateQuerySetProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return QuerySet(ret)
}

func DeviceCreateRenderPipelineAsync(device Device, descriptor *RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallbackInfo) {
	_, _, err := deviceCreateRenderPipelineAsyncProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
}

func DeviceCreateRenderBundleEncoder(device Device, descriptor *RenderBundleEncoderDescriptor) RenderBundleEncoder {
	ret, _, err := deviceCreateRenderBundleEncoderProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return RenderBundleEncoder(ret)
}

func DeviceCreateRenderPipeline(device Device, descriptor *RenderPipelineDescriptor) RenderPipeline {
	ret, _, err := deviceCreateRenderPipelineProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return RenderPipeline(ret)
}

func DeviceCreateSampler(device Device, descriptor *SamplerDescriptor) Sampler {
	ret, _, err := deviceCreateSamplerProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return Sampler(ret)
}

func DeviceCreateShaderModule(device Device, descriptor *ShaderModuleDescriptor) ShaderModule {
	ret, _, err := deviceCreateShaderModuleProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return ShaderModule(ret)
}

func DeviceCreateTexture(device Device, descriptor *TextureDescriptor) Texture {
	ret, _, err := deviceCreateTextureProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return Texture(ret)
}

func DeviceDestroy(device Device) {
	_, _, err := deviceDestroyProc.Call(
		uintptr(device),
	)
	if err != nil {
		panic(err)
	}
}

func DeviceGetLostFuture(device Device) Future {
	ret, _, err := deviceGetLostFutureProc.Call(
		uintptr(device),
	)
	if err != nil {
		panic(err)
	}
	return *(*Future)(unsafe.Pointer(new(uint64(ret))))
}

func DeviceGetLimits(device Device, limits *Limits) Status {
	ret, _, err := deviceGetLimitsProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(limits)),
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func DeviceHasFeature(device Device, feature FeatureName) Bool {
	ret, _, err := deviceHasFeatureProc.Call(
		uintptr(device),
		uintptr(feature), // enum.feature_name
	)
	if err != nil {
		panic(err)
	}
	return Bool(ret)
}

func DeviceGetFeatures(device Device, features *SupportedFeatures) {
	_, _, err := deviceGetFeaturesProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(features)),
	)
	if err != nil {
		panic(err)
	}
}

func DeviceGetAdapterInfo(device Device, adapterInfo *AdapterInfo) Status {
	ret, _, err := deviceGetAdapterInfoProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(adapterInfo)),
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func DeviceGetQueue(device Device) Queue {
	ret, _, err := deviceGetQueueProc.Call(
		uintptr(device),
	)
	if err != nil {
		panic(err)
	}
	return Queue(ret)
}

func DevicePushErrorScope(device Device, filter ErrorFilter) {
	_, _, err := devicePushErrorScopeProc.Call(
		uintptr(device),
		uintptr(filter), // enum.error_filter
	)
	if err != nil {
		panic(err)
	}
}

func DevicePopErrorScope(device Device, callback PopErrorScopeCallbackInfo) {
	_, _, err := devicePopErrorScopeProc.Call(
		uintptr(device),
	)
	if err != nil {
		panic(err)
	}
}

func DeviceSetLabel(device Device, label StringView) {
	_, _, err := deviceSetLabelProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := externalTextureSetLabelProc.Call(
		uintptr(externalTexture),
		uintptr(unsafe.Pointer(&label)),
	)
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
	ret, _, err := instanceCreateSurfaceProc.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return Surface(ret)
}

func InstanceGetWGSLLanguageFeatures(instance Instance, features *SupportedWGSLLanguageFeatures) {
	_, _, err := instanceGetWGSLLanguageFeaturesProc.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(features)),
	)
	if err != nil {
		panic(err)
	}
}

func InstanceHasWGSLLanguageFeature(instance Instance, feature WGSLLanguageFeatureName) Bool {
	ret, _, err := instanceHasWGSLLanguageFeatureProc.Call(
		uintptr(instance),
		uintptr(feature), // enum.WGSL_language_feature_name
	)
	if err != nil {
		panic(err)
	}
	return Bool(ret)
}

func InstanceProcessEvents(instance Instance) {
	_, _, err := instanceProcessEventsProc.Call(
		uintptr(instance),
	)
	if err != nil {
		panic(err)
	}
}

func InstanceRequestAdapter(instance Instance, options *RequestAdapterOptions, callback RequestAdapterCallbackInfo) {
	_, _, err := instanceRequestAdapterProc.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(options)),
	)
	if err != nil {
		panic(err)
	}
}

func InstanceWaitAny(instance Instance, futureCount uintptr, futures *FutureWaitInfo, timeoutNS uint64) WaitStatus {
	ret, _, err := instanceWaitAnyProc.Call(
		uintptr(instance),
		uintptr(futureCount), // usize
		uintptr(unsafe.Pointer(futures)),
		uintptr(timeoutNS), // uint64
	)
	if err != nil {
		panic(err)
	}
	return WaitStatus(ret)
}

func InstanceAddRef(instance Instance) {
	instanceAddRefProc.Call(uintptr(instance))
}

func InstanceRelease(instance Instance) {
	instanceReleaseProc.Call(uintptr(instance))
}

func PipelineLayoutSetLabel(pipelineLayout PipelineLayout, label StringView) {
	_, _, err := pipelineLayoutSetLabelProc.Call(
		uintptr(pipelineLayout),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := querySetSetLabelProc.Call(
		uintptr(querySet),
		uintptr(unsafe.Pointer(&label)),
	)
	if err != nil {
		panic(err)
	}
}

func QuerySetGetType(querySet QuerySet) QueryType {
	ret, _, err := querySetGetTypeProc.Call(
		uintptr(querySet),
	)
	if err != nil {
		panic(err)
	}
	return QueryType(ret)
}

func QuerySetGetCount(querySet QuerySet) uint32 {
	ret, _, err := querySetGetCountProc.Call(
		uintptr(querySet),
	)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}

func QuerySetDestroy(querySet QuerySet) {
	_, _, err := querySetDestroyProc.Call(
		uintptr(querySet),
	)
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
	_, _, err := queueSubmitProc.Call(
		uintptr(queue),
		uintptr(len(commands)),
		uintptr(unsafe.Pointer(unsafe.SliceData(commands))),
	)
	if err != nil {
		panic(err)
	}
}

func QueueOnSubmittedWorkDone(queue Queue, callback QueueWorkDoneCallbackInfo) {
	_, _, err := queueOnSubmittedWorkDoneProc.Call(
		uintptr(queue),
	)
	if err != nil {
		panic(err)
	}
}

func QueueWriteBuffer(queue Queue, buffer Buffer, bufferOffset uint64, data unsafe.Pointer, size uintptr) {
	_, _, err := queueWriteBufferProc.Call(
		uintptr(queue),
		uintptr(buffer),       // object.buffer
		uintptr(bufferOffset), // uint64
		uintptr(data),         // c_void_data_ptr
		uintptr(size),         // usize
	)
	if err != nil {
		panic(err)
	}
}

func QueueWriteTexture(queue Queue, destination *TexelCopyTextureInfo, data unsafe.Pointer, dataSize uintptr, dataLayout *TexelCopyBufferLayout, writeSize *Extent3D) {
	_, _, err := queueWriteTextureProc.Call(
		uintptr(queue),
		uintptr(unsafe.Pointer(destination)),
		uintptr(data),     // c_void_data_ptr
		uintptr(dataSize), // usize
		uintptr(unsafe.Pointer(dataLayout)),
		uintptr(unsafe.Pointer(writeSize)),
	)
	if err != nil {
		panic(err)
	}
}

func QueueSetLabel(queue Queue, label StringView) {
	_, _, err := queueSetLabelProc.Call(
		uintptr(queue),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := renderBundleSetLabelProc.Call(
		uintptr(renderBundle),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := renderBundleEncoderSetPipelineProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(pipeline), // object.render_pipeline
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetBindGroup(renderBundleEncoder RenderBundleEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {
	_, _, err := renderBundleEncoderSetBindGroupProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(groupIndex), // uint32
		uintptr(group),      // object.bind_group
		uintptr(len(dynamicOffsets)),
		uintptr(unsafe.Pointer(unsafe.SliceData(dynamicOffsets))),
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetImmediates(renderBundleEncoder RenderBundleEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	_, _, err := renderBundleEncoderSetImmediatesProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(offset), // uint32
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDraw(renderBundleEncoder RenderBundleEncoder, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	_, _, err := renderBundleEncoderDrawProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(vertexCount),   // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstVertex),   // uint32
		uintptr(firstInstance), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDrawIndexed(renderBundleEncoder RenderBundleEncoder, indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	_, _, err := renderBundleEncoderDrawIndexedProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(indexCount),    // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstIndex),    // uint32
		uintptr(baseVertex),    // int32
		uintptr(firstInstance), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDrawIndirect(renderBundleEncoder RenderBundleEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, err := renderBundleEncoderDrawIndirectProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderDrawIndexedIndirect(renderBundleEncoder RenderBundleEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, err := renderBundleEncoderDrawIndexedIndirectProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderInsertDebugMarker(renderBundleEncoder RenderBundleEncoder, markerLabel StringView) {
	_, _, err := renderBundleEncoderInsertDebugMarkerProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderPopDebugGroup(renderBundleEncoder RenderBundleEncoder) {
	_, _, err := renderBundleEncoderPopDebugGroupProc.Call(
		uintptr(renderBundleEncoder),
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderPushDebugGroup(renderBundleEncoder RenderBundleEncoder, groupLabel StringView) {
	_, _, err := renderBundleEncoderPushDebugGroupProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetVertexBuffer(renderBundleEncoder RenderBundleEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {
	_, _, err := renderBundleEncoderSetVertexBufferProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(slot),   // uint32
		uintptr(buffer), // object.buffer
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderSetIndexBuffer(renderBundleEncoder RenderBundleEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	_, _, err := renderBundleEncoderSetIndexBufferProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(buffer), // object.buffer
		uintptr(format), // enum.index_format
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderBundleEncoderFinish(renderBundleEncoder RenderBundleEncoder, descriptor *RenderBundleDescriptor) RenderBundle {
	ret, _, err := renderBundleEncoderFinishProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return RenderBundle(ret)
}

func RenderBundleEncoderSetLabel(renderBundleEncoder RenderBundleEncoder, label StringView) {
	_, _, err := renderBundleEncoderSetLabelProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := renderPassEncoderSetPipelineProc.Call(
		uintptr(renderPassEncoder),
		uintptr(pipeline), // object.render_pipeline
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetBindGroup(renderPassEncoder RenderPassEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {
	_, _, err := renderPassEncoderSetBindGroupProc.Call(
		uintptr(renderPassEncoder),
		uintptr(groupIndex), // uint32
		uintptr(group),      // object.bind_group
		uintptr(len(dynamicOffsets)),
		uintptr(unsafe.Pointer(unsafe.SliceData(dynamicOffsets))),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetImmediates(renderPassEncoder RenderPassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	_, _, err := renderPassEncoderSetImmediatesProc.Call(
		uintptr(renderPassEncoder),
		uintptr(offset), // uint32
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDraw(renderPassEncoder RenderPassEncoder, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	_, _, err := renderPassEncoderDrawProc.Call(
		uintptr(renderPassEncoder),
		uintptr(vertexCount),   // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstVertex),   // uint32
		uintptr(firstInstance), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDrawIndexed(renderPassEncoder RenderPassEncoder, indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	_, _, err := renderPassEncoderDrawIndexedProc.Call(
		uintptr(renderPassEncoder),
		uintptr(indexCount),    // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstIndex),    // uint32
		uintptr(baseVertex),    // int32
		uintptr(firstInstance), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDrawIndirect(renderPassEncoder RenderPassEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, err := renderPassEncoderDrawIndirectProc.Call(
		uintptr(renderPassEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderDrawIndexedIndirect(renderPassEncoder RenderPassEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, err := renderPassEncoderDrawIndexedIndirectProc.Call(
		uintptr(renderPassEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderExecuteBundles(renderPassEncoder RenderPassEncoder, bundles []RenderBundle) {
	_, _, err := renderPassEncoderExecuteBundlesProc.Call(
		uintptr(renderPassEncoder),
		uintptr(len(bundles)),
		uintptr(unsafe.Pointer(unsafe.SliceData(bundles))),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderInsertDebugMarker(renderPassEncoder RenderPassEncoder, markerLabel StringView) {
	_, _, err := renderPassEncoderInsertDebugMarkerProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderPopDebugGroup(renderPassEncoder RenderPassEncoder) {
	_, _, err := renderPassEncoderPopDebugGroupProc.Call(
		uintptr(renderPassEncoder),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderPushDebugGroup(renderPassEncoder RenderPassEncoder, groupLabel StringView) {
	_, _, err := renderPassEncoderPushDebugGroupProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetStencilReference(renderPassEncoder RenderPassEncoder, reference uint32) {
	_, _, err := renderPassEncoderSetStencilReferenceProc.Call(
		uintptr(renderPassEncoder),
		uintptr(reference), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetBlendConstant(renderPassEncoder RenderPassEncoder, color *Color) {
	_, _, err := renderPassEncoderSetBlendConstantProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(color)),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetViewport(renderPassEncoder RenderPassEncoder, x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) {
	_, _, err := renderPassEncoderSetViewportProc.Call(
		uintptr(renderPassEncoder),
		uintptr(x),        // float32
		uintptr(y),        // float32
		uintptr(width),    // float32
		uintptr(height),   // float32
		uintptr(minDepth), // float32
		uintptr(maxDepth), // float32
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetScissorRect(renderPassEncoder RenderPassEncoder, x uint32, y uint32, width uint32, height uint32) {
	_, _, err := renderPassEncoderSetScissorRectProc.Call(
		uintptr(renderPassEncoder),
		uintptr(x),      // uint32
		uintptr(y),      // uint32
		uintptr(width),  // uint32
		uintptr(height), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetVertexBuffer(renderPassEncoder RenderPassEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {
	_, _, err := renderPassEncoderSetVertexBufferProc.Call(
		uintptr(renderPassEncoder),
		uintptr(slot),   // uint32
		uintptr(buffer), // object.buffer
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetIndexBuffer(renderPassEncoder RenderPassEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	_, _, err := renderPassEncoderSetIndexBufferProc.Call(
		uintptr(renderPassEncoder),
		uintptr(buffer), // object.buffer
		uintptr(format), // enum.index_format
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderBeginOcclusionQuery(renderPassEncoder RenderPassEncoder, queryIndex uint32) {
	_, _, err := renderPassEncoderBeginOcclusionQueryProc.Call(
		uintptr(renderPassEncoder),
		uintptr(queryIndex), // uint32
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderEndOcclusionQuery(renderPassEncoder RenderPassEncoder) {
	_, _, err := renderPassEncoderEndOcclusionQueryProc.Call(
		uintptr(renderPassEncoder),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderEnd(renderPassEncoder RenderPassEncoder) {
	_, _, err := renderPassEncoderEndProc.Call(
		uintptr(renderPassEncoder),
	)
	if err != nil {
		panic(err)
	}
}

func RenderPassEncoderSetLabel(renderPassEncoder RenderPassEncoder, label StringView) {
	_, _, err := renderPassEncoderSetLabelProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
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
	ret, _, err := renderPipelineGetBindGroupLayoutProc.Call(
		uintptr(renderPipeline),
		uintptr(groupIndex), // uint32
	)
	if err != nil {
		panic(err)
	}
	return BindGroupLayout(ret)
}

func RenderPipelineSetLabel(renderPipeline RenderPipeline, label StringView) {
	_, _, err := renderPipelineSetLabelProc.Call(
		uintptr(renderPipeline),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := samplerSetLabelProc.Call(
		uintptr(sampler),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := shaderModuleGetCompilationInfoProc.Call(
		uintptr(shaderModule),
	)
	if err != nil {
		panic(err)
	}
}

func ShaderModuleSetLabel(shaderModule ShaderModule, label StringView) {
	_, _, err := shaderModuleSetLabelProc.Call(
		uintptr(shaderModule),
		uintptr(unsafe.Pointer(&label)),
	)
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
	_, _, err := surfaceConfigureProc.Call(
		uintptr(surface),
		uintptr(unsafe.Pointer(config)),
	)
	if err != nil {
		panic(err)
	}
}

func SurfaceGetCapabilities(surface Surface, adapter Adapter, capabilities *SurfaceCapabilities) Status {
	ret, _, err := surfaceGetCapabilitiesProc.Call(
		uintptr(surface),
		uintptr(adapter), // object.adapter
		uintptr(unsafe.Pointer(capabilities)),
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func SurfaceGetCurrentTexture(surface Surface, surfaceTexture *SurfaceTexture) {
	_, _, err := surfaceGetCurrentTextureProc.Call(
		uintptr(surface),
		uintptr(unsafe.Pointer(surfaceTexture)),
	)
	if err != nil {
		panic(err)
	}
}

func SurfacePresent(surface Surface) Status {
	ret, _, err := surfacePresentProc.Call(
		uintptr(surface),
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

func SurfaceUnconfigure(surface Surface) {
	_, _, err := surfaceUnconfigureProc.Call(
		uintptr(surface),
	)
	if err != nil {
		panic(err)
	}
}

func SurfaceSetLabel(surface Surface, label StringView) {
	_, _, err := surfaceSetLabelProc.Call(
		uintptr(surface),
		uintptr(unsafe.Pointer(&label)),
	)
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
	ret, _, err := textureCreateViewProc.Call(
		uintptr(texture),
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return TextureView(ret)
}

func TextureSetLabel(texture Texture, label StringView) {
	_, _, err := textureSetLabelProc.Call(
		uintptr(texture),
		uintptr(unsafe.Pointer(&label)),
	)
	if err != nil {
		panic(err)
	}
}

func TextureGetWidth(texture Texture) uint32 {
	ret, _, err := textureGetWidthProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}

func TextureGetHeight(texture Texture) uint32 {
	ret, _, err := textureGetHeightProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}

func TextureGetDepthOrArrayLayers(texture Texture) uint32 {
	ret, _, err := textureGetDepthOrArrayLayersProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}

func TextureGetMipLevelCount(texture Texture) uint32 {
	ret, _, err := textureGetMipLevelCountProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}

func TextureGetSampleCount(texture Texture) uint32 {
	ret, _, err := textureGetSampleCountProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}

func TextureGetDimension(texture Texture) TextureDimension {
	ret, _, err := textureGetDimensionProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return TextureDimension(ret)
}

func TextureGetTextureBindingViewDimension(texture Texture) TextureViewDimension {
	ret, _, err := textureGetTextureBindingViewDimensionProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return TextureViewDimension(ret)
}

func TextureGetFormat(texture Texture) TextureFormat {
	ret, _, err := textureGetFormatProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return TextureFormat(ret)
}

func TextureGetUsage(texture Texture) TextureUsage {
	ret, _, err := textureGetUsageProc.Call(
		uintptr(texture),
	)
	if err != nil {
		panic(err)
	}
	return TextureUsage(ret)
}

func TextureDestroy(texture Texture) {
	_, _, err := textureDestroyProc.Call(
		uintptr(texture),
	)
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
	_, _, err := textureViewSetLabelProc.Call(
		uintptr(textureView),
		uintptr(unsafe.Pointer(&label)),
	)
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
