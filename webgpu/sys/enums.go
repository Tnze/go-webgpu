// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package sys

type AdapterType int32

const (
	AdapterTypeDiscreteGPU AdapterType = 0x00000001

	AdapterTypeIntegratedGPU AdapterType = 0x00000002

	AdapterTypeCPU AdapterType = 0x00000003

	AdapterTypeUnknown AdapterType = 0x00000004
)

type AddressMode int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	AddressModeUndefined AddressMode = 0x00000000

	AddressModeClampToEdge AddressMode = 0x00000001

	AddressModeRepeat AddressMode = 0x00000002

	AddressModeMirrorRepeat AddressMode = 0x00000003
)

type BackendType int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	BackendTypeUndefined BackendType = 0x00000000

	BackendTypeNull BackendType = 0x00000001

	BackendTypeWebGPU BackendType = 0x00000002

	BackendTypeD3D11 BackendType = 0x00000003

	BackendTypeD3D12 BackendType = 0x00000004

	BackendTypeMetal BackendType = 0x00000005

	BackendTypeVulkan BackendType = 0x00000006

	BackendTypeOpenGL BackendType = 0x00000007

	BackendTypeOpenGLES BackendType = 0x00000008
)

type BlendFactor int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	BlendFactorUndefined BlendFactor = 0x00000000

	BlendFactorZero BlendFactor = 0x00000001

	BlendFactorOne BlendFactor = 0x00000002

	BlendFactorSrc BlendFactor = 0x00000003

	BlendFactorOneMinusSrc BlendFactor = 0x00000004

	BlendFactorSrcAlpha BlendFactor = 0x00000005

	BlendFactorOneMinusSrcAlpha BlendFactor = 0x00000006

	BlendFactorDst BlendFactor = 0x00000007

	BlendFactorOneMinusDst BlendFactor = 0x00000008

	BlendFactorDstAlpha BlendFactor = 0x00000009

	BlendFactorOneMinusDstAlpha BlendFactor = 0x0000000A

	BlendFactorSrcAlphaSaturated BlendFactor = 0x0000000B

	BlendFactorConstant BlendFactor = 0x0000000C

	BlendFactorOneMinusConstant BlendFactor = 0x0000000D

	BlendFactorSrc1 BlendFactor = 0x0000000E

	BlendFactorOneMinusSrc1 BlendFactor = 0x0000000F

	BlendFactorSrc1Alpha BlendFactor = 0x00000010

	BlendFactorOneMinusSrc1Alpha BlendFactor = 0x00000011
)

type BlendOperation int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	BlendOperationUndefined BlendOperation = 0x00000000

	BlendOperationAdd BlendOperation = 0x00000001

	BlendOperationSubtract BlendOperation = 0x00000002

	BlendOperationReverseSubtract BlendOperation = 0x00000003

	BlendOperationMin BlendOperation = 0x00000004

	BlendOperationMax BlendOperation = 0x00000005
)

type BufferBindingType int32

const (
	// Indicates that this @ref WGPUBufferBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	BufferBindingTypeBindingNotUsed BufferBindingType = 0x00000000

	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	BufferBindingTypeUndefined BufferBindingType = 0x00000001

	BufferBindingTypeUniform BufferBindingType = 0x00000002

	BufferBindingTypeStorage BufferBindingType = 0x00000003

	BufferBindingTypeReadOnlyStorage BufferBindingType = 0x00000004
)

type BufferMapState int32

const (
	BufferMapStateUnmapped BufferMapState = 0x00000001

	BufferMapStatePending BufferMapState = 0x00000002

	BufferMapStateMapped BufferMapState = 0x00000003
)

// The callback mode controls how a callback for an asynchronous operation may be fired. See @ref Asynchronous-Operations for how these are used.
type CallbackMode int32

const (
	// Callbacks created with `WGPUCallbackMode_WaitAnyOnly`:
	// - fire when the asynchronous operation's future is passed to a call to @ref wgpuInstanceWaitAny
	// AND the operation has already completed or it completes inside the call to @ref wgpuInstanceWaitAny.
	CallbackModeWaitAnyOnly CallbackMode = 0x00000001

	// Callbacks created with `WGPUCallbackMode_AllowProcessEvents`:
	// - fire for the same reasons as callbacks created with `WGPUCallbackMode_WaitAnyOnly`
	// - fire inside a call to @ref wgpuInstanceProcessEvents if the asynchronous operation is complete.
	CallbackModeAllowProcessEvents CallbackMode = 0x00000002

	// Callbacks created with `WGPUCallbackMode_AllowSpontaneous`:
	// - fire for the same reasons as callbacks created with `WGPUCallbackMode_AllowProcessEvents`
	// - **may** fire spontaneously on an arbitrary or application thread, when the WebGPU implementations discovers that the asynchronous operation is complete.
	//
	// Implementations _should_ fire spontaneous callbacks as soon as possible.
	//
	// @note Because spontaneous callbacks may fire at an arbitrary time on an arbitrary thread, applications should take extra care when acquiring locks or mutating state inside the callback. It undefined behavior to re-entrantly call into the webgpu.h API if the callback fires while inside the callstack of another webgpu.h function that is not `wgpuInstanceWaitAny` or `wgpuInstanceProcessEvents`.
	CallbackModeAllowSpontaneous CallbackMode = 0x00000003
)

