// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import "github.com/Tnze/go-webgpu/webgpu/sys"

type AdapterType = sys.AdapterType

const (
	AdapterTypeDiscreteGPU   AdapterType = sys.AdapterTypeDiscreteGPU
	AdapterTypeIntegratedGPU AdapterType = sys.AdapterTypeIntegratedGPU
	AdapterTypeCPU           AdapterType = sys.AdapterTypeCPU
	AdapterTypeUnknown       AdapterType = sys.AdapterTypeUnknown
)

type AddressMode = sys.AddressMode

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	AddressModeUndefined    AddressMode = sys.AddressModeUndefined
	AddressModeClampToEdge  AddressMode = sys.AddressModeClampToEdge
	AddressModeRepeat       AddressMode = sys.AddressModeRepeat
	AddressModeMirrorRepeat AddressMode = sys.AddressModeMirrorRepeat
)

type BackendType = sys.BackendType

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	BackendTypeUndefined BackendType = sys.BackendTypeUndefined
	BackendTypeNull      BackendType = sys.BackendTypeNull
	BackendTypeWebGPU    BackendType = sys.BackendTypeWebGPU
	BackendTypeD3D11     BackendType = sys.BackendTypeD3D11
	BackendTypeD3D12     BackendType = sys.BackendTypeD3D12
	BackendTypeMetal     BackendType = sys.BackendTypeMetal
	BackendTypeVulkan    BackendType = sys.BackendTypeVulkan
	BackendTypeOpenGL    BackendType = sys.BackendTypeOpenGL
	BackendTypeOpenGLES  BackendType = sys.BackendTypeOpenGLES
)

type BlendFactor = sys.BlendFactor

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	BlendFactorUndefined         BlendFactor = sys.BlendFactorUndefined
	BlendFactorZero              BlendFactor = sys.BlendFactorZero
	BlendFactorOne               BlendFactor = sys.BlendFactorOne
	BlendFactorSrc               BlendFactor = sys.BlendFactorSrc
	BlendFactorOneMinusSrc       BlendFactor = sys.BlendFactorOneMinusSrc
	BlendFactorSrcAlpha          BlendFactor = sys.BlendFactorSrcAlpha
	BlendFactorOneMinusSrcAlpha  BlendFactor = sys.BlendFactorOneMinusSrcAlpha
	BlendFactorDst               BlendFactor = sys.BlendFactorDst
	BlendFactorOneMinusDst       BlendFactor = sys.BlendFactorOneMinusDst
	BlendFactorDstAlpha          BlendFactor = sys.BlendFactorDstAlpha
	BlendFactorOneMinusDstAlpha  BlendFactor = sys.BlendFactorOneMinusDstAlpha
	BlendFactorSrcAlphaSaturated BlendFactor = sys.BlendFactorSrcAlphaSaturated
	BlendFactorConstant          BlendFactor = sys.BlendFactorConstant
	BlendFactorOneMinusConstant  BlendFactor = sys.BlendFactorOneMinusConstant
	BlendFactorSrc1              BlendFactor = sys.BlendFactorSrc1
	BlendFactorOneMinusSrc1      BlendFactor = sys.BlendFactorOneMinusSrc1
	BlendFactorSrc1Alpha         BlendFactor = sys.BlendFactorSrc1Alpha
	BlendFactorOneMinusSrc1Alpha BlendFactor = sys.BlendFactorOneMinusSrc1Alpha
)

type BlendOperation = sys.BlendOperation

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	BlendOperationUndefined       BlendOperation = sys.BlendOperationUndefined
	BlendOperationAdd             BlendOperation = sys.BlendOperationAdd
	BlendOperationSubtract        BlendOperation = sys.BlendOperationSubtract
	BlendOperationReverseSubtract BlendOperation = sys.BlendOperationReverseSubtract
	BlendOperationMin             BlendOperation = sys.BlendOperationMin
	BlendOperationMax             BlendOperation = sys.BlendOperationMax
)

type BufferBindingType = sys.BufferBindingType

const (
	// Indicates that this @ref WGPUBufferBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	BufferBindingTypeBindingNotUsed BufferBindingType = sys.BufferBindingTypeBindingNotUsed
	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	BufferBindingTypeUndefined       BufferBindingType = sys.BufferBindingTypeUndefined
	BufferBindingTypeUniform         BufferBindingType = sys.BufferBindingTypeUniform
	BufferBindingTypeStorage         BufferBindingType = sys.BufferBindingTypeStorage
	BufferBindingTypeReadOnlyStorage BufferBindingType = sys.BufferBindingTypeReadOnlyStorage
)

type BufferMapState = sys.BufferMapState

const (
	BufferMapStateUnmapped BufferMapState = sys.BufferMapStateUnmapped
	BufferMapStatePending  BufferMapState = sys.BufferMapStatePending
	BufferMapStateMapped   BufferMapState = sys.BufferMapStateMapped
)

// The callback mode controls how a callback for an asynchronous operation may be fired. See @ref Asynchronous-Operations for how these are used.
type CallbackMode = sys.CallbackMode

const (
	// Callbacks created with `WGPUCallbackMode_WaitAnyOnly`:
	// - fire when the asynchronous operation's future is passed to a call to @ref wgpuInstanceWaitAny
	// AND the operation has already completed or it completes inside the call to @ref wgpuInstanceWaitAny.
	CallbackModeWaitAnyOnly CallbackMode = sys.CallbackModeWaitAnyOnly
	// Callbacks created with `WGPUCallbackMode_AllowProcessEvents`:
	// - fire for the same reasons as callbacks created with `WGPUCallbackMode_WaitAnyOnly`
	// - fire inside a call to @ref wgpuInstanceProcessEvents if the asynchronous operation is complete.
	CallbackModeAllowProcessEvents CallbackMode = sys.CallbackModeAllowProcessEvents
	// Callbacks created with `WGPUCallbackMode_AllowSpontaneous`:
	// - fire for the same reasons as callbacks created with `WGPUCallbackMode_AllowProcessEvents`
	// - **may** fire spontaneously on an arbitrary or application thread, when the WebGPU implementations discovers that the asynchronous operation is complete.
	//
	// Implementations _should_ fire spontaneous callbacks as soon as possible.
	//
	// @note Because spontaneous callbacks may fire at an arbitrary time on an arbitrary thread, applications should take extra care when acquiring locks or mutating state inside the callback. It undefined behavior to re-entrantly call into the webgpu.h API if the callback fires while inside the callstack of another webgpu.h function that is not `wgpuInstanceWaitAny` or `wgpuInstanceProcessEvents`.
	CallbackModeAllowSpontaneous CallbackMode = sys.CallbackModeAllowSpontaneous
)

