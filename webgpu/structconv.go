// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import (
	"strings"
	"unsafe"

	"github.com/Tnze/go-webgpu/webgpu/sys"
)

func (a *AdapterInfo) unwrap() (out sys.AdapterInfo) {
	out.Chain = a.Chain
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
	out.BackendType = a.BackendType
	out.AdapterType = a.AdapterType
	out.VendorID = a.VendorID
	out.DeviceID = a.DeviceID
	out.SubgroupMinSize = a.SubgroupMinSize
	out.SubgroupMaxSize = a.SubgroupMaxSize
	return
}

func (a *AdapterInfo) wrap(in *sys.AdapterInfo) {
	a.Chain = in.Chain
	a.Vendor = strings.Clone(unsafe.String(
		in.Vendor.Data,
		in.Vendor.Length,
	))
	a.Architecture = strings.Clone(unsafe.String(
		in.Architecture.Data,
		in.Architecture.Length,
	))
	a.Device = strings.Clone(unsafe.String(
		in.Device.Data,
		in.Device.Length,
	))
	a.Description = strings.Clone(unsafe.String(
		in.Description.Data,
		in.Description.Length,
	))
	a.BackendType = in.BackendType
	a.AdapterType = in.AdapterType
	a.VendorID = in.VendorID
	a.DeviceID = in.DeviceID
	a.SubgroupMinSize = in.SubgroupMinSize
	a.SubgroupMaxSize = in.SubgroupMaxSize
}

func (b *BindGroupDescriptor) unwrap() (out sys.BindGroupDescriptor) {
	out.Chain = b.Chain
	out.Label = sys.StringView{
		Length: uint(len(b.Label)),
		Data:   unsafe.StringData(b.Label),
	}
	if b.Layout != nil {
		out.Layout = b.Layout.inner
	}
	if b.Entries != nil {
		out.EntriesCount = uint(len(b.Entries))
		entries := make([]sys.BindGroupEntry, len(b.Entries))
		for i := range entries {
			entries[i] = b.Entries[i].unwrap()
		}
		out.Entries = unsafe.SliceData(entries)
	}
	return
}

func (b *BindGroupDescriptor) wrap(in *sys.BindGroupDescriptor) {
	b.Chain = in.Chain
	b.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	b.Layout = new(BindGroupLayout{inner: in.Layout}).owned()
	_entries := unsafe.Slice(in.Entries, in.EntriesCount)
	entries := make([]BindGroupEntry, in.EntriesCount)
	for i := uint(0); i < in.EntriesCount; i++ {
		entries[i].wrap(&_entries[i])
	}
	b.Entries = entries

}

func (b *BindGroupEntry) unwrap() (out sys.BindGroupEntry) {
	out.Chain = b.Chain
	out.Binding = b.Binding
	if b.Buffer != nil {
		out.Buffer = b.Buffer.inner
	}
	out.Offset = b.Offset
	out.Size = b.Size
	if b.Sampler != nil {
		out.Sampler = b.Sampler.inner
	}
	if b.TextureView != nil {
		out.TextureView = b.TextureView.inner
	}
	return
}

func (b *BindGroupEntry) wrap(in *sys.BindGroupEntry) {
	b.Chain = in.Chain
	b.Binding = in.Binding
	b.Buffer = new(Buffer{inner: in.Buffer}).owned()
	b.Offset = in.Offset
	b.Size = in.Size
	b.Sampler = new(Sampler{inner: in.Sampler}).owned()
	b.TextureView = new(TextureView{inner: in.TextureView}).owned()
}

func (b *BindGroupLayoutDescriptor) unwrap() (out sys.BindGroupLayoutDescriptor) {
	out.Chain = b.Chain
	out.Label = sys.StringView{
		Length: uint(len(b.Label)),
		Data:   unsafe.StringData(b.Label),
	}
	if b.Entries != nil {
		out.EntriesCount = uint(len(b.Entries))
		out.Entries = unsafe.SliceData(b.Entries)
	}
	return
}

func (b *BindGroupLayoutDescriptor) wrap(in *sys.BindGroupLayoutDescriptor) {
	b.Chain = in.Chain
	b.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	b.Entries = unsafe.Slice(in.Entries, in.EntriesCount)
}

func (b *BufferDescriptor) unwrap() (out sys.BufferDescriptor) {
	out.Chain = b.Chain
	out.Label = sys.StringView{
		Length: uint(len(b.Label)),
		Data:   unsafe.StringData(b.Label),
	}
	out.Usage = b.Usage
	out.Size = b.Size
	out.MappedAtCreation = b.MappedAtCreation
	return
}

