// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package sys

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
	return BindGroupLayoutEntry{
		Visibility: ShaderStageNone,
	}
}

func DefaultBlendComponent() BlendComponent {
	return BlendComponent{}
}

func DefaultBlendState() BlendState {
	return BlendState{}
}

func DefaultBufferBindingLayout() BufferBindingLayout {
	return BufferBindingLayout{}
}

func DefaultBufferDescriptor() BufferDescriptor {
	return BufferDescriptor{
		Usage: BufferUsageNone,
	}
}

func DefaultColor() Color {
	return Color{}
}

func DefaultColorTargetState() ColorTargetState {
	return ColorTargetState{
		WriteMask: ColorWriteMaskAll,
	}
}

func DefaultCommandBufferDescriptor() CommandBufferDescriptor {
	return CommandBufferDescriptor{}
}

func DefaultCommandEncoderDescriptor() CommandEncoderDescriptor {
	return CommandEncoderDescriptor{}
}

func DefaultCompatibilityModeLimits() CompatibilityModeLimits {
	return CompatibilityModeLimits{
		MaxStorageBuffersInVertexStage:    LimitU32Undefined,
		MaxStorageTexturesInVertexStage:   LimitU32Undefined,
		MaxStorageBuffersInFragmentStage:  LimitU32Undefined,
		MaxStorageTexturesInFragmentStage: LimitU32Undefined,
	}
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
	return DepthStencilState{
		StencilReadMask:  0xFFFFFFFF,
		StencilWriteMask: 0xFFFFFFFF,
	}
}

func DefaultDeviceDescriptor() DeviceDescriptor {
	return DeviceDescriptor{}
}

func DefaultExtent3D() Extent3D {
	return Extent3D{
		Height:             1,
		DepthOrArrayLayers: 1,
	}
}

func DefaultExternalTextureBindingEntry() ExternalTextureBindingEntry {
	return ExternalTextureBindingEntry{}
}

func DefaultExternalTextureBindingLayout() ExternalTextureBindingLayout {
	return ExternalTextureBindingLayout{}
}

func DefaultFragmentState() FragmentState {
	return FragmentState{}
}

func DefaultFuture() Future {
	return Future{}
}

func DefaultFutureWaitInfo() FutureWaitInfo {
	return FutureWaitInfo{}
}

func DefaultInstanceDescriptor() InstanceDescriptor {
	return InstanceDescriptor{}
}

func DefaultInstanceLimits() InstanceLimits {
	return InstanceLimits{}
}

func DefaultLimits() Limits {
	return Limits{
		MaxTextureDimension1D:                     LimitU32Undefined,
		MaxTextureDimension2D:                     LimitU32Undefined,
		MaxTextureDimension3D:                     LimitU32Undefined,
		MaxTextureArrayLayers:                     LimitU32Undefined,
		MaxBindGroups:                             LimitU32Undefined,
		MaxBindGroupsPlusVertexBuffers:            LimitU32Undefined,
		MaxBindingsPerBindGroup:                   LimitU32Undefined,
		MaxDynamicUniformBuffersPerPipelineLayout: LimitU32Undefined,
		MaxDynamicStorageBuffersPerPipelineLayout: LimitU32Undefined,
		MaxSampledTexturesPerShaderStage:          LimitU32Undefined,
		MaxSamplersPerShaderStage:                 LimitU32Undefined,
		MaxStorageBuffersPerShaderStage:           LimitU32Undefined,
		MaxStorageTexturesPerShaderStage:          LimitU32Undefined,
		MaxUniformBuffersPerShaderStage:           LimitU32Undefined,
		MaxUniformBufferBindingSize:               LimitU64Undefined,
		MaxStorageBufferBindingSize:               LimitU64Undefined,
		MinUniformBufferOffsetAlignment:           LimitU32Undefined,
		MinStorageBufferOffsetAlignment:           LimitU32Undefined,
		MaxVertexBuffers:                          LimitU32Undefined,
		MaxBufferSize:                             LimitU64Undefined,
		MaxVertexAttributes:                       LimitU32Undefined,
		MaxVertexBufferArrayStride:                LimitU32Undefined,
		MaxInterStageShaderVariables:              LimitU32Undefined,
		MaxColorAttachments:                       LimitU32Undefined,
		MaxColorAttachmentBytesPerSample:          LimitU32Undefined,
		MaxComputeWorkgroupStorageSize:            LimitU32Undefined,
		MaxComputeInvocationsPerWorkgroup:         LimitU32Undefined,
		MaxComputeWorkgroupSizeX:                  LimitU32Undefined,
		MaxComputeWorkgroupSizeY:                  LimitU32Undefined,
		MaxComputeWorkgroupSizeZ:                  LimitU32Undefined,
		MaxComputeWorkgroupsPerDimension:          LimitU32Undefined,
		MaxImmediateSize:                          LimitU32Undefined,
	}
}