type CompareFunction int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	CompareFunctionUndefined CompareFunction = 0x00000000

	CompareFunctionNever CompareFunction = 0x00000001

	CompareFunctionLess CompareFunction = 0x00000002

	CompareFunctionEqual CompareFunction = 0x00000003

	CompareFunctionLessEqual CompareFunction = 0x00000004

	CompareFunctionGreater CompareFunction = 0x00000005

	CompareFunctionNotEqual CompareFunction = 0x00000006

	CompareFunctionGreaterEqual CompareFunction = 0x00000007

	CompareFunctionAlways CompareFunction = 0x00000008
)

type CompilationInfoRequestStatus int32

const (
	CompilationInfoRequestStatusSuccess CompilationInfoRequestStatus = 0x00000001

	// See @ref CallbackStatuses.
	CompilationInfoRequestStatusCallbackCancelled CompilationInfoRequestStatus = 0x00000002
)

type CompilationMessageType int32

const (
	CompilationMessageTypeError CompilationMessageType = 0x00000001

	CompilationMessageTypeWarning CompilationMessageType = 0x00000002

	CompilationMessageTypeInfo CompilationMessageType = 0x00000003
)

type ComponentSwizzle int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	ComponentSwizzleUndefined ComponentSwizzle = 0x00000000

	// Force its value to 0.
	ComponentSwizzleZero ComponentSwizzle = 0x00000001

	// Force its value to 1.
	ComponentSwizzleOne ComponentSwizzle = 0x00000002

	// Take its value from the red channel of the texture.
	ComponentSwizzleR ComponentSwizzle = 0x00000003

	// Take its value from the green channel of the texture.
	ComponentSwizzleG ComponentSwizzle = 0x00000004

	// Take its value from the blue channel of the texture.
	ComponentSwizzleB ComponentSwizzle = 0x00000005

	// Take its value from the alpha channel of the texture.
	ComponentSwizzleA ComponentSwizzle = 0x00000006
)

// Describes how frames are composited with other contents on the screen when @ref wgpuSurfacePresent is called.
type CompositeAlphaMode int32

const (
	// Lets the WebGPU implementation choose the best mode (supported, and with the best performance) between @ref WGPUCompositeAlphaMode_Opaque or @ref WGPUCompositeAlphaMode_Inherit.
	CompositeAlphaModeAuto CompositeAlphaMode = 0x00000000

	// The alpha component of the image is ignored and teated as if it is always 1.0.
	CompositeAlphaModeOpaque CompositeAlphaMode = 0x00000001

	// The alpha component is respected and non-alpha components are assumed to be already multiplied with the alpha component. For example, (0.5, 0, 0, 0.5) is semi-transparent bright red.
	CompositeAlphaModePremultiplied CompositeAlphaMode = 0x00000002

	// The alpha component is respected and non-alpha components are assumed to NOT be already multiplied with the alpha component. For example, (1.0, 0, 0, 0.5) is semi-transparent bright red.
	CompositeAlphaModeUnpremultiplied CompositeAlphaMode = 0x00000003

	// The handling of the alpha component is unknown to WebGPU and should be handled by the application using system-specific APIs. This mode may be unavailable (for example on Wasm).
	CompositeAlphaModeInherit CompositeAlphaMode = 0x00000004
)

type CreatePipelineAsyncStatus int32

const (
	CreatePipelineAsyncStatusSuccess CreatePipelineAsyncStatus = 0x00000001

	// See @ref CallbackStatuses.
	CreatePipelineAsyncStatusCallbackCancelled CreatePipelineAsyncStatus = 0x00000002

	CreatePipelineAsyncStatusValidationError CreatePipelineAsyncStatus = 0x00000003

	CreatePipelineAsyncStatusInternalError CreatePipelineAsyncStatus = 0x00000004
)

type CullMode int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	CullModeUndefined CullMode = 0x00000000

	CullModeNone CullMode = 0x00000001

	CullModeFront CullMode = 0x00000002

	CullModeBack CullMode = 0x00000003
)

type DeviceLostReason int32

const (
	DeviceLostReasonUnknown DeviceLostReason = 0x00000001

	DeviceLostReasonDestroyed DeviceLostReason = 0x00000002

	// See @ref CallbackStatuses.
	DeviceLostReasonCallbackCancelled DeviceLostReason = 0x00000003

	DeviceLostReasonFailedCreation DeviceLostReason = 0x00000004
)

type ErrorFilter int32

