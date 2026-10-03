// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package sys

import (
	"structs"
	"unsafe"
)

type AdapterInfo struct {
	_               structs.HostLayout
	Chain           *ChainedStruct
	Vendor          StringView
	Architecture    StringView
	Device          StringView
	Description     StringView
	BackendType     BackendType
	AdapterType     AdapterType
	VendorID        uint32
	DeviceID        uint32
	SubgroupMinSize uint32
	SubgroupMaxSize uint32
}

type BindGroupDescriptor struct {
	_      structs.HostLayout
	Chain  *ChainedStruct
	Label  StringView
	Layout BindGroupLayout
	// Array count for Entries
	EntriesCount uint
	Entries      *BindGroupEntry
}

type BindGroupEntry struct {
	_     structs.HostLayout
	Chain *ChainedStruct
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
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
	// Array count for Entries
	EntriesCount uint
	Entries      *BindGroupLayoutEntry
}

type BindGroupLayoutEntry struct {
	_          structs.HostLayout
	Chain      *ChainedStruct
	Binding    uint32
	Visibility ShaderStage
	// If non-zero, this entry defines a binding array with this size.
	BindingArraySize uint32
	Buffer           BufferBindingLayout
	Sampler          SamplerBindingLayout
	Texture          TextureBindingLayout
	StorageTexture   StorageTextureBindingLayout
}

type BlendComponent struct {
	_ structs.HostLayout

	// If set to @ref WGPUBlendOperation_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUBlendOperation_Add.
	Operation BlendOperation
	// If set to @ref WGPUBlendFactor_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUBlendFactor_One.
	SrcFactor BlendFactor
	// If set to @ref WGPUBlendFactor_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUBlendFactor_Zero.
	DstFactor BlendFactor
}

type BlendState struct {
	_ structs.HostLayout

	Color BlendComponent
	Alpha BlendComponent
}

type BufferBindingLayout struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// If set to @ref WGPUBufferBindingType_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUBufferBindingType_Uniform.
	Type             BufferBindingType
	HasDynamicOffset Bool
	MinBindingSize   uint64
}

type BufferDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
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
type Color struct {
	_ structs.HostLayout

	R float64
	G float64
	B float64
	A float64
}

type ColorTargetState struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// The texture format of the target. If @ref WGPUTextureFormat_Undefined,
	// indicates a "hole" in the parent @ref WGPUFragmentState `targets` array:
	// the pipeline does not output a value at this `location`.
	Format    TextureFormat
	Blend     *BlendState
	WriteMask ColorWriteMask
}

type CommandBufferDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
}

type CommandEncoderDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
}

// Note: While Compatibility Mode is optional to implement, this extension struct
// is required to be supported (for both queries and requests) and behave as
// defined in the WebGPU spec.
type CompatibilityModeLimits struct {
	_                                 structs.HostLayout
	NextInChain                       ChainedStruct
	MaxStorageBuffersInVertexStage    uint32
	MaxStorageTexturesInVertexStage   uint32
	MaxStorageBuffersInFragmentStage  uint32
	MaxStorageTexturesInFragmentStage uint32
}

type CompilationInfo struct {
	_ structs.HostLayout

	// Array count for Messages
	MessagesCount uint
	Messages      *CompilationMessage
}

type CompilationMessage struct {
	_ structs.HostLayout

	// A @ref LocalizableHumanReadableMessageString.
	Message StringView
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
	_               structs.HostLayout
	Chain           *ChainedStruct
	Label           StringView
	TimestampWrites *PassTimestampWrites
}

type ComputePipelineDescriptor struct {
	_       structs.HostLayout
	Chain   *ChainedStruct
	Label   StringView
	Layout  PipelineLayout
	Compute ComputeState
}

type ComputeState struct {
	_          structs.HostLayout
	Chain      *ChainedStruct
	Module     ShaderModule
	EntryPoint StringView
	// Array count for Constants
	ConstantsCount uint
	Constants      *ConstantEntry
}

type ConstantEntry struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Key   StringView
	// Represents a WGSL numeric or boolean value using @ref DoubleAsSupertype.
	//
	// If non-finite, produces a @ref NonFiniteFloatValueError.
	Value float64
}

type DepthStencilState struct {
	_                 structs.HostLayout
	Chain             *ChainedStruct
	Format            TextureFormat
	DepthWriteEnabled OptionalBool
	DepthCompare      CompareFunction
	StencilFront      StencilFaceState
	StencilBack       StencilFaceState
	StencilReadMask   uint32
	StencilWriteMask  uint32
	DepthBias         int32
	// TODO
	//
	// If non-finite, produces a @ref NonFiniteFloatValueError.
	DepthBiasSlopeScale float32
	// TODO
	//
	// If non-finite, produces a @ref NonFiniteFloatValueError.
	DepthBiasClamp float32
}

type DeviceDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
	// Array count for RequiredFeatures
	RequiredFeaturesCount  uint
	RequiredFeatures       *FeatureName
	RequiredLimits         *Limits
	DefaultQueue           QueueDescriptor
	DeviceLostCallbackInfo DeviceLostCallbackInfo
	// Called when there is an uncaptured error on this device, from any thread.
	// See @ref ErrorScopes.
	//
	// **Important:** This callback does not have a configurable @ref WGPUCallbackMode; it may be called at any time (like @ref WGPUCallbackMode_AllowSpontaneous). As such, calls into the `webgpu.h` API from this callback are unsafe. See @ref CallbackReentrancy.
	UncapturedErrorCallbackInfo UncapturedErrorCallbackInfo
}

type Extent3D struct {
	_ structs.HostLayout

	Width              uint32
	Height             uint32
	DepthOrArrayLayers uint32
}

// Chained in an @ref WGPUBindGroupEntry to set it to an @ref WGPUExternalTexture. This must have a corresponding @ref WGPUExternalTextureBindingLayout in the @ref WGPUBindGroupLayout.
type ExternalTextureBindingEntry struct {
	_               structs.HostLayout
	NextInChain     ChainedStruct
	ExternalTexture ExternalTexture
}

// Chained in @ref WGPUBindGroupLayoutEntry to specify that the corresponding entries in an @ref WGPUBindGroup will contain an @ref WGPUExternalTexture.
type ExternalTextureBindingLayout struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
}

type FragmentState struct {
	_          structs.HostLayout
	Chain      *ChainedStruct
	Module     ShaderModule
	EntryPoint StringView
	// Array count for Constants
	ConstantsCount uint
	Constants      *ConstantEntry
	// Array count for Targets
	TargetsCount uint
	Targets      *ColorTargetState
}

// Opaque handle to an asynchronous operation. See @ref Asynchronous-Operations for more information.
type Future struct {
	_ structs.HostLayout

	// Opaque id of the @ref WGPUFuture
	Id uint64
}

// Struct holding a future to wait on, and a `completed` boolean flag.
type FutureWaitInfo struct {
	_ structs.HostLayout

	// The future to wait on.
	Future Future
	// Whether or not the future completed.
	Completed Bool
}

type InstanceDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// Array count for RequiredFeatures
	RequiredFeaturesCount uint
	RequiredFeatures      *InstanceFeatureName
	RequiredLimits        *InstanceLimits
}

type InstanceLimits struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// The maximum number @ref WGPUFutureWaitInfo supported in a call to ::wgpuInstanceWaitAny with `timeoutNS > 0`.
	TimedWaitAnyMaxCount uintptr
}

type Limits struct {
	_                                         structs.HostLayout
	Chain                                     *ChainedStruct
	MaxTextureDimension1D                     uint32
	MaxTextureDimension2D                     uint32
	MaxTextureDimension3D                     uint32
	MaxTextureArrayLayers                     uint32
	MaxBindGroups                             uint32
	MaxBindGroupsPlusVertexBuffers            uint32
	MaxBindingsPerBindGroup                   uint32
	MaxDynamicUniformBuffersPerPipelineLayout uint32
	MaxDynamicStorageBuffersPerPipelineLayout uint32
	MaxSampledTexturesPerShaderStage          uint32
	MaxSamplersPerShaderStage                 uint32
	MaxStorageBuffersPerShaderStage           uint32
	MaxStorageTexturesPerShaderStage          uint32
	MaxUniformBuffersPerShaderStage           uint32
	MaxUniformBufferBindingSize               uint64
	MaxStorageBufferBindingSize               uint64
	MinUniformBufferOffsetAlignment           uint32
	MinStorageBufferOffsetAlignment           uint32
	MaxVertexBuffers                          uint32
	MaxBufferSize                             uint64
	MaxVertexAttributes                       uint32
	MaxVertexBufferArrayStride                uint32
	MaxInterStageShaderVariables              uint32
	MaxColorAttachments                       uint32
	MaxColorAttachmentBytesPerSample          uint32
	MaxComputeWorkgroupStorageSize            uint32
	MaxComputeInvocationsPerWorkgroup         uint32
	MaxComputeWorkgroupSizeX                  uint32
	MaxComputeWorkgroupSizeY                  uint32
	MaxComputeWorkgroupSizeZ                  uint32
	MaxComputeWorkgroupsPerDimension          uint32
	MaxImmediateSize                          uint32
}

