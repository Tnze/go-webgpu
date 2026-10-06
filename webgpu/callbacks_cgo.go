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
	"runtime/cgo"
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu/sys"
)

//export gowebgpuBufferMapCallback
func gowebgpuBufferMapCallback(status C.WGPUMapAsyncStatus, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- bufferMapCallbackResult) <- bufferMapCallbackResult{
		status:  MapAsyncStatus(status),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

func newBufferMapCallback(done chan<- bufferMapCallbackResult) (callbackInfo sys.BufferMapCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.BufferMapCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuBufferMapCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuCompilationInfoCallback
func gowebgpuCompilationInfoCallback(status C.WGPUCompilationInfoRequestStatus, compilationInfo *C.WGPUCompilationInfo, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- compilationInfoCallbackResult) <- compilationInfoCallbackResult{
		status:          CompilationInfoRequestStatus(status),
		compilationInfo: (*CompilationInfo)(unsafe.Pointer(compilationInfo)),
	}
}

func newCompilationInfoCallback(done chan<- compilationInfoCallbackResult) (callbackInfo sys.CompilationInfoCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.CompilationInfoCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuCompilationInfoCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuCreateComputePipelineAsyncCallback
func gowebgpuCreateComputePipelineAsyncCallback(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPUComputePipeline, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- createComputePipelineAsyncCallbackResult) <- createComputePipelineAsyncCallbackResult{
		status:   CreatePipelineAsyncStatus(status),
		pipeline: new(ComputePipeline{inner: sys.ComputePipeline(pipeline)}).owned(),
		message:  C.GoStringN(message.data, C.int(message.length)),
	}
}

func newCreateComputePipelineAsyncCallback(done chan<- createComputePipelineAsyncCallbackResult) (callbackInfo sys.CreateComputePipelineAsyncCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.CreateComputePipelineAsyncCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuCreateComputePipelineAsyncCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuCreateRenderPipelineAsyncCallback
func gowebgpuCreateRenderPipelineAsyncCallback(status C.WGPUCreatePipelineAsyncStatus, pipeline C.WGPURenderPipeline, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- createRenderPipelineAsyncCallbackResult) <- createRenderPipelineAsyncCallbackResult{
		status:   CreatePipelineAsyncStatus(status),
		pipeline: new(RenderPipeline{inner: sys.RenderPipeline(pipeline)}).owned(),
		message:  C.GoStringN(message.data, C.int(message.length)),
	}
}

func newCreateRenderPipelineAsyncCallback(done chan<- createRenderPipelineAsyncCallbackResult) (callbackInfo sys.CreateRenderPipelineAsyncCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.CreateRenderPipelineAsyncCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuCreateRenderPipelineAsyncCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuDeviceLostCallback
func gowebgpuDeviceLostCallback(device *C.WGPUDevice, reason C.WGPUDeviceLostReason, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- deviceLostCallbackResult) <- deviceLostCallbackResult{
		device:  new(Device{inner: sys.Device(device)}).owned(),
		reason:  DeviceLostReason(reason),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

func newDeviceLostCallback(done chan<- deviceLostCallbackResult) (callbackInfo sys.DeviceLostCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.DeviceLostCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuDeviceLostCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuPopErrorScopeCallback
func gowebgpuPopErrorScopeCallback(status C.WGPUPopErrorScopeStatus, errorType C.WGPUErrorType, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- popErrorScopeCallbackResult) <- popErrorScopeCallbackResult{
		status:    PopErrorScopeStatus(status),
		errorType: ErrorType(errorType),
		message:   C.GoStringN(message.data, C.int(message.length)),
	}
}

func newPopErrorScopeCallback(done chan<- popErrorScopeCallbackResult) (callbackInfo sys.PopErrorScopeCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.PopErrorScopeCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuPopErrorScopeCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuQueueWorkDoneCallback
func gowebgpuQueueWorkDoneCallback(status C.WGPUQueueWorkDoneStatus, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- queueWorkDoneCallbackResult) <- queueWorkDoneCallbackResult{
		status:  QueueWorkDoneStatus(status),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

func newQueueWorkDoneCallback(done chan<- queueWorkDoneCallbackResult) (callbackInfo sys.QueueWorkDoneCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.QueueWorkDoneCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuQueueWorkDoneCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuRequestAdapterCallback
func gowebgpuRequestAdapterCallback(status C.WGPURequestAdapterStatus, adapter C.WGPUAdapter, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- requestAdapterCallbackResult) <- requestAdapterCallbackResult{
		status:  RequestAdapterStatus(status),
		adapter: new(Adapter{inner: sys.Adapter(adapter)}).owned(),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

func newRequestAdapterCallback(done chan<- requestAdapterCallbackResult) (callbackInfo sys.RequestAdapterCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.RequestAdapterCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuRequestAdapterCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuRequestDeviceCallback
func gowebgpuRequestDeviceCallback(status C.WGPURequestDeviceStatus, device C.WGPUDevice, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	(*cgo.Handle)(userdata1).Value().(chan<- requestDeviceCallbackResult) <- requestDeviceCallbackResult{
		status:  RequestDeviceStatus(status),
		device:  new(Device{inner: sys.Device(device)}).owned(),
		message: C.GoStringN(message.data, C.int(message.length)),
	}
}

func newRequestDeviceCallback(done chan<- requestDeviceCallbackResult) (callbackInfo sys.RequestDeviceCallbackInfo, delete func()) {
	handle := new(cgo.NewHandle(done))
	info := sys.RequestDeviceCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  C.gowebgpuRequestDeviceCallback,
		Userdata1: unsafe.Pointer(handle),
	}
	return info, handle.Delete
}

//export gowebgpuUncapturedErrorCallback
func gowebgpuUncapturedErrorCallback(device *C.WGPUDevice, errorType C.WGPUErrorType, message C.WGPUStringView, userdata1, userdata2 unsafe.Pointer) {
	callback := (*cgo.Handle)(userdata1).Value().(func(device sys.Device, errorType ErrorType, message string, userdata2 unsafe.Pointer))
	callback(sys.Device(device), ErrorType(errorType), unsafe.String((*byte)(unsafe.Pointer(message.data)), message.length), userdata2)
}

func NewUncapturedErrorCallback(callback func(device sys.Device, errorType ErrorType, message string, userdata2 unsafe.Pointer), userdata2 unsafe.Pointer) sys.UncapturedErrorCallbackInfo {
	handle := new(cgo.NewHandle(callback))
	return sys.UncapturedErrorCallbackInfo{
		Callback:  C.gowebgpuUncapturedErrorCallback,
		Userdata1: unsafe.Pointer(handle),
		Userdata2: userdata2,
	}
}
