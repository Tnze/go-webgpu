// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import "github.com/Tnze/go-webgpu/webgpu/sys"

func DefaultAdapterInfo() AdapterInfo {
	return AdapterInfo{}
}

func DefaultBindGroupDescriptor() BindGroupDescriptor {
	return BindGroupDescriptor{}
}

func DefaultBindGroupEntry() BindGroupEntry {
	return BindGroupEntry{
		Size: WholeSize,
	}
}

func DefaultBindGroupLayoutDescriptor() BindGroupLayoutDescriptor {
	return BindGroupLayoutDescriptor{}
}

func DefaultBindGroupLayoutEntry() BindGroupLayoutEntry {
	return sys.DefaultBindGroupLayoutEntry()
}

func DefaultBlendComponent() BlendComponent {
	return sys.DefaultBlendComponent()
}

func DefaultBlendState() BlendState {
	return sys.DefaultBlendState()
}

func DefaultBufferBindingLayout() BufferBindingLayout {
	return sys.DefaultBufferBindingLayout()
}

func DefaultBufferDescriptor() BufferDescriptor {
	return BufferDescriptor{
		Usage: BufferUsageNone,
	}
}

func DefaultColor() Color {
	return sys.DefaultColor()
}

func DefaultColorTargetState() ColorTargetState {
	return sys.DefaultColorTargetState()
}

func DefaultCommandBufferDescriptor() CommandBufferDescriptor {
	return CommandBufferDescriptor{}
}

func DefaultCommandEncoderDescriptor() CommandEncoderDescriptor {
	return CommandEncoderDescriptor{}
}

func DefaultCompatibilityModeLimits() CompatibilityModeLimits {
	return sys.DefaultCompatibilityModeLimits()
}

func DefaultCompilationInfo() CompilationInfo {
	return CompilationInfo{}
}

func DefaultCompilationMessage() CompilationMessage {
	return CompilationMessage{}
}

func DefaultComputePassDescriptor() ComputePassDescriptor {
	return ComputePassDescriptor{}
}

func DefaultComputePipelineDescriptor() ComputePipelineDescriptor {
	return ComputePipelineDescriptor{}
}

func DefaultComputeState() ComputeState {
	return ComputeState{}
}

func DefaultConstantEntry() ConstantEntry {
	return ConstantEntry{}
}

func DefaultDepthStencilState() DepthStencilState {
	return sys.DefaultDepthStencilState()
}

func DefaultDeviceDescriptor() DeviceDescriptor {
	return DeviceDescriptor{}
}

func DefaultExtent3D() Extent3D {
	return sys.DefaultExtent3D()
}

func DefaultExternalTextureBindingEntry() ExternalTextureBindingEntry {
	return ExternalTextureBindingEntry{}
}

func DefaultExternalTextureBindingLayout() ExternalTextureBindingLayout {
	return sys.DefaultExternalTextureBindingLayout()
}

func DefaultFragmentState() FragmentState {
	return FragmentState{}
}

func DefaultFuture() Future {
	return sys.DefaultFuture()
}

func DefaultFutureWaitInfo() FutureWaitInfo {
	return sys.DefaultFutureWaitInfo()
}

func DefaultInstanceDescriptor() InstanceDescriptor {
	return InstanceDescriptor{}
}

func DefaultInstanceLimits() InstanceLimits {
	return sys.DefaultInstanceLimits()
}

func DefaultLimits() Limits {
	return sys.DefaultLimits()
}

func DefaultMultisampleState() MultisampleState {
	return sys.DefaultMultisampleState()
}

func DefaultOrigin3D() Origin3D {
	return sys.DefaultOrigin3D()
}

func DefaultPassTimestampWrites() PassTimestampWrites {
	return PassTimestampWrites{
		BeginningOfPassWriteIndex: QuerySetIndexUndefined,
		EndOfPassWriteIndex:       QuerySetIndexUndefined,
	}
}

func DefaultPipelineLayoutDescriptor() PipelineLayoutDescriptor {
	return PipelineLayoutDescriptor{}
}

func DefaultPrimitiveState() PrimitiveState {
	return sys.DefaultPrimitiveState()
}

func DefaultQuerySetDescriptor() QuerySetDescriptor {
	return QuerySetDescriptor{}
}

func DefaultQueueDescriptor() QueueDescriptor {
	return QueueDescriptor{}
}

func DefaultRenderBundleDescriptor() RenderBundleDescriptor {
	return RenderBundleDescriptor{}
}

func DefaultRenderBundleEncoderDescriptor() RenderBundleEncoderDescriptor {
	return RenderBundleEncoderDescriptor{
		SampleCount: 1,
	}
}

func DefaultRenderPassColorAttachment() RenderPassColorAttachment {
	return RenderPassColorAttachment{
		DepthSlice: DepthSliceUndefined,
	}
}

func DefaultRenderPassDepthStencilAttachment() RenderPassDepthStencilAttachment {
	return RenderPassDepthStencilAttachment{
		DepthClearValue: DepthClearValueUndefined,
	}
}

func DefaultRenderPassDescriptor() RenderPassDescriptor {
	return RenderPassDescriptor{}
}