type MultisampleState struct {
	_                      structs.HostLayout
	Chain                  *ChainedStruct
	Count                  uint32
	Mask                   uint32
	AlphaToCoverageEnabled Bool
}

type Origin3D struct {
	_ structs.HostLayout

	X uint32
	Y uint32
	Z uint32
}

type PassTimestampWrites struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// Query set to write timestamps to.
	QuerySet                  QuerySet
	BeginningOfPassWriteIndex uint32
	EndOfPassWriteIndex       uint32
}

type PipelineLayoutDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
	// Array count for BindGroupLayouts
	BindGroupLayoutsCount uint
	BindGroupLayouts      *BindGroupLayout
	ImmediateSize         uint32
}

type PrimitiveState struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// If set to @ref WGPUPrimitiveTopology_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUPrimitiveTopology_TriangleList.
	Topology         PrimitiveTopology
	StripIndexFormat IndexFormat
	// If set to @ref WGPUFrontFace_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUFrontFace_CCW.
	FrontFace FrontFace
	// If set to @ref WGPUCullMode_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUCullMode_None.
	CullMode       CullMode
	UnclippedDepth Bool
}

type QuerySetDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
	Type  QueryType
	Count uint32
}

type QueueDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
}

type RenderBundleDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
}

type RenderBundleEncoderDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
	// Array count for ColorFormats
	ColorFormatsCount  uint
	ColorFormats       *TextureFormat
	DepthStencilFormat TextureFormat
	SampleCount        uint32
	DepthReadOnly      Bool
	StencilReadOnly    Bool
}

type RenderPassColorAttachment struct {
	_     structs.HostLayout
	Chain *ChainedStruct
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
	_            structs.HostLayout
	Chain        *ChainedStruct
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
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
	// Array count for ColorAttachments
	ColorAttachmentsCount  uint
	ColorAttachments       *RenderPassColorAttachment
	DepthStencilAttachment *RenderPassDepthStencilAttachment
	OcclusionQuerySet      QuerySet
	TimestampWrites        *PassTimestampWrites
}

type RenderPassMaxDrawCount struct {
	_            structs.HostLayout
	NextInChain  ChainedStruct
	MaxDrawCount uint64
}

type RenderPipelineDescriptor struct {
	_            structs.HostLayout
	Chain        *ChainedStruct
	Label        StringView
	Layout       PipelineLayout
	Vertex       VertexState
	Primitive    PrimitiveState
	DepthStencil *DepthStencilState
	Multisample  MultisampleState
	Fragment     *FragmentState
}

type RequestAdapterOptions struct {
	_     structs.HostLayout
	Chain *ChainedStruct
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
type RequestAdapterWebXROptions struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	// Sets the `xrCompatible` option in the JS API.
	XrCompatible Bool
}

type SamplerBindingLayout struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// If set to @ref WGPUSamplerBindingType_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUSamplerBindingType_Filtering.
	Type SamplerBindingType
}

type SamplerDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
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
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
}

type ShaderSourceSPIRV struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	CodeSize    uint32
	Code        *uint32
}

type ShaderSourceWGSL struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	Code        StringView
}

type StencilFaceState struct {
	_ structs.HostLayout

	// If set to @ref WGPUCompareFunction_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUCompareFunction_Always.
	Compare CompareFunction
	// If set to @ref WGPUStencilOperation_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUStencilOperation_Keep.
	FailOp StencilOperation
	// If set to @ref WGPUStencilOperation_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUStencilOperation_Keep.
	DepthFailOp StencilOperation
	// If set to @ref WGPUStencilOperation_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUStencilOperation_Keep.
	PassOp StencilOperation
}

type StorageTextureBindingLayout struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// If set to @ref WGPUStorageTextureAccess_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUStorageTextureAccess_WriteOnly.
	Access StorageTextureAccess
	Format TextureFormat
	// If set to @ref WGPUTextureViewDimension_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureViewDimension_2D.
	ViewDimension TextureViewDimension
}

type SupportedFeatures struct {
	_ structs.HostLayout

	// Array count for Features
	FeaturesCount uint
	Features      *FeatureName
}

type SupportedInstanceFeatures struct {
	_ structs.HostLayout

	// Array count for Features
	FeaturesCount uint
	Features      *InstanceFeatureName
}

type SupportedWGSLLanguageFeatures struct {
	_ structs.HostLayout

	// Array count for Features
	FeaturesCount uint
	Features      *WGSLLanguageFeatureName
}

