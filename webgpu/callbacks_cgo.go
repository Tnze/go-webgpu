// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo

package webgpu

// #cgo CFLAGS: -I${SRCDIR}/../webgpu-headers
// #include <webgpu.h>
//
// void gowebgpuBufferMapCallback(WGPUMapAsyncStatus, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuCompilationInfoCallback(WGPUCompilationInfoRequestStatus, WGPUCompilationInfo*, void* userdata1, void* userdata2);
// void gowebgpuCreateComputePipelineAsyncCallback(WGPUCreatePipelineAsyncStatus, WGPUComputePipeline, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuCreateRenderPipelineAsyncCallback(WGPUCreatePipelineAsyncStatus, WGPURenderPipeline, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuDeviceLostCallback(WGPUDevice*, WGPUDeviceLostReason, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuPopErrorScopeCallback(WGPUPopErrorScopeStatus, WGPUErrorType, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuQueueWorkDoneCallback(WGPUQueueWorkDoneStatus, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuRequestAdapterCallback(WGPURequestAdapterStatus, WGPUAdapter, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuRequestDeviceCallback(WGPURequestDeviceStatus, WGPUDevice, WGPUStringView, void* userdata1, void* userdata2);
// void gowebgpuUncapturedErrorCallback(WGPUDevice*, WGPUErrorType, WGPUStringView, void* userdata1, void* userdata2);
import "C"

import (
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu/sys"
)

//export gowebgpuBufferMapCallback
func gowebgpuBufferMapCallback(status C.WGPUMapAsyncStatus, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*bufferMapCallback)(userdata1).C <- bufferMapCallbackResult{
		status:  MapAsyncStatus(status),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

//export gowebgpuCompilationInfoCallback
func gowebgpuCompilationInfoCallback(status C.WGPUCompilationInfoRequestStatus, compilationInfo *C.WGPUCompilationInfo, userdata1, userdata2 unsafe.Pointer) {
	(*compilationInfoCallback)(userdata1).C <- compilationInfoCallbackResult{
		status:          CompilationInfoRequestStatus(status),
		compilationInfo: (*CompilationInfo)(unsafe.Pointer(compilationInfo)),
	}
}

//export gowebgpuCreateComputePipelineAsyncCallback
func gowebgpuCreateComputePipelineAsyncCallback(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPUComputePipeline, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*createComputePipelineAsyncCallback)(userdata1).C <- createComputePipelineAsyncCallbackResult{
		status:   CreatePipelineAsyncStatus(status),
		pipeline: new(ComputePipeline{inner: sys.ComputePipeline(pipeline)}).owned(),
		message:  C.GoStringN(message.data, C.int(message.length)),
	}
}

//export gowebgpuCreateRenderPipelineAsyncCallback
func gowebgpuCreateRenderPipelineAsyncCallback(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPURenderPipeline, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*createRenderPipelineAsyncCallback)(userdata1).C <- createRenderPipelineAsyncCallbackResult{
		status:   CreatePipelineAsyncStatus(status),
		pipeline: new(RenderPipeline{inner: sys.RenderPipeline(pipeline)}).owned(),
		message:  C.GoStringN(message.data, C.int(message.length)),
	}
}

//export gowebgpuDeviceLostCallback
func gowebgpuDeviceLostCallback(device *C.WGPUDevice, reason C.WGPUDeviceLostReason, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*deviceLostCallback)(userdata1).C <- deviceLostCallbackResult{
		device:  new(Device{inner: sys.Device(device)}).owned(),
		reason:  DeviceLostReason(reason),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

//export gowebgpuPopErrorScopeCallback
func gowebgpuPopErrorScopeCallback(status C.WGPUPopErrorScopeStatus, errorType C.WGPUErrorType, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*popErrorScopeCallback)(userdata1).C <- popErrorScopeCallbackResult{
		status:    PopErrorScopeStatus(status),
		errorType: ErrorType(errorType),
		message:   C.GoStringN(message.data, C.int(message.length)),
	}
}