func (b *BufferDescriptor) wrap(in *sys.BufferDescriptor) {
	b.Chain = in.Chain
	b.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	b.Usage = in.Usage
	b.Size = in.Size
	b.MappedAtCreation = in.MappedAtCreation
}

func (c *CommandBufferDescriptor) unwrap() (out sys.CommandBufferDescriptor) {
	out.Chain = c.Chain
	out.Label = sys.StringView{
		Length: uint(len(c.Label)),
		Data:   unsafe.StringData(c.Label),
	}
	return
}

func (c *CommandBufferDescriptor) wrap(in *sys.CommandBufferDescriptor) {
	c.Chain = in.Chain
	c.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
}

func (c *CommandEncoderDescriptor) unwrap() (out sys.CommandEncoderDescriptor) {
	out.Chain = c.Chain
	out.Label = sys.StringView{
		Length: uint(len(c.Label)),
		Data:   unsafe.StringData(c.Label),
	}
	return
}

func (c *CommandEncoderDescriptor) wrap(in *sys.CommandEncoderDescriptor) {
	c.Chain = in.Chain
	c.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
}

func (c *CompilationInfo) unwrap() (out sys.CompilationInfo) {
	if c.Messages != nil {
		out.MessagesCount = uint(len(c.Messages))
		messages := make([]sys.CompilationMessage, len(c.Messages))
		for i := range messages {
			messages[i] = c.Messages[i].unwrap()
		}
		out.Messages = unsafe.SliceData(messages)
	}
	return
}

func (c *CompilationInfo) wrap(in *sys.CompilationInfo) {
	_messages := unsafe.Slice(in.Messages, in.MessagesCount)
	messages := make([]CompilationMessage, in.MessagesCount)
	for i := uint(0); i < in.MessagesCount; i++ {
		messages[i].wrap(&_messages[i])
	}
	c.Messages = messages

}

func (c *CompilationMessage) unwrap() (out sys.CompilationMessage) {
	out.Message = sys.StringView{
		Length: uint(len(c.Message)),
		Data:   unsafe.StringData(c.Message),
	}
	out.Type = c.Type
	out.LineNum = c.LineNum
	out.LinePos = c.LinePos
	out.Offset = c.Offset
	out.Length = c.Length
	return
}

func (c *CompilationMessage) wrap(in *sys.CompilationMessage) {
	c.Message = strings.Clone(unsafe.String(
		in.Message.Data,
		in.Message.Length,
	))
	c.Type = in.Type
	c.LineNum = in.LineNum
	c.LinePos = in.LinePos
	c.Offset = in.Offset
	c.Length = in.Length
}

func (c *ComputePassDescriptor) unwrap() (out sys.ComputePassDescriptor) {
	out.Chain = c.Chain
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
	c.Chain = in.Chain
	c.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	if in.TimestampWrites != nil {
		c.TimestampWrites = new(PassTimestampWrites)
		c.TimestampWrites.wrap(in.TimestampWrites)
	}
}

func (c *ComputePipelineDescriptor) unwrap() (out sys.ComputePipelineDescriptor) {
	out.Chain = c.Chain
	out.Label = sys.StringView{
		Length: uint(len(c.Label)),
		Data:   unsafe.StringData(c.Label),
	}
	if c.Layout != nil {
		out.Layout = c.Layout.inner
	}
	out.Compute = c.Compute.unwrap()
	return
}

func (c *ComputePipelineDescriptor) wrap(in *sys.ComputePipelineDescriptor) {
	c.Chain = in.Chain
	c.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	c.Layout = new(PipelineLayout{inner: in.Layout}).owned()
	c.Compute.wrap(&in.Compute)
}

func (c *ComputeState) unwrap() (out sys.ComputeState) {
	out.Chain = c.Chain
	if c.Module != nil {
		out.Module = c.Module.inner
	}
	out.EntryPoint = sys.StringView{
		Length: uint(len(c.EntryPoint)),
		Data:   unsafe.StringData(c.EntryPoint),
	}
	if c.Constants != nil {
		out.ConstantsCount = uint(len(c.Constants))
		constants := make([]sys.ConstantEntry, len(c.Constants))
		for i := range constants {
			constants[i] = c.Constants[i].unwrap()
		}
		out.Constants = unsafe.SliceData(constants)
	}
	return
}

