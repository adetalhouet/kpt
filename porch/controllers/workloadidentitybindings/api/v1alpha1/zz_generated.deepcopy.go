//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 The kpt Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpServiceAccountRef) DeepCopyInto(out *GcpServiceAccountRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpServiceAccountRef.
func (in *GcpServiceAccountRef) DeepCopy() *GcpServiceAccountRef {
	if in == nil {
		return nil
	}
	out := new(GcpServiceAccountRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountRef) DeepCopyInto(out *KubernetesServiceAccountRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountRef.
func (in *KubernetesServiceAccountRef) DeepCopy() *KubernetesServiceAccountRef {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityBinding) DeepCopyInto(out *WorkloadIdentityBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityBinding.
func (in *WorkloadIdentityBinding) DeepCopy() *WorkloadIdentityBinding {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadIdentityBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityBindingList) DeepCopyInto(out *WorkloadIdentityBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadIdentityBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityBindingList.
func (in *WorkloadIdentityBindingList) DeepCopy() *WorkloadIdentityBindingList {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadIdentityBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityBindingSpec) DeepCopyInto(out *WorkloadIdentityBindingSpec) {
	*out = *in
	out.KubernetesServiceAccountRef = in.KubernetesServiceAccountRef
	out.GcpServiceAccountRef = in.GcpServiceAccountRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityBindingSpec.
func (in *WorkloadIdentityBindingSpec) DeepCopy() *WorkloadIdentityBindingSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityBindingStatus) DeepCopyInto(out *WorkloadIdentityBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityBindingStatus.
func (in *WorkloadIdentityBindingStatus) DeepCopy() *WorkloadIdentityBindingStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityBindingStatus)
	in.DeepCopyInto(out)
	return out
}
