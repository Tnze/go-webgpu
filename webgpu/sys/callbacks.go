// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package sys

import (
	"runtime"
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

func (b *BufferMapCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(b)
	for next := b.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(b.Callback)
	if b.Userdata1 != nil {
		pinner.Pin(b.Userdata1)
	}
	if b.Userdata2 != nil {
		pinner.Pin(b.Userdata2)
	}
}

type CompilationInfoCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (c *CompilationInfoCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(c)
	for next := c.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(c.Callback)
	if c.Userdata1 != nil {
		pinner.Pin(c.Userdata1)
	}
	if c.Userdata2 != nil {
		pinner.Pin(c.Userdata2)
	}
}

type CreateComputePipelineAsyncCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (c *CreateComputePipelineAsyncCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(c)
	for next := c.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(c.Callback)
	if c.Userdata1 != nil {
		pinner.Pin(c.Userdata1)
	}
	if c.Userdata2 != nil {
		pinner.Pin(c.Userdata2)
	}
}

type CreateRenderPipelineAsyncCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (c *CreateRenderPipelineAsyncCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(c)
	for next := c.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(c.Callback)
	if c.Userdata1 != nil {
		pinner.Pin(c.Userdata1)
	}
	if c.Userdata2 != nil {
		pinner.Pin(c.Userdata2)
	}
}

type DeviceLostCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (d *DeviceLostCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(d)
	for next := d.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(d.Callback)
	if d.Userdata1 != nil {
		pinner.Pin(d.Userdata1)
	}
	if d.Userdata2 != nil {
		pinner.Pin(d.Userdata2)
	}
}

type PopErrorScopeCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (p *PopErrorScopeCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(p)
	for next := p.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(p.Callback)
	if p.Userdata1 != nil {
		pinner.Pin(p.Userdata1)
	}
	if p.Userdata2 != nil {
		pinner.Pin(p.Userdata2)
	}
}

type QueueWorkDoneCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (q *QueueWorkDoneCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(q)
	for next := q.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(q.Callback)
	if q.Userdata1 != nil {
		pinner.Pin(q.Userdata1)
	}
	if q.Userdata2 != nil {
		pinner.Pin(q.Userdata2)
	}
}

type RequestAdapterCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (r *RequestAdapterCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(r)
	for next := r.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(r.Callback)
	if r.Userdata1 != nil {
		pinner.Pin(r.Userdata1)
	}
	if r.Userdata2 != nil {
		pinner.Pin(r.Userdata2)
	}
}

type RequestDeviceCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Mode                 CallbackMode
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (r *RequestDeviceCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(r)
	for next := r.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(r.Callback)
	if r.Userdata1 != nil {
		pinner.Pin(r.Userdata1)
	}
	if r.Userdata2 != nil {
		pinner.Pin(r.Userdata2)
	}
}

type UncapturedErrorCallbackInfo struct {
	_                    structs.HostLayout
	NextInChain          *ChainedStruct
	Callback             unsafe.Pointer
	Userdata1, Userdata2 unsafe.Pointer
}

func (u *UncapturedErrorCallbackInfo) pin(pinner *runtime.Pinner) {
	pinner.Pin(u)
	for next := u.NextInChain; next != nil; next = next.Next {
		pinner.Pin(next)
	}
	pinner.Pin(u.Callback)
	if u.Userdata1 != nil {
		pinner.Pin(u.Userdata1)
	}
	if u.Userdata2 != nil {
		pinner.Pin(u.Userdata2)
	}
}
