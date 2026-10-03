// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo

package sys

import (
	"structs"
	"unsafe"
)

type BufferMapCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type CompilationInfoCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type CreateComputePipelineAsyncCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type CreateRenderPipelineAsyncCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type DeviceLostCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type PopErrorScopeCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type QueueWorkDoneCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type RequestAdapterCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type RequestDeviceCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

type UncapturedErrorCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}
