// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import "github.com/Tnze/go-webgpu/webgpu/sys"

type BufferUsage = sys.BufferUsage

const (
	BufferUsageNone BufferUsage = sys.BufferUsageNone
	// The buffer can be *mapped* on the CPU side in *read* mode (using @ref WGPUMapMode_Read).
	BufferUsageMapRead BufferUsage = sys.BufferUsageMapRead
	// The buffer can be *mapped* on the CPU side in *write* mode (using @ref WGPUMapMode_Write).
	//
	// @note This usage is **not** required to set `mappedAtCreation` to `true` in @ref WGPUBufferDescriptor.
	BufferUsageMapWrite BufferUsage = sys.BufferUsageMapWrite
	// The buffer can be used as the *source* of a GPU-side copy operation.
	BufferUsageCopySrc BufferUsage = sys.BufferUsageCopySrc
	// The buffer can be used as the *destination* of a GPU-side copy operation.
	BufferUsageCopyDst BufferUsage = sys.BufferUsageCopyDst
	// The buffer can be used as an Index buffer when doing indexed drawing in a render pipeline.
	BufferUsageIndex BufferUsage = sys.BufferUsageIndex
	// The buffer can be used as a Vertex buffer when using a render pipeline.
	BufferUsageVertex BufferUsage = sys.BufferUsageVertex
	// The buffer can be bound to a shader as a uniform buffer.
	BufferUsageUniform BufferUsage = sys.BufferUsageUniform
	// The buffer can be bound to a shader as a storage buffer.
	BufferUsageStorage BufferUsage = sys.BufferUsageStorage
	// The buffer can store arguments for an indirect draw call.
	BufferUsageIndirect BufferUsage = sys.BufferUsageIndirect
	// The buffer can store the result of a timestamp or occlusion query.
	BufferUsageQueryResolve BufferUsage = sys.BufferUsageQueryResolve
)

type ColorWriteMask = sys.ColorWriteMask

const (
	ColorWriteMaskNone  ColorWriteMask = sys.ColorWriteMaskNone
	ColorWriteMaskRed   ColorWriteMask = sys.ColorWriteMaskRed
	ColorWriteMaskGreen ColorWriteMask = sys.ColorWriteMaskGreen
	ColorWriteMaskBlue  ColorWriteMask = sys.ColorWriteMaskBlue
	ColorWriteMaskAlpha ColorWriteMask = sys.ColorWriteMaskAlpha
	ColorWriteMaskAll   ColorWriteMask = sys.ColorWriteMaskAll
)

type MapMode = sys.MapMode

const (
	MapModeNone  MapMode = sys.MapModeNone
	MapModeRead  MapMode = sys.MapModeRead
	MapModeWrite MapMode = sys.MapModeWrite
)

type ShaderStage = sys.ShaderStage

const (
	ShaderStageNone     ShaderStage = sys.ShaderStageNone
	ShaderStageVertex   ShaderStage = sys.ShaderStageVertex
	ShaderStageFragment ShaderStage = sys.ShaderStageFragment
	ShaderStageCompute  ShaderStage = sys.ShaderStageCompute
)

type TextureUsage = sys.TextureUsage

const (
	TextureUsageNone                TextureUsage = sys.TextureUsageNone
	TextureUsageCopySrc             TextureUsage = sys.TextureUsageCopySrc
	TextureUsageCopyDst             TextureUsage = sys.TextureUsageCopyDst
	TextureUsageTextureBinding      TextureUsage = sys.TextureUsageTextureBinding
	TextureUsageStorageBinding      TextureUsage = sys.TextureUsageStorageBinding
	TextureUsageRenderAttachment    TextureUsage = sys.TextureUsageRenderAttachment
	TextureUsageTransientAttachment TextureUsage = sys.TextureUsageTransientAttachment
)