type CompareFunction = sys.CompareFunction

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	CompareFunctionUndefined    CompareFunction = sys.CompareFunctionUndefined
	CompareFunctionNever        CompareFunction = sys.CompareFunctionNever
	CompareFunctionLess         CompareFunction = sys.CompareFunctionLess
	CompareFunctionEqual        CompareFunction = sys.CompareFunctionEqual
	CompareFunctionLessEqual    CompareFunction = sys.CompareFunctionLessEqual
	CompareFunctionGreater      CompareFunction = sys.CompareFunctionGreater
	CompareFunctionNotEqual     CompareFunction = sys.CompareFunctionNotEqual
	CompareFunctionGreaterEqual CompareFunction = sys.CompareFunctionGreaterEqual
	CompareFunctionAlways       CompareFunction = sys.CompareFunctionAlways
)

type CompilationInfoRequestStatus = sys.CompilationInfoRequestStatus

const (
	CompilationInfoRequestStatusSuccess CompilationInfoRequestStatus = sys.CompilationInfoRequestStatusSuccess
	// See @ref CallbackStatuses.
	CompilationInfoRequestStatusCallbackCancelled CompilationInfoRequestStatus = sys.CompilationInfoRequestStatusCallbackCancelled
)

type CompilationMessageType = sys.CompilationMessageType

const (
	CompilationMessageTypeError   CompilationMessageType = sys.CompilationMessageTypeError
	CompilationMessageTypeWarning CompilationMessageType = sys.CompilationMessageTypeWarning
	CompilationMessageTypeInfo    CompilationMessageType = sys.CompilationMessageTypeInfo
)

type ComponentSwizzle = sys.ComponentSwizzle

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	ComponentSwizzleUndefined ComponentSwizzle = sys.ComponentSwizzleUndefined
	// Force its value to 0.
	ComponentSwizzleZero ComponentSwizzle = sys.ComponentSwizzleZero
	// Force its value to 1.
	ComponentSwizzleOne ComponentSwizzle = sys.ComponentSwizzleOne
	// Take its value from the red channel of the texture.
	ComponentSwizzleR ComponentSwizzle = sys.ComponentSwizzleR
	// Take its value from the green channel of the texture.
	ComponentSwizzleG ComponentSwizzle = sys.ComponentSwizzleG
	// Take its value from the blue channel of the texture.
	ComponentSwizzleB ComponentSwizzle = sys.ComponentSwizzleB
	// Take its value from the alpha channel of the texture.
	ComponentSwizzleA ComponentSwizzle = sys.ComponentSwizzleA
)

// Describes how frames are composited with other contents on the screen when @ref wgpuSurfacePresent is called.
type CompositeAlphaMode = sys.CompositeAlphaMode

const (
	// Lets the WebGPU implementation choose the best mode (supported, and with the best performance) between @ref WGPUCompositeAlphaMode_Opaque or @ref WGPUCompositeAlphaMode_Inherit.
	CompositeAlphaModeAuto CompositeAlphaMode = sys.CompositeAlphaModeAuto
	// The alpha component of the image is ignored and teated as if it is always 1.0.
	CompositeAlphaModeOpaque CompositeAlphaMode = sys.CompositeAlphaModeOpaque
	// The alpha component is respected and non-alpha components are assumed to be already multiplied with the alpha component. For example, (0.5, 0, 0, 0.5) is semi-transparent bright red.
	CompositeAlphaModePremultiplied CompositeAlphaMode = sys.CompositeAlphaModePremultiplied
	// The alpha component is respected and non-alpha components are assumed to NOT be already multiplied with the alpha component. For example, (1.0, 0, 0, 0.5) is semi-transparent bright red.
	CompositeAlphaModeUnpremultiplied CompositeAlphaMode = sys.CompositeAlphaModeUnpremultiplied
	// The handling of the alpha component is unknown to WebGPU and should be handled by the application using system-specific APIs. This mode may be unavailable (for example on Wasm).
	CompositeAlphaModeInherit CompositeAlphaMode = sys.CompositeAlphaModeInherit
)

type CreatePipelineAsyncStatus = sys.CreatePipelineAsyncStatus

const (
	CreatePipelineAsyncStatusSuccess CreatePipelineAsyncStatus = sys.CreatePipelineAsyncStatusSuccess
	// See @ref CallbackStatuses.
	CreatePipelineAsyncStatusCallbackCancelled CreatePipelineAsyncStatus = sys.CreatePipelineAsyncStatusCallbackCancelled
	CreatePipelineAsyncStatusValidationError   CreatePipelineAsyncStatus = sys.CreatePipelineAsyncStatusValidationError
	CreatePipelineAsyncStatusInternalError     CreatePipelineAsyncStatus = sys.CreatePipelineAsyncStatusInternalError
)

type CullMode = sys.CullMode

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	CullModeUndefined CullMode = sys.CullModeUndefined
	CullModeNone      CullMode = sys.CullModeNone
	CullModeFront     CullMode = sys.CullModeFront
	CullModeBack      CullMode = sys.CullModeBack
)

type DeviceLostReason = sys.DeviceLostReason

const (
	DeviceLostReasonUnknown   DeviceLostReason = sys.DeviceLostReasonUnknown
	DeviceLostReasonDestroyed DeviceLostReason = sys.DeviceLostReasonDestroyed
	// See @ref CallbackStatuses.
	DeviceLostReasonCallbackCancelled DeviceLostReason = sys.DeviceLostReasonCallbackCancelled
	DeviceLostReasonFailedCreation    DeviceLostReason = sys.DeviceLostReasonFailedCreation
)

type ErrorFilter = sys.ErrorFilter

const (
	ErrorFilterValidation  ErrorFilter = sys.ErrorFilterValidation
	ErrorFilterOutOfMemory ErrorFilter = sys.ErrorFilterOutOfMemory
	ErrorFilterInternal    ErrorFilter = sys.ErrorFilterInternal
)

type ErrorType = sys.ErrorType