func (c *ComputeState) wrap(in *sys.ComputeState) {
	c.Chain = in.Chain
	c.Module = new(ShaderModule{inner: in.Module}).owned()
	c.EntryPoint = strings.Clone(unsafe.String(
		in.EntryPoint.Data,
		in.EntryPoint.Length,
	))
	_constants := unsafe.Slice(in.Constants, in.ConstantsCount)
	constants := make([]ConstantEntry, in.ConstantsCount)
	for i := uint(0); i < in.ConstantsCount; i++ {
		constants[i].wrap(&_constants[i])
	}
	c.Constants = constants

}

func (c *ConstantEntry) unwrap() (out sys.ConstantEntry) {
	out.Chain = c.Chain
	out.Key = sys.StringView{
		Length: uint(len(c.Key)),
		Data:   unsafe.StringData(c.Key),
	}
	out.Value = c.Value
	return
}

func (c *ConstantEntry) wrap(in *sys.ConstantEntry) {
	c.Chain = in.Chain
	c.Key = strings.Clone(unsafe.String(
		in.Key.Data,
		in.Key.Length,
	))
	c.Value = in.Value
}

func (d *DeviceDescriptor) unwrap() (out sys.DeviceDescriptor) {
	out.Chain = d.Chain
	out.Label = sys.StringView{
		Length: uint(len(d.Label)),
		Data:   unsafe.StringData(d.Label),
	}
	if d.RequiredFeatures != nil {
		out.RequiredFeaturesCount = uint(len(d.RequiredFeatures))
		out.RequiredFeatures = unsafe.SliceData(d.RequiredFeatures)
	}
	if d.RequiredLimits != nil {
		out.RequiredLimits = d.RequiredLimits
	}
	out.DefaultQueue = d.DefaultQueue.unwrap()
	out.DeviceLostCallbackInfo = d.DeviceLostCallbackInfo
	out.UncapturedErrorCallbackInfo = d.UncapturedErrorCallbackInfo
	return
}

func (d *DeviceDescriptor) wrap(in *sys.DeviceDescriptor) {
	d.Chain = in.Chain
	d.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	if in.RequiredLimits != nil {
		d.RequiredLimits = in.RequiredLimits
	}
	d.DefaultQueue.wrap(&in.DefaultQueue)
	d.DeviceLostCallbackInfo = in.DeviceLostCallbackInfo
	d.UncapturedErrorCallbackInfo = in.UncapturedErrorCallbackInfo
}

func (e *ExternalTextureBindingEntry) unwrap() (out sys.ExternalTextureBindingEntry) {
	if e.ExternalTexture != nil {
		out.ExternalTexture = e.ExternalTexture.inner
	}
	return
}

func (e *ExternalTextureBindingEntry) wrap(in *sys.ExternalTextureBindingEntry) {
	e.ExternalTexture = new(ExternalTexture{inner: in.ExternalTexture}).owned()
}

func (f *FragmentState) unwrap() (out sys.FragmentState) {
	out.Chain = f.Chain
	if f.Module != nil {
		out.Module = f.Module.inner
	}
	out.EntryPoint = sys.StringView{
		Length: uint(len(f.EntryPoint)),
		Data:   unsafe.StringData(f.EntryPoint),
	}
	if f.Constants != nil {
		out.ConstantsCount = uint(len(f.Constants))
		constants := make([]sys.ConstantEntry, len(f.Constants))
		for i := range constants {
			constants[i] = f.Constants[i].unwrap()
		}
		out.Constants = unsafe.SliceData(constants)
	}
	if f.Targets != nil {
		out.TargetsCount = uint(len(f.Targets))
		out.Targets = unsafe.SliceData(f.Targets)
	}
	return
}

func (f *FragmentState) wrap(in *sys.FragmentState) {
	f.Chain = in.Chain
	f.Module = new(ShaderModule{inner: in.Module}).owned()
	f.EntryPoint = strings.Clone(unsafe.String(
		in.EntryPoint.Data,
		in.EntryPoint.Length,
	))
	_constants := unsafe.Slice(in.Constants, in.ConstantsCount)
	constants := make([]ConstantEntry, in.ConstantsCount)
	for i := uint(0); i < in.ConstantsCount; i++ {
		constants[i].wrap(&_constants[i])
	}
	f.Constants = constants

	f.Targets = unsafe.Slice(in.Targets, in.TargetsCount)
}

func (i *InstanceDescriptor) unwrap() (out sys.InstanceDescriptor) {
	out.Chain = i.Chain
	if i.RequiredFeatures != nil {
		out.RequiredFeaturesCount = uint(len(i.RequiredFeatures))
		out.RequiredFeatures = unsafe.SliceData(i.RequiredFeatures)
	}
	if i.RequiredLimits != nil {
		out.RequiredLimits = i.RequiredLimits
	}
	return
}

