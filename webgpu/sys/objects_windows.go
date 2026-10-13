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
	webgpuDLL                                        = syscall.NewLazyDLL("wgpu_native.dll")
	adapterGetLimitsProc                             = webgpuDLL.NewProc("wgpuAdapterGetLimits")
	adapterHasFeatureProc                            = webgpuDLL.NewProc("wgpuAdapterHasFeature")
	adapterGetFeaturesProc                           = webgpuDLL.NewProc("wgpuAdapterGetFeatures")
	adapterGetInfoProc                               = webgpuDLL.NewProc("wgpuAdapterGetInfo")
	adapterRequestDeviceProc                         = webgpuDLL.NewProc("wgpuAdapterRequestDevice")
	adapterAddRefProc                                = webgpuDLL.NewProc("wgpuAdapterAddRef")
	adapterReleaseProc                               = webgpuDLL.NewProc("wgpuAdapterRelease")
	bindGroupSetLabelProc                            = webgpuDLL.NewProc("wgpuBindGroupSetLabel")
	bindGroupAddRefProc                              = webgpuDLL.NewProc("wgpuBindGroupAddRef")
	bindGroupReleaseProc                             = webgpuDLL.NewProc("wgpuBindGroupRelease")
	bindGroupLayoutSetLabelProc                      = webgpuDLL.NewProc("wgpuBindGroupLayoutSetLabel")
	bindGroupLayoutAddRefProc                        = webgpuDLL.NewProc("wgpuBindGroupLayoutAddRef")
	bindGroupLayoutReleaseProc                       = webgpuDLL.NewProc("wgpuBindGroupLayoutRelease")
	bufferMapAsyncProc                               = webgpuDLL.NewProc("wgpuBufferMapAsync")
	bufferGetMappedRangeProc                         = webgpuDLL.NewProc("wgpuBufferGetMappedRange")
	bufferGetConstMappedRangeProc                    = webgpuDLL.NewProc("wgpuBufferGetConstMappedRange")
	bufferReadMappedRangeProc                        = webgpuDLL.NewProc("wgpuBufferReadMappedRange")
	bufferWriteMappedRangeProc                       = webgpuDLL.NewProc("wgpuBufferWriteMappedRange")
	bufferSetLabelProc                               = webgpuDLL.NewProc("wgpuBufferSetLabel")
	bufferGetUsageProc                               = webgpuDLL.NewProc("wgpuBufferGetUsage")
	bufferGetSizeProc                                = webgpuDLL.NewProc("wgpuBufferGetSize")
	bufferGetMapStateProc                            = webgpuDLL.NewProc("wgpuBufferGetMapState")
	bufferUnmapProc                                  = webgpuDLL.NewProc("wgpuBufferUnmap")
	bufferDestroyProc                                = webgpuDLL.NewProc("wgpuBufferDestroy")
	bufferAddRefProc                                 = webgpuDLL.NewProc("wgpuBufferAddRef")
	bufferReleaseProc                                = webgpuDLL.NewProc("wgpuBufferRelease")
	commandBufferSetLabelProc                        = webgpuDLL.NewProc("wgpuCommandBufferSetLabel")
	commandBufferAddRefProc                          = webgpuDLL.NewProc("wgpuCommandBufferAddRef")
	commandBufferReleaseProc                         = webgpuDLL.NewProc("wgpuCommandBufferRelease")
	commandEncoderFinishProc                         = webgpuDLL.NewProc("wgpuCommandEncoderFinish")
	commandEncoderBeginComputePassProc               = webgpuDLL.NewProc("wgpuCommandEncoderBeginComputePass")
	commandEncoderBeginRenderPassProc                = webgpuDLL.NewProc("wgpuCommandEncoderBeginRenderPass")
	commandEncoderCopyBufferToBufferProc             = webgpuDLL.NewProc("wgpuCommandEncoderCopyBufferToBuffer")
	commandEncoderCopyBufferToTextureProc            = webgpuDLL.NewProc("wgpuCommandEncoderCopyBufferToTexture")
	commandEncoderCopyTextureToBufferProc            = webgpuDLL.NewProc("wgpuCommandEncoderCopyTextureToBuffer")
	commandEncoderCopyTextureToTextureProc           = webgpuDLL.NewProc("wgpuCommandEncoderCopyTextureToTexture")
	commandEncoderClearBufferProc                    = webgpuDLL.NewProc("wgpuCommandEncoderClearBuffer")
	commandEncoderInsertDebugMarkerProc              = webgpuDLL.NewProc("wgpuCommandEncoderInsertDebugMarker")
	commandEncoderPopDebugGroupProc                  = webgpuDLL.NewProc("wgpuCommandEncoderPopDebugGroup")
	commandEncoderPushDebugGroupProc                 = webgpuDLL.NewProc("wgpuCommandEncoderPushDebugGroup")
	commandEncoderResolveQuerySetProc                = webgpuDLL.NewProc("wgpuCommandEncoderResolveQuerySet")
	commandEncoderWriteTimestampProc                 = webgpuDLL.NewProc("wgpuCommandEncoderWriteTimestamp")
	commandEncoderSetLabelProc                       = webgpuDLL.NewProc("wgpuCommandEncoderSetLabel")
	commandEncoderAddRefProc                         = webgpuDLL.NewProc("wgpuCommandEncoderAddRef")
	commandEncoderReleaseProc                        = webgpuDLL.NewProc("wgpuCommandEncoderRelease")
	computePassEncoderInsertDebugMarkerProc          = webgpuDLL.NewProc("wgpuComputePassEncoderInsertDebugMarker")
	computePassEncoderPopDebugGroupProc              = webgpuDLL.NewProc("wgpuComputePassEncoderPopDebugGroup")
	computePassEncoderPushDebugGroupProc             = webgpuDLL.NewProc("wgpuComputePassEncoderPushDebugGroup")
	computePassEncoderSetPipelineProc                = webgpuDLL.NewProc("wgpuComputePassEncoderSetPipeline")
	computePassEncoderSetBindGroupProc               = webgpuDLL.NewProc("wgpuComputePassEncoderSetBindGroup")
	computePassEncoderSetImmediatesProc              = webgpuDLL.NewProc("wgpuComputePassEncoderSetImmediates")
	computePassEncoderDispatchWorkgroupsProc         = webgpuDLL.NewProc("wgpuComputePassEncoderDispatchWorkgroups")
	computePassEncoderDispatchWorkgroupsIndirectProc = webgpuDLL.NewProc("wgpuComputePassEncoderDispatchWorkgroupsIndirect")
	computePassEncoderEndProc                        = webgpuDLL.NewProc("wgpuComputePassEncoderEnd")
	computePassEncoderSetLabelProc                   = webgpuDLL.NewProc("wgpuComputePassEncoderSetLabel")
	computePassEncoderAddRefProc                     = webgpuDLL.NewProc("wgpuComputePassEncoderAddRef")
	computePassEncoderReleaseProc                    = webgpuDLL.NewProc("wgpuComputePassEncoderRelease")
	computePipelineGetBindGroupLayoutProc            = webgpuDLL.NewProc("wgpuComputePipelineGetBindGroupLayout")
	computePipelineSetLabelProc                      = webgpuDLL.NewProc("wgpuComputePipelineSetLabel")
	computePipelineAddRefProc                        = webgpuDLL.NewProc("wgpuComputePipelineAddRef")
	computePipelineReleaseProc                       = webgpuDLL.NewProc("wgpuComputePipelineRelease")
	deviceCreateBindGroupProc                        = webgpuDLL.NewProc("wgpuDeviceCreateBindGroup")
	deviceCreateBindGroupLayoutProc                  = webgpuDLL.NewProc("wgpuDeviceCreateBindGroupLayout")
	deviceCreateBufferProc                           = webgpuDLL.NewProc("wgpuDeviceCreateBuffer")
	deviceCreateCommandEncoderProc                   = webgpuDLL.NewProc("wgpuDeviceCreateCommandEncoder")
	deviceCreateComputePipelineProc                  = webgpuDLL.NewProc("wgpuDeviceCreateComputePipeline")
	deviceCreateComputePipelineAsyncProc             = webgpuDLL.NewProc("wgpuDeviceCreateComputePipelineAsync")
	deviceCreatePipelineLayoutProc                   = webgpuDLL.NewProc("wgpuDeviceCreatePipelineLayout")
	deviceCreateQuerySetProc                         = webgpuDLL.NewProc("wgpuDeviceCreateQuerySet")
	deviceCreateRenderPipelineAsyncProc              = webgpuDLL.NewProc("wgpuDeviceCreateRenderPipelineAsync")
	deviceCreateRenderBundleEncoderProc              = webgpuDLL.NewProc("wgpuDeviceCreateRenderBundleEncoder")
	deviceCreateRenderPipelineProc                   = webgpuDLL.NewProc("wgpuDeviceCreateRenderPipeline")
	deviceCreateSamplerProc                          = webgpuDLL.NewProc("wgpuDeviceCreateSampler")
	deviceCreateShaderModuleProc                     = webgpuDLL.NewProc("wgpuDeviceCreateShaderModule")
	deviceCreateTextureProc                          = webgpuDLL.NewProc("wgpuDeviceCreateTexture")
	deviceDestroyProc                                = webgpuDLL.NewProc("wgpuDeviceDestroy")
	deviceGetLostFutureProc                          = webgpuDLL.NewProc("wgpuDeviceGetLostFuture")
	deviceGetLimitsProc                              = webgpuDLL.NewProc("wgpuDeviceGetLimits")
	deviceHasFeatureProc                             = webgpuDLL.NewProc("wgpuDeviceHasFeature")
	deviceGetFeaturesProc                            = webgpuDLL.NewProc("wgpuDeviceGetFeatures")
	deviceGetAdapterInfoProc                         = webgpuDLL.NewProc("wgpuDeviceGetAdapterInfo")
	deviceGetQueueProc                               = webgpuDLL.NewProc("wgpuDeviceGetQueue")
	devicePushErrorScopeProc                         = webgpuDLL.NewProc("wgpuDevicePushErrorScope")
	devicePopErrorScopeProc                          = webgpuDLL.NewProc("wgpuDevicePopErrorScope")
	deviceSetLabelProc                               = webgpuDLL.NewProc("wgpuDeviceSetLabel")
	deviceAddRefProc                                 = webgpuDLL.NewProc("wgpuDeviceAddRef")
	deviceReleaseProc                                = webgpuDLL.NewProc("wgpuDeviceRelease")
	externalTextureSetLabelProc                      = webgpuDLL.NewProc("wgpuExternalTextureSetLabel")
	externalTextureAddRefProc                        = webgpuDLL.NewProc("wgpuExternalTextureAddRef")
	externalTextureReleaseProc                       = webgpuDLL.NewProc("wgpuExternalTextureRelease")
	instanceCreateSurfaceProc                        = webgpuDLL.NewProc("wgpuInstanceCreateSurface")
	instanceGetWGSLLanguageFeaturesProc              = webgpuDLL.NewProc("wgpuInstanceGetWGSLLanguageFeatures")
	instanceHasWGSLLanguageFeatureProc               = webgpuDLL.NewProc("wgpuInstanceHasWGSLLanguageFeature")
	instanceProcessEventsProc                        = webgpuDLL.NewProc("wgpuInstanceProcessEvents")
	instanceRequestAdapterProc                       = webgpuDLL.NewProc("wgpuInstanceRequestAdapter")
	instanceWaitAnyProc                              = webgpuDLL.NewProc("wgpuInstanceWaitAny")
	instanceAddRefProc                               = webgpuDLL.NewProc("wgpuInstanceAddRef")
	instanceReleaseProc                              = webgpuDLL.NewProc("wgpuInstanceRelease")
	pipelineLayoutSetLabelProc                       = webgpuDLL.NewProc("wgpuPipelineLayoutSetLabel")
	pipelineLayoutAddRefProc                         = webgpuDLL.NewProc("wgpuPipelineLayoutAddRef")
	pipelineLayoutReleaseProc                        = webgpuDLL.NewProc("wgpuPipelineLayoutRelease")
	querySetSetLabelProc                             = webgpuDLL.NewProc("wgpuQuerySetSetLabel")
	querySetGetTypeProc                              = webgpuDLL.NewProc("wgpuQuerySetGetType")
	querySetGetCountProc                             = webgpuDLL.NewProc("wgpuQuerySetGetCount")
	querySetDestroyProc                              = webgpuDLL.NewProc("wgpuQuerySetDestroy")
	querySetAddRefProc                               = webgpuDLL.NewProc("wgpuQuerySetAddRef")
	querySetReleaseProc                              = webgpuDLL.NewProc("wgpuQuerySetRelease")
	queueSubmitProc                                  = webgpuDLL.NewProc("wgpuQueueSubmit")
	queueOnSubmittedWorkDoneProc                     = webgpuDLL.NewProc("wgpuQueueOnSubmittedWorkDone")
	queueWriteBufferProc                             = webgpuDLL.NewProc("wgpuQueueWriteBuffer")
	queueWriteTextureProc                            = webgpuDLL.NewProc("wgpuQueueWriteTexture")
	queueSetLabelProc                                = webgpuDLL.NewProc("wgpuQueueSetLabel")
	queueAddRefProc                                  = webgpuDLL.NewProc("wgpuQueueAddRef")
	queueReleaseProc                                 = webgpuDLL.NewProc("wgpuQueueRelease")
	renderBundleSetLabelProc                         = webgpuDLL.NewProc("wgpuRenderBundleSetLabel")
	renderBundleAddRefProc                           = webgpuDLL.NewProc("wgpuRenderBundleAddRef")
	renderBundleReleaseProc                          = webgpuDLL.NewProc("wgpuRenderBundleRelease")
	renderBundleEncoderSetPipelineProc               = webgpuDLL.NewProc("wgpuRenderBundleEncoderSetPipeline")
	renderBundleEncoderSetBindGroupProc              = webgpuDLL.NewProc("wgpuRenderBundleEncoderSetBindGroup")
	renderBundleEncoderSetImmediatesProc             = webgpuDLL.NewProc("wgpuRenderBundleEncoderSetImmediates")
	renderBundleEncoderDrawProc                      = webgpuDLL.NewProc("wgpuRenderBundleEncoderDraw")
	renderBundleEncoderDrawIndexedProc               = webgpuDLL.NewProc("wgpuRenderBundleEncoderDrawIndexed")
	renderBundleEncoderDrawIndirectProc              = webgpuDLL.NewProc("wgpuRenderBundleEncoderDrawIndirect")
	renderBundleEncoderDrawIndexedIndirectProc       = webgpuDLL.NewProc("wgpuRenderBundleEncoderDrawIndexedIndirect")
	renderBundleEncoderInsertDebugMarkerProc         = webgpuDLL.NewProc("wgpuRenderBundleEncoderInsertDebugMarker")
	renderBundleEncoderPopDebugGroupProc             = webgpuDLL.NewProc("wgpuRenderBundleEncoderPopDebugGroup")
	renderBundleEncoderPushDebugGroupProc            = webgpuDLL.NewProc("wgpuRenderBundleEncoderPushDebugGroup")
	renderBundleEncoderSetVertexBufferProc           = webgpuDLL.NewProc("wgpuRenderBundleEncoderSetVertexBuffer")
	renderBundleEncoderSetIndexBufferProc            = webgpuDLL.NewProc("wgpuRenderBundleEncoderSetIndexBuffer")
	renderBundleEncoderFinishProc                    = webgpuDLL.NewProc("wgpuRenderBundleEncoderFinish")
	renderBundleEncoderSetLabelProc                  = webgpuDLL.NewProc("wgpuRenderBundleEncoderSetLabel")
	renderBundleEncoderAddRefProc                    = webgpuDLL.NewProc("wgpuRenderBundleEncoderAddRef")
	renderBundleEncoderReleaseProc                   = webgpuDLL.NewProc("wgpuRenderBundleEncoderRelease")
	renderPassEncoderSetPipelineProc                 = webgpuDLL.NewProc("wgpuRenderPassEncoderSetPipeline")
	renderPassEncoderSetBindGroupProc                = webgpuDLL.NewProc("wgpuRenderPassEncoderSetBindGroup")
	renderPassEncoderSetImmediatesProc               = webgpuDLL.NewProc("wgpuRenderPassEncoderSetImmediates")
	renderPassEncoderDrawProc                        = webgpuDLL.NewProc("wgpuRenderPassEncoderDraw")
	renderPassEncoderDrawIndexedProc                 = webgpuDLL.NewProc("wgpuRenderPassEncoderDrawIndexed")
	renderPassEncoderDrawIndirectProc                = webgpuDLL.NewProc("wgpuRenderPassEncoderDrawIndirect")
	renderPassEncoderDrawIndexedIndirectProc         = webgpuDLL.NewProc("wgpuRenderPassEncoderDrawIndexedIndirect")
	renderPassEncoderExecuteBundlesProc              = webgpuDLL.NewProc("wgpuRenderPassEncoderExecuteBundles")
	renderPassEncoderInsertDebugMarkerProc           = webgpuDLL.NewProc("wgpuRenderPassEncoderInsertDebugMarker")
	renderPassEncoderPopDebugGroupProc               = webgpuDLL.NewProc("wgpuRenderPassEncoderPopDebugGroup")
	renderPassEncoderPushDebugGroupProc              = webgpuDLL.NewProc("wgpuRenderPassEncoderPushDebugGroup")
	renderPassEncoderSetStencilReferenceProc         = webgpuDLL.NewProc("wgpuRenderPassEncoderSetStencilReference")
	renderPassEncoderSetBlendConstantProc            = webgpuDLL.NewProc("wgpuRenderPassEncoderSetBlendConstant")
	renderPassEncoderSetViewportProc                 = webgpuDLL.NewProc("wgpuRenderPassEncoderSetViewport")
	renderPassEncoderSetScissorRectProc              = webgpuDLL.NewProc("wgpuRenderPassEncoderSetScissorRect")
	renderPassEncoderSetVertexBufferProc             = webgpuDLL.NewProc("wgpuRenderPassEncoderSetVertexBuffer")
	renderPassEncoderSetIndexBufferProc              = webgpuDLL.NewProc("wgpuRenderPassEncoderSetIndexBuffer")
	renderPassEncoderBeginOcclusionQueryProc         = webgpuDLL.NewProc("wgpuRenderPassEncoderBeginOcclusionQuery")
	renderPassEncoderEndOcclusionQueryProc           = webgpuDLL.NewProc("wgpuRenderPassEncoderEndOcclusionQuery")
	renderPassEncoderEndProc                         = webgpuDLL.NewProc("wgpuRenderPassEncoderEnd")
	renderPassEncoderSetLabelProc                    = webgpuDLL.NewProc("wgpuRenderPassEncoderSetLabel")
	renderPassEncoderAddRefProc                      = webgpuDLL.NewProc("wgpuRenderPassEncoderAddRef")
	renderPassEncoderReleaseProc                     = webgpuDLL.NewProc("wgpuRenderPassEncoderRelease")
	renderPipelineGetBindGroupLayoutProc             = webgpuDLL.NewProc("wgpuRenderPipelineGetBindGroupLayout")
	renderPipelineSetLabelProc                       = webgpuDLL.NewProc("wgpuRenderPipelineSetLabel")
	renderPipelineAddRefProc                         = webgpuDLL.NewProc("wgpuRenderPipelineAddRef")
	renderPipelineReleaseProc                        = webgpuDLL.NewProc("wgpuRenderPipelineRelease")
	samplerSetLabelProc                              = webgpuDLL.NewProc("wgpuSamplerSetLabel")
	samplerAddRefProc                                = webgpuDLL.NewProc("wgpuSamplerAddRef")
	samplerReleaseProc                               = webgpuDLL.NewProc("wgpuSamplerRelease")
	shaderModuleGetCompilationInfoProc               = webgpuDLL.NewProc("wgpuShaderModuleGetCompilationInfo")
	shaderModuleSetLabelProc                         = webgpuDLL.NewProc("wgpuShaderModuleSetLabel")
	shaderModuleAddRefProc                           = webgpuDLL.NewProc("wgpuShaderModuleAddRef")
	shaderModuleReleaseProc                          = webgpuDLL.NewProc("wgpuShaderModuleRelease")
	surfaceConfigureProc                             = webgpuDLL.NewProc("wgpuSurfaceConfigure")
	surfaceGetCapabilitiesProc                       = webgpuDLL.NewProc("wgpuSurfaceGetCapabilities")
	surfaceGetCurrentTextureProc                     = webgpuDLL.NewProc("wgpuSurfaceGetCurrentTexture")
	surfacePresentProc                               = webgpuDLL.NewProc("wgpuSurfacePresent")
	surfaceUnconfigureProc                           = webgpuDLL.NewProc("wgpuSurfaceUnconfigure")
	surfaceSetLabelProc                              = webgpuDLL.NewProc("wgpuSurfaceSetLabel")
	surfaceAddRefProc                                = webgpuDLL.NewProc("wgpuSurfaceAddRef")
	surfaceReleaseProc                               = webgpuDLL.NewProc("wgpuSurfaceRelease")
	textureCreateViewProc                            = webgpuDLL.NewProc("wgpuTextureCreateView")
	textureSetLabelProc                              = webgpuDLL.NewProc("wgpuTextureSetLabel")
	textureGetWidthProc                              = webgpuDLL.NewProc("wgpuTextureGetWidth")
	textureGetHeightProc                             = webgpuDLL.NewProc("wgpuTextureGetHeight")
	textureGetDepthOrArrayLayersProc                 = webgpuDLL.NewProc("wgpuTextureGetDepthOrArrayLayers")
	textureGetMipLevelCountProc                      = webgpuDLL.NewProc("wgpuTextureGetMipLevelCount")
	textureGetSampleCountProc                        = webgpuDLL.NewProc("wgpuTextureGetSampleCount")
	textureGetDimensionProc                          = webgpuDLL.NewProc("wgpuTextureGetDimension")
	textureGetTextureBindingViewDimensionProc        = webgpuDLL.NewProc("wgpuTextureGetTextureBindingViewDimension")
	textureGetFormatProc                             = webgpuDLL.NewProc("wgpuTextureGetFormat")
	textureGetUsageProc                              = webgpuDLL.NewProc("wgpuTextureGetUsage")
	textureDestroyProc                               = webgpuDLL.NewProc("wgpuTextureDestroy")
	textureAddRefProc                                = webgpuDLL.NewProc("wgpuTextureAddRef")
	textureReleaseProc                               = webgpuDLL.NewProc("wgpuTextureRelease")
	textureViewSetLabelProc                          = webgpuDLL.NewProc("wgpuTextureViewSetLabel")
	textureViewAddRefProc                            = webgpuDLL.NewProc("wgpuTextureViewAddRef")
	textureViewReleaseProc                           = webgpuDLL.NewProc("wgpuTextureViewRelease")
)

