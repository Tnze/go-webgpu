// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package sys

import "unsafe"

type Adapter unsafe.Pointer

type BindGroup unsafe.Pointer

type BindGroupLayout unsafe.Pointer

type Buffer unsafe.Pointer

type CommandBuffer unsafe.Pointer

type CommandEncoder unsafe.Pointer

type ComputePassEncoder unsafe.Pointer

type ComputePipeline unsafe.Pointer

// TODO
//
// Releasing the last ref to a `WGPUDevice` also calls @ref wgpuDeviceDestroy.
// For more info, see @ref DeviceRelease.
type Device unsafe.Pointer

// A sampleable 2D texture that may perform 0-copy YUV sampling internally. Creation of @ref WGPUExternalTexture is extremely implementation-dependent and not defined in this header.
type ExternalTexture unsafe.Pointer

type Instance unsafe.Pointer

type PipelineLayout unsafe.Pointer

type QuerySet unsafe.Pointer

type Queue unsafe.Pointer

type RenderBundle unsafe.Pointer

type RenderBundleEncoder unsafe.Pointer

type RenderPassEncoder unsafe.Pointer

type RenderPipeline unsafe.Pointer

type Sampler unsafe.Pointer

type ShaderModule unsafe.Pointer

// An object used to continuously present image data to the user, see @ref Surfaces for more details.
type Surface unsafe.Pointer

type Texture unsafe.Pointer

type TextureView unsafe.Pointer