const (
	ErrorTypeNoError     ErrorType = sys.ErrorTypeNoError
	ErrorTypeValidation  ErrorType = sys.ErrorTypeValidation
	ErrorTypeOutOfMemory ErrorType = sys.ErrorTypeOutOfMemory
	ErrorTypeInternal    ErrorType = sys.ErrorTypeInternal
	ErrorTypeUnknown     ErrorType = sys.ErrorTypeUnknown
)

// See @ref WGPURequestAdapterOptions::featureLevel.
type FeatureLevel = sys.FeatureLevel

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	FeatureLevelUndefined FeatureLevel = sys.FeatureLevelUndefined
	// "Compatibility" profile which can be supported on OpenGL ES 3.1 and D3D11.
	FeatureLevelCompatibility FeatureLevel = sys.FeatureLevelCompatibility
	// "Core" profile which can be supported on Vulkan/Metal/D3D12 (at least).
	FeatureLevelCore FeatureLevel = sys.FeatureLevelCore
)

type FeatureName = sys.FeatureName

const (
	FeatureNameCoreFeaturesAndLimits          FeatureName = sys.FeatureNameCoreFeaturesAndLimits
	FeatureNameDepthClipControl               FeatureName = sys.FeatureNameDepthClipControl
	FeatureNameDepth32FloatStencil8           FeatureName = sys.FeatureNameDepth32FloatStencil8
	FeatureNameTextureCompressionBC           FeatureName = sys.FeatureNameTextureCompressionBC
	FeatureNameTextureCompressionBCSliced3D   FeatureName = sys.FeatureNameTextureCompressionBCSliced3D
	FeatureNameTextureCompressionETC2         FeatureName = sys.FeatureNameTextureCompressionETC2
	FeatureNameTextureCompressionASTC         FeatureName = sys.FeatureNameTextureCompressionASTC
	FeatureNameTextureCompressionASTCSliced3D FeatureName = sys.FeatureNameTextureCompressionASTCSliced3D
	FeatureNameTimestampQuery                 FeatureName = sys.FeatureNameTimestampQuery
	FeatureNameIndirectFirstInstance          FeatureName = sys.FeatureNameIndirectFirstInstance
	FeatureNameShaderF16                      FeatureName = sys.FeatureNameShaderF16
	FeatureNameRG11B10UfloatRenderable        FeatureName = sys.FeatureNameRG11B10UfloatRenderable
	FeatureNameBGRA8UnormStorage              FeatureName = sys.FeatureNameBGRA8UnormStorage
	FeatureNameFloat32Filterable              FeatureName = sys.FeatureNameFloat32Filterable
	FeatureNameFloat32Blendable               FeatureName = sys.FeatureNameFloat32Blendable
	FeatureNameClipDistances                  FeatureName = sys.FeatureNameClipDistances
	FeatureNameDualSourceBlending             FeatureName = sys.FeatureNameDualSourceBlending
	FeatureNameSubgroups                      FeatureName = sys.FeatureNameSubgroups
	FeatureNameTextureFormatsTier1            FeatureName = sys.FeatureNameTextureFormatsTier1
	FeatureNameTextureFormatsTier2            FeatureName = sys.FeatureNameTextureFormatsTier2
	FeatureNamePrimitiveIndex                 FeatureName = sys.FeatureNamePrimitiveIndex
	FeatureNameTextureComponentSwizzle        FeatureName = sys.FeatureNameTextureComponentSwizzle
	FeatureNameSubgroupSizeControl            FeatureName = sys.FeatureNameSubgroupSizeControl
	FeatureNameTextureCompressionUnaligned    FeatureName = sys.FeatureNameTextureCompressionUnaligned
)

type FilterMode = sys.FilterMode

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	FilterModeUndefined FilterMode = sys.FilterModeUndefined
	FilterModeNearest   FilterMode = sys.FilterModeNearest
	FilterModeLinear    FilterMode = sys.FilterModeLinear
)

type FrontFace = sys.FrontFace

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	FrontFaceUndefined FrontFace = sys.FrontFaceUndefined
	FrontFaceCCW       FrontFace = sys.FrontFaceCCW
	FrontFaceCW        FrontFace = sys.FrontFaceCW
)

type IndexFormat = sys.IndexFormat

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	IndexFormatUndefined IndexFormat = sys.IndexFormatUndefined
	IndexFormatUint16    IndexFormat = sys.IndexFormatUint16
	IndexFormatUint32    IndexFormat = sys.IndexFormatUint32
)

type InstanceFeatureName = sys.InstanceFeatureName

const (
	// Enable use of ::wgpuInstanceWaitAny with `timeoutNS > 0`.
	InstanceFeatureNameTimedWaitAny InstanceFeatureName = sys.InstanceFeatureNameTimedWaitAny
	// Enable passing SPIR-V shaders to @ref wgpuDeviceCreateShaderModule,
	// via @ref WGPUShaderSourceSPIRV.
	InstanceFeatureNameShaderSourceSPIRV InstanceFeatureName = sys.InstanceFeatureNameShaderSourceSPIRV
	// Normally, a @ref WGPUAdapter can only create a single device. If this is
	// available and enabled, then adapters won't immediately expire when they
	// create a device, so can be reused to make multiple devices. They may
	// still expire for other reasons.
	InstanceFeatureNameMultipleDevicesPerAdapter InstanceFeatureName = sys.InstanceFeatureNameMultipleDevicesPerAdapter
)

type LoadOp = sys.LoadOp

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	LoadOpUndefined LoadOp = sys.LoadOpUndefined
	LoadOpLoad      LoadOp = sys.LoadOpLoad
	LoadOpClear     LoadOp = sys.LoadOpClear
)

type MapAsyncStatus = sys.MapAsyncStatus

const (
	MapAsyncStatusSuccess MapAsyncStatus = sys.MapAsyncStatusSuccess
	// See @ref CallbackStatuses.
	MapAsyncStatusCallbackCancelled MapAsyncStatus = sys.MapAsyncStatusCallbackCancelled
	MapAsyncStatusError             MapAsyncStatus = sys.MapAsyncStatusError
	MapAsyncStatusAborted           MapAsyncStatus = sys.MapAsyncStatusAborted
)

type MipmapFilterMode = sys.MipmapFilterMode

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	MipmapFilterModeUndefined MipmapFilterMode = sys.MipmapFilterModeUndefined
	MipmapFilterModeNearest   MipmapFilterMode = sys.MipmapFilterModeNearest
	MipmapFilterModeLinear    MipmapFilterMode = sys.MipmapFilterModeLinear
)

