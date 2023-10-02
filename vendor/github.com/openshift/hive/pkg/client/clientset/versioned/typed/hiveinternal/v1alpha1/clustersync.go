// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/openshift/hive/apis/hiveinternal/v1alpha1"
	hiveinternalv1alpha1 "github.com/openshift/hive/pkg/client/applyconfiguration/hiveinternal/v1alpha1"
	scheme "github.com/openshift/hive/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterSyncsGetter has a method to return a ClusterSyncInterface.
// A group's client should implement this interface.
type ClusterSyncsGetter interface {
	ClusterSyncs(namespace string) ClusterSyncInterface
}

// ClusterSyncInterface has methods to work with ClusterSync resources.
type ClusterSyncInterface interface {
	Create(ctx context.Context, clusterSync *v1alpha1.ClusterSync, opts v1.CreateOptions) (*v1alpha1.ClusterSync, error)
	Update(ctx context.Context, clusterSync *v1alpha1.ClusterSync, opts v1.UpdateOptions) (*v1alpha1.ClusterSync, error)
	UpdateStatus(ctx context.Context, clusterSync *v1alpha1.ClusterSync, opts v1.UpdateOptions) (*v1alpha1.ClusterSync, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterSync, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterSyncList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSync, err error)
	Apply(ctx context.Context, clusterSync *hiveinternalv1alpha1.ClusterSyncApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSync, err error)
	ApplyStatus(ctx context.Context, clusterSync *hiveinternalv1alpha1.ClusterSyncApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSync, err error)
	ClusterSyncExpansion
}

// clusterSyncs implements ClusterSyncInterface
type clusterSyncs struct {
	client rest.Interface
	ns     string
}

// newClusterSyncs returns a ClusterSyncs
func newClusterSyncs(c *HiveinternalV1alpha1Client, namespace string) *clusterSyncs {
	return &clusterSyncs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterSync, and returns the corresponding clusterSync object, and an error if there is any.
func (c *clusterSyncs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterSync, err error) {
	result = &v1alpha1.ClusterSync{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustersyncs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterSyncs that match those selectors.
func (c *clusterSyncs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterSyncList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterSyncList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustersyncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterSyncs.
func (c *clusterSyncs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clustersyncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterSync and creates it.  Returns the server's representation of the clusterSync, and an error, if there is any.
func (c *clusterSyncs) Create(ctx context.Context, clusterSync *v1alpha1.ClusterSync, opts v1.CreateOptions) (result *v1alpha1.ClusterSync, err error) {
	result = &v1alpha1.ClusterSync{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clustersyncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSync).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterSync and updates it. Returns the server's representation of the clusterSync, and an error, if there is any.
func (c *clusterSyncs) Update(ctx context.Context, clusterSync *v1alpha1.ClusterSync, opts v1.UpdateOptions) (result *v1alpha1.ClusterSync, err error) {
	result = &v1alpha1.ClusterSync{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustersyncs").
		Name(clusterSync.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSync).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterSyncs) UpdateStatus(ctx context.Context, clusterSync *v1alpha1.ClusterSync, opts v1.UpdateOptions) (result *v1alpha1.ClusterSync, err error) {
	result = &v1alpha1.ClusterSync{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustersyncs").
		Name(clusterSync.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSync).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterSync and deletes it. Returns an error if one occurs.
func (c *clusterSyncs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustersyncs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterSyncs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustersyncs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterSync.
func (c *clusterSyncs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSync, err error) {
	result = &v1alpha1.ClusterSync{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clustersyncs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterSync.
func (c *clusterSyncs) Apply(ctx context.Context, clusterSync *hiveinternalv1alpha1.ClusterSyncApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSync, err error) {
	if clusterSync == nil {
		return nil, fmt.Errorf("clusterSync provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterSync)
	if err != nil {
		return nil, err
	}
	name := clusterSync.Name
	if name == nil {
		return nil, fmt.Errorf("clusterSync.Name must be provided to Apply")
	}
	result = &v1alpha1.ClusterSync{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("clustersyncs").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *clusterSyncs) ApplyStatus(ctx context.Context, clusterSync *hiveinternalv1alpha1.ClusterSyncApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ClusterSync, err error) {
	if clusterSync == nil {
		return nil, fmt.Errorf("clusterSync provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterSync)
	if err != nil {
		return nil, err
	}

	name := clusterSync.Name
	if name == nil {
		return nil, fmt.Errorf("clusterSync.Name must be provided to Apply")
	}

	result = &v1alpha1.ClusterSync{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("clustersyncs").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
