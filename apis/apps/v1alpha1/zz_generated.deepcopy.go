// +build !ignore_autogenerated

/*
Copyright 2020 Red Hat.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIcast) DeepCopyInto(out *APIcast) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIcast.
func (in *APIcast) DeepCopy() *APIcast {
	if in == nil {
		return nil
	}
	out := new(APIcast)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIcast) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIcastCondition) DeepCopyInto(out *APIcastCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIcastCondition.
func (in *APIcastCondition) DeepCopy() *APIcastCondition {
	if in == nil {
		return nil
	}
	out := new(APIcastCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIcastExposedHost) DeepCopyInto(out *APIcastExposedHost) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]v1beta1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIcastExposedHost.
func (in *APIcastExposedHost) DeepCopy() *APIcastExposedHost {
	if in == nil {
		return nil
	}
	out := new(APIcastExposedHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIcastList) DeepCopyInto(out *APIcastList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIcast, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIcastList.
func (in *APIcastList) DeepCopy() *APIcastList {
	if in == nil {
		return nil
	}
	out := new(APIcastList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIcastList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIcastSpec) DeepCopyInto(out *APIcastSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int64)
		**out = **in
	}
	if in.AdminPortalCredentialsRef != nil {
		in, out := &in.AdminPortalCredentialsRef, &out.AdminPortalCredentialsRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.EmbeddedConfigurationSecretRef != nil {
		in, out := &in.EmbeddedConfigurationSecretRef, &out.EmbeddedConfigurationSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ExposedHost != nil {
		in, out := &in.ExposedHost, &out.ExposedHost
		*out = new(APIcastExposedHost)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentEnvironment != nil {
		in, out := &in.DeploymentEnvironment, &out.DeploymentEnvironment
		*out = new(DeploymentEnvironmentType)
		**out = **in
	}
	if in.DNSResolverAddress != nil {
		in, out := &in.DNSResolverAddress, &out.DNSResolverAddress
		*out = new(string)
		**out = **in
	}
	if in.EnabledServices != nil {
		in, out := &in.EnabledServices, &out.EnabledServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigurationLoadMode != nil {
		in, out := &in.ConfigurationLoadMode, &out.ConfigurationLoadMode
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.PathRoutingEnabled != nil {
		in, out := &in.PathRoutingEnabled, &out.PathRoutingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResponseCodesIncluded != nil {
		in, out := &in.ResponseCodesIncluded, &out.ResponseCodesIncluded
		*out = new(bool)
		**out = **in
	}
	if in.CacheConfigurationSeconds != nil {
		in, out := &in.CacheConfigurationSeconds, &out.CacheConfigurationSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ManagementAPIScope != nil {
		in, out := &in.ManagementAPIScope, &out.ManagementAPIScope
		*out = new(string)
		**out = **in
	}
	if in.OpenSSLPeerVerificationEnabled != nil {
		in, out := &in.OpenSSLPeerVerificationEnabled, &out.OpenSSLPeerVerificationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.UpstreamRetryCases != nil {
		in, out := &in.UpstreamRetryCases, &out.UpstreamRetryCases
		*out = new(string)
		**out = **in
	}
	if in.CacheMaxTime != nil {
		in, out := &in.CacheMaxTime, &out.CacheMaxTime
		*out = new(string)
		**out = **in
	}
	if in.CacheStatusCodes != nil {
		in, out := &in.CacheStatusCodes, &out.CacheStatusCodes
		*out = new(string)
		**out = **in
	}
	if in.OidcLogLevel != nil {
		in, out := &in.OidcLogLevel, &out.OidcLogLevel
		*out = new(string)
		**out = **in
	}
	if in.LoadServicesWhenNeeded != nil {
		in, out := &in.LoadServicesWhenNeeded, &out.LoadServicesWhenNeeded
		*out = new(bool)
		**out = **in
	}
	if in.ServicesFilterByURL != nil {
		in, out := &in.ServicesFilterByURL, &out.ServicesFilterByURL
		*out = new(string)
		**out = **in
	}
	if in.ServiceConfigurationVersionOverride != nil {
		in, out := &in.ServiceConfigurationVersionOverride, &out.ServiceConfigurationVersionOverride
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HTTPSPort != nil {
		in, out := &in.HTTPSPort, &out.HTTPSPort
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIcastSpec.
func (in *APIcastSpec) DeepCopy() *APIcastSpec {
	if in == nil {
		return nil
	}
	out := new(APIcastSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIcastStatus) DeepCopyInto(out *APIcastStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]APIcastCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIcastStatus.
func (in *APIcastStatus) DeepCopy() *APIcastStatus {
	if in == nil {
		return nil
	}
	out := new(APIcastStatus)
	in.DeepCopyInto(out)
	return out
}
