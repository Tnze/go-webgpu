// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
// 
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

type Flags uint64

type BufferUsage Flags

var (
	BufferUsageNone BufferUsage = 0

	// The buffer can be *mapped* on the CPU side in *read* mode (using @ref WGPUMapMode_Read).
	BufferUsageMapRead BufferUsage = 1 << 0

	// The buffer can be *mapped* on the CPU side in *write* mode (using @ref WGPUMapMode_Write).
	// 
	// @note This usage is **not** required to set `mappedAtCreation` to `true` in @ref WGPUBufferDescriptor.
	BufferUsageMapWrite BufferUsage = 1 << 1

	// The buffer can be used as the *source* of a GPU-side copy operation.
	BufferUsageCopySrc BufferUsage = 1 << 2

	// The buffer can be used as the *destination* of a GPU-side copy operation.
	BufferUsageCopyDst BufferUsage = 1 << 3

	// The buffer can be used as an Index buffer when doing indexed drawing in a render pipeline.
	BufferUsageIndex BufferUsage = 1 << 4

	// The buffer can be used as a Vertex buffer when using a render pipeline.
	BufferUsageVertex BufferUsage = 1 << 5

	// The buffer can be bound to a shader as a uniform buffer.
	BufferUsageUniform BufferUsage = 1 << 6

	// The buffer can be bound to a shader as a storage buffer.
	BufferUsageStorage BufferUsage = 1 << 7

	// The buffer can store arguments for an indirect draw call.
	BufferUsageIndirect BufferUsage = 1 << 8

	// The buffer can store the result of a timestamp or occlusion query.
	BufferUsageQueryResolve BufferUsage = 1 << 9
)

type ColorWriteMask Flags

var (
	ColorWriteMaskNone ColorWriteMask = 0

	ColorWriteMaskRed ColorWriteMask = 1 << 0

	ColorWriteMaskGreen ColorWriteMask = 1 << 1

	ColorWriteMaskBlue ColorWriteMask = 1 << 2

	ColorWriteMaskAlpha ColorWriteMask = 1 << 3

	ColorWriteMaskAll ColorWriteMask = ColorWriteMaskRed | ColorWriteMaskGreen | ColorWriteMaskBlue | ColorWriteMaskAlpha
)

type MapMode Flags

var (
	MapModeNone MapMode = 0

	MapModeRead MapMode = 1 << 0

	MapModeWrite MapMode = 1 << 1
)

type ShaderStage Flags

var (
	ShaderStageNone ShaderStage = 0

	ShaderStageVertex ShaderStage = 1 << 0

	ShaderStageFragment ShaderStage = 1 << 1

	ShaderStageCompute ShaderStage = 1 << 2
)

type TextureUsage Flags

var (
	TextureUsageNone TextureUsage = 0

	TextureUsageCopySrc TextureUsage = 1 << 0

	TextureUsageCopyDst TextureUsage = 1 << 1

	TextureUsageTextureBinding TextureUsage = 1 << 2

	TextureUsageStorageBinding TextureUsage = 1 << 3

	TextureUsageRenderAttachment TextureUsage = 1 << 4

	TextureUsageTransientAttachment TextureUsage = 1 << 5
)