const (
	ErrorFilterValidation ErrorFilter = 0x00000001

	ErrorFilterOutOfMemory ErrorFilter = 0x00000002

	ErrorFilterInternal ErrorFilter = 0x00000003
)

type ErrorType int32

const (
	ErrorTypeNoError ErrorType = 0x00000001

	ErrorTypeValidation ErrorType = 0x00000002

	ErrorTypeOutOfMemory ErrorType = 0x00000003

	ErrorTypeInternal ErrorType = 0x00000004

	ErrorTypeUnknown ErrorType = 0x00000005
)

// See @ref WGPURequestAdapterOptions::featureLevel.
type FeatureLevel int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	FeatureLevelUndefined FeatureLevel = 0x00000000

	// "Compatibility" profile which can be supported on OpenGL ES 3.1 and D3D11.
	FeatureLevelCompatibility FeatureLevel = 0x00000001

	// "Core" profile which can be supported on Vulkan/Metal/D3D12 (at least).
	FeatureLevelCore FeatureLevel = 0x00000002
)

type FeatureName int32

const (
	FeatureNameCoreFeaturesAndLimits FeatureName = 0x00000001

	FeatureNameDepthClipControl FeatureName = 0x00000002

	FeatureNameDepth32FloatStencil8 FeatureName = 0x00000003

	FeatureNameTextureCompressionBC FeatureName = 0x00000004

	FeatureNameTextureCompressionBCSliced3D FeatureName = 0x00000005

	FeatureNameTextureCompressionETC2 FeatureName = 0x00000006

	FeatureNameTextureCompressionASTC FeatureName = 0x00000007

	FeatureNameTextureCompressionASTCSliced3D FeatureName = 0x00000008

	FeatureNameTimestampQuery FeatureName = 0x00000009

	FeatureNameIndirectFirstInstance FeatureName = 0x0000000A

	FeatureNameShaderF16 FeatureName = 0x0000000B

	FeatureNameRG11B10UfloatRenderable FeatureName = 0x0000000C

	FeatureNameBGRA8UnormStorage FeatureName = 0x0000000D

	FeatureNameFloat32Filterable FeatureName = 0x0000000E

	FeatureNameFloat32Blendable FeatureName = 0x0000000F

	FeatureNameClipDistances FeatureName = 0x00000010

	FeatureNameDualSourceBlending FeatureName = 0x00000011

	FeatureNameSubgroups FeatureName = 0x00000012

	FeatureNameTextureFormatsTier1 FeatureName = 0x00000013

	FeatureNameTextureFormatsTier2 FeatureName = 0x00000014

	FeatureNamePrimitiveIndex FeatureName = 0x00000015

	FeatureNameTextureComponentSwizzle FeatureName = 0x00000016

	FeatureNameSubgroupSizeControl FeatureName = 0x00000017

	FeatureNameTextureCompressionUnaligned FeatureName = 0x00000018
)

type FilterMode int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	FilterModeUndefined FilterMode = 0x00000000

	FilterModeNearest FilterMode = 0x00000001

	FilterModeLinear FilterMode = 0x00000002
)

type FrontFace int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	FrontFaceUndefined FrontFace = 0x00000000

	FrontFaceCCW FrontFace = 0x00000001

	FrontFaceCW FrontFace = 0x00000002
)

type IndexFormat int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	IndexFormatUndefined IndexFormat = 0x00000000

	IndexFormatUint16 IndexFormat = 0x00000001

	IndexFormatUint32 IndexFormat = 0x00000002
)

type InstanceFeatureName int32

const (
	// Enable use of ::wgpuInstanceWaitAny with `timeoutNS > 0`.
	InstanceFeatureNameTimedWaitAny InstanceFeatureName = 0x00000001

	// Enable passing SPIR-V shaders to @ref wgpuDeviceCreateShaderModule,
	// via @ref WGPUShaderSourceSPIRV.
	InstanceFeatureNameShaderSourceSPIRV InstanceFeatureName = 0x00000002

	// Normally, a @ref WGPUAdapter can only create a single device. If this is
	// available and enabled, then adapters won't immediately expire when they
	// create a device, so can be reused to make multiple devices. They may
	// still expire for other reasons.
	InstanceFeatureNameMultipleDevicesPerAdapter InstanceFeatureName = 0x00000003
)

type LoadOp int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	LoadOpUndefined LoadOp = 0x00000000

	LoadOpLoad LoadOp = 0x00000001

	LoadOpClear LoadOp = 0x00000002
)

type MapAsyncStatus int32

const (
	MapAsyncStatusSuccess MapAsyncStatus = 0x00000001

	// See @ref CallbackStatuses.
	MapAsyncStatusCallbackCancelled MapAsyncStatus = 0x00000002

	MapAsyncStatusError MapAsyncStatus = 0x00000003

	MapAsyncStatusAborted MapAsyncStatus = 0x00000004
)

