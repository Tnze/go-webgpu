// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import (
	"github.com/Tnze/go-webgpu/webgpu/sys"
	"unsafe"
)

func (a *AdapterInfo) unwrap() (out sys.AdapterInfo) {
	out.Vendor = sys.StringView{
		Length: uint(len(a.Vendor)),
		Data:   unsafe.StringData(a.Vendor),
	}
	out.Architecture = sys.StringView{
		Length: uint(len(a.Architecture)),
		Data:   unsafe.StringData(a.Architecture),
	}
	out.Device = sys.StringView{
		Length: uint(len(a.Device)),
		Data:   unsafe.StringData(a.Device),
	}
	out.Description = sys.StringView{
		Length: uint(len(a.Description)),
		Data:   unsafe.StringData(a.Description),
	}
	out.BackendType = a.BackendType         // enum.backend_type
	out.AdapterType = a.AdapterType         // enum.adapter_type
	out.VendorID = a.VendorID               // uint32
	out.DeviceID = a.DeviceID               // uint32
	out.SubgroupMinSize = a.SubgroupMinSize // uint32
	out.SubgroupMaxSize = a.SubgroupMaxSize // uint32
	return
}

func (a *AdapterInfo) wrap(in *sys.AdapterInfo) {
	panic("TODO")
}

func (b *BindGroupDescriptor) unwrap() (out sys.BindGroupDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(b.Label)),
		Data:   unsafe.StringData(b.Label),
	}
	out.Layout = b.Layout.inner
	out.EntriesCount = uint(len(b.Entries))
	entries := make([]sys.BindGroupEntry, len(b.Entries))
	for i := range entries {
		entries[i] = b.Entries[i].unwrap()
	}
	out.Entries = &entries[0]
	return
}

func (b *BindGroupDescriptor) wrap(in *sys.BindGroupDescriptor) {
	panic("TODO")
}

func (b *BindGroupEntry) unwrap() (out sys.BindGroupEntry) {
	out.Binding = b.Binding // uint32
	out.Buffer = b.Buffer.inner
	out.Offset = b.Offset // uint64
	out.Size = b.Size     // uint64
	out.Sampler = b.Sampler.inner
	out.TextureView = b.TextureView.inner
	return
}

func (b *BindGroupEntry) wrap(in *sys.BindGroupEntry) {
	panic("TODO")
}

func (b *BindGroupLayoutDescriptor) unwrap() (out sys.BindGroupLayoutDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(b.Label)),
		Data:   unsafe.StringData(b.Label),
	}
	out.EntriesCount = uint(len(b.Entries))

	out.Entries = &b.Entries[0]
	return
}

func (b *BindGroupLayoutDescriptor) wrap(in *sys.BindGroupLayoutDescriptor) {
	panic("TODO")
}

func (b *BufferDescriptor) unwrap() (out sys.BufferDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(b.Label)),
		Data:   unsafe.StringData(b.Label),
	}
	out.Usage = b.Usage                       // bitflag.buffer_usage
	out.Size = b.Size                         // uint64
	out.MappedAtCreation = b.MappedAtCreation // bool
	return
}

func (b *BufferDescriptor) wrap(in *sys.BufferDescriptor) {
	panic("TODO")
}

func (c *CommandBufferDescriptor) unwrap() (out sys.CommandBufferDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(c.Label)),
		Data:   unsafe.StringData(c.Label),
	}
	return
}

func (c *CommandBufferDescriptor) wrap(in *sys.CommandBufferDescriptor) {
	panic("TODO")
}

func (c *CommandEncoderDescriptor) unwrap() (out sys.CommandEncoderDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(c.Label)),
		Data:   unsafe.StringData(c.Label),
	}
	return
}

func (c *CommandEncoderDescriptor) wrap(in *sys.CommandEncoderDescriptor) {
	panic("TODO")
}

func (c *CompilationInfo) unwrap() (out sys.CompilationInfo) {
	out.MessagesCount = uint(len(c.Messages))
	messages := make([]sys.CompilationMessage, len(c.Messages))
	for i := range messages {
		messages[i] = c.Messages[i].unwrap()
	}
	out.Messages = &messages[0]
	return
}

func (c *CompilationInfo) wrap(in *sys.CompilationInfo) {
	panic("TODO")
}

func (c *CompilationMessage) unwrap() (out sys.CompilationMessage) {
	out.Message = sys.StringView{
		Length: uint(len(c.Message)),
		Data:   unsafe.StringData(c.Message),
	}
	out.Type = c.Type       // enum.compilation_message_type
	out.LineNum = c.LineNum // uint64
	out.LinePos = c.LinePos // uint64
	out.Offset = c.Offset   // uint64
	out.Length = c.Length   // uint64
	return
}