func (i *InstanceDescriptor) wrap(in *sys.InstanceDescriptor) {
	i.Chain = in.Chain
	if in.RequiredLimits != nil {
		i.RequiredLimits = in.RequiredLimits
	}
}

func (p *PassTimestampWrites) unwrap() (out sys.PassTimestampWrites) {
	out.Chain = p.Chain
	if p.QuerySet != nil {
		out.QuerySet = p.QuerySet.inner
	}
	out.BeginningOfPassWriteIndex = p.BeginningOfPassWriteIndex
	out.EndOfPassWriteIndex = p.EndOfPassWriteIndex
	return
}

func (p *PassTimestampWrites) wrap(in *sys.PassTimestampWrites) {
	p.Chain = in.Chain
	p.QuerySet = new(QuerySet{inner: in.QuerySet}).owned()
	p.BeginningOfPassWriteIndex = in.BeginningOfPassWriteIndex
	p.EndOfPassWriteIndex = in.EndOfPassWriteIndex
}

func (p *PipelineLayoutDescriptor) unwrap() (out sys.PipelineLayoutDescriptor) {
	out.Chain = p.Chain
	out.Label = sys.StringView{
		Length: uint(len(p.Label)),
		Data:   unsafe.StringData(p.Label),
	}
	if p.BindGroupLayouts != nil {
		out.BindGroupLayoutsCount = uint(len(p.BindGroupLayouts))
		bindGroupLayouts := make([]sys.BindGroupLayout, len(p.BindGroupLayouts))
		for i := range bindGroupLayouts {
			bindGroupLayouts[i] = p.BindGroupLayouts[i].inner
		}
		out.BindGroupLayouts = unsafe.SliceData(bindGroupLayouts)
	}
	out.ImmediateSize = p.ImmediateSize
	return
}

func (p *PipelineLayoutDescriptor) wrap(in *sys.PipelineLayoutDescriptor) {
	p.Chain = in.Chain
	p.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	_bindGroupLayouts := unsafe.Slice(in.BindGroupLayouts, in.BindGroupLayoutsCount)
	bindGroupLayouts := make([]*BindGroupLayout, in.BindGroupLayoutsCount)
	for i := uint(0); i < in.BindGroupLayoutsCount; i++ {
		bindGroupLayouts[i] = new(BindGroupLayout{inner: _bindGroupLayouts[i]}).owned()
	}
	p.BindGroupLayouts = bindGroupLayouts
	p.ImmediateSize = in.ImmediateSize
}

func (q *QuerySetDescriptor) unwrap() (out sys.QuerySetDescriptor) {
	out.Chain = q.Chain
	out.Label = sys.StringView{
		Length: uint(len(q.Label)),
		Data:   unsafe.StringData(q.Label),
	}
	out.Type = q.Type
	out.Count = q.Count
	return
}

func (q *QuerySetDescriptor) wrap(in *sys.QuerySetDescriptor) {
	q.Chain = in.Chain
	q.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	q.Type = in.Type
	q.Count = in.Count
}

func (q *QueueDescriptor) unwrap() (out sys.QueueDescriptor) {
	out.Chain = q.Chain
	out.Label = sys.StringView{
		Length: uint(len(q.Label)),
		Data:   unsafe.StringData(q.Label),
	}
	return
}

func (q *QueueDescriptor) wrap(in *sys.QueueDescriptor) {
	q.Chain = in.Chain
	q.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
}

func (r *RenderBundleDescriptor) unwrap() (out sys.RenderBundleDescriptor) {
	out.Chain = r.Chain
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}
	return
}

func (r *RenderBundleDescriptor) wrap(in *sys.RenderBundleDescriptor) {
	r.Chain = in.Chain
	r.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
}

func (r *RenderBundleEncoderDescriptor) unwrap() (out sys.RenderBundleEncoderDescriptor) {
	out.Chain = r.Chain
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}
	if r.ColorFormats != nil {
		out.ColorFormatsCount = uint(len(r.ColorFormats))
		out.ColorFormats = unsafe.SliceData(r.ColorFormats)
	}
	out.DepthStencilFormat = r.DepthStencilFormat
	out.SampleCount = r.SampleCount
	out.DepthReadOnly = r.DepthReadOnly
	out.StencilReadOnly = r.StencilReadOnly
	return
}

func (r *RenderBundleEncoderDescriptor) wrap(in *sys.RenderBundleEncoderDescriptor) {
	r.Chain = in.Chain
	r.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	r.DepthStencilFormat = in.DepthStencilFormat
	r.SampleCount = in.SampleCount
	r.DepthReadOnly = in.DepthReadOnly
	r.StencilReadOnly = in.StencilReadOnly
}