type MipmapFilterMode int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	MipmapFilterModeUndefined MipmapFilterMode = 0x00000000

	MipmapFilterModeNearest MipmapFilterMode = 0x00000001

	MipmapFilterModeLinear MipmapFilterMode = 0x00000002
)

type OptionalBool int32

const (
	OptionalBoolFalse OptionalBool = 0x00000000

	OptionalBoolTrue OptionalBool = 0x00000001

	OptionalBoolUndefined OptionalBool = 0x00000002
)

type PopErrorScopeStatus int32

const (
	// The error scope stack was successfully popped and a result was reported.
	PopErrorScopeStatusSuccess PopErrorScopeStatus = 0x00000001

	// See @ref CallbackStatuses.
	PopErrorScopeStatusCallbackCancelled PopErrorScopeStatus = 0x00000002

	// The error scope stack could not be popped, because it was empty.
	PopErrorScopeStatusError PopErrorScopeStatus = 0x00000003
)

type PowerPreference int32

const (
	// No preference. (See also @ref SentinelValues.)
	PowerPreferenceUndefined PowerPreference = 0x00000000

	PowerPreferenceLowPower PowerPreference = 0x00000001

	PowerPreferenceHighPerformance PowerPreference = 0x00000002
)

type PredefinedColorSpace int32

const (
	PredefinedColorSpaceSRGB PredefinedColorSpace = 0x00000001

	PredefinedColorSpaceDisplayP3 PredefinedColorSpace = 0x00000002
)

// Describes when and in which order frames are presented on the screen when @ref wgpuSurfacePresent is called.
type PresentMode int32

const (
	// Present mode is not specified. Use the default.
	PresentModeUndefined PresentMode = 0x00000000

	// The presentation of the image to the user waits for the next vertical blanking period to update in a first-in, first-out manner.
	// Tearing cannot be observed and frame-loop will be limited to the display's refresh rate.
	// This is the only mode that's always available.
	PresentModeFifo PresentMode = 0x00000001

	// The presentation of the image to the user tries to wait for the next vertical blanking period but may decide to not wait if a frame is presented late.
	// Tearing can sometimes be observed but late-frame don't produce a full-frame stutter in the presentation.
	// This is still a first-in, first-out mechanism so a frame-loop will be limited to the display's refresh rate.
	PresentModeFifoRelaxed PresentMode = 0x00000002

	// The presentation of the image to the user is updated immediately without waiting for a vertical blank.
	// Tearing can be observed but latency is minimized.
	PresentModeImmediate PresentMode = 0x00000003

	// The presentation of the image to the user waits for the next vertical blanking period to update to the latest provided image.
	// Tearing cannot be observed and a frame-loop is not limited to the display's refresh rate.
	PresentModeMailbox PresentMode = 0x00000004
)

type PrimitiveTopology int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	PrimitiveTopologyUndefined PrimitiveTopology = 0x00000000

	PrimitiveTopologyPointList PrimitiveTopology = 0x00000001

	PrimitiveTopologyLineList PrimitiveTopology = 0x00000002

	PrimitiveTopologyLineStrip PrimitiveTopology = 0x00000003

	PrimitiveTopologyTriangleList PrimitiveTopology = 0x00000004

	PrimitiveTopologyTriangleStrip PrimitiveTopology = 0x00000005
)

type QueryType int32

const (
	QueryTypeOcclusion QueryType = 0x00000001

	QueryTypeTimestamp QueryType = 0x00000002
)

type QueueWorkDoneStatus int32

const (
	QueueWorkDoneStatusSuccess QueueWorkDoneStatus = 0x00000001

	// See @ref CallbackStatuses.
	QueueWorkDoneStatusCallbackCancelled QueueWorkDoneStatus = 0x00000002

	// There was some deterministic error. (Note this is currently never used,
	// but it will be relevant when it's possible to create a queue object.)
	QueueWorkDoneStatusError QueueWorkDoneStatus = 0x00000003
)

type RequestAdapterStatus int32

const (
	RequestAdapterStatusSuccess RequestAdapterStatus = 0x00000001

	// See @ref CallbackStatuses.
	RequestAdapterStatusCallbackCancelled RequestAdapterStatus = 0x00000002

	RequestAdapterStatusUnavailable RequestAdapterStatus = 0x00000003

	RequestAdapterStatusError RequestAdapterStatus = 0x00000004
)

type RequestDeviceStatus int32

const (
	RequestDeviceStatusSuccess RequestDeviceStatus = 0x00000001

	// See @ref CallbackStatuses.
	RequestDeviceStatusCallbackCancelled RequestDeviceStatus = 0x00000002

	RequestDeviceStatusError RequestDeviceStatus = 0x00000003
)

type SType int32