func (c *CompilationMessage) wrap(in *sys.CompilationMessage) {
	panic("TODO")
}

func (c *ComputePassDescriptor) unwrap() (out sys.ComputePassDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(c.Label)),
		Data:   unsafe.StringData(c.Label),
	}
	if c.TimestampWrites != nil {
		out.TimestampWrites = new(c.TimestampWrites.unwrap())
	}
	return
}

func (c *ComputePassDescriptor) wrap(in *sys.ComputePassDescriptor) {
	panic("TODO")
}

func (c *ComputePipelineDescriptor) unwrap() (out sys.ComputePipelineDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(c.Label)),
		Data:   unsafe.StringData(c.Label),
	}
	out.Layout = c.Layout.inner
	out.Compute = c.Compute.unwrap()
	return
}

func (c *ComputePipelineDescriptor) wrap(in *sys.ComputePipelineDescriptor) {
	panic("TODO")
}

func (c *ComputeState) unwrap() (out sys.ComputeState) {
	out.Module = c.Module.inner
	out.EntryPoint = sys.StringView{
		Length: uint(len(c.EntryPoint)),
		Data:   unsafe.StringData(c.EntryPoint),
	}
	out.ConstantsCount = uint(len(c.Constants))
	constants := make([]sys.ConstantEntry, len(c.Constants))
	for i := range constants {
		constants[i] = c.Constants[i].unwrap()
	}
	out.Constants = &constants[0]
	return
}

func (c *ComputeState) wrap(in *sys.ComputeState) {
	panic("TODO")
}

func (c *ConstantEntry) unwrap() (out sys.ConstantEntry) {
	out.Key = sys.StringView{
		Length: uint(len(c.Key)),
		Data:   unsafe.StringData(c.Key),
	}
	out.Value = c.Value // float64_supertype
	return
}

func (c *ConstantEntry) wrap(in *sys.ConstantEntry) {
	panic("TODO")
}

func (d *DeviceDescriptor) unwrap() (out sys.DeviceDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(d.Label)),
		Data:   unsafe.StringData(d.Label),
	}

	out.RequiredFeatures = &d.RequiredFeatures[0]
	if d.RequiredLimits != nil {
		out.RequiredLimits = d.RequiredLimits
	}
	out.DefaultQueue = d.DefaultQueue.unwrap()
	out.DeviceLostCallbackInfo = d.DeviceLostCallbackInfo           // callback.device_lost
	out.UncapturedErrorCallbackInfo = d.UncapturedErrorCallbackInfo // callback.uncaptured_error
	return
}

func (d *DeviceDescriptor) wrap(in *sys.DeviceDescriptor) {
	panic("TODO")
}

func (e *ExternalTextureBindingEntry) unwrap() (out sys.ExternalTextureBindingEntry) {
	out.ExternalTexture = e.ExternalTexture.inner
	return
}

func (e *ExternalTextureBindingEntry) wrap(in *sys.ExternalTextureBindingEntry) {
	panic("TODO")
}

func (f *FragmentState) unwrap() (out sys.FragmentState) {
	out.Module = f.Module.inner
	out.EntryPoint = sys.StringView{
		Length: uint(len(f.EntryPoint)),
		Data:   unsafe.StringData(f.EntryPoint),
	}
	out.ConstantsCount = uint(len(f.Constants))
	constants := make([]sys.ConstantEntry, len(f.Constants))
	for i := range constants {
		constants[i] = f.Constants[i].unwrap()
	}
	out.Constants = &constants[0]
	out.TargetsCount = uint(len(f.Targets))

	out.Targets = &f.Targets[0]
	return
}

func (f *FragmentState) wrap(in *sys.FragmentState) {
	panic("TODO")
}

func (i *InstanceDescriptor) unwrap() (out sys.InstanceDescriptor) {

	out.RequiredFeatures = &i.RequiredFeatures[0]
	if i.RequiredLimits != nil {
		out.RequiredLimits = i.RequiredLimits
	}
	return
}

func (i *InstanceDescriptor) wrap(in *sys.InstanceDescriptor) {
	panic("TODO")
}

func (p *PassTimestampWrites) unwrap() (out sys.PassTimestampWrites) {
	out.QuerySet = p.QuerySet.inner
	out.BeginningOfPassWriteIndex = p.BeginningOfPassWriteIndex // uint32
	out.EndOfPassWriteIndex = p.EndOfPassWriteIndex             // uint32
	return
}