//export gowebgpuQueueWorkDoneCallback
func gowebgpuQueueWorkDoneCallback(status C.WGPUQueueWorkDoneStatus, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*queueWorkDoneCallback)(userdata1).C <- queueWorkDoneCallbackResult{
		status:  QueueWorkDoneStatus(status),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

//export gowebgpuRequestAdapterCallback
func gowebgpuRequestAdapterCallback(status C.WGPURequestAdapterStatus, adapter C.WGPUAdapter, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*requestAdapterCallback)(userdata1).C <- requestAdapterCallbackResult{
		status:  RequestAdapterStatus(status),
		adapter: new(Adapter{inner: sys.Adapter(adapter)}).owned(),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

//export gowebgpuRequestDeviceCallback
func gowebgpuRequestDeviceCallback(status C.WGPURequestDeviceStatus, device C.WGPUDevice, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*requestDeviceCallback)(userdata1).C <- requestDeviceCallbackResult{
		status:  RequestDeviceStatus(status),
		device:  new(Device{inner: sys.Device(device)}).owned(),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

type bufferMapCallback struct {
	C chan bufferMapCallbackResult
}

func (b *bufferMapCallback) info() sys.BufferMapCallbackInfo {
	return sys.BufferMapCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuBufferMapCallback,
		Userdata1: unsafe.Pointer(b),
	}
}

type compilationInfoCallback struct {
	C chan compilationInfoCallbackResult
}

func (c *compilationInfoCallback) info() sys.CompilationInfoCallbackInfo {
	return sys.CompilationInfoCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuCompilationInfoCallback,
		Userdata1: unsafe.Pointer(c),
	}
}

type createComputePipelineAsyncCallback struct {
	C chan createComputePipelineAsyncCallbackResult
}

func (c *createComputePipelineAsyncCallback) info() sys.CreateComputePipelineAsyncCallbackInfo {
	return sys.CreateComputePipelineAsyncCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuCreateComputePipelineAsyncCallback,
		Userdata1: unsafe.Pointer(c),
	}
}

type createRenderPipelineAsyncCallback struct {
	C chan createRenderPipelineAsyncCallbackResult
}

func (c *createRenderPipelineAsyncCallback) info() sys.CreateRenderPipelineAsyncCallbackInfo {
	return sys.CreateRenderPipelineAsyncCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuCreateRenderPipelineAsyncCallback,
		Userdata1: unsafe.Pointer(c),
	}
}

type deviceLostCallback struct {
	C chan deviceLostCallbackResult
}

func (d *deviceLostCallback) info() sys.DeviceLostCallbackInfo {
	return sys.DeviceLostCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuDeviceLostCallback,
		Userdata1: unsafe.Pointer(d),
	}
}

type popErrorScopeCallback struct {
	C chan popErrorScopeCallbackResult
}

func (p *popErrorScopeCallback) info() sys.PopErrorScopeCallbackInfo {
	return sys.PopErrorScopeCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuPopErrorScopeCallback,
		Userdata1: unsafe.Pointer(p),
	}
}

type queueWorkDoneCallback struct {
	C chan queueWorkDoneCallbackResult
}

func (q *queueWorkDoneCallback) info() sys.QueueWorkDoneCallbackInfo {
	return sys.QueueWorkDoneCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuQueueWorkDoneCallback,
		Userdata1: unsafe.Pointer(q),
	}
}

type requestAdapterCallback struct {
	C chan requestAdapterCallbackResult
}

func (r *requestAdapterCallback) info() sys.RequestAdapterCallbackInfo {
	return sys.RequestAdapterCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuRequestAdapterCallback,
		Userdata1: unsafe.Pointer(r),
	}
}

type requestDeviceCallback struct {
	C chan requestDeviceCallbackResult
}

func (r *requestDeviceCallback) info() sys.RequestDeviceCallbackInfo {
	return sys.RequestDeviceCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuRequestDeviceCallback,
		Userdata1: unsafe.Pointer(r),
	}
}
