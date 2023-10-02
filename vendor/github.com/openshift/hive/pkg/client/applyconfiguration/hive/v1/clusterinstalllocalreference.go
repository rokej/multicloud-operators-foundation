// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ClusterInstallLocalReferenceApplyConfiguration represents an declarative configuration of the ClusterInstallLocalReference type for use
// with apply.
type ClusterInstallLocalReferenceApplyConfiguration struct {
	Group   *string `json:"group,omitempty"`
	Version *string `json:"version,omitempty"`
	Kind    *string `json:"kind,omitempty"`
	Name    *string `json:"name,omitempty"`
}

// ClusterInstallLocalReferenceApplyConfiguration constructs an declarative configuration of the ClusterInstallLocalReference type for use with
// apply.
func ClusterInstallLocalReference() *ClusterInstallLocalReferenceApplyConfiguration {
	return &ClusterInstallLocalReferenceApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *ClusterInstallLocalReferenceApplyConfiguration) WithGroup(value string) *ClusterInstallLocalReferenceApplyConfiguration {
	b.Group = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ClusterInstallLocalReferenceApplyConfiguration) WithVersion(value string) *ClusterInstallLocalReferenceApplyConfiguration {
	b.Version = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ClusterInstallLocalReferenceApplyConfiguration) WithKind(value string) *ClusterInstallLocalReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ClusterInstallLocalReferenceApplyConfiguration) WithName(value string) *ClusterInstallLocalReferenceApplyConfiguration {
	b.Name = &value
	return b
}