func (p *PassTimestampWrites) wrap(in *sys.PassTimestampWrites) {
	panic("TODO")
}

func (p *PipelineLayoutDescriptor) unwrap() (out sys.PipelineLayoutDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(p.Label)),
		Data:   unsafe.StringData(p.Label),
	}
	bindGroupLayouts := make([]sys.BindGroupLayout, len(p.BindGroupLayouts))
	for i := range bindGroupLayouts {
		bindGroupLayouts[i] = p.BindGroupLayouts[i].inner
	}
	out.BindGroupLayouts = &bindGroupLayouts[0]
	out.ImmediateSize = p.ImmediateSize // uint32
	return
}

func (p *PipelineLayoutDescriptor) wrap(in *sys.PipelineLayoutDescriptor) {
	panic("TODO")
}

func (q *QuerySetDescriptor) unwrap() (out sys.QuerySetDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(q.Label)),
		Data:   unsafe.StringData(q.Label),
	}
	out.Type = q.Type   // enum.query_type
	out.Count = q.Count // uint32
	return
}

func (q *QuerySetDescriptor) wrap(in *sys.QuerySetDescriptor) {
	panic("TODO")
}

func (q *QueueDescriptor) unwrap() (out sys.QueueDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(q.Label)),
		Data:   unsafe.StringData(q.Label),
	}
	return
}

func (q *QueueDescriptor) wrap(in *sys.QueueDescriptor) {
	panic("TODO")
}

func (r *RenderBundleDescriptor) unwrap() (out sys.RenderBundleDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}
	return
}

func (r *RenderBundleDescriptor) wrap(in *sys.RenderBundleDescriptor) {
	panic("TODO")
}

func (r *RenderBundleEncoderDescriptor) unwrap() (out sys.RenderBundleEncoderDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}

	out.ColorFormats = &r.ColorFormats[0]
	out.DepthStencilFormat = r.DepthStencilFormat // enum.texture_format
	out.SampleCount = r.SampleCount               // uint32
	out.DepthReadOnly = r.DepthReadOnly           // bool
	out.StencilReadOnly = r.StencilReadOnly       // bool
	return
}

func (r *RenderBundleEncoderDescriptor) wrap(in *sys.RenderBundleEncoderDescriptor) {
	panic("TODO")
}

func (r *RenderPassColorAttachment) unwrap() (out sys.RenderPassColorAttachment) {
	out.View = r.View.inner
	out.DepthSlice = r.DepthSlice // uint32
	out.ResolveTarget = r.ResolveTarget.inner
	out.LoadOp = r.LoadOp   // enum.load_op
	out.StoreOp = r.StoreOp // enum.store_op
	out.ClearValue = r.ClearValue
	return
}

func (r *RenderPassColorAttachment) wrap(in *sys.RenderPassColorAttachment) {
	panic("TODO")
}

func (r *RenderPassDepthStencilAttachment) unwrap() (out sys.RenderPassDepthStencilAttachment) {
	out.View = r.View.inner
	out.DepthLoadOp = r.DepthLoadOp             // enum.load_op
	out.DepthStoreOp = r.DepthStoreOp           // enum.store_op
	out.DepthClearValue = r.DepthClearValue     // nullable_float32
	out.DepthReadOnly = r.DepthReadOnly         // bool
	out.StencilLoadOp = r.StencilLoadOp         // enum.load_op
	out.StencilStoreOp = r.StencilStoreOp       // enum.store_op
	out.StencilClearValue = r.StencilClearValue // uint32
	out.StencilReadOnly = r.StencilReadOnly     // bool
	return
}

func (r *RenderPassDepthStencilAttachment) wrap(in *sys.RenderPassDepthStencilAttachment) {
	panic("TODO")
}

func (r *RenderPassDescriptor) unwrap() (out sys.RenderPassDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}
	out.ColorAttachmentsCount = uint(len(r.ColorAttachments))
	colorAttachments := make([]sys.RenderPassColorAttachment, len(r.ColorAttachments))
	for i := range colorAttachments {
		colorAttachments[i] = r.ColorAttachments[i].unwrap()
	}
	out.ColorAttachments = &colorAttachments[0]
	if r.DepthStencilAttachment != nil {
		out.DepthStencilAttachment = new(r.DepthStencilAttachment.unwrap())
	}
	out.OcclusionQuerySet = r.OcclusionQuerySet.inner
	if r.TimestampWrites != nil {
		out.TimestampWrites = new(r.TimestampWrites.unwrap())
	}
	return
}

