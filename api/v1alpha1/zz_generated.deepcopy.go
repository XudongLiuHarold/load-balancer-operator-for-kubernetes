// +build !ignore_autogenerated

// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1alpha3"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfig) DeepCopyInto(out *AKODeploymentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfig.
func (in *AKODeploymentConfig) DeepCopy() *AKODeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKODeploymentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfigList) DeepCopyInto(out *AKODeploymentConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKODeploymentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfigList.
func (in *AKODeploymentConfigList) DeepCopy() *AKODeploymentConfigList {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKODeploymentConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfigSpec) DeepCopyInto(out *AKODeploymentConfigSpec) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
	if in.WorkloadCredentialRef != nil {
		in, out := &in.WorkloadCredentialRef, &out.WorkloadCredentialRef
		*out = new(SecretRef)
		**out = **in
	}
	if in.AdminCredentialRef != nil {
		in, out := &in.AdminCredentialRef, &out.AdminCredentialRef
		*out = new(SecretRef)
		**out = **in
	}
	if in.CertificateAuthorityRef != nil {
		in, out := &in.CertificateAuthorityRef, &out.CertificateAuthorityRef
		*out = new(SecretRef)
		**out = **in
	}
	out.Tenant = in.Tenant
	in.DataNetwork.DeepCopyInto(&out.DataNetwork)
	out.ExtraConfigs = in.ExtraConfigs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfigSpec.
func (in *AKODeploymentConfigSpec) DeepCopy() *AKODeploymentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfigStatus) DeepCopyInto(out *AKODeploymentConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1alpha3.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfigStatus.
func (in *AKODeploymentConfigStatus) DeepCopy() *AKODeploymentConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOImageConfig) DeepCopyInto(out *AKOImageConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOImageConfig.
func (in *AKOImageConfig) DeepCopy() *AKOImageConfig {
	if in == nil {
		return nil
	}
	out := new(AKOImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOIngressConfig) DeepCopyInto(out *AKOIngressConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOIngressConfig.
func (in *AKOIngressConfig) DeepCopy() *AKOIngressConfig {
	if in == nil {
		return nil
	}
	out := new(AKOIngressConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOLogConfig) DeepCopyInto(out *AKOLogConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOLogConfig.
func (in *AKOLogConfig) DeepCopy() *AKOLogConfig {
	if in == nil {
		return nil
	}
	out := new(AKOLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKORbacConfig) DeepCopyInto(out *AKORbacConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKORbacConfig.
func (in *AKORbacConfig) DeepCopy() *AKORbacConfig {
	if in == nil {
		return nil
	}
	out := new(AKORbacConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AVITenant) DeepCopyInto(out *AVITenant) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AVITenant.
func (in *AVITenant) DeepCopy() *AVITenant {
	if in == nil {
		return nil
	}
	out := new(AVITenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNetwork) DeepCopyInto(out *DataNetwork) {
	*out = *in
	if in.IPPools != nil {
		in, out := &in.IPPools, &out.IPPools
		*out = make([]IPPool, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNetwork.
func (in *DataNetwork) DeepCopy() *DataNetwork {
	if in == nil {
		return nil
	}
	out := new(DataNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraConfigs) DeepCopyInto(out *ExtraConfigs) {
	*out = *in
	out.Image = in.Image
	out.Log = in.Log
	out.Rbac = in.Rbac
	out.IngressConfigs = in.IngressConfigs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraConfigs.
func (in *ExtraConfigs) DeepCopy() *ExtraConfigs {
	if in == nil {
		return nil
	}
	out := new(ExtraConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPool) DeepCopyInto(out *IPPool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPool.
func (in *IPPool) DeepCopy() *IPPool {
	if in == nil {
		return nil
	}
	out := new(IPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}
