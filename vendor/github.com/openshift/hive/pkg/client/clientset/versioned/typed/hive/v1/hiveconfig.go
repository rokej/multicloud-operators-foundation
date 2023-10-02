// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/hive/apis/hive/v1"
	hivev1 "github.com/openshift/hive/pkg/client/applyconfiguration/hive/v1"
	scheme "github.com/openshift/hive/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HiveConfigsGetter has a method to return a HiveConfigInterface.
// A group's client should implement this interface.
type HiveConfigsGetter interface {
	HiveConfigs() HiveConfigInterface
}

// HiveConfigInterface has methods to work with HiveConfig resources.
type HiveConfigInterface interface {
	Create(ctx context.Context, hiveConfig *v1.HiveConfig, opts metav1.CreateOptions) (*v1.HiveConfig, error)
	Update(ctx context.Context, hiveConfig *v1.HiveConfig, opts metav1.UpdateOptions) (*v1.HiveConfig, error)
	UpdateStatus(ctx context.Context, hiveConfig *v1.HiveConfig, opts metav1.UpdateOptions) (*v1.HiveConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.HiveConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HiveConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HiveConfig, err error)
	Apply(ctx context.Context, hiveConfig *hivev1.HiveConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HiveConfig, err error)
	ApplyStatus(ctx context.Context, hiveConfig *hivev1.HiveConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HiveConfig, err error)
	HiveConfigExpansion
}

// hiveConfigs implements HiveConfigInterface
type hiveConfigs struct {
	client rest.Interface
}

// newHiveConfigs returns a HiveConfigs
func newHiveConfigs(c *HiveV1Client) *hiveConfigs {
	return &hiveConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the hiveConfig, and returns the corresponding hiveConfig object, and an error if there is any.
func (c *hiveConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HiveConfig, err error) {
	result = &v1.HiveConfig{}
	err = c.client.Get().
		Resource("hiveconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HiveConfigs that match those selectors.
func (c *hiveConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HiveConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.HiveConfigList{}
	err = c.client.Get().
		Resource("hiveconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hiveConfigs.
func (c *hiveConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("hiveconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hiveConfig and creates it.  Returns the server's representation of the hiveConfig, and an error, if there is any.
func (c *hiveConfigs) Create(ctx context.Context, hiveConfig *v1.HiveConfig, opts metav1.CreateOptions) (result *v1.HiveConfig, err error) {
	result = &v1.HiveConfig{}
	err = c.client.Post().
		Resource("hiveconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hiveConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hiveConfig and updates it. Returns the server's representation of the hiveConfig, and an error, if there is any.
func (c *hiveConfigs) Update(ctx context.Context, hiveConfig *v1.HiveConfig, opts metav1.UpdateOptions) (result *v1.HiveConfig, err error) {
	result = &v1.HiveConfig{}
	err = c.client.Put().
		Resource("hiveconfigs").
		Name(hiveConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hiveConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hiveConfigs) UpdateStatus(ctx context.Context, hiveConfig *v1.HiveConfig, opts metav1.UpdateOptions) (result *v1.HiveConfig, err error) {
	result = &v1.HiveConfig{}
	err = c.client.Put().
		Resource("hiveconfigs").
		Name(hiveConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hiveConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hiveConfig and deletes it. Returns an error if one occurs.
func (c *hiveConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("hiveconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hiveConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("hiveconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hiveConfig.
func (c *hiveConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HiveConfig, err error) {
	result = &v1.HiveConfig{}
	err = c.client.Patch(pt).
		Resource("hiveconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied hiveConfig.
func (c *hiveConfigs) Apply(ctx context.Context, hiveConfig *hivev1.HiveConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HiveConfig, err error) {
	if hiveConfig == nil {
		return nil, fmt.Errorf("hiveConfig provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(hiveConfig)
	if err != nil {
		return nil, err
	}
	name := hiveConfig.Name
	if name == nil {
		return nil, fmt.Errorf("hiveConfig.Name must be provided to Apply")
	}
	result = &v1.HiveConfig{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("hiveconfigs").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *hiveConfigs) ApplyStatus(ctx context.Context, hiveConfig *hivev1.HiveConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HiveConfig, err error) {
	if hiveConfig == nil {
		return nil, fmt.Errorf("hiveConfig provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(hiveConfig)
	if err != nil {
		return nil, err
	}

	name := hiveConfig.Name
	if name == nil {
		return nil, fmt.Errorf("hiveConfig.Name must be provided to Apply")
	}

	result = &v1.HiveConfig{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("hiveconfigs").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