func AdapterGetLimits(adapter Adapter, limits *Limits) Status {
	ret, _, _ := adapterGetLimitsProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(limits)),
	)
	return Status(ret)
}

func AdapterHasFeature(adapter Adapter, feature FeatureName) Bool {
	ret, _, _ := adapterHasFeatureProc.Call(
		uintptr(adapter),
		uintptr(feature), // enum.feature_name
	)
	return Bool(ret)
}

func AdapterGetFeatures(adapter Adapter, features *SupportedFeatures) {
	_, _, _ = adapterGetFeaturesProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(features)),
	)
}

func AdapterGetInfo(adapter Adapter, info *AdapterInfo) Status {
	ret, _, _ := adapterGetInfoProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(info)),
	)
	return Status(ret)
}

func AdapterRequestDevice(adapter Adapter, descriptor *DeviceDescriptor, callback RequestDeviceCallbackInfo) {
	_, _, _ = adapterRequestDeviceProc.Call(
		uintptr(adapter),
		uintptr(unsafe.Pointer(descriptor)),
		uintptr(unsafe.Pointer(&callback)),
	)
}

func AdapterAddRef(adapter Adapter) {
	adapterAddRefProc.Call(uintptr(adapter))
}

func AdapterRelease(adapter Adapter) {
	adapterReleaseProc.Call(uintptr(adapter))
}

