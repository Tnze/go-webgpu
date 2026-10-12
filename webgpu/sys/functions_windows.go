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
	ret, _, err := createInstanceProc.Call(
		uintptr(unsafe.Pointer(descriptor)),
	)
	if err != nil {
		panic(err)
	}
	return Instance(ret)
}

// Get the list of @ref WGPUInstanceFeatureName values supported by the instance.
func GetInstanceFeatures(features *SupportedInstanceFeatures) {
	_, _, err := getInstanceFeaturesProc.Call(
		uintptr(unsafe.Pointer(features)),
	)
	if err != nil {
		panic(err)
	}

}

// Get the limits supported by the instance.
func GetInstanceLimits(limits *InstanceLimits) Status {
	ret, _, err := getInstanceLimitsProc.Call(
		uintptr(unsafe.Pointer(limits)),
	)
	if err != nil {
		panic(err)
	}
	return Status(ret)
}

// Check whether a particular @ref WGPUInstanceFeatureName is supported by the instance.
func HasInstanceFeature(feature InstanceFeatureName) Bool {
	ret, _, err := hasInstanceFeatureProc.Call(
		uintptr(feature),
	)
	if err != nil {
		panic(err)
	}
	return Bool(ret)
}
