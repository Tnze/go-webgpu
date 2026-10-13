// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build !cgo && amd64

package webgpu

import (
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu/sys"
)

var gowebgpuBufferMapCallback = syscall.NewCallback(func(status sys.MapAsyncStatus, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- bufferMapCallbackResult)(userdata1) <- bufferMapCallbackResult{
		status:  MapAsyncStatus(status),
		message: strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newBufferMapCallback(done chan<- bufferMapCallbackResult) (callbackInfo sys.BufferMapCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.BufferMapCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuBufferMapCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuCompilationInfoCallback = syscall.NewCallback(func(status sys.CompilationInfoRequestStatus, compilationInfo *sys.CompilationInfo, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- compilationInfoCallbackResult)(userdata1) <- compilationInfoCallbackResult{
		status:          CompilationInfoRequestStatus(status),
		compilationInfo: (*CompilationInfo)(unsafe.Pointer(compilationInfo)),
	}
	return 0
})

func newCompilationInfoCallback(done chan<- compilationInfoCallbackResult) (callbackInfo sys.CompilationInfoCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.CompilationInfoCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuCompilationInfoCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuCreateComputePipelineAsyncCallback = syscall.NewCallback(func(status sys.CreatePipelineAsyncStatus, pipeline sys.ComputePipeline, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- createComputePipelineAsyncCallbackResult)(userdata1) <- createComputePipelineAsyncCallbackResult{
		status:   CreatePipelineAsyncStatus(status),
		pipeline: new(ComputePipeline{inner: sys.ComputePipeline(pipeline)}).owned(),
		message:  strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newCreateComputePipelineAsyncCallback(done chan<- createComputePipelineAsyncCallbackResult) (callbackInfo sys.CreateComputePipelineAsyncCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.CreateComputePipelineAsyncCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuCreateComputePipelineAsyncCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuCreateRenderPipelineAsyncCallback = syscall.NewCallback(func(status sys.CreatePipelineAsyncStatus, pipeline sys.RenderPipeline, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- createRenderPipelineAsyncCallbackResult)(userdata1) <- createRenderPipelineAsyncCallbackResult{
		status:   CreatePipelineAsyncStatus(status),
		pipeline: new(RenderPipeline{inner: sys.RenderPipeline(pipeline)}).owned(),
		message:  strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newCreateRenderPipelineAsyncCallback(done chan<- createRenderPipelineAsyncCallbackResult) (callbackInfo sys.CreateRenderPipelineAsyncCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.CreateRenderPipelineAsyncCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuCreateRenderPipelineAsyncCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuDeviceLostCallback = syscall.NewCallback(func(device *sys.Device, reason sys.DeviceLostReason, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- deviceLostCallbackResult)(userdata1) <- deviceLostCallbackResult{
		device:  new(Device{inner: sys.Device(device)}).owned(),
		reason:  DeviceLostReason(reason),
		message: strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newDeviceLostCallback(done chan<- deviceLostCallbackResult) (callbackInfo sys.DeviceLostCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.DeviceLostCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuDeviceLostCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuPopErrorScopeCallback = syscall.NewCallback(func(status sys.PopErrorScopeStatus, errorType sys.ErrorType, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- popErrorScopeCallbackResult)(userdata1) <- popErrorScopeCallbackResult{
		status:    PopErrorScopeStatus(status),
		errorType: ErrorType(errorType),
		message:   strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newPopErrorScopeCallback(done chan<- popErrorScopeCallbackResult) (callbackInfo sys.PopErrorScopeCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.PopErrorScopeCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuPopErrorScopeCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuQueueWorkDoneCallback = syscall.NewCallback(func(status sys.QueueWorkDoneStatus, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- queueWorkDoneCallbackResult)(userdata1) <- queueWorkDoneCallbackResult{
		status:  QueueWorkDoneStatus(status),
		message: strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newQueueWorkDoneCallback(done chan<- queueWorkDoneCallbackResult) (callbackInfo sys.QueueWorkDoneCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.QueueWorkDoneCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuQueueWorkDoneCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuRequestAdapterCallback = syscall.NewCallback(func(status sys.RequestAdapterStatus, adapter sys.Adapter, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- requestAdapterCallbackResult)(userdata1) <- requestAdapterCallbackResult{
		status:  RequestAdapterStatus(status),
		adapter: new(Adapter{inner: sys.Adapter(adapter)}).owned(),
		message: strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newRequestAdapterCallback(done chan<- requestAdapterCallbackResult) (callbackInfo sys.RequestAdapterCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.RequestAdapterCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuRequestAdapterCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuRequestDeviceCallback = syscall.NewCallback(func(status sys.RequestDeviceStatus, device sys.Device, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	*(*chan<- requestDeviceCallbackResult)(userdata1) <- requestDeviceCallbackResult{
		status:  RequestDeviceStatus(status),
		device:  new(Device{inner: sys.Device(device)}).owned(),
		message: strings.Clone(unsafe.String(message.Data, message.Length)),
	}
	return 0
})

func newRequestDeviceCallback(done chan<- requestDeviceCallbackResult) (callbackInfo sys.RequestDeviceCallbackInfo, delete func()) {
	var pinner runtime.Pinner
	pinner.Pin(&done)
	info := sys.RequestDeviceCallbackInfo{
		Mode:      CallbackModeAllowProcessEvents,
		Callback:  unsafe.Pointer(gowebgpuRequestDeviceCallback),
		Userdata1: unsafe.Pointer(&done),
	}
	return info, pinner.Unpin
}

var gowebgpuUncapturedErrorCallback = syscall.NewCallback(func(device *Device, errorType ErrorType, message *sys.StringView, userdata1, userdata2 unsafe.Pointer) uintptr {
	panic("TODO")
	return 0
})

func NewUncapturedErrorCallback(callback func(device sys.Device, errorType ErrorType, message string, userdata2 unsafe.Pointer), userdata2 unsafe.Pointer) sys.UncapturedErrorCallbackInfo {
	return sys.UncapturedErrorCallbackInfo{
		Callback:  unsafe.Pointer(gowebgpuUncapturedErrorCallback),
		Userdata1: unsafe.Pointer(nil),
		Userdata2: userdata2,
	}
}
