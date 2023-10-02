// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apishivev1 "github.com/openshift/hive/apis/hive/v1"
	v1 "k8s.io/api/core/v1"
)

// AWSPrivateLinkConfigApplyConfiguration represents an declarative configuration of the AWSPrivateLinkConfig type for use
// with apply.
type AWSPrivateLinkConfigApplyConfiguration struct {
	CredentialsSecretRef *v1.LocalObjectReference                    `json:"credentialsSecretRef,omitempty"`
	EndpointVPCInventory []AWSPrivateLinkInventoryApplyConfiguration `json:"endpointVPCInventory,omitempty"`
	AssociatedVPCs       []AWSAssociatedVPCApplyConfiguration        `json:"associatedVPCs,omitempty"`
	DNSRecordType        *apishivev1.AWSPrivateLinkDNSRecordType     `json:"dnsRecordType,omitempty"`
}

// AWSPrivateLinkConfigApplyConfiguration constructs an declarative configuration of the AWSPrivateLinkConfig type for use with
// apply.
func AWSPrivateLinkConfig() *AWSPrivateLinkConfigApplyConfiguration {
	return &AWSPrivateLinkConfigApplyConfiguration{}
}

// WithCredentialsSecretRef sets the CredentialsSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CredentialsSecretRef field is set to the value of the last call.
func (b *AWSPrivateLinkConfigApplyConfiguration) WithCredentialsSecretRef(value v1.LocalObjectReference) *AWSPrivateLinkConfigApplyConfiguration {
	b.CredentialsSecretRef = &value
	return b
}

// WithEndpointVPCInventory adds the given value to the EndpointVPCInventory field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EndpointVPCInventory field.
func (b *AWSPrivateLinkConfigApplyConfiguration) WithEndpointVPCInventory(values ...*AWSPrivateLinkInventoryApplyConfiguration) *AWSPrivateLinkConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEndpointVPCInventory")
		}
		b.EndpointVPCInventory = append(b.EndpointVPCInventory, *values[i])
	}
	return b
}

// WithAssociatedVPCs adds the given value to the AssociatedVPCs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AssociatedVPCs field.
func (b *AWSPrivateLinkConfigApplyConfiguration) WithAssociatedVPCs(values ...*AWSAssociatedVPCApplyConfiguration) *AWSPrivateLinkConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAssociatedVPCs")
		}
		b.AssociatedVPCs = append(b.AssociatedVPCs, *values[i])
	}
	return b
}

// WithDNSRecordType sets the DNSRecordType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSRecordType field is set to the value of the last call.
func (b *AWSPrivateLinkConfigApplyConfiguration) WithDNSRecordType(value apishivev1.AWSPrivateLinkDNSRecordType) *AWSPrivateLinkConfigApplyConfiguration {
	b.DNSRecordType = &value
	return b
}
