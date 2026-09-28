// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
// 
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import "math"

// Indicates no array layer count is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const ARRAY_LAYER_COUNT_UNDEFINED = math.MaxUint32

// Indicates no copy stride is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const COPY_STRIDE_UNDEFINED = math.MaxUint32

// Indicates no depth clear value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
// const DEPTH_CLEAR_VALUE_UNDEFINED = math.NaN()

// Indicates no depth slice is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const DEPTH_SLICE_UNDEFINED = math.MaxUint32

// For `uint32_t` limits, indicates no limit value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const LIMIT_U32_UNDEFINED = math.MaxUint32

// For `uint64_t` limits, indicates no limit value is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const LIMIT_U64_UNDEFINED = math.MaxUint64

// Indicates no mip level count is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const MIP_LEVEL_COUNT_UNDEFINED = math.MaxUint32

// Indicates no query set index is specified. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const QUERY_SET_INDEX_UNDEFINED = math.MaxUint32

// Sentinel value used in @ref WGPUStringView to indicate that the pointer
// is to a null-terminated string, rather than an explicitly-sized string.
const STRLEN = math.MaxUint

// Indicates a size extending to the end of the buffer. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const WHOLE_MAP_SIZE = math.MaxUint

// Indicates a size extending to the end of the buffer. For more info,
// see @ref SentinelValues and the places that use this sentinel value.
const WHOLE_SIZE = math.MaxUint64
