// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package sys

import "math"

// Indicates no array layer count is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const ArrayLayerCountUndefined = math.MaxUint32

// Indicates no copy stride is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const CopyStrideUndefined = math.MaxUint32

// Indicates no depth clear value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
var DepthClearValueUndefined = float32(math.NaN())

// Indicates no depth slice is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const DepthSliceUndefined = math.MaxUint32

// For `uint32_t` limits, indicates no limit value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const LimitU32Undefined = math.MaxUint32

// For `uint64_t` limits, indicates no limit value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const LimitU64Undefined = math.MaxUint64

// Indicates no mip level count is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const MipLevelCountUndefined = math.MaxUint32

// Indicates no query set index is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const QuerySetIndexUndefined = math.MaxUint32

// Sentinel value used in @ref WGPUStringView to indicate that the pointer
// is to a null-terminated string, rather than an explicitly-sized string.
const Strlen = math.MaxUint

// Indicates a size extending to the end of the buffer. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const WholeMapSize = math.MaxUint

// Indicates a size extending to the end of the buffer. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const WholeSize = math.MaxUint64