func (r *RenderPassColorAttachment) unwrap() (out sys.RenderPassColorAttachment) {
	out.Chain = r.Chain
	if r.View != nil {
		out.View = r.View.inner
	}
	out.DepthSlice = r.DepthSlice
	if r.ResolveTarget != nil {
		out.ResolveTarget = r.ResolveTarget.inner
	}
	out.LoadOp = r.LoadOp
	out.StoreOp = r.StoreOp
	out.ClearValue = r.ClearValue
	return
}

func (r *RenderPassColorAttachment) wrap(in *sys.RenderPassColorAttachment) {
	r.Chain = in.Chain
	r.View = new(TextureView{inner: in.View}).owned()
	r.DepthSlice = in.DepthSlice
	r.ResolveTarget = new(TextureView{inner: in.ResolveTarget}).owned()
	r.LoadOp = in.LoadOp
	r.StoreOp = in.StoreOp
	r.ClearValue = in.ClearValue
}

func (r *RenderPassDepthStencilAttachment) unwrap() (out sys.RenderPassDepthStencilAttachment) {
	out.Chain = r.Chain
	if r.View != nil {
		out.View = r.View.inner
	}
	out.DepthLoadOp = r.DepthLoadOp
	out.DepthStoreOp = r.DepthStoreOp
	out.DepthClearValue = r.DepthClearValue
	out.DepthReadOnly = r.DepthReadOnly
	out.StencilLoadOp = r.StencilLoadOp
	out.StencilStoreOp = r.StencilStoreOp
	out.StencilClearValue = r.StencilClearValue
	out.StencilReadOnly = r.StencilReadOnly
	return
}

func (r *RenderPassDepthStencilAttachment) wrap(in *sys.RenderPassDepthStencilAttachment) {
	r.Chain = in.Chain
	r.View = new(TextureView{inner: in.View}).owned()
	r.DepthLoadOp = in.DepthLoadOp
	r.DepthStoreOp = in.DepthStoreOp
	r.DepthClearValue = in.DepthClearValue
	r.DepthReadOnly = in.DepthReadOnly
	r.StencilLoadOp = in.StencilLoadOp
	r.StencilStoreOp = in.StencilStoreOp
	r.StencilClearValue = in.StencilClearValue
	r.StencilReadOnly = in.StencilReadOnly
}

func (r *RenderPassDescriptor) unwrap() (out sys.RenderPassDescriptor) {
	out.Chain = r.Chain
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}
	if r.ColorAttachments != nil {
		out.ColorAttachmentsCount = uint(len(r.ColorAttachments))
		colorAttachments := make([]sys.RenderPassColorAttachment, len(r.ColorAttachments))
		for i := range colorAttachments {
			colorAttachments[i] = r.ColorAttachments[i].unwrap()
		}
		out.ColorAttachments = unsafe.SliceData(colorAttachments)
	}
	if r.DepthStencilAttachment != nil {
		out.DepthStencilAttachment = new(r.DepthStencilAttachment.unwrap())
	}
	if r.OcclusionQuerySet != nil {
		out.OcclusionQuerySet = r.OcclusionQuerySet.inner
	}
	if r.TimestampWrites != nil {
		out.TimestampWrites = new(r.TimestampWrites.unwrap())
	}
	return
}

func (r *RenderPassDescriptor) wrap(in *sys.RenderPassDescriptor) {
	r.Chain = in.Chain
	r.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	_colorAttachments := unsafe.Slice(in.ColorAttachments, in.ColorAttachmentsCount)
	colorAttachments := make([]RenderPassColorAttachment, in.ColorAttachmentsCount)
	for i := uint(0); i < in.ColorAttachmentsCount; i++ {
		colorAttachments[i].wrap(&_colorAttachments[i])
	}
	r.ColorAttachments = colorAttachments

	if in.DepthStencilAttachment != nil {
		r.DepthStencilAttachment = new(RenderPassDepthStencilAttachment)
		r.DepthStencilAttachment.wrap(in.DepthStencilAttachment)
	}
	r.OcclusionQuerySet = new(QuerySet{inner: in.OcclusionQuerySet}).owned()
	if in.TimestampWrites != nil {
		r.TimestampWrites = new(PassTimestampWrites)
		r.TimestampWrites.wrap(in.TimestampWrites)
	}
}

