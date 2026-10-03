// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo

package sys

// #cgo CFLAGS: -I${SRCDIR}/../../webgpu-headers
// #include <webgpu.h>
import "C"

import "unsafe"

// Create a WGPUInstance
func CreateInstance(descriptor *InstanceDescriptor) Instance {
	return Instance(C.wgpuCreateInstance(
		(*C.WGPUInstanceDescriptor)(unsafe.Pointer(descriptor)),
	))
}

// Get the list of @ref WGPUInstanceFeatureName values supported by the instance.
func GetInstanceFeatures(features *SupportedInstanceFeatures) {
	C.wgpuGetInstanceFeatures(
		(*C.WGPUSupportedInstanceFeatures)(unsafe.Pointer(features)),
	)
}

// Get the limits supported by the instance.
func GetInstanceLimits(limits *InstanceLimits) Status {
	return Status(C.wgpuGetInstanceLimits(
		(*C.WGPUInstanceLimits)(unsafe.Pointer(limits)),
	))
}

// Check whether a particular @ref WGPUInstanceFeatureName is supported by the instance.
func HasInstanceFeature(feature InstanceFeatureName) Bool {
	return Bool(C.wgpuHasInstanceFeature(
		C.WGPUInstanceFeatureName(feature),
	))
}