func BindGroupSetLabel(bindGroup BindGroup, label StringView) {
	_, _, _ = bindGroupSetLabelProc.Call(
		uintptr(bindGroup),
		uintptr(unsafe.Pointer(&label)),
	)
}

func BindGroupAddRef(bindGroup BindGroup) {
	bindGroupAddRefProc.Call(uintptr(bindGroup))
}

func BindGroupRelease(bindGroup BindGroup) {
	bindGroupReleaseProc.Call(uintptr(bindGroup))
}

func BindGroupLayoutSetLabel(bindGroupLayout BindGroupLayout, label StringView) {
	_, _, _ = bindGroupLayoutSetLabelProc.Call(
		uintptr(bindGroupLayout),
		uintptr(unsafe.Pointer(&label)),
	)
}

func BindGroupLayoutAddRef(bindGroupLayout BindGroupLayout) {
	bindGroupLayoutAddRefProc.Call(uintptr(bindGroupLayout))
}

func BindGroupLayoutRelease(bindGroupLayout BindGroupLayout) {
	bindGroupLayoutReleaseProc.Call(uintptr(bindGroupLayout))
}

func BufferMapAsync(buffer Buffer, mode MapMode, offset uintptr, size uintptr, callback BufferMapCallbackInfo) {
	_, _, _ = bufferMapAsyncProc.Call(
		uintptr(buffer),
		uintptr(mode),   // bitflag.map_mode
		uintptr(offset), // usize
		uintptr(size),   // usize
		uintptr(unsafe.Pointer(&callback)),
	)
}

func BufferGetMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {
	ret, _, _ := bufferGetMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(size),   // usize
	)
	return unsafe.Pointer(ret)
}

func BufferGetConstMappedRange(buffer Buffer, offset uintptr, size uintptr) unsafe.Pointer {
	ret, _, _ := bufferGetConstMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(size),   // usize
	)
	return unsafe.Pointer(ret)
}

func BufferReadMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {
	ret, _, _ := bufferReadMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
	return Status(ret)
}

func BufferWriteMappedRange(buffer Buffer, offset uintptr, data unsafe.Pointer, size uintptr) Status {
	ret, _, _ := bufferWriteMappedRangeProc.Call(
		uintptr(buffer),
		uintptr(offset), // usize
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
	return Status(ret)
}

func BufferSetLabel(buffer Buffer, label StringView) {
	_, _, _ = bufferSetLabelProc.Call(
		uintptr(buffer),
		uintptr(unsafe.Pointer(&label)),
	)
}

func BufferGetUsage(buffer Buffer) BufferUsage {
	ret, _, _ := bufferGetUsageProc.Call(
		uintptr(buffer),
	)
	return BufferUsage(ret)
}

func BufferGetSize(buffer Buffer) uint64 {
	ret, _, _ := bufferGetSizeProc.Call(
		uintptr(buffer),
	)
	return uint64(ret)
}

func BufferGetMapState(buffer Buffer) BufferMapState {
	ret, _, _ := bufferGetMapStateProc.Call(
		uintptr(buffer),
	)
	return BufferMapState(ret)
}

func BufferUnmap(buffer Buffer) {
	_, _, _ = bufferUnmapProc.Call(
		uintptr(buffer),
	)
}

func BufferDestroy(buffer Buffer) {
	_, _, _ = bufferDestroyProc.Call(
		uintptr(buffer),
	)
}

func BufferAddRef(buffer Buffer) {
	bufferAddRefProc.Call(uintptr(buffer))
}

func BufferRelease(buffer Buffer) {
	bufferReleaseProc.Call(uintptr(buffer))
}

func CommandBufferSetLabel(commandBuffer CommandBuffer, label StringView) {
	_, _, _ = commandBufferSetLabelProc.Call(
		uintptr(commandBuffer),
		uintptr(unsafe.Pointer(&label)),
	)
}

func CommandBufferAddRef(commandBuffer CommandBuffer) {
	commandBufferAddRefProc.Call(uintptr(commandBuffer))
}

func CommandBufferRelease(commandBuffer CommandBuffer) {
	commandBufferReleaseProc.Call(uintptr(commandBuffer))
}

func CommandEncoderFinish(commandEncoder CommandEncoder, descriptor *CommandBufferDescriptor) CommandBuffer {
	ret, _, _ := commandEncoderFinishProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return CommandBuffer(ret)
}

func CommandEncoderBeginComputePass(commandEncoder CommandEncoder, descriptor *ComputePassDescriptor) ComputePassEncoder {
	ret, _, _ := commandEncoderBeginComputePassProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return ComputePassEncoder(ret)
}

func CommandEncoderBeginRenderPass(commandEncoder CommandEncoder, descriptor *RenderPassDescriptor) RenderPassEncoder {
	ret, _, _ := commandEncoderBeginRenderPassProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return RenderPassEncoder(ret)
}

func CommandEncoderCopyBufferToBuffer(commandEncoder CommandEncoder, source Buffer, sourceOffset uint64, destination Buffer, destinationOffset uint64, size uint64) {
	_, _, _ = commandEncoderCopyBufferToBufferProc.Call(
		uintptr(commandEncoder),
		uintptr(source),            // object.buffer
		uintptr(sourceOffset),      // uint64
		uintptr(destination),       // object.buffer
		uintptr(destinationOffset), // uint64
		uintptr(size),              // uint64
	)
}

func CommandEncoderCopyBufferToTexture(commandEncoder CommandEncoder, source *TexelCopyBufferInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {
	_, _, _ = commandEncoderCopyBufferToTextureProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(copySize)),
	)
}

func CommandEncoderCopyTextureToBuffer(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyBufferInfo, copySize *Extent3D) {
	_, _, _ = commandEncoderCopyTextureToBufferProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(copySize)),
	)
}

func CommandEncoderCopyTextureToTexture(commandEncoder CommandEncoder, source *TexelCopyTextureInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {
	_, _, _ = commandEncoderCopyTextureToTextureProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(source)),
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(copySize)),
	)
}

