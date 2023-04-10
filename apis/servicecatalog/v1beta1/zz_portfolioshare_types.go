/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PortfolioShareObservation struct {

	// Whether the shared portfolio is imported by the recipient account. If the recipient is organizational, the share is automatically imported, and the field is always set to true.
	Accepted *bool `json:"accepted,omitempty" tf:"accepted,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PortfolioShareParameters struct {

	// Language code. Valid values: en (English), jp (Japanese), zh (Chinese). Default value is en.
	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// Portfolio identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/servicecatalog/v1beta1.Portfolio
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PortfolioID *string `json:"portfolioId,omitempty" tf:"portfolio_id,omitempty"`

	// Reference to a Portfolio in servicecatalog to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDRef *v1.Reference `json:"portfolioIdRef,omitempty" tf:"-"`

	// Selector for a Portfolio in servicecatalog to populate portfolioId.
	// +kubebuilder:validation:Optional
	PortfolioIDSelector *v1.Selector `json:"portfolioIdSelector,omitempty" tf:"-"`

	// Identifier of the principal with whom you will share the portfolio. Valid values AWS account IDs and ARNs of AWS Organizations and organizational units.
	// +kubebuilder:validation:Required
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Enables or disables Principal sharing when creating the portfolio share. If this flag is not provided, principal sharing is disabled.
	// +kubebuilder:validation:Optional
	SharePrincipals *bool `json:"sharePrincipals,omitempty" tf:"share_principals,omitempty"`

	// Whether to enable sharing of aws_servicecatalog_tag_option resources when creating the portfolio share.
	// +kubebuilder:validation:Optional
	ShareTagOptions *bool `json:"shareTagOptions,omitempty" tf:"share_tag_options,omitempty"`

	// Type of portfolio share. Valid values are ACCOUNT (an external account), ORGANIZATION (a share to every account in an organization), ORGANIZATIONAL_UNIT, ORGANIZATION_MEMBER_ACCOUNT (a share to an account in an organization).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Whether to wait (up to the timeout) for the share to be accepted. Organizational shares are automatically accepted.
	// +kubebuilder:validation:Optional
	WaitForAcceptance *bool `json:"waitForAcceptance,omitempty" tf:"wait_for_acceptance,omitempty"`
}

// PortfolioShareSpec defines the desired state of PortfolioShare
type PortfolioShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortfolioShareParameters `json:"forProvider"`
}

// PortfolioShareStatus defines the observed state of PortfolioShare.
type PortfolioShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortfolioShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PortfolioShare is the Schema for the PortfolioShares API. Manages a Service Catalog Portfolio Share
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PortfolioShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortfolioShareSpec   `json:"spec"`
	Status            PortfolioShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortfolioShareList contains a list of PortfolioShares
type PortfolioShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortfolioShare `json:"items"`
}

// Repository type metadata.
var (
	PortfolioShare_Kind             = "PortfolioShare"
	PortfolioShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PortfolioShare_Kind}.String()
	PortfolioShare_KindAPIVersion   = PortfolioShare_Kind + "." + CRDGroupVersion.String()
	PortfolioShare_GroupVersionKind = CRDGroupVersion.WithKind(PortfolioShare_Kind)
)

func init() {
	SchemeBuilder.Register(&PortfolioShare{}, &PortfolioShareList{})
}