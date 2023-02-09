//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociation) DeepCopyInto(out *LicenseAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociation.
func (in *LicenseAssociation) DeepCopy() *LicenseAssociation {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationList) DeepCopyInto(out *LicenseAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationList.
func (in *LicenseAssociationList) DeepCopy() *LicenseAssociationList {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationObservation) DeepCopyInto(out *LicenseAssociationObservation) {
	*out = *in
	if in.FreeTrialExpiration != nil {
		in, out := &in.FreeTrialExpiration, &out.FreeTrialExpiration
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LicenseExpiration != nil {
		in, out := &in.LicenseExpiration, &out.LicenseExpiration
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationObservation.
func (in *LicenseAssociationObservation) DeepCopy() *LicenseAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationParameters) DeepCopyInto(out *LicenseAssociationParameters) {
	*out = *in
	if in.LicenseType != nil {
		in, out := &in.LicenseType, &out.LicenseType
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceID != nil {
		in, out := &in.WorkspaceID, &out.WorkspaceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationParameters.
func (in *LicenseAssociationParameters) DeepCopy() *LicenseAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationSpec) DeepCopyInto(out *LicenseAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationSpec.
func (in *LicenseAssociationSpec) DeepCopy() *LicenseAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAssociationStatus) DeepCopyInto(out *LicenseAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAssociationStatus.
func (in *LicenseAssociationStatus) DeepCopy() *LicenseAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseAssociationStatus)
	in.DeepCopyInto(out)
	return out
}