type OptionalBool = sys.OptionalBool

const (
	OptionalBoolFalse     OptionalBool = sys.OptionalBoolFalse
	OptionalBoolTrue      OptionalBool = sys.OptionalBoolTrue
	OptionalBoolUndefined OptionalBool = sys.OptionalBoolUndefined
)

type PopErrorScopeStatus = sys.PopErrorScopeStatus

const (
	// The error scope stack was successfully popped and a result was reported.
	PopErrorScopeStatusSuccess PopErrorScopeStatus = sys.PopErrorScopeStatusSuccess
	// See @ref CallbackStatuses.
	PopErrorScopeStatusCallbackCancelled PopErrorScopeStatus = sys.PopErrorScopeStatusCallbackCancelled
	// The error scope stack could not be popped, because it was empty.
	PopErrorScopeStatusError PopErrorScopeStatus = sys.PopErrorScopeStatusError
)

type PowerPreference = sys.PowerPreference

const (
	// No preference. (See also @ref SentinelValues.)
	PowerPreferenceUndefined       PowerPreference = sys.PowerPreferenceUndefined
	PowerPreferenceLowPower        PowerPreference = sys.PowerPreferenceLowPower
	PowerPreferenceHighPerformance PowerPreference = sys.PowerPreferenceHighPerformance
)

type PredefinedColorSpace = sys.PredefinedColorSpace

const (
	PredefinedColorSpaceSRGB      PredefinedColorSpace = sys.PredefinedColorSpaceSRGB
	PredefinedColorSpaceDisplayP3 PredefinedColorSpace = sys.PredefinedColorSpaceDisplayP3
)

// Describes when and in which order frames are presented on the screen when @ref wgpuSurfacePresent is called.
type PresentMode = sys.PresentMode

const (
	// Present mode is not specified. Use the default.
	PresentModeUndefined PresentMode = sys.PresentModeUndefined
	// The presentation of the image to the user waits for the next vertical blanking period to update in a first-in, first-out manner.
	// Tearing cannot be observed and frame-loop will be limited to the display's refresh rate.
	// This is the only mode that's always available.
	PresentModeFifo PresentMode = sys.PresentModeFifo
	// The presentation of the image to the user tries to wait for the next vertical blanking period but may decide to not wait if a frame is presented late.
	// Tearing can sometimes be observed but late-frame don't produce a full-frame stutter in the presentation.
	// This is still a first-in, first-out mechanism so a frame-loop will be limited to the display's refresh rate.
	PresentModeFifoRelaxed PresentMode = sys.PresentModeFifoRelaxed
	// The presentation of the image to the user is updated immediately without waiting for a vertical blank.
	// Tearing can be observed but latency is minimized.
	PresentModeImmediate PresentMode = sys.PresentModeImmediate
	// The presentation of the image to the user waits for the next vertical blanking period to update to the latest provided image.
	// Tearing cannot be observed and a frame-loop is not limited to the display's refresh rate.
	PresentModeMailbox PresentMode = sys.PresentModeMailbox
)

type PrimitiveTopology = sys.PrimitiveTopology

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	PrimitiveTopologyUndefined     PrimitiveTopology = sys.PrimitiveTopologyUndefined
	PrimitiveTopologyPointList     PrimitiveTopology = sys.PrimitiveTopologyPointList
	PrimitiveTopologyLineList      PrimitiveTopology = sys.PrimitiveTopologyLineList
	PrimitiveTopologyLineStrip     PrimitiveTopology = sys.PrimitiveTopologyLineStrip
	PrimitiveTopologyTriangleList  PrimitiveTopology = sys.PrimitiveTopologyTriangleList
	PrimitiveTopologyTriangleStrip PrimitiveTopology = sys.PrimitiveTopologyTriangleStrip
)

type QueryType = sys.QueryType

const (
	QueryTypeOcclusion QueryType = sys.QueryTypeOcclusion
	QueryTypeTimestamp QueryType = sys.QueryTypeTimestamp
)

type QueueWorkDoneStatus = sys.QueueWorkDoneStatus

const (
	QueueWorkDoneStatusSuccess QueueWorkDoneStatus = sys.QueueWorkDoneStatusSuccess
	// See @ref CallbackStatuses.
	QueueWorkDoneStatusCallbackCancelled QueueWorkDoneStatus = sys.QueueWorkDoneStatusCallbackCancelled
	// There was some deterministic error. (Note this is currently never used,
	// but it will be relevant when it's possible to create a queue object.)
	QueueWorkDoneStatusError QueueWorkDoneStatus = sys.QueueWorkDoneStatusError
)

type RequestAdapterStatus = sys.RequestAdapterStatus

const (
	RequestAdapterStatusSuccess RequestAdapterStatus = sys.RequestAdapterStatusSuccess
	// See @ref CallbackStatuses.
	RequestAdapterStatusCallbackCancelled RequestAdapterStatus = sys.RequestAdapterStatusCallbackCancelled
	RequestAdapterStatusUnavailable       RequestAdapterStatus = sys.RequestAdapterStatusUnavailable
	RequestAdapterStatusError             RequestAdapterStatus = sys.RequestAdapterStatusError
)

type RequestDeviceStatus = sys.RequestDeviceStatus

const (
	RequestDeviceStatusSuccess RequestDeviceStatus = sys.RequestDeviceStatusSuccess
	// See @ref CallbackStatuses.
	RequestDeviceStatusCallbackCancelled RequestDeviceStatus = sys.RequestDeviceStatusCallbackCancelled
	RequestDeviceStatusError             RequestDeviceStatus = sys.RequestDeviceStatusError
)

type SType = sys.SType

