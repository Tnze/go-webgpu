// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import "github.com/Tnze/go-webgpu/webgpu/sys"

// Indicates no array layer count is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const ArrayLayerCountUndefined = sys.ArrayLayerCountUndefined

// Indicates no copy stride is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const CopyStrideUndefined = sys.CopyStrideUndefined

// Indicates no depth clear value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
var DepthClearValueUndefined = sys.DepthClearValueUndefined

// Indicates no depth slice is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const DepthSliceUndefined = sys.DepthSliceUndefined

// For `uint32_t` limits, indicates no limit value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const LimitU32Undefined = sys.LimitU32Undefined

// For `uint64_t` limits, indicates no limit value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const LimitU64Undefined = sys.LimitU64Undefined

// Indicates no mip level count is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const MipLevelCountUndefined = sys.MipLevelCountUndefined

// Indicates no query set index is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const QuerySetIndexUndefined = sys.QuerySetIndexUndefined

// Sentinel value used in @ref WGPUStringView to indicate that the pointer
// is to a null-terminated string, rather than an explicitly-sized string.
const Strlen = sys.Strlen

// Indicates a size extending to the end of the buffer. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const WholeMapSize = sys.WholeMapSize

// Indicates a size extending to the end of the buffer. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const WholeSize = sys.WholeSize