func DefaultRenderPassMaxDrawCount() RenderPassMaxDrawCount {
	return sys.DefaultRenderPassMaxDrawCount()
}

func DefaultRenderPipelineDescriptor() RenderPipelineDescriptor {
	return RenderPipelineDescriptor{}
}

func DefaultRequestAdapterOptions() RequestAdapterOptions {
	return RequestAdapterOptions{}
}

func DefaultRequestAdapterWebXROptions() RequestAdapterWebXROptions {
	return sys.DefaultRequestAdapterWebXROptions()
}

func DefaultSamplerBindingLayout() SamplerBindingLayout {
	return sys.DefaultSamplerBindingLayout()
}

func DefaultSamplerDescriptor() SamplerDescriptor {
	return SamplerDescriptor{
		LodMaxClamp:   32.0,
		MaxAnisotropy: 1,
	}
}

func DefaultShaderModuleDescriptor() ShaderModuleDescriptor {
	return ShaderModuleDescriptor{}
}

func DefaultShaderSourceSPIRV() ShaderSourceSPIRV {
	return sys.DefaultShaderSourceSPIRV()
}

func DefaultShaderSourceWGSL() ShaderSourceWGSL {
	return ShaderSourceWGSL{}
}

func DefaultStencilFaceState() StencilFaceState {
	return sys.DefaultStencilFaceState()
}

func DefaultStorageTextureBindingLayout() StorageTextureBindingLayout {
	return sys.DefaultStorageTextureBindingLayout()
}

func DefaultSupportedFeatures() SupportedFeatures {
	return SupportedFeatures{}
}

func DefaultSupportedInstanceFeatures() SupportedInstanceFeatures {
	return SupportedInstanceFeatures{}
}

func DefaultSupportedWGSLLanguageFeatures() SupportedWGSLLanguageFeatures {
	return SupportedWGSLLanguageFeatures{}
}

func DefaultSurfaceCapabilities() SurfaceCapabilities {
	return SurfaceCapabilities{}
}

func DefaultSurfaceColorManagement() SurfaceColorManagement {
	return sys.DefaultSurfaceColorManagement()
}

func DefaultSurfaceConfiguration() SurfaceConfiguration {
	return SurfaceConfiguration{
		Usage:     TextureUsageRenderAttachment,
		AlphaMode: CompositeAlphaModeAuto,
	}
}

func DefaultSurfaceDescriptor() SurfaceDescriptor {
	return SurfaceDescriptor{}
}

func DefaultSurfaceSourceAndroidNativeWindow() SurfaceSourceAndroidNativeWindow {
	return sys.DefaultSurfaceSourceAndroidNativeWindow()
}

func DefaultSurfaceSourceMetalLayer() SurfaceSourceMetalLayer {
	return sys.DefaultSurfaceSourceMetalLayer()
}

func DefaultSurfaceSourceWaylandSurface() SurfaceSourceWaylandSurface {
	return sys.DefaultSurfaceSourceWaylandSurface()
}

func DefaultSurfaceSourceWindowsHWND() SurfaceSourceWindowsHWND {
	return sys.DefaultSurfaceSourceWindowsHWND()
}

func DefaultSurfaceSourceXCBWindow() SurfaceSourceXCBWindow {
	return sys.DefaultSurfaceSourceXCBWindow()
}

func DefaultSurfaceSourceXlibWindow() SurfaceSourceXlibWindow {
	return sys.DefaultSurfaceSourceXlibWindow()
}

func DefaultSurfaceTexture() SurfaceTexture {
	return SurfaceTexture{}
}

func DefaultTexelCopyBufferInfo() TexelCopyBufferInfo {
	return TexelCopyBufferInfo{}
}

func DefaultTexelCopyBufferLayout() TexelCopyBufferLayout {
	return sys.DefaultTexelCopyBufferLayout()
}

func DefaultTexelCopyTextureInfo() TexelCopyTextureInfo {
	return TexelCopyTextureInfo{}
}

func DefaultTextureBindingLayout() TextureBindingLayout {
	return sys.DefaultTextureBindingLayout()
}

func DefaultTextureBindingViewDimension() TextureBindingViewDimension {
	return sys.DefaultTextureBindingViewDimension()
}

func DefaultTextureComponentSwizzle() TextureComponentSwizzle {
	return sys.DefaultTextureComponentSwizzle()
}

func DefaultTextureComponentSwizzleDescriptor() TextureComponentSwizzleDescriptor {
	return sys.DefaultTextureComponentSwizzleDescriptor()
}

func DefaultTextureDescriptor() TextureDescriptor {
	return TextureDescriptor{
		Usage:         TextureUsageNone,
		MipLevelCount: 1,
		SampleCount:   1,
	}
}

func DefaultTextureViewDescriptor() TextureViewDescriptor {
	return TextureViewDescriptor{
		MipLevelCount:   MipLevelCountUndefined,
		ArrayLayerCount: ArrayLayerCountUndefined,
		Usage:           TextureUsageNone,
	}
}

func DefaultVertexAttribute() VertexAttribute {
	return sys.DefaultVertexAttribute()
}

func DefaultVertexBufferLayout() VertexBufferLayout {
	return VertexBufferLayout{}
}

func DefaultVertexState() VertexState {
	return VertexState{}
}