const (
	STypeShaderSourceSPIRV                 SType = sys.STypeShaderSourceSPIRV
	STypeShaderSourceWGSL                  SType = sys.STypeShaderSourceWGSL
	STypeRenderPassMaxDrawCount            SType = sys.STypeRenderPassMaxDrawCount
	STypeSurfaceSourceMetalLayer           SType = sys.STypeSurfaceSourceMetalLayer
	STypeSurfaceSourceWindowsHWND          SType = sys.STypeSurfaceSourceWindowsHWND
	STypeSurfaceSourceXlibWindow           SType = sys.STypeSurfaceSourceXlibWindow
	STypeSurfaceSourceWaylandSurface       SType = sys.STypeSurfaceSourceWaylandSurface
	STypeSurfaceSourceAndroidNativeWindow  SType = sys.STypeSurfaceSourceAndroidNativeWindow
	STypeSurfaceSourceXCBWindow            SType = sys.STypeSurfaceSourceXCBWindow
	STypeSurfaceColorManagement            SType = sys.STypeSurfaceColorManagement
	STypeRequestAdapterWebXROptions        SType = sys.STypeRequestAdapterWebXROptions
	STypeTextureComponentSwizzleDescriptor SType = sys.STypeTextureComponentSwizzleDescriptor
	STypeExternalTextureBindingLayout      SType = sys.STypeExternalTextureBindingLayout
	STypeExternalTextureBindingEntry       SType = sys.STypeExternalTextureBindingEntry
	STypeCompatibilityModeLimits           SType = sys.STypeCompatibilityModeLimits
	STypeTextureBindingViewDimension       SType = sys.STypeTextureBindingViewDimension
)

type SamplerBindingType = sys.SamplerBindingType

const (
	// Indicates that this @ref WGPUSamplerBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	SamplerBindingTypeBindingNotUsed SamplerBindingType = sys.SamplerBindingTypeBindingNotUsed
	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	SamplerBindingTypeUndefined    SamplerBindingType = sys.SamplerBindingTypeUndefined
	SamplerBindingTypeFiltering    SamplerBindingType = sys.SamplerBindingTypeFiltering
	SamplerBindingTypeNonFiltering SamplerBindingType = sys.SamplerBindingTypeNonFiltering
	SamplerBindingTypeComparison   SamplerBindingType = sys.SamplerBindingTypeComparison
)

// Status code returned (synchronously) from many operations. Generally
// indicates an invalid input like an unknown enum value or @ref OutStructChainError.
// Read the function's documentation for specific error conditions.
type Status = sys.Status

const (
	StatusSuccess Status = sys.StatusSuccess
	StatusError   Status = sys.StatusError
)

type StencilOperation = sys.StencilOperation

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	StencilOperationUndefined      StencilOperation = sys.StencilOperationUndefined
	StencilOperationKeep           StencilOperation = sys.StencilOperationKeep
	StencilOperationZero           StencilOperation = sys.StencilOperationZero
	StencilOperationReplace        StencilOperation = sys.StencilOperationReplace
	StencilOperationInvert         StencilOperation = sys.StencilOperationInvert
	StencilOperationIncrementClamp StencilOperation = sys.StencilOperationIncrementClamp
	StencilOperationDecrementClamp StencilOperation = sys.StencilOperationDecrementClamp
	StencilOperationIncrementWrap  StencilOperation = sys.StencilOperationIncrementWrap
	StencilOperationDecrementWrap  StencilOperation = sys.StencilOperationDecrementWrap
)

type StorageTextureAccess = sys.StorageTextureAccess

const (
	// Indicates that this @ref WGPUStorageTextureBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	StorageTextureAccessBindingNotUsed StorageTextureAccess = sys.StorageTextureAccessBindingNotUsed
	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	StorageTextureAccessUndefined StorageTextureAccess = sys.StorageTextureAccessUndefined
	StorageTextureAccessWriteOnly StorageTextureAccess = sys.StorageTextureAccessWriteOnly
	StorageTextureAccessReadOnly  StorageTextureAccess = sys.StorageTextureAccessReadOnly
	StorageTextureAccessReadWrite StorageTextureAccess = sys.StorageTextureAccessReadWrite
)

type StoreOp = sys.StoreOp

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	StoreOpUndefined StoreOp = sys.StoreOpUndefined
	StoreOpStore     StoreOp = sys.StoreOpStore
	StoreOpDiscard   StoreOp = sys.StoreOpDiscard
)

// The status enum for @ref wgpuSurfaceGetCurrentTexture.
type SurfaceGetCurrentTextureStatus = sys.SurfaceGetCurrentTextureStatus

const (
	// Yay! Everything is good and we can render this frame.
	SurfaceGetCurrentTextureStatusSuccessOptimal SurfaceGetCurrentTextureStatus = sys.SurfaceGetCurrentTextureStatusSuccessOptimal
	// Still OK - the surface can present the frame, but in a suboptimal way. The surface may need reconfiguration.
	SurfaceGetCurrentTextureStatusSuccessSuboptimal SurfaceGetCurrentTextureStatus = sys.SurfaceGetCurrentTextureStatusSuccessSuboptimal
	// Some operation timed out while trying to acquire the frame.
	SurfaceGetCurrentTextureStatusTimeout SurfaceGetCurrentTextureStatus = sys.SurfaceGetCurrentTextureStatusTimeout
	// The surface is too different to be used, compared to when it was originally created.
	SurfaceGetCurrentTextureStatusOutdated SurfaceGetCurrentTextureStatus = sys.SurfaceGetCurrentTextureStatusOutdated
	// The connection to whatever owns the surface was lost, or generally needs to be fully reinitialized.
	SurfaceGetCurrentTextureStatusLost SurfaceGetCurrentTextureStatus = sys.SurfaceGetCurrentTextureStatusLost
	// There was some deterministic error (for example, the surface is not configured, or there was an @ref OutStructChainError). Should produce @ref ImplementationDefinedLogging containing details.
	SurfaceGetCurrentTextureStatusError SurfaceGetCurrentTextureStatus = sys.SurfaceGetCurrentTextureStatusError
)

type TextureAspect = sys.TextureAspect

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureAspectUndefined   TextureAspect = sys.TextureAspectUndefined
	TextureAspectAll         TextureAspect = sys.TextureAspectAll
	TextureAspectStencilOnly TextureAspect = sys.TextureAspectStencilOnly
	TextureAspectDepthOnly   TextureAspect = sys.TextureAspectDepthOnly
)

type TextureDimension = sys.TextureDimension

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureDimensionUndefined TextureDimension = sys.TextureDimensionUndefined
	TextureDimension1D        TextureDimension = sys.TextureDimension1D
	TextureDimension2D        TextureDimension = sys.TextureDimension2D
	TextureDimension3D        TextureDimension = sys.TextureDimension3D
)