const (
	STypeShaderSourceSPIRV SType = 0x00000001

	STypeShaderSourceWGSL SType = 0x00000002

	STypeRenderPassMaxDrawCount SType = 0x00000003

	STypeSurfaceSourceMetalLayer SType = 0x00000004

	STypeSurfaceSourceWindowsHWND SType = 0x00000005

	STypeSurfaceSourceXlibWindow SType = 0x00000006

	STypeSurfaceSourceWaylandSurface SType = 0x00000007

	STypeSurfaceSourceAndroidNativeWindow SType = 0x00000008

	STypeSurfaceSourceXCBWindow SType = 0x00000009

	STypeSurfaceColorManagement SType = 0x0000000A

	STypeRequestAdapterWebXROptions SType = 0x0000000B

	STypeTextureComponentSwizzleDescriptor SType = 0x0000000C

	STypeExternalTextureBindingLayout SType = 0x0000000D

	STypeExternalTextureBindingEntry SType = 0x0000000E

	STypeCompatibilityModeLimits SType = 0x0000000F

	STypeTextureBindingViewDimension SType = 0x00000010
)

type SamplerBindingType int32

const (
	// Indicates that this @ref WGPUSamplerBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	SamplerBindingTypeBindingNotUsed SamplerBindingType = 0x00000000

	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	SamplerBindingTypeUndefined SamplerBindingType = 0x00000001

	SamplerBindingTypeFiltering SamplerBindingType = 0x00000002

	SamplerBindingTypeNonFiltering SamplerBindingType = 0x00000003

	SamplerBindingTypeComparison SamplerBindingType = 0x00000004
)

// Status code returned (synchronously) from many operations. Generally
// indicates an invalid input like an unknown enum value or @ref OutStructChainError.
// Read the function's documentation for specific error conditions.
type Status int32

const (
	StatusSuccess Status = 0x00000001

	StatusError Status = 0x00000002
)

type StencilOperation int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	StencilOperationUndefined StencilOperation = 0x00000000

	StencilOperationKeep StencilOperation = 0x00000001

	StencilOperationZero StencilOperation = 0x00000002

	StencilOperationReplace StencilOperation = 0x00000003

	StencilOperationInvert StencilOperation = 0x00000004

	StencilOperationIncrementClamp StencilOperation = 0x00000005

	StencilOperationDecrementClamp StencilOperation = 0x00000006

	StencilOperationIncrementWrap StencilOperation = 0x00000007

	StencilOperationDecrementWrap StencilOperation = 0x00000008
)

type StorageTextureAccess int32

const (
	// Indicates that this @ref WGPUStorageTextureBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	StorageTextureAccessBindingNotUsed StorageTextureAccess = 0x00000000

	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	StorageTextureAccessUndefined StorageTextureAccess = 0x00000001

	StorageTextureAccessWriteOnly StorageTextureAccess = 0x00000002

	StorageTextureAccessReadOnly StorageTextureAccess = 0x00000003

	StorageTextureAccessReadWrite StorageTextureAccess = 0x00000004
)

type StoreOp int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	StoreOpUndefined StoreOp = 0x00000000

	StoreOpStore StoreOp = 0x00000001

	StoreOpDiscard StoreOp = 0x00000002
)

// The status enum for @ref wgpuSurfaceGetCurrentTexture.
type SurfaceGetCurrentTextureStatus int32

const (
	// Yay! Everything is good and we can render this frame.
	SurfaceGetCurrentTextureStatusSuccessOptimal SurfaceGetCurrentTextureStatus = 0x00000001

	// Still OK - the surface can present the frame, but in a suboptimal way. The surface may need reconfiguration.
	SurfaceGetCurrentTextureStatusSuccessSuboptimal SurfaceGetCurrentTextureStatus = 0x00000002

	// Some operation timed out while trying to acquire the frame.
	SurfaceGetCurrentTextureStatusTimeout SurfaceGetCurrentTextureStatus = 0x00000003

	// The surface is too different to be used, compared to when it was originally created.
	SurfaceGetCurrentTextureStatusOutdated SurfaceGetCurrentTextureStatus = 0x00000004

	// The connection to whatever owns the surface was lost, or generally needs to be fully reinitialized.
	SurfaceGetCurrentTextureStatusLost SurfaceGetCurrentTextureStatus = 0x00000005

	// There was some deterministic error (for example, the surface is not configured, or there was an @ref OutStructChainError). Should produce @ref ImplementationDefinedLogging containing details.
	SurfaceGetCurrentTextureStatusError SurfaceGetCurrentTextureStatus = 0x00000006
)

type TextureAspect int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureAspectUndefined TextureAspect = 0x00000000

	TextureAspectAll TextureAspect = 0x00000001

	TextureAspectStencilOnly TextureAspect = 0x00000002

	TextureAspectDepthOnly TextureAspect = 0x00000003
)

type TextureDimension int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureDimensionUndefined TextureDimension = 0x00000000

	TextureDimension1D TextureDimension = 0x00000001

	TextureDimension2D TextureDimension = 0x00000002

	TextureDimension3D TextureDimension = 0x00000003
)