func (r *RenderPassDescriptor) wrap(in *sys.RenderPassDescriptor) {
	panic("TODO")
}

func (r *RenderPipelineDescriptor) unwrap() (out sys.RenderPipelineDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}
	out.Layout = r.Layout.inner
	out.Vertex = r.Vertex.unwrap()
	out.Primitive = r.Primitive
	if r.DepthStencil != nil {
		out.DepthStencil = r.DepthStencil
	}
	out.Multisample = r.Multisample
	if r.Fragment != nil {
		out.Fragment = new(r.Fragment.unwrap())
	}
	return
}

func (r *RenderPipelineDescriptor) wrap(in *sys.RenderPipelineDescriptor) {
	panic("TODO")
}

func (r *RequestAdapterOptions) unwrap() (out sys.RequestAdapterOptions) {
	out.FeatureLevel = r.FeatureLevel                 // enum.feature_level
	out.PowerPreference = r.PowerPreference           // enum.power_preference
	out.ForceFallbackAdapter = r.ForceFallbackAdapter // bool
	out.BackendType = r.BackendType                   // enum.backend_type
	out.CompatibleSurface = r.CompatibleSurface.inner
	return
}

func (r *RequestAdapterOptions) wrap(in *sys.RequestAdapterOptions) {
	panic("TODO")
}

func (s *SamplerDescriptor) unwrap() (out sys.SamplerDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(s.Label)),
		Data:   unsafe.StringData(s.Label),
	}
	out.AddressModeU = s.AddressModeU   // enum.address_mode
	out.AddressModeV = s.AddressModeV   // enum.address_mode
	out.AddressModeW = s.AddressModeW   // enum.address_mode
	out.MagFilter = s.MagFilter         // enum.filter_mode
	out.MinFilter = s.MinFilter         // enum.filter_mode
	out.MipmapFilter = s.MipmapFilter   // enum.mipmap_filter_mode
	out.LodMinClamp = s.LodMinClamp     // float32
	out.LodMaxClamp = s.LodMaxClamp     // float32
	out.Compare = s.Compare             // enum.compare_function
	out.MaxAnisotropy = s.MaxAnisotropy // uint16
	return
}

func (s *SamplerDescriptor) wrap(in *sys.SamplerDescriptor) {
	panic("TODO")
}

func (s *ShaderModuleDescriptor) unwrap() (out sys.ShaderModuleDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(s.Label)),
		Data:   unsafe.StringData(s.Label),
	}
	return
}

func (s *ShaderModuleDescriptor) wrap(in *sys.ShaderModuleDescriptor) {
	panic("TODO")
}

func (s *ShaderSourceWGSL) unwrap() (out sys.ShaderSourceWGSL) {
	out.Code = sys.StringView{
		Length: uint(len(s.Code)),
		Data:   unsafe.StringData(s.Code),
	}
	return
}

func (s *ShaderSourceWGSL) wrap(in *sys.ShaderSourceWGSL) {
	panic("TODO")
}

func (s *SupportedFeatures) unwrap() (out sys.SupportedFeatures) {

	out.Features = &s.Features[0]
	return
}

func (s *SupportedFeatures) wrap(in *sys.SupportedFeatures) {
	panic("TODO")
}

func (s *SupportedInstanceFeatures) unwrap() (out sys.SupportedInstanceFeatures) {

	out.Features = &s.Features[0]
	return
}

func (s *SupportedInstanceFeatures) wrap(in *sys.SupportedInstanceFeatures) {
	panic("TODO")
}

func (s *SupportedWGSLLanguageFeatures) unwrap() (out sys.SupportedWGSLLanguageFeatures) {

	out.Features = &s.Features[0]
	return
}

func (s *SupportedWGSLLanguageFeatures) wrap(in *sys.SupportedWGSLLanguageFeatures) {
	panic("TODO")
}

func (s *SurfaceCapabilities) unwrap() (out sys.SurfaceCapabilities) {
	out.Usages = s.Usages // bitflag.texture_usage

	out.Formats = &s.Formats[0]

	out.PresentModes = &s.PresentModes[0]

	out.AlphaModes = &s.AlphaModes[0]
	return
}

func (s *SurfaceCapabilities) wrap(in *sys.SurfaceCapabilities) {
	panic("TODO")
}

func (s *SurfaceConfiguration) unwrap() (out sys.SurfaceConfiguration) {
	out.Device = s.Device.inner
	out.Format = s.Format // enum.texture_format
	out.Usage = s.Usage   // bitflag.texture_usage
	out.Width = s.Width   // uint32
	out.Height = s.Height // uint32

	out.ViewFormats = &s.ViewFormats[0]
	out.AlphaMode = s.AlphaMode     // enum.composite_alpha_mode
	out.PresentMode = s.PresentMode // enum.present_mode
	return
}