func (r *RenderPipelineDescriptor) unwrap() (out sys.RenderPipelineDescriptor) {
	out.Chain = r.Chain
	out.Label = sys.StringView{
		Length: uint(len(r.Label)),
		Data:   unsafe.StringData(r.Label),
	}
	if r.Layout != nil {
		out.Layout = r.Layout.inner
	}
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
	r.Chain = in.Chain
	r.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	r.Layout = new(PipelineLayout{inner: in.Layout}).owned()
	r.Vertex.wrap(&in.Vertex)
	r.Primitive = in.Primitive
	if in.DepthStencil != nil {
		r.DepthStencil = in.DepthStencil
	}
	r.Multisample = in.Multisample
	if in.Fragment != nil {
		r.Fragment = new(FragmentState)
		r.Fragment.wrap(in.Fragment)
	}
}

func (r *RequestAdapterOptions) unwrap() (out sys.RequestAdapterOptions) {
	out.Chain = r.Chain
	out.FeatureLevel = r.FeatureLevel
	out.PowerPreference = r.PowerPreference
	out.ForceFallbackAdapter = r.ForceFallbackAdapter
	out.BackendType = r.BackendType
	if r.CompatibleSurface != nil {
		out.CompatibleSurface = r.CompatibleSurface.inner
	}
	return
}

func (r *RequestAdapterOptions) wrap(in *sys.RequestAdapterOptions) {
	r.Chain = in.Chain
	r.FeatureLevel = in.FeatureLevel
	r.PowerPreference = in.PowerPreference
	r.ForceFallbackAdapter = in.ForceFallbackAdapter
	r.BackendType = in.BackendType
	r.CompatibleSurface = new(Surface{inner: in.CompatibleSurface}).owned()
}

func (s *SamplerDescriptor) unwrap() (out sys.SamplerDescriptor) {
	out.Chain = s.Chain
	out.Label = sys.StringView{
		Length: uint(len(s.Label)),
		Data:   unsafe.StringData(s.Label),
	}
	out.AddressModeU = s.AddressModeU
	out.AddressModeV = s.AddressModeV
	out.AddressModeW = s.AddressModeW
	out.MagFilter = s.MagFilter
	out.MinFilter = s.MinFilter
	out.MipmapFilter = s.MipmapFilter
	out.LodMinClamp = s.LodMinClamp
	out.LodMaxClamp = s.LodMaxClamp
	out.Compare = s.Compare
	out.MaxAnisotropy = s.MaxAnisotropy
	return
}

func (s *SamplerDescriptor) wrap(in *sys.SamplerDescriptor) {
	s.Chain = in.Chain
	s.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	s.AddressModeU = in.AddressModeU
	s.AddressModeV = in.AddressModeV
	s.AddressModeW = in.AddressModeW
	s.MagFilter = in.MagFilter
	s.MinFilter = in.MinFilter
	s.MipmapFilter = in.MipmapFilter
	s.LodMinClamp = in.LodMinClamp
	s.LodMaxClamp = in.LodMaxClamp
	s.Compare = in.Compare
	s.MaxAnisotropy = in.MaxAnisotropy
}

func (s *ShaderModuleDescriptor) unwrap() (out sys.ShaderModuleDescriptor) {
	out.Chain = s.Chain
	out.Label = sys.StringView{
		Length: uint(len(s.Label)),
		Data:   unsafe.StringData(s.Label),
	}
	return
}

func (s *ShaderModuleDescriptor) wrap(in *sys.ShaderModuleDescriptor) {
	s.Chain = in.Chain
	s.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
}

func (s *ShaderSourceWGSL) unwrap() (out sys.ShaderSourceWGSL) {
	out.Code = sys.StringView{
		Length: uint(len(s.Code)),
		Data:   unsafe.StringData(s.Code),
	}
	return
}

func (s *ShaderSourceWGSL) wrap(in *sys.ShaderSourceWGSL) {
	s.Code = strings.Clone(unsafe.String(
		in.Code.Data,
		in.Code.Length,
	))
}

func (s *SupportedFeatures) unwrap() (out sys.SupportedFeatures) {
	if s.Features != nil {
		out.FeaturesCount = uint(len(s.Features))
		out.Features = unsafe.SliceData(s.Features)
	}
	return
}

func (s *SupportedFeatures) wrap(in *sys.SupportedFeatures) {
}

func (s *SupportedInstanceFeatures) unwrap() (out sys.SupportedInstanceFeatures) {
	if s.Features != nil {
		out.FeaturesCount = uint(len(s.Features))
		out.Features = unsafe.SliceData(s.Features)
	}
	return
}