type TextureFormat = sys.TextureFormat

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureFormatUndefined            TextureFormat = sys.TextureFormatUndefined
	TextureFormatR8Unorm              TextureFormat = sys.TextureFormatR8Unorm
	TextureFormatR8Snorm              TextureFormat = sys.TextureFormatR8Snorm
	TextureFormatR8Uint               TextureFormat = sys.TextureFormatR8Uint
	TextureFormatR8Sint               TextureFormat = sys.TextureFormatR8Sint
	TextureFormatR16Unorm             TextureFormat = sys.TextureFormatR16Unorm
	TextureFormatR16Snorm             TextureFormat = sys.TextureFormatR16Snorm
	TextureFormatR16Uint              TextureFormat = sys.TextureFormatR16Uint
	TextureFormatR16Sint              TextureFormat = sys.TextureFormatR16Sint
	TextureFormatR16Float             TextureFormat = sys.TextureFormatR16Float
	TextureFormatRG8Unorm             TextureFormat = sys.TextureFormatRG8Unorm
	TextureFormatRG8Snorm             TextureFormat = sys.TextureFormatRG8Snorm
	TextureFormatRG8Uint              TextureFormat = sys.TextureFormatRG8Uint
	TextureFormatRG8Sint              TextureFormat = sys.TextureFormatRG8Sint
	TextureFormatR32Float             TextureFormat = sys.TextureFormatR32Float
	TextureFormatR32Uint              TextureFormat = sys.TextureFormatR32Uint
	TextureFormatR32Sint              TextureFormat = sys.TextureFormatR32Sint
	TextureFormatRG16Unorm            TextureFormat = sys.TextureFormatRG16Unorm
	TextureFormatRG16Snorm            TextureFormat = sys.TextureFormatRG16Snorm
	TextureFormatRG16Uint             TextureFormat = sys.TextureFormatRG16Uint
	TextureFormatRG16Sint             TextureFormat = sys.TextureFormatRG16Sint
	TextureFormatRG16Float            TextureFormat = sys.TextureFormatRG16Float
	TextureFormatRGBA8Unorm           TextureFormat = sys.TextureFormatRGBA8Unorm
	TextureFormatRGBA8UnormSrgb       TextureFormat = sys.TextureFormatRGBA8UnormSrgb
	TextureFormatRGBA8Snorm           TextureFormat = sys.TextureFormatRGBA8Snorm
	TextureFormatRGBA8Uint            TextureFormat = sys.TextureFormatRGBA8Uint
	TextureFormatRGBA8Sint            TextureFormat = sys.TextureFormatRGBA8Sint
	TextureFormatBGRA8Unorm           TextureFormat = sys.TextureFormatBGRA8Unorm
	TextureFormatBGRA8UnormSrgb       TextureFormat = sys.TextureFormatBGRA8UnormSrgb
	TextureFormatRGB10A2Uint          TextureFormat = sys.TextureFormatRGB10A2Uint
	TextureFormatRGB10A2Unorm         TextureFormat = sys.TextureFormatRGB10A2Unorm
	TextureFormatRG11B10Ufloat        TextureFormat = sys.TextureFormatRG11B10Ufloat
	TextureFormatRGB9E5Ufloat         TextureFormat = sys.TextureFormatRGB9E5Ufloat
	TextureFormatRG32Float            TextureFormat = sys.TextureFormatRG32Float
	TextureFormatRG32Uint             TextureFormat = sys.TextureFormatRG32Uint
	TextureFormatRG32Sint             TextureFormat = sys.TextureFormatRG32Sint
	TextureFormatRGBA16Unorm          TextureFormat = sys.TextureFormatRGBA16Unorm
	TextureFormatRGBA16Snorm          TextureFormat = sys.TextureFormatRGBA16Snorm
	TextureFormatRGBA16Uint           TextureFormat = sys.TextureFormatRGBA16Uint
	TextureFormatRGBA16Sint           TextureFormat = sys.TextureFormatRGBA16Sint
	TextureFormatRGBA16Float          TextureFormat = sys.TextureFormatRGBA16Float
	TextureFormatRGBA32Float          TextureFormat = sys.TextureFormatRGBA32Float
	TextureFormatRGBA32Uint           TextureFormat = sys.TextureFormatRGBA32Uint
	TextureFormatRGBA32Sint           TextureFormat = sys.TextureFormatRGBA32Sint
	TextureFormatStencil8             TextureFormat = sys.TextureFormatStencil8
	TextureFormatDepth16Unorm         TextureFormat = sys.TextureFormatDepth16Unorm
	TextureFormatDepth24Plus          TextureFormat = sys.TextureFormatDepth24Plus
	TextureFormatDepth24PlusStencil8  TextureFormat = sys.TextureFormatDepth24PlusStencil8
	TextureFormatDepth32Float         TextureFormat = sys.TextureFormatDepth32Float
	TextureFormatDepth32FloatStencil8 TextureFormat = sys.TextureFormatDepth32FloatStencil8
	TextureFormatBC1RGBAUnorm         TextureFormat = sys.TextureFormatBC1RGBAUnorm
	TextureFormatBC1RGBAUnormSrgb     TextureFormat = sys.TextureFormatBC1RGBAUnormSrgb
	TextureFormatBC2RGBAUnorm         TextureFormat = sys.TextureFormatBC2RGBAUnorm
	TextureFormatBC2RGBAUnormSrgb     TextureFormat = sys.TextureFormatBC2RGBAUnormSrgb
	TextureFormatBC3RGBAUnorm         TextureFormat = sys.TextureFormatBC3RGBAUnorm
	TextureFormatBC3RGBAUnormSrgb     TextureFormat = sys.TextureFormatBC3RGBAUnormSrgb
	TextureFormatBC4RUnorm            TextureFormat = sys.TextureFormatBC4RUnorm
	TextureFormatBC4RSnorm            TextureFormat = sys.TextureFormatBC4RSnorm
	TextureFormatBC5RGUnorm           TextureFormat = sys.TextureFormatBC5RGUnorm
	TextureFormatBC5RGSnorm           TextureFormat = sys.TextureFormatBC5RGSnorm
	TextureFormatBC6HRGBUfloat        TextureFormat = sys.TextureFormatBC6HRGBUfloat
	TextureFormatBC6HRGBFloat         TextureFormat = sys.TextureFormatBC6HRGBFloat
	TextureFormatBC7RGBAUnorm         TextureFormat = sys.TextureFormatBC7RGBAUnorm
	TextureFormatBC7RGBAUnormSrgb     TextureFormat = sys.TextureFormatBC7RGBAUnormSrgb
	TextureFormatETC2RGB8Unorm        TextureFormat = sys.TextureFormatETC2RGB8Unorm
	TextureFormatETC2RGB8UnormSrgb    TextureFormat = sys.TextureFormatETC2RGB8UnormSrgb
	TextureFormatETC2RGB8A1Unorm      TextureFormat = sys.TextureFormatETC2RGB8A1Unorm
	TextureFormatETC2RGB8A1UnormSrgb  TextureFormat = sys.TextureFormatETC2RGB8A1UnormSrgb
	TextureFormatETC2RGBA8Unorm       TextureFormat = sys.TextureFormatETC2RGBA8Unorm
	TextureFormatETC2RGBA8UnormSrgb   TextureFormat = sys.TextureFormatETC2RGBA8UnormSrgb
	TextureFormatEACR11Unorm          TextureFormat = sys.TextureFormatEACR11Unorm
	TextureFormatEACR11Snorm          TextureFormat = sys.TextureFormatEACR11Snorm
	TextureFormatEACRG11Unorm         TextureFormat = sys.TextureFormatEACRG11Unorm
	TextureFormatEACRG11Snorm         TextureFormat = sys.TextureFormatEACRG11Snorm
	TextureFormatASTC4x4Unorm         TextureFormat = sys.TextureFormatASTC4x4Unorm
	TextureFormatASTC4x4UnormSrgb     TextureFormat = sys.TextureFormatASTC4x4UnormSrgb
	TextureFormatASTC5x4Unorm         TextureFormat = sys.TextureFormatASTC5x4Unorm
	TextureFormatASTC5x4UnormSrgb     TextureFormat = sys.TextureFormatASTC5x4UnormSrgb
	TextureFormatASTC5x5Unorm         TextureFormat = sys.TextureFormatASTC5x5Unorm
	TextureFormatASTC5x5UnormSrgb     TextureFormat = sys.TextureFormatASTC5x5UnormSrgb
	TextureFormatASTC6x5Unorm         TextureFormat = sys.TextureFormatASTC6x5Unorm
	TextureFormatASTC6x5UnormSrgb     TextureFormat = sys.TextureFormatASTC6x5UnormSrgb
	TextureFormatASTC6x6Unorm         TextureFormat = sys.TextureFormatASTC6x6Unorm
	TextureFormatASTC6x6UnormSrgb     TextureFormat = sys.TextureFormatASTC6x6UnormSrgb
	TextureFormatASTC8x5Unorm         TextureFormat = sys.TextureFormatASTC8x5Unorm
	TextureFormatASTC8x5UnormSrgb     TextureFormat = sys.TextureFormatASTC8x5UnormSrgb
	TextureFormatASTC8x6Unorm         TextureFormat = sys.TextureFormatASTC8x6Unorm
	TextureFormatASTC8x6UnormSrgb     TextureFormat = sys.TextureFormatASTC8x6UnormSrgb
	TextureFormatASTC8x8Unorm         TextureFormat = sys.TextureFormatASTC8x8Unorm
	TextureFormatASTC8x8UnormSrgb     TextureFormat = sys.TextureFormatASTC8x8UnormSrgb
	TextureFormatASTC10x5Unorm        TextureFormat = sys.TextureFormatASTC10x5Unorm
	TextureFormatASTC10x5UnormSrgb    TextureFormat = sys.TextureFormatASTC10x5UnormSrgb
	TextureFormatASTC10x6Unorm        TextureFormat = sys.TextureFormatASTC10x6Unorm
	TextureFormatASTC10x6UnormSrgb    TextureFormat = sys.TextureFormatASTC10x6UnormSrgb
	TextureFormatASTC10x8Unorm        TextureFormat = sys.TextureFormatASTC10x8Unorm
	TextureFormatASTC10x8UnormSrgb    TextureFormat = sys.TextureFormatASTC10x8UnormSrgb
	TextureFormatASTC10x10Unorm       TextureFormat = sys.TextureFormatASTC10x10Unorm
	TextureFormatASTC10x10UnormSrgb   TextureFormat = sys.TextureFormatASTC10x10UnormSrgb
	TextureFormatASTC12x10Unorm       TextureFormat = sys.TextureFormatASTC12x10Unorm
	TextureFormatASTC12x10UnormSrgb   TextureFormat = sys.TextureFormatASTC12x10UnormSrgb
	TextureFormatASTC12x12Unorm       TextureFormat = sys.TextureFormatASTC12x12Unorm
	TextureFormatASTC12x12UnormSrgb   TextureFormat = sys.TextureFormatASTC12x12UnormSrgb
)

