// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import "github.com/Tnze/go-webgpu/webgpu/sys"

type AdapterInfo struct {
	Vendor          string
	Architecture    string
	Device          string
	Description     string
	BackendType     BackendType
	AdapterType     AdapterType
	VendorID        uint32
	DeviceID        uint32
	SubgroupMinSize uint32
	SubgroupMaxSize uint32
}

type BindGroupDescriptor struct {
	Label   string
	Layout  BindGroupLayout
	Entries []BindGroupEntry
}

type BindGroupEntry struct {
	// Binding index in the bind group.
	Binding uint32
	// Set this if the binding is a buffer object.
	// Otherwise must be null.
	Buffer Buffer
	// If the binding is a buffer, this is the byte offset of the binding range.
	// Otherwise ignored.
	Offset uint64
	// If the binding is a buffer, this is the byte size of the binding range
	// (@ref WGPU_WHOLE_SIZE means the binding ends at the end of the buffer).
	// Otherwise ignored.
	Size uint64
	// Set this if the binding is a sampler object.
	// Otherwise must be null.
	Sampler Sampler
	// Set this if the binding is a texture view object.
	// Otherwise must be null.
	TextureView TextureView
}

type BindGroupLayoutDescriptor struct {
	Label   string
	Entries []BindGroupLayoutEntry
}

type BindGroupLayoutEntry = sys.BindGroupLayoutEntry

type BlendComponent = sys.BlendComponent

type BlendState = sys.BlendState

type BufferBindingLayout = sys.BufferBindingLayout

type BufferDescriptor struct {
	Label string
	Usage BufferUsage
	Size  uint64
	// When true, the buffer is mapped in write mode at creation. It should thus be unmapped once its initial data has been written.
	//
	// @note Mapping at creation does **not** require the usage @ref WGPUBufferUsage_MapWrite.
	MappedAtCreation Bool
}

// An RGBA color. Represents a `f32`, `i32`, or `u32` color using @ref DoubleAsSupertype.
//
// If any channel is non-finite, produces a @ref NonFiniteFloatValueError.
type Color = sys.Color

type ColorTargetState = sys.ColorTargetState

type CommandBufferDescriptor struct {
	Label string
}

type CommandEncoderDescriptor struct {
	Label string
}

// Note: While Compatibility Mode is optional to implement, this extension struct
// is required to be supported (for both queries and requests) and behave as
// defined in the WebGPU spec.
type CompatibilityModeLimits = sys.CompatibilityModeLimits

type CompilationInfo struct {
	Messages []CompilationMessage
}

type CompilationMessage struct {
	// A @ref LocalizableHumanReadableMessageString.
	Message string
	// Severity level of the message.
	Type CompilationMessageType
	// Line number where the message is attached, starting at 1.
	LineNum uint64
	// Offset in UTF-8 code units (bytes) from the beginning of the line, starting at 1.
	LinePos uint64
	// Offset in UTF-8 code units (bytes) from the beginning of the shader code, starting at 0.
	Offset uint64
	// Length in UTF-8 code units (bytes) of the span the message corresponds to.
	Length uint64
}

type ComputePassDescriptor struct {
	Label           string
	TimestampWrites *PassTimestampWrites
}

type ComputePipelineDescriptor struct {
	Label   string
	Layout  PipelineLayout
	Compute ComputeState
}

type ComputeState struct {
	Module     ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
}

type ConstantEntry struct {
	Key string
	// Represents a WGSL numeric or boolean value using @ref DoubleAsSupertype.
	//
	// If non-finite, produces a @ref NonFiniteFloatValueError.
	Value float64
}

type DepthStencilState = sys.DepthStencilState

type DeviceDescriptor struct {
	Label                  string
	RequiredFeatures       []FeatureName
	RequiredLimits         *Limits
	DefaultQueue           QueueDescriptor
	DeviceLostCallbackInfo sys.DeviceLostCallbackInfo
	// Called when there is an uncaptured error on this device, from any thread.
	// See @ref ErrorScopes.
	//
	// **Important:** This callback does not have a configurable @ref WGPUCallbackMode; it may be called at any time (like @ref WGPUCallbackMode_AllowSpontaneous). As such, calls into the `webgpu.h` API from this callback are unsafe. See @ref CallbackReentrancy.
	UncapturedErrorCallbackInfo sys.UncapturedErrorCallbackInfo
}

type Extent3D = sys.Extent3D

// Chained in an @ref WGPUBindGroupEntry to set it to an @ref WGPUExternalTexture. This must have a corresponding @ref WGPUExternalTextureBindingLayout in the @ref WGPUBindGroupLayout.
type ExternalTextureBindingEntry struct {
	ExternalTexture ExternalTexture
}