func CommandEncoderClearBuffer(commandEncoder CommandEncoder, buffer Buffer, offset uint64, size uint64) {
	_, _, _ = commandEncoderClearBufferProc.Call(
		uintptr(commandEncoder),
		uintptr(buffer), // object.buffer
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
}

func CommandEncoderInsertDebugMarker(commandEncoder CommandEncoder, markerLabel StringView) {
	_, _, _ = commandEncoderInsertDebugMarkerProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
}

func CommandEncoderPopDebugGroup(commandEncoder CommandEncoder) {
	_, _, _ = commandEncoderPopDebugGroupProc.Call(
		uintptr(commandEncoder),
	)
}

func CommandEncoderPushDebugGroup(commandEncoder CommandEncoder, groupLabel StringView) {
	_, _, _ = commandEncoderPushDebugGroupProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
}

func CommandEncoderResolveQuerySet(commandEncoder CommandEncoder, querySet QuerySet, firstQuery uint32, queryCount uint32, destination Buffer, destinationOffset uint64) {
	_, _, _ = commandEncoderResolveQuerySetProc.Call(
		uintptr(commandEncoder),
		uintptr(querySet),          // object.query_set
		uintptr(firstQuery),        // uint32
		uintptr(queryCount),        // uint32
		uintptr(destination),       // object.buffer
		uintptr(destinationOffset), // uint64
	)
}

func CommandEncoderWriteTimestamp(commandEncoder CommandEncoder, querySet QuerySet, queryIndex uint32) {
	_, _, _ = commandEncoderWriteTimestampProc.Call(
		uintptr(commandEncoder),
		uintptr(querySet),   // object.query_set
		uintptr(queryIndex), // uint32
	)
}

func CommandEncoderSetLabel(commandEncoder CommandEncoder, label StringView) {
	_, _, _ = commandEncoderSetLabelProc.Call(
		uintptr(commandEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
}

func CommandEncoderAddRef(commandEncoder CommandEncoder) {
	commandEncoderAddRefProc.Call(uintptr(commandEncoder))
}

func CommandEncoderRelease(commandEncoder CommandEncoder) {
	commandEncoderReleaseProc.Call(uintptr(commandEncoder))
}

func ComputePassEncoderInsertDebugMarker(computePassEncoder ComputePassEncoder, markerLabel StringView) {
	_, _, _ = computePassEncoderInsertDebugMarkerProc.Call(
		uintptr(computePassEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
}

func ComputePassEncoderPopDebugGroup(computePassEncoder ComputePassEncoder) {
	_, _, _ = computePassEncoderPopDebugGroupProc.Call(
		uintptr(computePassEncoder),
	)
}

func ComputePassEncoderPushDebugGroup(computePassEncoder ComputePassEncoder, groupLabel StringView) {
	_, _, _ = computePassEncoderPushDebugGroupProc.Call(
		uintptr(computePassEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
}

func ComputePassEncoderSetPipeline(computePassEncoder ComputePassEncoder, pipeline ComputePipeline) {
	_, _, _ = computePassEncoderSetPipelineProc.Call(
		uintptr(computePassEncoder),
		uintptr(pipeline), // object.compute_pipeline
	)
}

func ComputePassEncoderSetBindGroup(computePassEncoder ComputePassEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {
	_, _, _ = computePassEncoderSetBindGroupProc.Call(
		uintptr(computePassEncoder),
		uintptr(groupIndex), // uint32
		uintptr(group),      // object.bind_group
		uintptr(len(dynamicOffsets)),
		uintptr(unsafe.Pointer(unsafe.SliceData(dynamicOffsets))),
	)
}

func ComputePassEncoderSetImmediates(computePassEncoder ComputePassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	_, _, _ = computePassEncoderSetImmediatesProc.Call(
		uintptr(computePassEncoder),
		uintptr(offset), // uint32
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
}

func ComputePassEncoderDispatchWorkgroups(computePassEncoder ComputePassEncoder, workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) {
	_, _, _ = computePassEncoderDispatchWorkgroupsProc.Call(
		uintptr(computePassEncoder),
		uintptr(workgroupCountX), // uint32
		uintptr(workgroupCountY), // uint32
		uintptr(workgroupCountZ), // uint32
	)
}

func ComputePassEncoderDispatchWorkgroupsIndirect(computePassEncoder ComputePassEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, _ = computePassEncoderDispatchWorkgroupsIndirectProc.Call(
		uintptr(computePassEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
}

func ComputePassEncoderEnd(computePassEncoder ComputePassEncoder) {
	_, _, _ = computePassEncoderEndProc.Call(
		uintptr(computePassEncoder),
	)
}

func ComputePassEncoderSetLabel(computePassEncoder ComputePassEncoder, label StringView) {
	_, _, _ = computePassEncoderSetLabelProc.Call(
		uintptr(computePassEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
}

func ComputePassEncoderAddRef(computePassEncoder ComputePassEncoder) {
	computePassEncoderAddRefProc.Call(uintptr(computePassEncoder))
}

func ComputePassEncoderRelease(computePassEncoder ComputePassEncoder) {
	computePassEncoderReleaseProc.Call(uintptr(computePassEncoder))
}

func ComputePipelineGetBindGroupLayout(computePipeline ComputePipeline, groupIndex uint32) BindGroupLayout {
	ret, _, _ := computePipelineGetBindGroupLayoutProc.Call(
		uintptr(computePipeline),
		uintptr(groupIndex), // uint32
	)
	return BindGroupLayout(ret)
}

func ComputePipelineSetLabel(computePipeline ComputePipeline, label StringView) {
	_, _, _ = computePipelineSetLabelProc.Call(
		uintptr(computePipeline),
		uintptr(unsafe.Pointer(&label)),
	)
}

func ComputePipelineAddRef(computePipeline ComputePipeline) {
	computePipelineAddRefProc.Call(uintptr(computePipeline))
}

func ComputePipelineRelease(computePipeline ComputePipeline) {
	computePipelineReleaseProc.Call(uintptr(computePipeline))
}

func DeviceCreateBindGroup(device Device, descriptor *BindGroupDescriptor) BindGroup {
	ret, _, _ := deviceCreateBindGroupProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return BindGroup(ret)
}

func DeviceCreateBindGroupLayout(device Device, descriptor *BindGroupLayoutDescriptor) BindGroupLayout {
	ret, _, _ := deviceCreateBindGroupLayoutProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return BindGroupLayout(ret)
}

func DeviceCreateBuffer(device Device, descriptor *BufferDescriptor) Buffer {
	ret, _, _ := deviceCreateBufferProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return Buffer(ret)
}

func DeviceCreateCommandEncoder(device Device, descriptor *CommandEncoderDescriptor) CommandEncoder {
	ret, _, _ := deviceCreateCommandEncoderProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return CommandEncoder(ret)
}

func DeviceCreateComputePipeline(device Device, descriptor *ComputePipelineDescriptor) ComputePipeline {
	ret, _, _ := deviceCreateComputePipelineProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return ComputePipeline(ret)
}

func DeviceCreateComputePipelineAsync(device Device, descriptor *ComputePipelineDescriptor, callback CreateComputePipelineAsyncCallbackInfo) {
	_, _, _ = deviceCreateComputePipelineAsyncProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
		uintptr(unsafe.Pointer(&callback)),
	)
}

func DeviceCreatePipelineLayout(device Device, descriptor *PipelineLayoutDescriptor) PipelineLayout {
	ret, _, _ := deviceCreatePipelineLayoutProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return PipelineLayout(ret)
}

func DeviceCreateQuerySet(device Device, descriptor *QuerySetDescriptor) QuerySet {
	ret, _, _ := deviceCreateQuerySetProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return QuerySet(ret)
}

func DeviceCreateRenderPipelineAsync(device Device, descriptor *RenderPipelineDescriptor, callback CreateRenderPipelineAsyncCallbackInfo) {
	_, _, _ = deviceCreateRenderPipelineAsyncProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
		uintptr(unsafe.Pointer(&callback)),
	)
}

func DeviceCreateRenderBundleEncoder(device Device, descriptor *RenderBundleEncoderDescriptor) RenderBundleEncoder {
	ret, _, _ := deviceCreateRenderBundleEncoderProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return RenderBundleEncoder(ret)
}

func DeviceCreateRenderPipeline(device Device, descriptor *RenderPipelineDescriptor) RenderPipeline {
	ret, _, _ := deviceCreateRenderPipelineProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return RenderPipeline(ret)
}

func DeviceCreateSampler(device Device, descriptor *SamplerDescriptor) Sampler {
	ret, _, _ := deviceCreateSamplerProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return Sampler(ret)
}

func DeviceCreateShaderModule(device Device, descriptor *ShaderModuleDescriptor) ShaderModule {
	ret, _, _ := deviceCreateShaderModuleProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return ShaderModule(ret)
}

func DeviceCreateTexture(device Device, descriptor *TextureDescriptor) Texture {
	ret, _, _ := deviceCreateTextureProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return Texture(ret)
}

func DeviceDestroy(device Device) {
	_, _, _ = deviceDestroyProc.Call(
		uintptr(device),
	)
}

func DeviceGetLostFuture(device Device) Future {
	ret, _, _ := deviceGetLostFutureProc.Call(
		uintptr(device),
	)
	return *(*Future)(unsafe.Pointer(new(uint64(ret))))
}

func DeviceGetLimits(device Device, limits *Limits) Status {
	ret, _, _ := deviceGetLimitsProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(limits)),
	)
	return Status(ret)
}

func DeviceHasFeature(device Device, feature FeatureName) Bool {
	ret, _, _ := deviceHasFeatureProc.Call(
		uintptr(device),
		uintptr(feature), // enum.feature_name
	)
	return Bool(ret)
}

func DeviceGetFeatures(device Device, features *SupportedFeatures) {
	_, _, _ = deviceGetFeaturesProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(features)),
	)
}

func DeviceGetAdapterInfo(device Device, adapterInfo *AdapterInfo) Status {
	ret, _, _ := deviceGetAdapterInfoProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(adapterInfo)),
	)
	return Status(ret)
}

func DeviceGetQueue(device Device) Queue {
	ret, _, _ := deviceGetQueueProc.Call(
		uintptr(device),
	)
	return Queue(ret)
}

func DevicePushErrorScope(device Device, filter ErrorFilter) {
	_, _, _ = devicePushErrorScopeProc.Call(
		uintptr(device),
		uintptr(filter), // enum.error_filter
	)
}

func DevicePopErrorScope(device Device, callback PopErrorScopeCallbackInfo) {
	_, _, _ = devicePopErrorScopeProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(&callback)),
	)
}

func DeviceSetLabel(device Device, label StringView) {
	_, _, _ = deviceSetLabelProc.Call(
		uintptr(device),
		uintptr(unsafe.Pointer(&label)),
	)
}

func DeviceAddRef(device Device) {
	deviceAddRefProc.Call(uintptr(device))
}

func DeviceRelease(device Device) {
	deviceReleaseProc.Call(uintptr(device))
}

func ExternalTextureSetLabel(externalTexture ExternalTexture, label StringView) {
	_, _, _ = externalTextureSetLabelProc.Call(
		uintptr(externalTexture),
		uintptr(unsafe.Pointer(&label)),
	)
}

func ExternalTextureAddRef(externalTexture ExternalTexture) {
	externalTextureAddRefProc.Call(uintptr(externalTexture))
}

func ExternalTextureRelease(externalTexture ExternalTexture) {
	externalTextureReleaseProc.Call(uintptr(externalTexture))
}

func InstanceCreateSurface(instance Instance, descriptor *SurfaceDescriptor) Surface {
	ret, _, _ := instanceCreateSurfaceProc.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return Surface(ret)
}

func InstanceGetWGSLLanguageFeatures(instance Instance, features *SupportedWGSLLanguageFeatures) {
	_, _, _ = instanceGetWGSLLanguageFeaturesProc.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(features)),
	)
}

func InstanceHasWGSLLanguageFeature(instance Instance, feature WGSLLanguageFeatureName) Bool {
	ret, _, _ := instanceHasWGSLLanguageFeatureProc.Call(
		uintptr(instance),
		uintptr(feature), // enum.WGSL_language_feature_name
	)
	return Bool(ret)
}

func InstanceProcessEvents(instance Instance) {
	_, _, _ = instanceProcessEventsProc.Call(
		uintptr(instance),
	)
}

func InstanceRequestAdapter(instance Instance, options *RequestAdapterOptions, callback RequestAdapterCallbackInfo) {
	_, _, _ = instanceRequestAdapterProc.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(&callback)),
	)
}

