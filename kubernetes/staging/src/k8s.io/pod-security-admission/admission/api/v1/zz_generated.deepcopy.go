//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityConfiguration) DeepCopyInto(out *PodSecurityConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Defaults = in.Defaults
	in.Exemptions.DeepCopyInto(&out.Exemptions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityConfiguration.
func (in *PodSecurityConfiguration) DeepCopy() *PodSecurityConfiguration {
	if in == nil {
		return nil
	}
	out := new(PodSecurityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityDefaults) DeepCopyInto(out *PodSecurityDefaults) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityDefaults.
func (in *PodSecurityDefaults) DeepCopy() *PodSecurityDefaults {
	if in == nil {
		return nil
	}
	out := new(PodSecurityDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityExemptions) DeepCopyInto(out *PodSecurityExemptions) {
	*out = *in
	if in.Usernames != nil {
		in, out := &in.Usernames, &out.Usernames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RuntimeClasses != nil {
		in, out := &in.RuntimeClasses, &out.RuntimeClasses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityExemptions.
func (in *PodSecurityExemptions) DeepCopy() *PodSecurityExemptions {
	if in == nil {
		return nil
	}
	out := new(PodSecurityExemptions)
	in.DeepCopyInto(out)
	return out
}