type TextureSampleType = sys.TextureSampleType

const (
	// Indicates that this @ref WGPUTextureBindingLayout member of
	// its parent @ref WGPUBindGroupLayoutEntry is not used.
	// (See also @ref SentinelValues.)
	TextureSampleTypeBindingNotUsed TextureSampleType = sys.TextureSampleTypeBindingNotUsed
	// `1`. Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureSampleTypeUndefined         TextureSampleType = sys.TextureSampleTypeUndefined
	TextureSampleTypeFloat             TextureSampleType = sys.TextureSampleTypeFloat
	TextureSampleTypeUnfilterableFloat TextureSampleType = sys.TextureSampleTypeUnfilterableFloat
	TextureSampleTypeDepth             TextureSampleType = sys.TextureSampleTypeDepth
	TextureSampleTypeSint              TextureSampleType = sys.TextureSampleTypeSint
	TextureSampleTypeUint              TextureSampleType = sys.TextureSampleTypeUint
)

type TextureViewDimension = sys.TextureViewDimension

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	TextureViewDimensionUndefined TextureViewDimension = sys.TextureViewDimensionUndefined
	TextureViewDimension1D        TextureViewDimension = sys.TextureViewDimension1D
	TextureViewDimension2D        TextureViewDimension = sys.TextureViewDimension2D
	TextureViewDimension2DArray   TextureViewDimension = sys.TextureViewDimension2DArray
	TextureViewDimensionCube      TextureViewDimension = sys.TextureViewDimensionCube
	TextureViewDimensionCubeArray TextureViewDimension = sys.TextureViewDimensionCubeArray
	TextureViewDimension3D        TextureViewDimension = sys.TextureViewDimension3D
)

type ToneMappingMode = sys.ToneMappingMode

const (
	ToneMappingModeStandard ToneMappingMode = sys.ToneMappingModeStandard
	ToneMappingModeExtended ToneMappingMode = sys.ToneMappingModeExtended
)

type VertexFormat = sys.VertexFormat