func (s *SurfaceConfiguration) wrap(in *sys.SurfaceConfiguration) {
	panic("TODO")
}

func (s *SurfaceDescriptor) unwrap() (out sys.SurfaceDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(s.Label)),
		Data:   unsafe.StringData(s.Label),
	}
	return
}

func (s *SurfaceDescriptor) wrap(in *sys.SurfaceDescriptor) {
	panic("TODO")
}

func (s *SurfaceTexture) unwrap() (out sys.SurfaceTexture) {
	out.Texture = s.Texture.inner
	out.Status = s.Status // enum.surface_get_current_texture_status
	return
}

func (s *SurfaceTexture) wrap(in *sys.SurfaceTexture) {
	panic("TODO")
}

func (t *TexelCopyBufferInfo) unwrap() (out sys.TexelCopyBufferInfo) {
	out.Layout = t.Layout
	out.Buffer = t.Buffer.inner
	return
}

func (t *TexelCopyBufferInfo) wrap(in *sys.TexelCopyBufferInfo) {
	panic("TODO")
}

func (t *TexelCopyTextureInfo) unwrap() (out sys.TexelCopyTextureInfo) {
	out.Texture = t.Texture.inner
	out.MipLevel = t.MipLevel // uint32
	out.Origin = t.Origin
	out.Aspect = t.Aspect // enum.texture_aspect
	return
}

func (t *TexelCopyTextureInfo) wrap(in *sys.TexelCopyTextureInfo) {
	panic("TODO")
}

func (t *TextureDescriptor) unwrap() (out sys.TextureDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(t.Label)),
		Data:   unsafe.StringData(t.Label),
	}
	out.Usage = t.Usage         // bitflag.texture_usage
	out.Dimension = t.Dimension // enum.texture_dimension
	out.Size = t.Size
	out.Format = t.Format               // enum.texture_format
	out.MipLevelCount = t.MipLevelCount // uint32
	out.SampleCount = t.SampleCount     // uint32

	out.ViewFormats = &t.ViewFormats[0]
	return
}

func (t *TextureDescriptor) wrap(in *sys.TextureDescriptor) {
	panic("TODO")
}

func (t *TextureViewDescriptor) unwrap() (out sys.TextureViewDescriptor) {
	out.Label = sys.StringView{
		Length: uint(len(t.Label)),
		Data:   unsafe.StringData(t.Label),
	}
	out.Format = t.Format                   // enum.texture_format
	out.Dimension = t.Dimension             // enum.texture_view_dimension
	out.BaseMipLevel = t.BaseMipLevel       // uint32
	out.MipLevelCount = t.MipLevelCount     // uint32
	out.BaseArrayLayer = t.BaseArrayLayer   // uint32
	out.ArrayLayerCount = t.ArrayLayerCount // uint32
	out.Aspect = t.Aspect                   // enum.texture_aspect
	out.Usage = t.Usage                     // bitflag.texture_usage
	return
}

func (t *TextureViewDescriptor) wrap(in *sys.TextureViewDescriptor) {
	panic("TODO")
}

func (v *VertexBufferLayout) unwrap() (out sys.VertexBufferLayout) {
	out.StepMode = v.StepMode       // enum.vertex_step_mode
	out.ArrayStride = v.ArrayStride // uint64
	out.AttributesCount = uint(len(v.Attributes))

	out.Attributes = &v.Attributes[0]
	return
}

func (v *VertexBufferLayout) wrap(in *sys.VertexBufferLayout) {
	panic("TODO")
}

func (v *VertexState) unwrap() (out sys.VertexState) {
	out.Module = v.Module.inner
	out.EntryPoint = sys.StringView{
		Length: uint(len(v.EntryPoint)),
		Data:   unsafe.StringData(v.EntryPoint),
	}
	out.ConstantsCount = uint(len(v.Constants))
	constants := make([]sys.ConstantEntry, len(v.Constants))
	for i := range constants {
		constants[i] = v.Constants[i].unwrap()
	}
	out.Constants = &constants[0]
	out.BuffersCount = uint(len(v.Buffers))
	buffers := make([]sys.VertexBufferLayout, len(v.Buffers))
	for i := range buffers {
		buffers[i] = v.Buffers[i].unwrap()
	}
	out.Buffers = &buffers[0]
	return
}

func (v *VertexState) wrap(in *sys.VertexState) {
	panic("TODO")
}
