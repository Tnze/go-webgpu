// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

type bufferMapCallbackResult struct {
	status  MapAsyncStatus
	message string
}

type compilationInfoCallbackResult struct {
	status          CompilationInfoRequestStatus
	compilationInfo *CompilationInfo
}

type createComputePipelineAsyncCallbackResult struct {
	status   CreatePipelineAsyncStatus
	pipeline *ComputePipeline
	message  string
}

type createRenderPipelineAsyncCallbackResult struct {
	status   CreatePipelineAsyncStatus
	pipeline *RenderPipeline
	message  string
}

type deviceLostCallbackResult struct {
	device  *Device
	reason  DeviceLostReason
	message string
}

type popErrorScopeCallbackResult struct {
	status    PopErrorScopeStatus
	errorType ErrorType
	message   string
}

type queueWorkDoneCallbackResult struct {
	status  QueueWorkDoneStatus
	message string
}

type requestAdapterCallbackResult struct {
	status  RequestAdapterStatus
	adapter *Adapter
	message string
}

type requestDeviceCallbackResult struct {
	status  RequestDeviceStatus
	device  *Device
	message string
}

type uncapturedErrorCallbackResult struct {
	device    *Device
	errorType ErrorType
	message   string
}
