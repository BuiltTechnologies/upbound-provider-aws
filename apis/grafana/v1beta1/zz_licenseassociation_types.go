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

type LicenseAssociationObservation struct {

	// If license_type is set to ENTERPRISE_FREE_TRIAL, this is the expiration date of the free trial.
	FreeTrialExpiration *string `json:"freeTrialExpiration,omitempty" tf:"free_trial_expiration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If license_type is set to ENTERPRISE, this is the expiration date of the enterprise license.
	LicenseExpiration *string `json:"licenseExpiration,omitempty" tf:"license_expiration,omitempty"`
}

type LicenseAssociationParameters struct {

	// The type of license for the workspace license association. Valid values are ENTERPRISE and ENTERPRISE_FREE_TRIAL.
	// +kubebuilder:validation:Required
	LicenseType *string `json:"licenseType" tf:"license_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The workspace id.
	// +kubebuilder:validation:Required
	WorkspaceID *string `json:"workspaceId" tf:"workspace_id,omitempty"`
}

// LicenseAssociationSpec defines the desired state of LicenseAssociation
type LicenseAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LicenseAssociationParameters `json:"forProvider"`
}

// LicenseAssociationStatus defines the observed state of LicenseAssociation.
type LicenseAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LicenseAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseAssociation is the Schema for the LicenseAssociations API. Provides an Amazon Managed Grafana workspace license association resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LicenseAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicenseAssociationSpec   `json:"spec"`
	Status            LicenseAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseAssociationList contains a list of LicenseAssociations
type LicenseAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseAssociation `json:"items"`
}

// Repository type metadata.
var (
	LicenseAssociation_Kind             = "LicenseAssociation"
	LicenseAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LicenseAssociation_Kind}.String()
	LicenseAssociation_KindAPIVersion   = LicenseAssociation_Kind + "." + CRDGroupVersion.String()
	LicenseAssociation_GroupVersionKind = CRDGroupVersion.WithKind(LicenseAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&LicenseAssociation{}, &LicenseAssociationList{})
}
