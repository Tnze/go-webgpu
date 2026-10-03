// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

// **Important:** *This documentation is a Work In Progress.*
//
// This is the home of WebGPU C API specification. We define here the standard
// `webgpu.h` header that all implementations should provide.
//
// For all details where behavior is not otherwise specified, `webgpu.h` has
// the same behavior as the WebGPU specification for JavaScript on the Web.
// The WebIDL-based Web specification is mapped into C as faithfully (and
// bidirectionally) as practical/possible.
// The working draft of WebGPU can be found at <https://www.w3.org/TR/webgpu/>.
//
// The standard include directive for this header is `#include <webgpu/webgpu.h>`
// (if it is provided in a system-wide or toolchain-wide include directory).
package sys

type StringView struct {
	Data   *byte
	Length uint
}

type Bool uint32

const (
	True  Bool = 1
	False Bool = 0
)

type ChainedStruct struct {
	Next  *ChainedStruct
	SType SType
}