type TextureFormat int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureFormatUndefined TextureFormat = 0x00000000

	TextureFormatR8Unorm TextureFormat = 0x00000001

	TextureFormatR8Snorm TextureFormat = 0x00000002

	TextureFormatR8Uint TextureFormat = 0x00000003

	TextureFormatR8Sint TextureFormat = 0x00000004

	TextureFormatR16Unorm TextureFormat = 0x00000005

	TextureFormatR16Snorm TextureFormat = 0x00000006

	TextureFormatR16Uint TextureFormat = 0x00000007

	TextureFormatR16Sint TextureFormat = 0x00000008

	TextureFormatR16Float TextureFormat = 0x00000009

	TextureFormatRG8Unorm TextureFormat = 0x0000000A

	TextureFormatRG8Snorm TextureFormat = 0x0000000B

	TextureFormatRG8Uint TextureFormat = 0x0000000C

	TextureFormatRG8Sint TextureFormat = 0x0000000D

	TextureFormatR32Float TextureFormat = 0x0000000E

	TextureFormatR32Uint TextureFormat = 0x0000000F

	TextureFormatR32Sint TextureFormat = 0x00000010

	TextureFormatRG16Unorm TextureFormat = 0x00000011

	TextureFormatRG16Snorm TextureFormat = 0x00000012

	TextureFormatRG16Uint TextureFormat = 0x00000013

	TextureFormatRG16Sint TextureFormat = 0x00000014

	TextureFormatRG16Float TextureFormat = 0x00000015

	TextureFormatRGBA8Unorm TextureFormat = 0x00000016

	TextureFormatRGBA8UnormSrgb TextureFormat = 0x00000017

	TextureFormatRGBA8Snorm TextureFormat = 0x00000018

	TextureFormatRGBA8Uint TextureFormat = 0x00000019

	TextureFormatRGBA8Sint TextureFormat = 0x0000001A

	TextureFormatBGRA8Unorm TextureFormat = 0x0000001B

	TextureFormatBGRA8UnormSrgb TextureFormat = 0x0000001C

	TextureFormatRGB10A2Uint TextureFormat = 0x0000001D

	TextureFormatRGB10A2Unorm TextureFormat = 0x0000001E

	TextureFormatRG11B10Ufloat TextureFormat = 0x0000001F

	TextureFormatRGB9E5Ufloat TextureFormat = 0x00000020

	TextureFormatRG32Float TextureFormat = 0x00000021

	TextureFormatRG32Uint TextureFormat = 0x00000022

	TextureFormatRG32Sint TextureFormat = 0x00000023

	TextureFormatRGBA16Unorm TextureFormat = 0x00000024

	TextureFormatRGBA16Snorm TextureFormat = 0x00000025

	TextureFormatRGBA16Uint TextureFormat = 0x00000026

	TextureFormatRGBA16Sint TextureFormat = 0x00000027

	TextureFormatRGBA16Float TextureFormat = 0x00000028

	TextureFormatRGBA32Float TextureFormat = 0x00000029

	TextureFormatRGBA32Uint TextureFormat = 0x0000002A

	TextureFormatRGBA32Sint TextureFormat = 0x0000002B

	TextureFormatStencil8 TextureFormat = 0x0000002C

	TextureFormatDepth16Unorm TextureFormat = 0x0000002D

	TextureFormatDepth24Plus TextureFormat = 0x0000002E

	TextureFormatDepth24PlusStencil8 TextureFormat = 0x0000002F

	TextureFormatDepth32Float TextureFormat = 0x00000030

	TextureFormatDepth32FloatStencil8 TextureFormat = 0x00000031

	TextureFormatBC1RGBAUnorm TextureFormat = 0x00000032

	TextureFormatBC1RGBAUnormSrgb TextureFormat = 0x00000033

	TextureFormatBC2RGBAUnorm TextureFormat = 0x00000034

	TextureFormatBC2RGBAUnormSrgb TextureFormat = 0x00000035

	TextureFormatBC3RGBAUnorm TextureFormat = 0x00000036

	TextureFormatBC3RGBAUnormSrgb TextureFormat = 0x00000037

	TextureFormatBC4RUnorm TextureFormat = 0x00000038

	TextureFormatBC4RSnorm TextureFormat = 0x00000039

	TextureFormatBC5RGUnorm TextureFormat = 0x0000003A

	TextureFormatBC5RGSnorm TextureFormat = 0x0000003B

	TextureFormatBC6HRGBUfloat TextureFormat = 0x0000003C

	TextureFormatBC6HRGBFloat TextureFormat = 0x0000003D

	TextureFormatBC7RGBAUnorm TextureFormat = 0x0000003E

	TextureFormatBC7RGBAUnormSrgb TextureFormat = 0x0000003F

	TextureFormatETC2RGB8Unorm TextureFormat = 0x00000040

	TextureFormatETC2RGB8UnormSrgb TextureFormat = 0x00000041

	TextureFormatETC2RGB8A1Unorm TextureFormat = 0x00000042

	TextureFormatETC2RGB8A1UnormSrgb TextureFormat = 0x00000043

	TextureFormatETC2RGBA8Unorm TextureFormat = 0x00000044

	TextureFormatETC2RGBA8UnormSrgb TextureFormat = 0x00000045

	TextureFormatEACR11Unorm TextureFormat = 0x00000046

	TextureFormatEACR11Snorm TextureFormat = 0x00000047

	TextureFormatEACRG11Unorm TextureFormat = 0x00000048

	TextureFormatEACRG11Snorm TextureFormat = 0x00000049

	TextureFormatASTC4x4Unorm TextureFormat = 0x0000004A

	TextureFormatASTC4x4UnormSrgb TextureFormat = 0x0000004B

	TextureFormatASTC5x4Unorm TextureFormat = 0x0000004C

	TextureFormatASTC5x4UnormSrgb TextureFormat = 0x0000004D

	TextureFormatASTC5x5Unorm TextureFormat = 0x0000004E

	TextureFormatASTC5x5UnormSrgb TextureFormat = 0x0000004F

	TextureFormatASTC6x5Unorm TextureFormat = 0x00000050

	TextureFormatASTC6x5UnormSrgb TextureFormat = 0x00000051

	TextureFormatASTC6x6Unorm TextureFormat = 0x00000052

	TextureFormatASTC6x6UnormSrgb TextureFormat = 0x00000053

	TextureFormatASTC8x5Unorm TextureFormat = 0x00000054

	TextureFormatASTC8x5UnormSrgb TextureFormat = 0x00000055

	TextureFormatASTC8x6Unorm TextureFormat = 0x00000056

	TextureFormatASTC8x6UnormSrgb TextureFormat = 0x00000057

	TextureFormatASTC8x8Unorm TextureFormat = 0x00000058

	TextureFormatASTC8x8UnormSrgb TextureFormat = 0x00000059

	TextureFormatASTC10x5Unorm TextureFormat = 0x0000005A

	TextureFormatASTC10x5UnormSrgb TextureFormat = 0x0000005B

	TextureFormatASTC10x6Unorm TextureFormat = 0x0000005C

	TextureFormatASTC10x6UnormSrgb TextureFormat = 0x0000005D

	TextureFormatASTC10x8Unorm TextureFormat = 0x0000005E

	TextureFormatASTC10x8UnormSrgb TextureFormat = 0x0000005F

	TextureFormatASTC10x10Unorm TextureFormat = 0x00000060

	TextureFormatASTC10x10UnormSrgb TextureFormat = 0x00000061

	TextureFormatASTC12x10Unorm TextureFormat = 0x00000062

	TextureFormatASTC12x10UnormSrgb TextureFormat = 0x00000063

	TextureFormatASTC12x12Unorm TextureFormat = 0x00000064

	TextureFormatASTC12x12UnormSrgb TextureFormat = 0x00000065
)

