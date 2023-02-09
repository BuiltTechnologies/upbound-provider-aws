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
func (in *EBSBlockDeviceObservation) DeepCopyInto(out *EBSBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBSBlockDeviceObservation.
func (in *EBSBlockDeviceObservation) DeepCopy() *EBSBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(EBSBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EBSBlockDeviceParameters) DeepCopyInto(out *EBSBlockDeviceParameters) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
		*out = new(float64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(float64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBSBlockDeviceParameters.
func (in *EBSBlockDeviceParameters) DeepCopy() *EBSBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(EBSBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralBlockDeviceObservation) DeepCopyInto(out *EphemeralBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralBlockDeviceObservation.
func (in *EphemeralBlockDeviceObservation) DeepCopy() *EphemeralBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(EphemeralBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralBlockDeviceParameters) DeepCopyInto(out *EphemeralBlockDeviceParameters) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(bool)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralBlockDeviceParameters.
func (in *EphemeralBlockDeviceParameters) DeepCopy() *EphemeralBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(EphemeralBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfiguration) DeepCopyInto(out *LaunchConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfiguration.
func (in *LaunchConfiguration) DeepCopy() *LaunchConfiguration {
	if in == nil {
		return nil
	}
	out := new(LaunchConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationList) DeepCopyInto(out *LaunchConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LaunchConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationList.
func (in *LaunchConfigurationList) DeepCopy() *LaunchConfigurationList {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationObservation) DeepCopyInto(out *LaunchConfigurationObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationObservation.
func (in *LaunchConfigurationObservation) DeepCopy() *LaunchConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationParameters) DeepCopyInto(out *LaunchConfigurationParameters) {
	*out = *in
	if in.AssociatePublicIPAddress != nil {
		in, out := &in.AssociatePublicIPAddress, &out.AssociatePublicIPAddress
		*out = new(bool)
		**out = **in
	}
	if in.EBSBlockDevice != nil {
		in, out := &in.EBSBlockDevice, &out.EBSBlockDevice
		*out = make([]EBSBlockDeviceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EBSOptimized != nil {
		in, out := &in.EBSOptimized, &out.EBSOptimized
		*out = new(bool)
		**out = **in
	}
	if in.EnableMonitoring != nil {
		in, out := &in.EnableMonitoring, &out.EnableMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]EphemeralBlockDeviceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IAMInstanceProfile != nil {
		in, out := &in.IAMInstanceProfile, &out.IAMInstanceProfile
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.MetadataOptions != nil {
		in, out := &in.MetadataOptions, &out.MetadataOptions
		*out = make([]MetadataOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlacementTenancy != nil {
		in, out := &in.PlacementTenancy, &out.PlacementTenancy
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RootBlockDevice != nil {
		in, out := &in.RootBlockDevice, &out.RootBlockDevice
		*out = make([]RootBlockDeviceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SpotPrice != nil {
		in, out := &in.SpotPrice, &out.SpotPrice
		*out = new(string)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.UserDataBase64 != nil {
		in, out := &in.UserDataBase64, &out.UserDataBase64
		*out = new(string)
		**out = **in
	}
	if in.VPCClassicLinkID != nil {
		in, out := &in.VPCClassicLinkID, &out.VPCClassicLinkID
		*out = new(string)
		**out = **in
	}
	if in.VPCClassicLinkSecurityGroups != nil {
		in, out := &in.VPCClassicLinkSecurityGroups, &out.VPCClassicLinkSecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationParameters.
func (in *LaunchConfigurationParameters) DeepCopy() *LaunchConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationSpec) DeepCopyInto(out *LaunchConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationSpec.
func (in *LaunchConfigurationSpec) DeepCopy() *LaunchConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchConfigurationStatus) DeepCopyInto(out *LaunchConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchConfigurationStatus.
func (in *LaunchConfigurationStatus) DeepCopy() *LaunchConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(LaunchConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptionsObservation) DeepCopyInto(out *MetadataOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptionsObservation.
func (in *MetadataOptionsObservation) DeepCopy() *MetadataOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(MetadataOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptionsParameters) DeepCopyInto(out *MetadataOptionsParameters) {
	*out = *in
	if in.HTTPEndpoint != nil {
		in, out := &in.HTTPEndpoint, &out.HTTPEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HTTPPutResponseHopLimit != nil {
		in, out := &in.HTTPPutResponseHopLimit, &out.HTTPPutResponseHopLimit
		*out = new(float64)
		**out = **in
	}
	if in.HTTPTokens != nil {
		in, out := &in.HTTPTokens, &out.HTTPTokens
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptionsParameters.
func (in *MetadataOptionsParameters) DeepCopy() *MetadataOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(MetadataOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootBlockDeviceObservation) DeepCopyInto(out *RootBlockDeviceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootBlockDeviceObservation.
func (in *RootBlockDeviceObservation) DeepCopy() *RootBlockDeviceObservation {
	if in == nil {
		return nil
	}
	out := new(RootBlockDeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootBlockDeviceParameters) DeepCopyInto(out *RootBlockDeviceParameters) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
		*out = new(float64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(float64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootBlockDeviceParameters.
func (in *RootBlockDeviceParameters) DeepCopy() *RootBlockDeviceParameters {
	if in == nil {
		return nil
	}
	out := new(RootBlockDeviceParameters)
	in.DeepCopyInto(out)
	return out
}
