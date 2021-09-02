// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "k8c.io/kubermatic/v2/pkg/crd/client/clientset/versioned/scheme"
	v1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MLAAdminSettingsGetter has a method to return a MLAAdminSettingInterface.
// A group's client should implement this interface.
type MLAAdminSettingsGetter interface {
	MLAAdminSettings(namespace string) MLAAdminSettingInterface
}

// MLAAdminSettingInterface has methods to work with MLAAdminSetting resources.
type MLAAdminSettingInterface interface {
	Create(ctx context.Context, mLAAdminSetting *v1.MLAAdminSetting, opts metav1.CreateOptions) (*v1.MLAAdminSetting, error)
	Update(ctx context.Context, mLAAdminSetting *v1.MLAAdminSetting, opts metav1.UpdateOptions) (*v1.MLAAdminSetting, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MLAAdminSetting, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MLAAdminSettingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MLAAdminSetting, err error)
	MLAAdminSettingExpansion
}

// mLAAdminSettings implements MLAAdminSettingInterface
type mLAAdminSettings struct {
	client rest.Interface
	ns     string
}

// newMLAAdminSettings returns a MLAAdminSettings
func newMLAAdminSettings(c *KubermaticV1Client, namespace string) *mLAAdminSettings {
	return &mLAAdminSettings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mLAAdminSetting, and returns the corresponding mLAAdminSetting object, and an error if there is any.
func (c *mLAAdminSettings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MLAAdminSetting, err error) {
	result = &v1.MLAAdminSetting{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlaadminsettings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MLAAdminSettings that match those selectors.
func (c *mLAAdminSettings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MLAAdminSettingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MLAAdminSettingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlaadminsettings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mLAAdminSettings.
func (c *mLAAdminSettings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mlaadminsettings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mLAAdminSetting and creates it.  Returns the server's representation of the mLAAdminSetting, and an error, if there is any.
func (c *mLAAdminSettings) Create(ctx context.Context, mLAAdminSetting *v1.MLAAdminSetting, opts metav1.CreateOptions) (result *v1.MLAAdminSetting, err error) {
	result = &v1.MLAAdminSetting{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mlaadminsettings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mLAAdminSetting).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mLAAdminSetting and updates it. Returns the server's representation of the mLAAdminSetting, and an error, if there is any.
func (c *mLAAdminSettings) Update(ctx context.Context, mLAAdminSetting *v1.MLAAdminSetting, opts metav1.UpdateOptions) (result *v1.MLAAdminSetting, err error) {
	result = &v1.MLAAdminSetting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mlaadminsettings").
		Name(mLAAdminSetting.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mLAAdminSetting).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mLAAdminSetting and deletes it. Returns an error if one occurs.
func (c *mLAAdminSettings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlaadminsettings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mLAAdminSettings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlaadminsettings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mLAAdminSetting.
func (c *mLAAdminSettings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MLAAdminSetting, err error) {
	result = &v1.MLAAdminSetting{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mlaadminsettings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