func DefaultMultisampleState() MultisampleState {
	return MultisampleState{
		Count: 1,
		Mask:  0xFFFFFFFF,
	}
}

func DefaultOrigin3D() Origin3D {
	return Origin3D{}
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
	return PrimitiveState{}
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
	return RenderPassMaxDrawCount{
		MaxDrawCount: 50000000,
	}
}

func DefaultRenderPipelineDescriptor() RenderPipelineDescriptor {
	return RenderPipelineDescriptor{}
}

func DefaultRequestAdapterOptions() RequestAdapterOptions {
	return RequestAdapterOptions{}
}

func DefaultRequestAdapterWebXROptions() RequestAdapterWebXROptions {
	return RequestAdapterWebXROptions{}
}

func DefaultSamplerBindingLayout() SamplerBindingLayout {
	return SamplerBindingLayout{}
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
	return ShaderSourceSPIRV{}
}

func DefaultShaderSourceWGSL() ShaderSourceWGSL {
	return ShaderSourceWGSL{}
}

func DefaultStencilFaceState() StencilFaceState {
	return StencilFaceState{}
}

func DefaultStorageTextureBindingLayout() StorageTextureBindingLayout {
	return StorageTextureBindingLayout{}
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
	return SurfaceColorManagement{}
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
	return SurfaceSourceAndroidNativeWindow{}
}

func DefaultSurfaceSourceMetalLayer() SurfaceSourceMetalLayer {
	return SurfaceSourceMetalLayer{}
}

func DefaultSurfaceSourceWaylandSurface() SurfaceSourceWaylandSurface {
	return SurfaceSourceWaylandSurface{}
}

func DefaultSurfaceSourceWindowsHWND() SurfaceSourceWindowsHWND {
	return SurfaceSourceWindowsHWND{}
}

func DefaultSurfaceSourceXCBWindow() SurfaceSourceXCBWindow {
	return SurfaceSourceXCBWindow{}
}

func DefaultSurfaceSourceXlibWindow() SurfaceSourceXlibWindow {
	return SurfaceSourceXlibWindow{}
}

func DefaultSurfaceTexture() SurfaceTexture {
	return SurfaceTexture{}
}

func DefaultTexelCopyBufferInfo() TexelCopyBufferInfo {
	return TexelCopyBufferInfo{}
}

func DefaultTexelCopyBufferLayout() TexelCopyBufferLayout {
	return TexelCopyBufferLayout{
		BytesPerRow:  CopyStrideUndefined,
		RowsPerImage: CopyStrideUndefined,
	}
}

func DefaultTexelCopyTextureInfo() TexelCopyTextureInfo {
	return TexelCopyTextureInfo{}
}

func DefaultTextureBindingLayout() TextureBindingLayout {
	return TextureBindingLayout{}
}

func DefaultTextureBindingViewDimension() TextureBindingViewDimension {
	return TextureBindingViewDimension{}
}

func DefaultTextureComponentSwizzle() TextureComponentSwizzle {
	return TextureComponentSwizzle{}
}

func DefaultTextureComponentSwizzleDescriptor() TextureComponentSwizzleDescriptor {
	return TextureComponentSwizzleDescriptor{}
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
	return VertexAttribute{}
}

func DefaultVertexBufferLayout() VertexBufferLayout {
	return VertexBufferLayout{}
}

func DefaultVertexState() VertexState {
	return VertexState{}
}