func (s *SupportedInstanceFeatures) wrap(in *sys.SupportedInstanceFeatures) {
}

func (s *SupportedWGSLLanguageFeatures) unwrap() (out sys.SupportedWGSLLanguageFeatures) {
	if s.Features != nil {
		out.FeaturesCount = uint(len(s.Features))
		out.Features = unsafe.SliceData(s.Features)
	}
	return
}

func (s *SupportedWGSLLanguageFeatures) wrap(in *sys.SupportedWGSLLanguageFeatures) {
}

func (s *SurfaceCapabilities) unwrap() (out sys.SurfaceCapabilities) {
	out.Chain = s.Chain
	out.Usages = s.Usages
	if s.Formats != nil {
		out.FormatsCount = uint(len(s.Formats))
		out.Formats = unsafe.SliceData(s.Formats)
	}
	if s.PresentModes != nil {
		out.PresentModesCount = uint(len(s.PresentModes))
		out.PresentModes = unsafe.SliceData(s.PresentModes)
	}
	if s.AlphaModes != nil {
		out.AlphaModesCount = uint(len(s.AlphaModes))
		out.AlphaModes = unsafe.SliceData(s.AlphaModes)
	}
	return
}

func (s *SurfaceCapabilities) wrap(in *sys.SurfaceCapabilities) {
	s.Chain = in.Chain
	s.Usages = in.Usages
}

func (s *SurfaceConfiguration) unwrap() (out sys.SurfaceConfiguration) {
	out.Chain = s.Chain
	if s.Device != nil {
		out.Device = s.Device.inner
	}
	out.Format = s.Format
	out.Usage = s.Usage
	out.Width = s.Width
	out.Height = s.Height
	if s.ViewFormats != nil {
		out.ViewFormatsCount = uint(len(s.ViewFormats))
		out.ViewFormats = unsafe.SliceData(s.ViewFormats)
	}
	out.AlphaMode = s.AlphaMode
	out.PresentMode = s.PresentMode
	return
}

func (s *SurfaceConfiguration) wrap(in *sys.SurfaceConfiguration) {
	s.Chain = in.Chain
	s.Device = new(Device{inner: in.Device}).owned()
	s.Format = in.Format
	s.Usage = in.Usage
	s.Width = in.Width
	s.Height = in.Height
	s.AlphaMode = in.AlphaMode
	s.PresentMode = in.PresentMode
}

func (s *SurfaceDescriptor) unwrap() (out sys.SurfaceDescriptor) {
	out.Chain = s.Chain
	out.Label = sys.StringView{
		Length: uint(len(s.Label)),
		Data:   unsafe.StringData(s.Label),
	}
	return
}

func (s *SurfaceDescriptor) wrap(in *sys.SurfaceDescriptor) {
	s.Chain = in.Chain
	s.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
}

func (s *SurfaceTexture) unwrap() (out sys.SurfaceTexture) {
	out.Chain = s.Chain
	if s.Texture != nil {
		out.Texture = s.Texture.inner
	}
	out.Status = s.Status
	return
}

func (s *SurfaceTexture) wrap(in *sys.SurfaceTexture) {
	s.Chain = in.Chain
	s.Texture = new(Texture{inner: in.Texture}).owned()
	s.Status = in.Status
}

func (t *TexelCopyBufferInfo) unwrap() (out sys.TexelCopyBufferInfo) {
	out.Layout = t.Layout
	if t.Buffer != nil {
		out.Buffer = t.Buffer.inner
	}
	return
}

func (t *TexelCopyBufferInfo) wrap(in *sys.TexelCopyBufferInfo) {
	t.Layout = in.Layout
	t.Buffer = new(Buffer{inner: in.Buffer}).owned()
}

func (t *TexelCopyTextureInfo) unwrap() (out sys.TexelCopyTextureInfo) {
	if t.Texture != nil {
		out.Texture = t.Texture.inner
	}
	out.MipLevel = t.MipLevel
	out.Origin = t.Origin
	out.Aspect = t.Aspect
	return
}

func (t *TexelCopyTextureInfo) wrap(in *sys.TexelCopyTextureInfo) {
	t.Texture = new(Texture{inner: in.Texture}).owned()
	t.MipLevel = in.MipLevel
	t.Origin = in.Origin
	t.Aspect = in.Aspect
}

