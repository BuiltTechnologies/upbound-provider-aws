/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ds/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BotAssociation.
func (mg *BotAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContactFlow.
func (mg *ContactFlow) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ContactFlowModule.
func (mg *ContactFlowModule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HoursOfOperation.
func (mg *HoursOfOperation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DirectoryID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DirectoryIDRef,
		Selector:     mg.Spec.ForProvider.DirectoryIDSelector,
		To: reference.To{
			List:    &v1beta1.DirectoryList{},
			Managed: &v1beta1.Directory{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DirectoryID")
	}
	mg.Spec.ForProvider.DirectoryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DirectoryIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LambdaFunctionAssociation.
func (mg *LambdaFunctionAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.FunctionArnRef,
		Selector:     mg.Spec.ForProvider.FunctionArnSelector,
		To: reference.To{
			List:    &v1beta11.FunctionList{},
			Managed: &v1beta11.Function{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionArn")
	}
	mg.Spec.ForProvider.FunctionArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Queue.
func (mg *Queue) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HoursOfOperationID),
		Extract:      resource.ExtractParamPath("hours_of_operation_id", true),
		Reference:    mg.Spec.ForProvider.HoursOfOperationIDRef,
		Selector:     mg.Spec.ForProvider.HoursOfOperationIDSelector,
		To: reference.To{
			List:    &HoursOfOperationList{},
			Managed: &HoursOfOperation{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HoursOfOperationID")
	}
	mg.Spec.ForProvider.HoursOfOperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HoursOfOperationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QuickConnect.
func (mg *QuickConnect) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoutingProfile.
func (mg *RoutingProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultOutboundQueueID),
		Extract:      resource.ExtractParamPath("queue_id", true),
		Reference:    mg.Spec.ForProvider.DefaultOutboundQueueIDRef,
		Selector:     mg.Spec.ForProvider.DefaultOutboundQueueIDSelector,
		To: reference.To{
			List:    &QueueList{},
			Managed: &Queue{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultOutboundQueueID")
	}
	mg.Spec.ForProvider.DefaultOutboundQueueID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DefaultOutboundQueueIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityProfile.
func (mg *SecurityProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserHierarchyStructure.
func (mg *UserHierarchyStructure) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}