// Chained in @ref WGPUBindGroupLayoutEntry to specify that the corresponding entries in an @ref WGPUBindGroup will contain an @ref WGPUExternalTexture.
type ExternalTextureBindingLayout = sys.ExternalTextureBindingLayout

type FragmentState struct {
	Module     ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Targets    []ColorTargetState
}

// Opaque handle to an asynchronous operation. See @ref Asynchronous-Operations for more information.
type Future = sys.Future

// Struct holding a future to wait on, and a `completed` boolean flag.
type FutureWaitInfo = sys.FutureWaitInfo

type InstanceDescriptor struct {
	RequiredFeatures []InstanceFeatureName
	RequiredLimits   *InstanceLimits
}

type InstanceLimits = sys.InstanceLimits

type Limits = sys.Limits

type MultisampleState = sys.MultisampleState

type Origin3D = sys.Origin3D

type PassTimestampWrites struct {
	// Query set to write timestamps to.
	QuerySet                  QuerySet
	BeginningOfPassWriteIndex uint32
	EndOfPassWriteIndex       uint32
}

type PipelineLayoutDescriptor struct {
	Label            string
	BindGroupLayouts []BindGroupLayout
	ImmediateSize    uint32
}

type PrimitiveState = sys.PrimitiveState

type QuerySetDescriptor struct {
	Label string
	Type  QueryType
	Count uint32
}

type QueueDescriptor struct {
	Label string
}

type RenderBundleDescriptor struct {
	Label string
}

type RenderBundleEncoderDescriptor struct {
	Label              string
	ColorFormats       []TextureFormat
	DepthStencilFormat TextureFormat
	SampleCount        uint32
	DepthReadOnly      Bool
	StencilReadOnly    Bool
}

type RenderPassColorAttachment struct {
	// If `NULL`, indicates a hole in the parent
	// @ref WGPURenderPassDescriptor::colorAttachments array.
	View          TextureView
	DepthSlice    uint32
	ResolveTarget TextureView
	LoadOp        LoadOp
	StoreOp       StoreOp
	ClearValue    Color
}

type RenderPassDepthStencilAttachment struct {
	View         TextureView
	DepthLoadOp  LoadOp
	DepthStoreOp StoreOp
	// This is a @ref NullableFloatingPointType.
	//
	// If `NaN`, indicates an `undefined` value (as defined by the JS spec).
	// Use @ref WGPU_DEPTH_CLEAR_VALUE_UNDEFINED to indicate this semantically.
	//
	// If infinite, produces a @ref NonFiniteFloatValueError.
	DepthClearValue   float32
	DepthReadOnly     Bool
	StencilLoadOp     LoadOp
	StencilStoreOp    StoreOp
	StencilClearValue uint32
	StencilReadOnly   Bool
}

type RenderPassDescriptor struct {
	Label                  string
	ColorAttachments       []RenderPassColorAttachment
	DepthStencilAttachment *RenderPassDepthStencilAttachment
	OcclusionQuerySet      QuerySet
	TimestampWrites        *PassTimestampWrites
}

type RenderPassMaxDrawCount = sys.RenderPassMaxDrawCount

type RenderPipelineDescriptor struct {
	Label        string
	Layout       PipelineLayout
	Vertex       VertexState
	Primitive    PrimitiveState
	DepthStencil *DepthStencilState
	Multisample  MultisampleState
	Fragment     *FragmentState
}

type RequestAdapterOptions struct {
	// "Feature level" for the adapter request. If an adapter is returned, it must support the features and limits in the requested feature level.
	//
	// If set to @ref WGPUFeatureLevel_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUFeatureLevel_Core.
	// Additionally, implementations may ignore @ref WGPUFeatureLevel_Compatibility
	// and provide @ref WGPUFeatureLevel_Core instead.
	FeatureLevel    FeatureLevel
	PowerPreference PowerPreference
	// If true, requires the adapter to be a "fallback" adapter as defined by the JS spec.
	// If this is not possible, the request returns null.
	ForceFallbackAdapter Bool
	// If set, requires the adapter to have a particular backend type.
	// If this is not possible, the request returns null.
	BackendType BackendType
	// If set, requires the adapter to be able to output to a particular surface.
	// If this is not possible, the request returns null.
	CompatibleSurface Surface
}

// Extension providing requestAdapter options for implementations with WebXR interop (i.e. Wasm).
type RequestAdapterWebXROptions = sys.RequestAdapterWebXROptions