// Filled by @ref wgpuSurfaceGetCapabilities with what's supported for @ref wgpuSurfaceConfigure for a pair of @ref WGPUSurface and @ref WGPUAdapter.
type SurfaceCapabilities struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// The bit set of supported @ref WGPUTextureUsage bits.
	// Guaranteed to contain @ref WGPUTextureUsage_RenderAttachment.
	Usages TextureUsage
	// Array count for Formats
	FormatsCount uint
	// A list of supported @ref WGPUTextureFormat values, in order of preference.
	Formats *TextureFormat
	// Array count for PresentModes
	PresentModesCount uint
	// A list of supported @ref WGPUPresentMode values.
	// Guaranteed to contain @ref WGPUPresentMode_Fifo.
	PresentModes *PresentMode
	// Array count for AlphaModes
	AlphaModesCount uint
	// A list of supported @ref WGPUCompositeAlphaMode values.
	// @ref WGPUCompositeAlphaMode_Auto will be an alias for the first element and will never be present in this array.
	AlphaModes *CompositeAlphaMode
}

// Extension of @ref WGPUSurfaceConfiguration for color spaces and HDR.
type SurfaceColorManagement struct {
	_               structs.HostLayout
	NextInChain     ChainedStruct
	ColorSpace      PredefinedColorSpace
	ToneMappingMode ToneMappingMode
}

// Options to @ref wgpuSurfaceConfigure for defining how a @ref WGPUSurface will be rendered to and presented to the user.
// See @ref Surface-Configuration for more details.
type SurfaceConfiguration struct {
	_     structs.HostLayout
	Chain *ChainedStruct
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
	// Array count for ViewFormats
	ViewFormatsCount uint
	// The additional @ref WGPUTextureFormat for @ref WGPUTextureView format reinterpretation of the surface's textures.
	ViewFormats *TextureFormat
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
	_     structs.HostLayout
	Chain *ChainedStruct
	// Label used to refer to the object.
	Label StringView
}

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping an Android [`ANativeWindow`](https://developer.android.com/ndk/reference/group/a-native-window).
type SurfaceSourceAndroidNativeWindow struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	// The pointer to the [`ANativeWindow`](https://developer.android.com/ndk/reference/group/a-native-window) that will be wrapped by the @ref WGPUSurface.
	Window unsafe.Pointer
}

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping a [`CAMetalLayer`](https://developer.apple.com/documentation/quartzcore/cametallayer?language=objc).
type SurfaceSourceMetalLayer struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	// The pointer to the [`CAMetalLayer`](https://developer.apple.com/documentation/quartzcore/cametallayer?language=objc) that will be wrapped by the @ref WGPUSurface.
	Layer unsafe.Pointer
}

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping a [Wayland](https://wayland.freedesktop.org/) [`wl_surface`](https://wayland.freedesktop.org/docs/html/apa.html#protocol-spec-wl_surface).
type SurfaceSourceWaylandSurface struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	// A [`wl_display`](https://wayland.freedesktop.org/docs/html/apa.html#protocol-spec-wl_display) for this Wayland instance.
	Display unsafe.Pointer
	// A [`wl_surface`](https://wayland.freedesktop.org/docs/html/apa.html#protocol-spec-wl_surface) that will be wrapped by the @ref WGPUSurface
	Surface unsafe.Pointer
}

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping a Windows [`HWND`](https://learn.microsoft.com/en-us/windows/apps/develop/ui-input/retrieve-hwnd).
type SurfaceSourceWindowsHWND struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	// The [`HINSTANCE`](https://learn.microsoft.com/en-us/windows/win32/learnwin32/winmain--the-application-entry-point) for this application.
	// Most commonly `GetModuleHandle(nullptr)`.
	Hinstance unsafe.Pointer
	// The [`HWND`](https://learn.microsoft.com/en-us/windows/apps/develop/ui-input/retrieve-hwnd) that will be wrapped by the @ref WGPUSurface.
	Hwnd unsafe.Pointer
}

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping an [XCB](https://xcb.freedesktop.org/) `xcb_window_t`.
type SurfaceSourceXCBWindow struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	// The `xcb_connection_t` for the connection to the X server.
	Connection unsafe.Pointer
	// The `xcb_window_t` for the window that will be wrapped by the @ref WGPUSurface.
	Window uint32
}

// Chained in @ref WGPUSurfaceDescriptor to make an @ref WGPUSurface wrapping an [Xlib](https://www.x.org/releases/current/doc/libX11/libX11/libX11.html) `Window`.
type SurfaceSourceXlibWindow struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	// A pointer to the [`Display`](https://www.x.org/releases/current/doc/libX11/libX11/libX11.html#Opening_the_Display) connected to the X server.
	Display unsafe.Pointer
	// The [`Window`](https://www.x.org/releases/current/doc/libX11/libX11/libX11.html#Creating_Windows) that will be wrapped by the @ref WGPUSurface.
	Window uint64
}

