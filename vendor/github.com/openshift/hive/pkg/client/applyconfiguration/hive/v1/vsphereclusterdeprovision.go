// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// VSphereClusterDeprovisionApplyConfiguration represents an declarative configuration of the VSphereClusterDeprovision type for use
// with apply.
type VSphereClusterDeprovisionApplyConfiguration struct {
	CredentialsSecretRef  *v1.LocalObjectReference `json:"credentialsSecretRef,omitempty"`
	CertificatesSecretRef *v1.LocalObjectReference `json:"certificatesSecretRef,omitempty"`
	VCenter               *string                  `json:"vCenter,omitempty"`
}

// VSphereClusterDeprovisionApplyConfiguration constructs an declarative configuration of the VSphereClusterDeprovision type for use with
// apply.
func VSphereClusterDeprovision() *VSphereClusterDeprovisionApplyConfiguration {
	return &VSphereClusterDeprovisionApplyConfiguration{}
}

// WithCredentialsSecretRef sets the CredentialsSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CredentialsSecretRef field is set to the value of the last call.
func (b *VSphereClusterDeprovisionApplyConfiguration) WithCredentialsSecretRef(value v1.LocalObjectReference) *VSphereClusterDeprovisionApplyConfiguration {
	b.CredentialsSecretRef = &value
	return b
}

// WithCertificatesSecretRef sets the CertificatesSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CertificatesSecretRef field is set to the value of the last call.
func (b *VSphereClusterDeprovisionApplyConfiguration) WithCertificatesSecretRef(value v1.LocalObjectReference) *VSphereClusterDeprovisionApplyConfiguration {
	b.CertificatesSecretRef = &value
	return b
}

// WithVCenter sets the VCenter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VCenter field is set to the value of the last call.
func (b *VSphereClusterDeprovisionApplyConfiguration) WithVCenter(value string) *VSphereClusterDeprovisionApplyConfiguration {
	b.VCenter = &value
	return b
}