type SamplerBindingLayout = sys.SamplerBindingLayout

type SamplerDescriptor struct {
	Label string
	// If set to @ref WGPUAddressMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUAddressMode_ClampToEdge.
	AddressModeU AddressMode
	// If set to @ref WGPUAddressMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUAddressMode_ClampToEdge.
	AddressModeV AddressMode
	// If set to @ref WGPUAddressMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUAddressMode_ClampToEdge.
	AddressModeW AddressMode
	// If set to @ref WGPUFilterMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUFilterMode_Nearest.
	MagFilter FilterMode
	// If set to @ref WGPUFilterMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUFilterMode_Nearest.
	MinFilter FilterMode
	// If set to @ref WGPUFilterMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUMipmapFilterMode_Nearest.
	MipmapFilter MipmapFilterMode
	// TODO
	//
	// If non-finite, produces a @ref NonFiniteFloatValueError.
	LodMinClamp float32
	// TODO
	//
	// If non-finite, produces a @ref NonFiniteFloatValueError.
	LodMaxClamp   float32
	Compare       CompareFunction
	MaxAnisotropy uint16
}

type ShaderModuleDescriptor struct {
	Label string
}

type ShaderSourceSPIRV = sys.ShaderSourceSPIRV

type ShaderSourceWGSL struct {
	Code string
}

type StencilFaceState = sys.StencilFaceState

type StorageTextureBindingLayout = sys.StorageTextureBindingLayout

type SupportedFeatures struct {
	Features []FeatureName
}

type SupportedInstanceFeatures struct {
	Features []InstanceFeatureName
}

type SupportedWGSLLanguageFeatures struct {
	Features []WGSLLanguageFeatureName
}

// Filled by @ref wgpuSurfaceGetCapabilities with what's supported for @ref wgpuSurfaceConfigure for a pair of @ref WGPUSurface and @ref WGPUAdapter.
type SurfaceCapabilities struct {
	// The bit set of supported @ref WGPUTextureUsage bits.
	// Guaranteed to contain @ref WGPUTextureUsage_RenderAttachment.
	Usages TextureUsage
	// A list of supported @ref WGPUTextureFormat values, in order of preference.
	Formats []TextureFormat
	// A list of supported @ref WGPUPresentMode values.
	// Guaranteed to contain @ref WGPUPresentMode_Fifo.
	PresentModes []PresentMode
	// A list of supported @ref WGPUCompositeAlphaMode values.
	// @ref WGPUCompositeAlphaMode_Auto will be an alias for the first element and will never be present in this array.
	AlphaModes []CompositeAlphaMode
}

// Extension of @ref WGPUSurfaceConfiguration for color spaces and HDR.
type SurfaceColorManagement = sys.SurfaceColorManagement

// Options to @ref wgpuSurfaceConfigure for defining how a @ref WGPUSurface will be rendered to and presented to the user.
// See @ref Surface-Configuration for more details.
type SurfaceConfiguration struct {
	// The @ref WGPUDevice to use to render to surface's textures.
	Device Device
	// The @ref WGPUTextureFormat of the surface's textures.
	Format TextureFormat
	// The @ref WGPUTextureUsage of the surface's textures.
	Usage TextureUsage
	// The width of the surface's textures.
	Width uint32
	// The height of the surface's textures.
	Height uint32
	// The additional @ref WGPUTextureFormat for @ref WGPUTextureView format reinterpretation of the surface's textures.
	ViewFormats []TextureFormat
	// How the surface's frames will be composited on the screen.
	//
	// If set to @ref WGPUCompositeAlphaMode_Auto,
	// [defaults] to @ref WGPUCompositeAlphaMode_Inherit in native (allowing the mode
	// to be configured externally), and to @ref WGPUCompositeAlphaMode_Opaque in Wasm.
	AlphaMode CompositeAlphaMode
	// When and in which order the surface's frames will be shown on the screen.
	//
	// If set to @ref WGPUPresentMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUPresentMode_Fifo.
	PresentMode PresentMode
}