// Queried each frame from a @ref WGPUSurface to get a @ref WGPUTexture to render to along with some metadata.
// See @ref Surface-Presenting for more details.
type SurfaceTexture struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// The @ref WGPUTexture representing the frame that will be shown on the surface.
	// It is @ref ReturnedWithOwnership from @ref wgpuSurfaceGetCurrentTexture.
	Texture Texture
	// Whether the call to @ref wgpuSurfaceGetCurrentTexture succeeded and a hint as to why it might not have.
	Status SurfaceGetCurrentTextureStatus
}

type TexelCopyBufferInfo struct {
	_ structs.HostLayout

	Layout TexelCopyBufferLayout
	Buffer Buffer
}

type TexelCopyBufferLayout struct {
	_ structs.HostLayout

	Offset       uint64
	BytesPerRow  uint32
	RowsPerImage uint32
}

type TexelCopyTextureInfo struct {
	_ structs.HostLayout

	Texture  Texture
	MipLevel uint32
	Origin   Origin3D
	// If set to @ref WGPUTextureAspect_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureAspect_All.
	Aspect TextureAspect
}

type TextureBindingLayout struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	// If set to @ref WGPUTextureSampleType_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureSampleType_Float.
	SampleType TextureSampleType
	// If set to @ref WGPUTextureViewDimension_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureViewDimension_2D.
	ViewDimension TextureViewDimension
	Multisampled  Bool
}

// Note: While Compatibility Mode is optional to implement, this extension struct
// is required to be accepted (but per the WebGPU spec, its contents are ignored
// on devices that have the @ref WGPUFeatureName_CoreFeaturesAndLimits feature).
type TextureBindingViewDimension struct {
	_                           structs.HostLayout
	NextInChain                 ChainedStruct
	TextureBindingViewDimension TextureViewDimension
}

// When accessed by a shader, the red/green/blue/alpha channels are replaced
// by the value corresponding to the component specified in r, g, b, and a,
// respectively unlike the JS API which uses a string of length four, with
// each character mapping to the texture view's red/green/blue/alpha channels.
type TextureComponentSwizzle struct {
	_ structs.HostLayout

	// The value that replaces the red channel in the shader.
	//
	// If set to @ref WGPUComponentSwizzle_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUComponentSwizzle_R.
	R ComponentSwizzle
	// The value that replaces the green channel in the shader.
	//
	// If set to @ref WGPUComponentSwizzle_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUComponentSwizzle_G.
	G ComponentSwizzle
	// The value that replaces the blue channel in the shader.
	//
	// If set to @ref WGPUComponentSwizzle_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUComponentSwizzle_B.
	B ComponentSwizzle
	// The value that replaces the alpha channel in the shader.
	//
	// If set to @ref WGPUComponentSwizzle_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUComponentSwizzle_A.
	A ComponentSwizzle
}

type TextureComponentSwizzleDescriptor struct {
	_           structs.HostLayout
	NextInChain ChainedStruct
	Swizzle     TextureComponentSwizzle
}

type TextureDescriptor struct {
	_     structs.HostLayout
	Chain *ChainedStruct
	Label StringView
	Usage TextureUsage
	// If set to @ref WGPUTextureDimension_Undefined,
	// [defaults](@ref SentinelValues) to @ref WGPUTextureDimension_2D.
	Dimension     TextureDimension
	Size          Extent3D
	Format        TextureFormat
	MipLevelCount uint32
	SampleCount   uint32
	// Array count for ViewFormats
	ViewFormatsCount uint
	ViewFormats      *TextureFormat
}

type TextureViewDescriptor struct {
	_               structs.HostLayout
	Chain           *ChainedStruct
	Label           StringView
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

type VertexAttribute struct {
	_              structs.HostLayout
	Chain          *ChainedStruct
	Format         VertexFormat
	Offset         uint64
	ShaderLocation uint32
}

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
	_           structs.HostLayout
	Chain       *ChainedStruct
	StepMode    VertexStepMode
	ArrayStride uint64
	// Array count for Attributes
	AttributesCount uint
	Attributes      *VertexAttribute
}

type VertexState struct {
	_          structs.HostLayout
	Chain      *ChainedStruct
	Module     ShaderModule
	EntryPoint StringView
	// Array count for Constants
	ConstantsCount uint
	Constants      *ConstantEntry
	// Array count for Buffers
	BuffersCount uint
	Buffers      *VertexBufferLayout
}
