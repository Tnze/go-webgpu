// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

//go:build !cgo && amd64

package sys

import "unsafe"

var (
	createInstanceProc      = webgpuDLL.NewProc("wgpuCreateInstance")
	getInstanceFeaturesProc = webgpuDLL.NewProc("wgpuGetInstanceFeatures")
	getInstanceLimitsProc   = webgpuDLL.NewProc("wgpuGetInstanceLimits")
	hasInstanceFeatureProc  = webgpuDLL.NewProc("wgpuHasInstanceFeature")
)

// Create a WGPUInstance
func CreateInstance(descriptor *InstanceDescriptor) Instance {
	ret, _, _ := createInstanceProc.Call(
		uintptr(unsafe.Pointer(descriptor)),
	)
	return Instance(ret)
}

// Get the list of @ref WGPUInstanceFeatureName values supported by the instance.
func GetInstanceFeatures(features *SupportedInstanceFeatures) {
	_, _, _ = getInstanceFeaturesProc.Call(
		uintptr(unsafe.Pointer(features)),
	)

}

// Get the limits supported by the instance.
func GetInstanceLimits(limits *InstanceLimits) Status {
	ret, _, _ := getInstanceLimitsProc.Call(
		uintptr(unsafe.Pointer(limits)),
	)
	return Status(ret)
}

// Check whether a particular @ref WGPUInstanceFeatureName is supported by the instance.
func HasInstanceFeature(feature InstanceFeatureName) Bool {
	ret, _, _ := hasInstanceFeatureProc.Call(
		uintptr(feature),
	)
	return Bool(ret)
}
