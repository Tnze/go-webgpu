// Code generated; DO NOT EDIT.
// Copyright 2019-2023 WebGPU-Native developers
//
// SPDX-License-Identifier: BSD-3-Clause

package webgpu

import "github.com/Tnze/go-webgpu/webgpu/sys"

// Create a WGPUInstance
func CreateInstance(descriptor *InstanceDescriptor) *Instance {
	var _descriptor *sys.InstanceDescriptor
	if descriptor != nil {
		_descriptor = new(descriptor.unwrap())
	}
	ret := sys.CreateInstance(_descriptor)
	return new(Instance{inner: ret}).owned()
}

// Get the list of @ref WGPUInstanceFeatureName values supported by the instance.
func GetInstanceFeatures(features *SupportedInstanceFeatures) {
	var _features *sys.SupportedInstanceFeatures
	if features != nil {
		_features = new(features.unwrap())
	}
	sys.GetInstanceFeatures(_features)
}

// Get the limits supported by the instance.
func GetInstanceLimits(limits *InstanceLimits) Status {
	ret := sys.GetInstanceLimits(limits)
	return ret
}

// Check whether a particular @ref WGPUInstanceFeatureName is supported by the instance.
func HasInstanceFeature(feature InstanceFeatureName) Bool {
	ret := sys.HasInstanceFeature(feature)
	return ret
}