type TextureSampleType int32

const (
	// Indicates that this @ref WGPUTextureBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	TextureSampleTypeBindingNotUsed TextureSampleType = 0x00000000

	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureSampleTypeUndefined TextureSampleType = 0x00000001

	TextureSampleTypeFloat TextureSampleType = 0x00000002

	TextureSampleTypeUnfilterableFloat TextureSampleType = 0x00000003

	TextureSampleTypeDepth TextureSampleType = 0x00000004

	TextureSampleTypeSint TextureSampleType = 0x00000005

	TextureSampleTypeUint TextureSampleType = 0x00000006
)

type TextureViewDimension int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureViewDimensionUndefined TextureViewDimension = 0x00000000

	TextureViewDimension1D TextureViewDimension = 0x00000001

	TextureViewDimension2D TextureViewDimension = 0x00000002

	TextureViewDimension2DArray TextureViewDimension = 0x00000003

	TextureViewDimensionCube TextureViewDimension = 0x00000004

	TextureViewDimensionCubeArray TextureViewDimension = 0x00000005

	TextureViewDimension3D TextureViewDimension = 0x00000006
)

type ToneMappingMode int32

const (
	ToneMappingModeStandard ToneMappingMode = 0x00000001

	ToneMappingModeExtended ToneMappingMode = 0x00000002
)

type VertexFormat int32

