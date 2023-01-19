//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	apisconfig "github.com/gardener/gardener/extensions/pkg/apis/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.NetworkProblemDetector != nil {
		in, out := &in.NetworkProblemDetector, &out.NetworkProblemDetector
		*out = new(NetworkProblemDetector)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheckConfig != nil {
		in, out := &in.HealthCheckConfig, &out.HealthCheckConfig
		*out = new(apisconfig.HealthCheckConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sExporter) DeepCopyInto(out *K8sExporter) {
	*out = *in
	if in.HeartbeatPeriod != nil {
		in, out := &in.HeartbeatPeriod, &out.HeartbeatPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sExporter.
func (in *K8sExporter) DeepCopy() *K8sExporter {
	if in == nil {
		return nil
	}
	out := new(K8sExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkProblemDetector) DeepCopyInto(out *NetworkProblemDetector) {
	*out = *in
	if in.DefaultPeriod != nil {
		in, out := &in.DefaultPeriod, &out.DefaultPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.MaxPeerNodes != nil {
		in, out := &in.MaxPeerNodes, &out.MaxPeerNodes
		*out = new(int)
		**out = **in
	}
	if in.PSPDisabled != nil {
		in, out := &in.PSPDisabled, &out.PSPDisabled
		*out = new(bool)
		**out = **in
	}
	if in.PingEnabled != nil {
		in, out := &in.PingEnabled, &out.PingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.K8sExporter != nil {
		in, out := &in.K8sExporter, &out.K8sExporter
		*out = new(K8sExporter)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkProblemDetector.
func (in *NetworkProblemDetector) DeepCopy() *NetworkProblemDetector {
	if in == nil {
		return nil
	}
	out := new(NetworkProblemDetector)
	in.DeepCopyInto(out)
	return out
}