func (t *TextureDescriptor) unwrap() (out sys.TextureDescriptor) {
	out.Chain = t.Chain
	out.Label = sys.StringView{
		Length: uint(len(t.Label)),
		Data:   unsafe.StringData(t.Label),
	}
	out.Usage = t.Usage
	out.Dimension = t.Dimension
	out.Size = t.Size
	out.Format = t.Format
	out.MipLevelCount = t.MipLevelCount
	out.SampleCount = t.SampleCount
	if t.ViewFormats != nil {
		out.ViewFormatsCount = uint(len(t.ViewFormats))
		out.ViewFormats = unsafe.SliceData(t.ViewFormats)
	}
	return
}

func (t *TextureDescriptor) wrap(in *sys.TextureDescriptor) {
	t.Chain = in.Chain
	t.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	t.Usage = in.Usage
	t.Dimension = in.Dimension
	t.Size = in.Size
	t.Format = in.Format
	t.MipLevelCount = in.MipLevelCount
	t.SampleCount = in.SampleCount
}

func (t *TextureViewDescriptor) unwrap() (out sys.TextureViewDescriptor) {
	out.Chain = t.Chain
	out.Label = sys.StringView{
		Length: uint(len(t.Label)),
		Data:   unsafe.StringData(t.Label),
	}
	out.Format = t.Format
	out.Dimension = t.Dimension
	out.BaseMipLevel = t.BaseMipLevel
	out.MipLevelCount = t.MipLevelCount
	out.BaseArrayLayer = t.BaseArrayLayer
	out.ArrayLayerCount = t.ArrayLayerCount
	out.Aspect = t.Aspect
	out.Usage = t.Usage
	return
}

func (t *TextureViewDescriptor) wrap(in *sys.TextureViewDescriptor) {
	t.Chain = in.Chain
	t.Label = strings.Clone(unsafe.String(
		in.Label.Data,
		in.Label.Length,
	))
	t.Format = in.Format
	t.Dimension = in.Dimension
	t.BaseMipLevel = in.BaseMipLevel
	t.MipLevelCount = in.MipLevelCount
	t.BaseArrayLayer = in.BaseArrayLayer
	t.ArrayLayerCount = in.ArrayLayerCount
	t.Aspect = in.Aspect
	t.Usage = in.Usage
}

func (v *VertexBufferLayout) unwrap() (out sys.VertexBufferLayout) {
	out.Chain = v.Chain
	out.StepMode = v.StepMode
	out.ArrayStride = v.ArrayStride
	if v.Attributes != nil {
		out.AttributesCount = uint(len(v.Attributes))
		out.Attributes = unsafe.SliceData(v.Attributes)
	}
	return
}

func (v *VertexBufferLayout) wrap(in *sys.VertexBufferLayout) {
	v.Chain = in.Chain
	v.StepMode = in.StepMode
	v.ArrayStride = in.ArrayStride
	v.Attributes = unsafe.Slice(in.Attributes, in.AttributesCount)
}

func (v *VertexState) unwrap() (out sys.VertexState) {
	out.Chain = v.Chain
	if v.Module != nil {
		out.Module = v.Module.inner
	}
	out.EntryPoint = sys.StringView{
		Length: uint(len(v.EntryPoint)),
		Data:   unsafe.StringData(v.EntryPoint),
	}
	if v.Constants != nil {
		out.ConstantsCount = uint(len(v.Constants))
		constants := make([]sys.ConstantEntry, len(v.Constants))
		for i := range constants {
			constants[i] = v.Constants[i].unwrap()
		}
		out.Constants = unsafe.SliceData(constants)
	}
	if v.Buffers != nil {
		out.BuffersCount = uint(len(v.Buffers))
		buffers := make([]sys.VertexBufferLayout, len(v.Buffers))
		for i := range buffers {
			buffers[i] = v.Buffers[i].unwrap()
		}
		out.Buffers = unsafe.SliceData(buffers)
	}
	return
}

func (v *VertexState) wrap(in *sys.VertexState) {
	v.Chain = in.Chain
	v.Module = new(ShaderModule{inner: in.Module}).owned()
	v.EntryPoint = strings.Clone(unsafe.String(
		in.EntryPoint.Data,
		in.EntryPoint.Length,
	))
	_constants := unsafe.Slice(in.Constants, in.ConstantsCount)
	constants := make([]ConstantEntry, in.ConstantsCount)
	for i := uint(0); i < in.ConstantsCount; i++ {
		constants[i].wrap(&_constants[i])
	}
	v.Constants = constants

	_buffers := unsafe.Slice(in.Buffers, in.BuffersCount)
	buffers := make([]VertexBufferLayout, in.BuffersCount)
	for i := uint(0); i < in.BuffersCount; i++ {
		buffers[i].wrap(&_buffers[i])
	}
	v.Buffers = buffers

}