func InstanceWaitAny(instance Instance, futureCount uintptr, futures *FutureWaitInfo, timeoutNS uint64) WaitStatus {
	ret, _, _ := instanceWaitAnyProc.Call(
		uintptr(instance),
		uintptr(futureCount), // usize
		uintptr(unsafe.Pointer(futures)),
		uintptr(timeoutNS), // uint64
	)
	return WaitStatus(ret)
}

func InstanceAddRef(instance Instance) {
	instanceAddRefProc.Call(uintptr(instance))
}

func InstanceRelease(instance Instance) {
	instanceReleaseProc.Call(uintptr(instance))
}

func PipelineLayoutSetLabel(pipelineLayout PipelineLayout, label StringView) {
	_, _, _ = pipelineLayoutSetLabelProc.Call(
		uintptr(pipelineLayout),
		uintptr(unsafe.Pointer(&label)),
	)
}

func PipelineLayoutAddRef(pipelineLayout PipelineLayout) {
	pipelineLayoutAddRefProc.Call(uintptr(pipelineLayout))
}

func PipelineLayoutRelease(pipelineLayout PipelineLayout) {
	pipelineLayoutReleaseProc.Call(uintptr(pipelineLayout))
}

func QuerySetSetLabel(querySet QuerySet, label StringView) {
	_, _, _ = querySetSetLabelProc.Call(
		uintptr(querySet),
		uintptr(unsafe.Pointer(&label)),
	)
}

func QuerySetGetType(querySet QuerySet) QueryType {
	ret, _, _ := querySetGetTypeProc.Call(
		uintptr(querySet),
	)
	return QueryType(ret)
}

func QuerySetGetCount(querySet QuerySet) uint32 {
	ret, _, _ := querySetGetCountProc.Call(
		uintptr(querySet),
	)
	return uint32(ret)
}

func QuerySetDestroy(querySet QuerySet) {
	_, _, _ = querySetDestroyProc.Call(
		uintptr(querySet),
	)
}

func QuerySetAddRef(querySet QuerySet) {
	querySetAddRefProc.Call(uintptr(querySet))
}

func QuerySetRelease(querySet QuerySet) {
	querySetReleaseProc.Call(uintptr(querySet))
}

func QueueSubmit(queue Queue, commands []CommandBuffer) {
	_, _, _ = queueSubmitProc.Call(
		uintptr(queue),
		uintptr(len(commands)),
		uintptr(unsafe.Pointer(unsafe.SliceData(commands))),
	)
}

func QueueOnSubmittedWorkDone(queue Queue, callback QueueWorkDoneCallbackInfo) {
	_, _, _ = queueOnSubmittedWorkDoneProc.Call(
		uintptr(queue),
		uintptr(unsafe.Pointer(&callback)),
	)
}

func QueueWriteBuffer(queue Queue, buffer Buffer, bufferOffset uint64, data unsafe.Pointer, size uintptr) {
	_, _, _ = queueWriteBufferProc.Call(
		uintptr(queue),
		uintptr(buffer),       // object.buffer
		uintptr(bufferOffset), // uint64
		uintptr(data),         // c_void_data_ptr
		uintptr(size),         // usize
	)
}

func QueueWriteTexture(queue Queue, destination *TexelCopyTextureInfo, data unsafe.Pointer, dataSize uintptr, dataLayout *TexelCopyBufferLayout, writeSize *Extent3D) {
	_, _, _ = queueWriteTextureProc.Call(
		uintptr(queue),
		uintptr(unsafe.Pointer(destination)),
		uintptr(data),     // c_void_data_ptr
		uintptr(dataSize), // usize
		uintptr(unsafe.Pointer(dataLayout)),
		uintptr(unsafe.Pointer(writeSize)),
	)
}

func QueueSetLabel(queue Queue, label StringView) {
	_, _, _ = queueSetLabelProc.Call(
		uintptr(queue),
		uintptr(unsafe.Pointer(&label)),
	)
}

func QueueAddRef(queue Queue) {
	queueAddRefProc.Call(uintptr(queue))
}

func QueueRelease(queue Queue) {
	queueReleaseProc.Call(uintptr(queue))
}

func RenderBundleSetLabel(renderBundle RenderBundle, label StringView) {
	_, _, _ = renderBundleSetLabelProc.Call(
		uintptr(renderBundle),
		uintptr(unsafe.Pointer(&label)),
	)
}

func RenderBundleAddRef(renderBundle RenderBundle) {
	renderBundleAddRefProc.Call(uintptr(renderBundle))
}

func RenderBundleRelease(renderBundle RenderBundle) {
	renderBundleReleaseProc.Call(uintptr(renderBundle))
}

func RenderBundleEncoderSetPipeline(renderBundleEncoder RenderBundleEncoder, pipeline RenderPipeline) {
	_, _, _ = renderBundleEncoderSetPipelineProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(pipeline), // object.render_pipeline
	)
}

func RenderBundleEncoderSetBindGroup(renderBundleEncoder RenderBundleEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {
	_, _, _ = renderBundleEncoderSetBindGroupProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(groupIndex), // uint32
		uintptr(group),      // object.bind_group
		uintptr(len(dynamicOffsets)),
		uintptr(unsafe.Pointer(unsafe.SliceData(dynamicOffsets))),
	)
}

func RenderBundleEncoderSetImmediates(renderBundleEncoder RenderBundleEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	_, _, _ = renderBundleEncoderSetImmediatesProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(offset), // uint32
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
}

func RenderBundleEncoderDraw(renderBundleEncoder RenderBundleEncoder, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	_, _, _ = renderBundleEncoderDrawProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(vertexCount),   // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstVertex),   // uint32
		uintptr(firstInstance), // uint32
	)
}