// The root descriptor for the creation of an @ref WGPUSurface with @ref wgpuInstanceCreateSurface.
// It isn't sufficient by itself and must have one of the `WGPUSurfaceSource*` in its chain.
// See @ref Surface-Creation for more details.
type SurfaceDescriptor struct {
	// Label used to refer to the object.
	Label string
}

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping an Android [`ANativeWindow`](https://developer.android.com/ndk/reference/group/a-native-window).
type SurfaceSourceAndroidNativeWindow = sys.SurfaceSourceAndroidNativeWindow

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping a [`CAMetalLayer`](https://developer.apple.com/documentation/quartzcore/cametallayer?language=objc).
type SurfaceSourceMetalLayer = sys.SurfaceSourceMetalLayer

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping a [Wayland](https://wayland.freedesktop.org/) [`wl_surface`](https://wayland.freedesktop.org/docs/html/apa.html#protocol-spec-wl_surface).
type SurfaceSourceWaylandSurface = sys.SurfaceSourceWaylandSurface

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping a Windows [`HWND`](https://learn.microsoft.com/en-us/windows/apps/develop/ui-input/retrieve-hwnd).
type SurfaceSourceWindowsHWND = sys.SurfaceSourceWindowsHWND

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping an [XCB](https://xcb.freedesktop.org/) `xcb_window_t`.
type SurfaceSourceXCBWindow = sys.SurfaceSourceXCBWindow

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping an [Xlib](https://www.x.org/releases/current/doc/libX11/libX11/libX11.html) `Window`.
type SurfaceSourceXlibWindow = sys.SurfaceSourceXlibWindow

// Queried each frame from a @ref WGPUSurface to get a @ref WGPUTexture to render to along with some metadata.
// See @ref Surface-Presenting for more details.
type SurfaceTexture struct {
	// The @ref WGPUTexture representing the frame that will be shown on the surface.
	// It is @ref ReturnedWithOwnership from @ref wgpuSurfaceGetCurrentTexture.
	Texture Texture
	// Whether the call to @ref wgpuSurfaceGetCurrentTexture succeeded and a hint as to why it might not have.
	Status SurfaceGetCurrentTextureStatus
}

type TexelCopyBufferInfo struct {
	Layout TexelCopyBufferLayout
	Buffer Buffer
}

type TexelCopyBufferLayout = sys.TexelCopyBufferLayout

type TexelCopyTextureInfo struct {
	Texture  Texture
	MipLevel uint32
	Origin   Origin3D
	// If set to @ref WGPUTextureAspect_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureAspect_All.
	Aspect TextureAspect
}

type TextureBindingLayout = sys.TextureBindingLayout

// Note: While Compatibility Mode is optional to implement, this extension struct
// is required to be accepted (but per the WebGPU spec, its contents are ignored
// on devices that have the @ref WGPUFeatureName_CoreFeaturesAndLimits feature).
type TextureBindingViewDimension = sys.TextureBindingViewDimension

// When accessed by a shader, the red/green/blue/alpha channels are replaced
// by the value corresponding to the component specified in r, g, b, and a,
// respectively unlike the JS API which uses a string of length four, with
// each character mapping to the texture view's red/green/blue/alpha channels.
type TextureComponentSwizzle = sys.TextureComponentSwizzle

type TextureComponentSwizzleDescriptor = sys.TextureComponentSwizzleDescriptor

type TextureDescriptor struct {
	Label string
	Usage TextureUsage
	// If set to @ref WGPUTextureDimension_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureDimension_2D.
	Dimension     TextureDimension
	Size          Extent3D
	Format        TextureFormat
	MipLevelCount uint32
	SampleCount   uint32
	ViewFormats   []TextureFormat
}

type TextureViewDescriptor struct {
	Label           string
	Format          TextureFormat
	Dimension       TextureViewDimension
	BaseMipLevel    uint32
	MipLevelCount   uint32
	BaseArrayLayer  uint32
	ArrayLayerCount uint32
	// If set to @ref WGPUTextureAspect_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureAspect_All.
	Aspect TextureAspect
	Usage  TextureUsage
}

type VertexAttribute = sys.VertexAttribute

// If `attributes` is empty *and* `stepMode` is @ref WGPUVertexStepMode_Undefined,
// indicates a "hole" in the parent @ref WGPUVertexState `buffers` array,
// with behavior equivalent to `null` in the JS API.
//
// If `attributes` is empty but `stepMode` is *not* @ref WGPUVertexStepMode_Undefined,
// indicates a vertex buffer with no attributes, with behavior equivalent to
// `{ attributes: [] }` in the JS API. (TODO: If the JS API changes not to
// distinguish these cases, then this distinction doesn't matter and we can
// remove this documentation.)
//
// If `stepMode` is @ref WGPUVertexStepMode_Undefined but `attributes` is *not* empty,
// `stepMode` [defaults](@ref SentinelValues) to @ref WGPUVertexStepMode_Vertex.
type VertexBufferLayout struct {
	StepMode    VertexStepMode
	ArrayStride uint64
	Attributes  []VertexAttribute
}

type VertexState struct {
	Module     ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Buffers    []VertexBufferLayout
}