const (
	VertexFormatUint8 VertexFormat = 0x00000001

	VertexFormatUint8x2 VertexFormat = 0x00000002

	VertexFormatUint8x4 VertexFormat = 0x00000003

	VertexFormatSint8 VertexFormat = 0x00000004

	VertexFormatSint8x2 VertexFormat = 0x00000005

	VertexFormatSint8x4 VertexFormat = 0x00000006

	VertexFormatUnorm8 VertexFormat = 0x00000007

	VertexFormatUnorm8x2 VertexFormat = 0x00000008

	VertexFormatUnorm8x4 VertexFormat = 0x00000009

	VertexFormatSnorm8 VertexFormat = 0x0000000A

	VertexFormatSnorm8x2 VertexFormat = 0x0000000B

	VertexFormatSnorm8x4 VertexFormat = 0x0000000C

	VertexFormatUint16 VertexFormat = 0x0000000D

	VertexFormatUint16x2 VertexFormat = 0x0000000E

	VertexFormatUint16x4 VertexFormat = 0x0000000F

	VertexFormatSint16 VertexFormat = 0x00000010

	VertexFormatSint16x2 VertexFormat = 0x00000011

	VertexFormatSint16x4 VertexFormat = 0x00000012

	VertexFormatUnorm16 VertexFormat = 0x00000013

	VertexFormatUnorm16x2 VertexFormat = 0x00000014

	VertexFormatUnorm16x4 VertexFormat = 0x00000015

	VertexFormatSnorm16 VertexFormat = 0x00000016

	VertexFormatSnorm16x2 VertexFormat = 0x00000017

	VertexFormatSnorm16x4 VertexFormat = 0x00000018

	VertexFormatFloat16 VertexFormat = 0x00000019

	VertexFormatFloat16x2 VertexFormat = 0x0000001A

	VertexFormatFloat16x4 VertexFormat = 0x0000001B

	VertexFormatFloat32 VertexFormat = 0x0000001C

	VertexFormatFloat32x2 VertexFormat = 0x0000001D

	VertexFormatFloat32x3 VertexFormat = 0x0000001E

	VertexFormatFloat32x4 VertexFormat = 0x0000001F

	VertexFormatUint32 VertexFormat = 0x00000020

	VertexFormatUint32x2 VertexFormat = 0x00000021

	VertexFormatUint32x3 VertexFormat = 0x00000022

	VertexFormatUint32x4 VertexFormat = 0x00000023

	VertexFormatSint32 VertexFormat = 0x00000024

	VertexFormatSint32x2 VertexFormat = 0x00000025

	VertexFormatSint32x3 VertexFormat = 0x00000026

	VertexFormatSint32x4 VertexFormat = 0x00000027

	VertexFormatUnorm10_10_10_2 VertexFormat = 0x00000028

	VertexFormatUnorm8x4BGRA VertexFormat = 0x00000029

	VertexFormatSnorm10_10_10_2 VertexFormat = 0x0000002A
)

type VertexStepMode int32

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	VertexStepModeUndefined VertexStepMode = 0x00000000

	VertexStepModeVertex VertexStepMode = 0x00000001

	VertexStepModeInstance VertexStepMode = 0x00000002
)

// Status returned from a call to ::wgpuInstanceWaitAny.
type WaitStatus int32

const (
	// At least one WGPUFuture completed successfully.
	WaitStatusSuccess WaitStatus = 0x00000001

	// The wait operation succeeded, but no WGPUFutures completed within the timeout.
	WaitStatusTimedOut WaitStatus = 0x00000002

	// The call was invalid for some reason (see @ref Wait-Any).
	// Should produce @ref ImplementationDefinedLogging containing details.
	WaitStatusError WaitStatus = 0x00000003
)

type WGSLLanguageFeatureName int32

const (
	WGSLLanguageFeatureNameReadonlyAndReadwriteStorageTextures WGSLLanguageFeatureName = 0x00000001

	WGSLLanguageFeatureNamePacked4x8IntegerDotProduct WGSLLanguageFeatureName = 0x00000002

	WGSLLanguageFeatureNameUnrestrictedPointerParameters WGSLLanguageFeatureName = 0x00000003

	WGSLLanguageFeatureNamePointerCompositeAccess WGSLLanguageFeatureName = 0x00000004

	WGSLLanguageFeatureNameUniformBufferStandardLayout WGSLLanguageFeatureName = 0x00000005

	WGSLLanguageFeatureNameSubgroupId WGSLLanguageFeatureName = 0x00000006

	WGSLLanguageFeatureNameTextureAndSamplerLet WGSLLanguageFeatureName = 0x00000007

	WGSLLanguageFeatureNameSubgroupUniformity WGSLLanguageFeatureName = 0x00000008

	WGSLLanguageFeatureNameTextureFormatsTier1 WGSLLanguageFeatureName = 0x00000009

	WGSLLanguageFeatureNameLinearIndexing WGSLLanguageFeatureName = 0x0000000A

	WGSLLanguageFeatureNameImmediateAddressSpace WGSLLanguageFeatureName = 0x0000000B

	WGSLLanguageFeatureNameBufferView WGSLLanguageFeatureName = 0x0000000C

	WGSLLanguageFeatureNameSwizzleAssignment WGSLLanguageFeatureName = 0x0000000D

	WGSLLanguageFeatureNameFragmentDepth WGSLLanguageFeatureName = 0x0000000E
)