func RenderBundleEncoderDrawIndexed(renderBundleEncoder RenderBundleEncoder, indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	_, _, _ = renderBundleEncoderDrawIndexedProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(indexCount),    // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstIndex),    // uint32
		uintptr(baseVertex),    // int32
		uintptr(firstInstance), // uint32
	)
}

func RenderBundleEncoderDrawIndirect(renderBundleEncoder RenderBundleEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, _ = renderBundleEncoderDrawIndirectProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
}

func RenderBundleEncoderDrawIndexedIndirect(renderBundleEncoder RenderBundleEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, _ = renderBundleEncoderDrawIndexedIndirectProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
}

func RenderBundleEncoderInsertDebugMarker(renderBundleEncoder RenderBundleEncoder, markerLabel StringView) {
	_, _, _ = renderBundleEncoderInsertDebugMarkerProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
}

func RenderBundleEncoderPopDebugGroup(renderBundleEncoder RenderBundleEncoder) {
	_, _, _ = renderBundleEncoderPopDebugGroupProc.Call(
		uintptr(renderBundleEncoder),
	)
}

func RenderBundleEncoderPushDebugGroup(renderBundleEncoder RenderBundleEncoder, groupLabel StringView) {
	_, _, _ = renderBundleEncoderPushDebugGroupProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
}

func RenderBundleEncoderSetVertexBuffer(renderBundleEncoder RenderBundleEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {
	_, _, _ = renderBundleEncoderSetVertexBufferProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(slot),   // uint32
		uintptr(buffer), // object.buffer
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
}

func RenderBundleEncoderSetIndexBuffer(renderBundleEncoder RenderBundleEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	_, _, _ = renderBundleEncoderSetIndexBufferProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(buffer), // object.buffer
		uintptr(format), // enum.index_format
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
}

func RenderBundleEncoderFinish(renderBundleEncoder RenderBundleEncoder, descriptor *RenderBundleDescriptor) RenderBundle {
	ret, _, _ := renderBundleEncoderFinishProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return RenderBundle(ret)
}

func RenderBundleEncoderSetLabel(renderBundleEncoder RenderBundleEncoder, label StringView) {
	_, _, _ = renderBundleEncoderSetLabelProc.Call(
		uintptr(renderBundleEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
}

func RenderBundleEncoderAddRef(renderBundleEncoder RenderBundleEncoder) {
	renderBundleEncoderAddRefProc.Call(uintptr(renderBundleEncoder))
}

func RenderBundleEncoderRelease(renderBundleEncoder RenderBundleEncoder) {
	renderBundleEncoderReleaseProc.Call(uintptr(renderBundleEncoder))
}

func RenderPassEncoderSetPipeline(renderPassEncoder RenderPassEncoder, pipeline RenderPipeline) {
	_, _, _ = renderPassEncoderSetPipelineProc.Call(
		uintptr(renderPassEncoder),
		uintptr(pipeline), // object.render_pipeline
	)
}

func RenderPassEncoderSetBindGroup(renderPassEncoder RenderPassEncoder, groupIndex uint32, group BindGroup, dynamicOffsets []uint32) {
	_, _, _ = renderPassEncoderSetBindGroupProc.Call(
		uintptr(renderPassEncoder),
		uintptr(groupIndex), // uint32
		uintptr(group),      // object.bind_group
		uintptr(len(dynamicOffsets)),
		uintptr(unsafe.Pointer(unsafe.SliceData(dynamicOffsets))),
	)
}

func RenderPassEncoderSetImmediates(renderPassEncoder RenderPassEncoder, offset uint32, data unsafe.Pointer, size uintptr) {
	_, _, _ = renderPassEncoderSetImmediatesProc.Call(
		uintptr(renderPassEncoder),
		uintptr(offset), // uint32
		uintptr(data),   // c_void_data_ptr
		uintptr(size),   // usize
	)
}

func RenderPassEncoderDraw(renderPassEncoder RenderPassEncoder, vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) {
	_, _, _ = renderPassEncoderDrawProc.Call(
		uintptr(renderPassEncoder),
		uintptr(vertexCount),   // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstVertex),   // uint32
		uintptr(firstInstance), // uint32
	)
}

func RenderPassEncoderDrawIndexed(renderPassEncoder RenderPassEncoder, indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	_, _, _ = renderPassEncoderDrawIndexedProc.Call(
		uintptr(renderPassEncoder),
		uintptr(indexCount),    // uint32
		uintptr(instanceCount), // uint32
		uintptr(firstIndex),    // uint32
		uintptr(baseVertex),    // int32
		uintptr(firstInstance), // uint32
	)
}

func RenderPassEncoderDrawIndirect(renderPassEncoder RenderPassEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, _ = renderPassEncoderDrawIndirectProc.Call(
		uintptr(renderPassEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
}

func RenderPassEncoderDrawIndexedIndirect(renderPassEncoder RenderPassEncoder, indirectBuffer Buffer, indirectOffset uint64) {
	_, _, _ = renderPassEncoderDrawIndexedIndirectProc.Call(
		uintptr(renderPassEncoder),
		uintptr(indirectBuffer), // object.buffer
		uintptr(indirectOffset), // uint64
	)
}

func RenderPassEncoderExecuteBundles(renderPassEncoder RenderPassEncoder, bundles []RenderBundle) {
	_, _, _ = renderPassEncoderExecuteBundlesProc.Call(
		uintptr(renderPassEncoder),
		uintptr(len(bundles)),
		uintptr(unsafe.Pointer(unsafe.SliceData(bundles))),
	)
}

func RenderPassEncoderInsertDebugMarker(renderPassEncoder RenderPassEncoder, markerLabel StringView) {
	_, _, _ = renderPassEncoderInsertDebugMarkerProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(&markerLabel)),
	)
}

func RenderPassEncoderPopDebugGroup(renderPassEncoder RenderPassEncoder) {
	_, _, _ = renderPassEncoderPopDebugGroupProc.Call(
		uintptr(renderPassEncoder),
	)
}

func RenderPassEncoderPushDebugGroup(renderPassEncoder RenderPassEncoder, groupLabel StringView) {
	_, _, _ = renderPassEncoderPushDebugGroupProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(&groupLabel)),
	)
}

func RenderPassEncoderSetStencilReference(renderPassEncoder RenderPassEncoder, reference uint32) {
	_, _, _ = renderPassEncoderSetStencilReferenceProc.Call(
		uintptr(renderPassEncoder),
		uintptr(reference), // uint32
	)
}

func RenderPassEncoderSetBlendConstant(renderPassEncoder RenderPassEncoder, color *Color) {
	_, _, _ = renderPassEncoderSetBlendConstantProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(color)),
	)
}

func RenderPassEncoderSetViewport(renderPassEncoder RenderPassEncoder, x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) {
	_, _, _ = renderPassEncoderSetViewportProc.Call(
		uintptr(renderPassEncoder),
		uintptr(x),        // float32
		uintptr(y),        // float32
		uintptr(width),    // float32
		uintptr(height),   // float32
		uintptr(minDepth), // float32
		uintptr(maxDepth), // float32
	)
}

func RenderPassEncoderSetScissorRect(renderPassEncoder RenderPassEncoder, x uint32, y uint32, width uint32, height uint32) {
	_, _, _ = renderPassEncoderSetScissorRectProc.Call(
		uintptr(renderPassEncoder),
		uintptr(x),      // uint32
		uintptr(y),      // uint32
		uintptr(width),  // uint32
		uintptr(height), // uint32
	)
}