const (
	VertexFormatUint8           VertexFormat = sys.VertexFormatUint8
	VertexFormatUint8x2         VertexFormat = sys.VertexFormatUint8x2
	VertexFormatUint8x4         VertexFormat = sys.VertexFormatUint8x4
	VertexFormatSint8           VertexFormat = sys.VertexFormatSint8
	VertexFormatSint8x2         VertexFormat = sys.VertexFormatSint8x2
	VertexFormatSint8x4         VertexFormat = sys.VertexFormatSint8x4
	VertexFormatUnorm8          VertexFormat = sys.VertexFormatUnorm8
	VertexFormatUnorm8x2        VertexFormat = sys.VertexFormatUnorm8x2
	VertexFormatUnorm8x4        VertexFormat = sys.VertexFormatUnorm8x4
	VertexFormatSnorm8          VertexFormat = sys.VertexFormatSnorm8
	VertexFormatSnorm8x2        VertexFormat = sys.VertexFormatSnorm8x2
	VertexFormatSnorm8x4        VertexFormat = sys.VertexFormatSnorm8x4
	VertexFormatUint16          VertexFormat = sys.VertexFormatUint16
	VertexFormatUint16x2        VertexFormat = sys.VertexFormatUint16x2
	VertexFormatUint16x4        VertexFormat = sys.VertexFormatUint16x4
	VertexFormatSint16          VertexFormat = sys.VertexFormatSint16
	VertexFormatSint16x2        VertexFormat = sys.VertexFormatSint16x2
	VertexFormatSint16x4        VertexFormat = sys.VertexFormatSint16x4
	VertexFormatUnorm16         VertexFormat = sys.VertexFormatUnorm16
	VertexFormatUnorm16x2       VertexFormat = sys.VertexFormatUnorm16x2
	VertexFormatUnorm16x4       VertexFormat = sys.VertexFormatUnorm16x4
	VertexFormatSnorm16         VertexFormat = sys.VertexFormatSnorm16
	VertexFormatSnorm16x2       VertexFormat = sys.VertexFormatSnorm16x2
	VertexFormatSnorm16x4       VertexFormat = sys.VertexFormatSnorm16x4
	VertexFormatFloat16         VertexFormat = sys.VertexFormatFloat16
	VertexFormatFloat16x2       VertexFormat = sys.VertexFormatFloat16x2
	VertexFormatFloat16x4       VertexFormat = sys.VertexFormatFloat16x4
	VertexFormatFloat32         VertexFormat = sys.VertexFormatFloat32
	VertexFormatFloat32x2       VertexFormat = sys.VertexFormatFloat32x2
	VertexFormatFloat32x3       VertexFormat = sys.VertexFormatFloat32x3
	VertexFormatFloat32x4       VertexFormat = sys.VertexFormatFloat32x4
	VertexFormatUint32          VertexFormat = sys.VertexFormatUint32
	VertexFormatUint32x2        VertexFormat = sys.VertexFormatUint32x2
	VertexFormatUint32x3        VertexFormat = sys.VertexFormatUint32x3
	VertexFormatUint32x4        VertexFormat = sys.VertexFormatUint32x4
	VertexFormatSint32          VertexFormat = sys.VertexFormatSint32
	VertexFormatSint32x2        VertexFormat = sys.VertexFormatSint32x2
	VertexFormatSint32x3        VertexFormat = sys.VertexFormatSint32x3
	VertexFormatSint32x4        VertexFormat = sys.VertexFormatSint32x4
	VertexFormatUnorm10_10_10_2 VertexFormat = sys.VertexFormatUnorm10_10_10_2
	VertexFormatUnorm8x4BGRA    VertexFormat = sys.VertexFormatUnorm8x4BGRA
	VertexFormatSnorm10_10_10_2 VertexFormat = sys.VertexFormatSnorm10_10_10_2
)

type VertexStepMode = sys.VertexStepMode

const (
	// Indicates no value is passed for this argument. See @ref SentinelValues.
	VertexStepModeUndefined VertexStepMode = sys.VertexStepModeUndefined
	VertexStepModeVertex    VertexStepMode = sys.VertexStepModeVertex
	VertexStepModeInstance  VertexStepMode = sys.VertexStepModeInstance
)

// Status returned from a call to ::wgpuInstanceWaitAny.
type WaitStatus = sys.WaitStatus

const (
	// At least one WGPUFuture completed successfully.
	WaitStatusSuccess WaitStatus = sys.WaitStatusSuccess
	// The wait operation succeeded, but no WGPUFutures completed within the timeout.
	WaitStatusTimedOut WaitStatus = sys.WaitStatusTimedOut
	// The call was invalid for some reason (see @ref Wait-Any).
	// Should produce @ref ImplementationDefinedLogging containing details.
	WaitStatusError WaitStatus = sys.WaitStatusError
)

type WGSLLanguageFeatureName = sys.WGSLLanguageFeatureName

const (
	WGSLLanguageFeatureNameReadonlyAndReadwriteStorageTextures WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameReadonlyAndReadwriteStorageTextures
	WGSLLanguageFeatureNamePacked4x8IntegerDotProduct          WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNamePacked4x8IntegerDotProduct
	WGSLLanguageFeatureNameUnrestrictedPointerParameters       WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameUnrestrictedPointerParameters
	WGSLLanguageFeatureNamePointerCompositeAccess              WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNamePointerCompositeAccess
	WGSLLanguageFeatureNameUniformBufferStandardLayout         WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameUniformBufferStandardLayout
	WGSLLanguageFeatureNameSubgroupId                          WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameSubgroupId
	WGSLLanguageFeatureNameTextureAndSamplerLet                WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameTextureAndSamplerLet
	WGSLLanguageFeatureNameSubgroupUniformity                  WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameSubgroupUniformity
	WGSLLanguageFeatureNameTextureFormatsTier1                 WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameTextureFormatsTier1
	WGSLLanguageFeatureNameLinearIndexing                      WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameLinearIndexing
	WGSLLanguageFeatureNameImmediateAddressSpace               WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameImmediateAddressSpace
	WGSLLanguageFeatureNameBufferView                          WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameBufferView
	WGSLLanguageFeatureNameSwizzleAssignment                   WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameSwizzleAssignment
	WGSLLanguageFeatureNameFragmentDepth                       WGSLLanguageFeatureName = sys.WGSLLanguageFeatureNameFragmentDepth
)