func RenderPassEncoderSetVertexBuffer(renderPassEncoder RenderPassEncoder, slot uint32, buffer Buffer, offset uint64, size uint64) {
	_, _, _ = renderPassEncoderSetVertexBufferProc.Call(
		uintptr(renderPassEncoder),
		uintptr(slot),   // uint32
		uintptr(buffer), // object.buffer
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
}

func RenderPassEncoderSetIndexBuffer(renderPassEncoder RenderPassEncoder, buffer Buffer, format IndexFormat, offset uint64, size uint64) {
	_, _, _ = renderPassEncoderSetIndexBufferProc.Call(
		uintptr(renderPassEncoder),
		uintptr(buffer), // object.buffer
		uintptr(format), // enum.index_format
		uintptr(offset), // uint64
		uintptr(size),   // uint64
	)
}

func RenderPassEncoderBeginOcclusionQuery(renderPassEncoder RenderPassEncoder, queryIndex uint32) {
	_, _, _ = renderPassEncoderBeginOcclusionQueryProc.Call(
		uintptr(renderPassEncoder),
		uintptr(queryIndex), // uint32
	)
}

func RenderPassEncoderEndOcclusionQuery(renderPassEncoder RenderPassEncoder) {
	_, _, _ = renderPassEncoderEndOcclusionQueryProc.Call(
		uintptr(renderPassEncoder),
	)
}

func RenderPassEncoderEnd(renderPassEncoder RenderPassEncoder) {
	_, _, _ = renderPassEncoderEndProc.Call(
		uintptr(renderPassEncoder),
	)
}

func RenderPassEncoderSetLabel(renderPassEncoder RenderPassEncoder, label StringView) {
	_, _, _ = renderPassEncoderSetLabelProc.Call(
		uintptr(renderPassEncoder),
		uintptr(unsafe.Pointer(&label)),
	)
}

func RenderPassEncoderAddRef(renderPassEncoder RenderPassEncoder) {
	renderPassEncoderAddRefProc.Call(uintptr(renderPassEncoder))
}

func RenderPassEncoderRelease(renderPassEncoder RenderPassEncoder) {
	renderPassEncoderReleaseProc.Call(uintptr(renderPassEncoder))
}

func RenderPipelineGetBindGroupLayout(renderPipeline RenderPipeline, groupIndex uint32) BindGroupLayout {
	ret, _, _ := renderPipelineGetBindGroupLayoutProc.Call(
		uintptr(renderPipeline),
		uintptr(groupIndex), // uint32
	)
	return BindGroupLayout(ret)
}

func RenderPipelineSetLabel(renderPipeline RenderPipeline, label StringView) {
	_, _, _ = renderPipelineSetLabelProc.Call(
		uintptr(renderPipeline),
		uintptr(unsafe.Pointer(&label)),
	)
}

func RenderPipelineAddRef(renderPipeline RenderPipeline) {
	renderPipelineAddRefProc.Call(uintptr(renderPipeline))
}

func RenderPipelineRelease(renderPipeline RenderPipeline) {
	renderPipelineReleaseProc.Call(uintptr(renderPipeline))
}

func SamplerSetLabel(sampler Sampler, label StringView) {
	_, _, _ = samplerSetLabelProc.Call(
		uintptr(sampler),
		uintptr(unsafe.Pointer(&label)),
	)
}

func SamplerAddRef(sampler Sampler) {
	samplerAddRefProc.Call(uintptr(sampler))
}

func SamplerRelease(sampler Sampler) {
	samplerReleaseProc.Call(uintptr(sampler))
}

func ShaderModuleGetCompilationInfo(shaderModule ShaderModule, callback CompilationInfoCallbackInfo) {
	_, _, _ = shaderModuleGetCompilationInfoProc.Call(
		uintptr(shaderModule),
		uintptr(unsafe.Pointer(&callback)),
	)
}

func ShaderModuleSetLabel(shaderModule ShaderModule, label StringView) {
	_, _, _ = shaderModuleSetLabelProc.Call(
		uintptr(shaderModule),
		uintptr(unsafe.Pointer(&label)),
	)
}

func ShaderModuleAddRef(shaderModule ShaderModule) {
	shaderModuleAddRefProc.Call(uintptr(shaderModule))
}

func ShaderModuleRelease(shaderModule ShaderModule) {
	shaderModuleReleaseProc.Call(uintptr(shaderModule))
}

func SurfaceConfigure(surface Surface, config *SurfaceConfiguration) {
	_, _, _ = surfaceConfigureProc.Call(
		uintptr(surface),
		uintptr(unsafe.Pointer(config)),
	)
}

func SurfaceGetCapabilities(surface Surface, adapter Adapter, capabilities *SurfaceCapabilities) Status {
	ret, _, _ := surfaceGetCapabilitiesProc.Call(
		uintptr(surface),
		uintptr(adapter), // object.adapter
		uintptr(unsafe.Pointer(capabilities)),
	)
	return Status(ret)
}

func SurfaceGetCurrentTexture(surface Surface, surfaceTexture *SurfaceTexture) {
	_, _, _ = surfaceGetCurrentTextureProc.Call(
		uintptr(surface),
		uintptr(unsafe.Pointer(surfaceTexture)),
	)
}

func SurfacePresent(surface Surface) Status {
	ret, _, _ := surfacePresentProc.Call(
		uintptr(surface),
	)
	return Status(ret)
}

func SurfaceUnconfigure(surface Surface) {
	_, _, _ = surfaceUnconfigureProc.Call(
		uintptr(surface),
	)
}

func SurfaceSetLabel(surface Surface, label StringView) {
	_, _, _ = surfaceSetLabelProc.Call(
		uintptr(surface),
		uintptr(unsafe.Pointer(&label)),
	)
}

func SurfaceAddRef(surface Surface) {
	surfaceAddRefProc.Call(uintptr(surface))
}

func SurfaceRelease(surface Surface) {
	surfaceReleaseProc.Call(uintptr(surface))
}

func TextureCreateView(texture Texture, descriptor *TextureViewDescriptor) TextureView {
	ret, _, _ := textureCreateViewProc.Call(
		uintptr(texture),
		uintptr(unsafe.Pointer(descriptor)),
	)
	return TextureView(ret)
}

func TextureSetLabel(texture Texture, label StringView) {
	_, _, _ = textureSetLabelProc.Call(
		uintptr(texture),
		uintptr(unsafe.Pointer(&label)),
	)
}

func TextureGetWidth(texture Texture) uint32 {
	ret, _, _ := textureGetWidthProc.Call(
		uintptr(texture),
	)
	return uint32(ret)
}

func TextureGetHeight(texture Texture) uint32 {
	ret, _, _ := textureGetHeightProc.Call(
		uintptr(texture),
	)
	return uint32(ret)
}

func TextureGetDepthOrArrayLayers(texture Texture) uint32 {
	ret, _, _ := textureGetDepthOrArrayLayersProc.Call(
		uintptr(texture),
	)
	return uint32(ret)
}

func TextureGetMipLevelCount(texture Texture) uint32 {
	ret, _, _ := textureGetMipLevelCountProc.Call(
		uintptr(texture),
	)
	return uint32(ret)
}

func TextureGetSampleCount(texture Texture) uint32 {
	ret, _, _ := textureGetSampleCountProc.Call(
		uintptr(texture),
	)
	return uint32(ret)
}

func TextureGetDimension(texture Texture) TextureDimension {
	ret, _, _ := textureGetDimensionProc.Call(
		uintptr(texture),
	)
	return TextureDimension(ret)
}

func TextureGetTextureBindingViewDimension(texture Texture) TextureViewDimension {
	ret, _, _ := textureGetTextureBindingViewDimensionProc.Call(
		uintptr(texture),
	)
	return TextureViewDimension(ret)
}

func TextureGetFormat(texture Texture) TextureFormat {
	ret, _, _ := textureGetFormatProc.Call(
		uintptr(texture),
	)
	return TextureFormat(ret)
}

func TextureGetUsage(texture Texture) TextureUsage {
	ret, _, _ := textureGetUsageProc.Call(
		uintptr(texture),
	)
	return TextureUsage(ret)
}

func TextureDestroy(texture Texture) {
	_, _, _ = textureDestroyProc.Call(
		uintptr(texture),
	)
}

func TextureAddRef(texture Texture) {
	textureAddRefProc.Call(uintptr(texture))
}

func TextureRelease(texture Texture) {
	textureReleaseProc.Call(uintptr(texture))
}

func TextureViewSetLabel(textureView TextureView, label StringView) {
	_, _, _ = textureViewSetLabelProc.Call(
		uintptr(textureView),
		uintptr(unsafe.Pointer(&label)),
	)
}

func TextureViewAddRef(textureView TextureView) {
	textureViewAddRefProc.Call(uintptr(textureView))
}

func TextureViewRelease(textureView TextureView) {
	textureViewReleaseProc.Call(uintptr(textureView))
}
