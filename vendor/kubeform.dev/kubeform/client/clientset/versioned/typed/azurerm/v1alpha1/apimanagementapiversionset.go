/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApiManagementAPIVersionSetsGetter has a method to return a ApiManagementAPIVersionSetInterface.
// A group's client should implement this interface.
type ApiManagementAPIVersionSetsGetter interface {
	ApiManagementAPIVersionSets(namespace string) ApiManagementAPIVersionSetInterface
}

// ApiManagementAPIVersionSetInterface has methods to work with ApiManagementAPIVersionSet resources.
type ApiManagementAPIVersionSetInterface interface {
	Create(*v1alpha1.ApiManagementAPIVersionSet) (*v1alpha1.ApiManagementAPIVersionSet, error)
	Update(*v1alpha1.ApiManagementAPIVersionSet) (*v1alpha1.ApiManagementAPIVersionSet, error)
	UpdateStatus(*v1alpha1.ApiManagementAPIVersionSet) (*v1alpha1.ApiManagementAPIVersionSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementAPIVersionSet, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementAPIVersionSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAPIVersionSet, err error)
	ApiManagementAPIVersionSetExpansion
}

// apiManagementAPIVersionSets implements ApiManagementAPIVersionSetInterface
type apiManagementAPIVersionSets struct {
	client rest.Interface
	ns     string
}

// newApiManagementAPIVersionSets returns a ApiManagementAPIVersionSets
func newApiManagementAPIVersionSets(c *AzurermV1alpha1Client, namespace string) *apiManagementAPIVersionSets {
	return &apiManagementAPIVersionSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementAPIVersionSet, and returns the corresponding apiManagementAPIVersionSet object, and an error if there is any.
func (c *apiManagementAPIVersionSets) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementAPIVersionSet, err error) {
	result = &v1alpha1.ApiManagementAPIVersionSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementAPIVersionSets that match those selectors.
func (c *apiManagementAPIVersionSets) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementAPIVersionSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementAPIVersionSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementAPIVersionSets.
func (c *apiManagementAPIVersionSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementAPIVersionSet and creates it.  Returns the server's representation of the apiManagementAPIVersionSet, and an error, if there is any.
func (c *apiManagementAPIVersionSets) Create(apiManagementAPIVersionSet *v1alpha1.ApiManagementAPIVersionSet) (result *v1alpha1.ApiManagementAPIVersionSet, err error) {
	result = &v1alpha1.ApiManagementAPIVersionSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		Body(apiManagementAPIVersionSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementAPIVersionSet and updates it. Returns the server's representation of the apiManagementAPIVersionSet, and an error, if there is any.
func (c *apiManagementAPIVersionSets) Update(apiManagementAPIVersionSet *v1alpha1.ApiManagementAPIVersionSet) (result *v1alpha1.ApiManagementAPIVersionSet, err error) {
	result = &v1alpha1.ApiManagementAPIVersionSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		Name(apiManagementAPIVersionSet.Name).
		Body(apiManagementAPIVersionSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementAPIVersionSets) UpdateStatus(apiManagementAPIVersionSet *v1alpha1.ApiManagementAPIVersionSet) (result *v1alpha1.ApiManagementAPIVersionSet, err error) {
	result = &v1alpha1.ApiManagementAPIVersionSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		Name(apiManagementAPIVersionSet.Name).
		SubResource("status").
		Body(apiManagementAPIVersionSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementAPIVersionSet and deletes it. Returns an error if one occurs.
func (c *apiManagementAPIVersionSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementAPIVersionSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementAPIVersionSet.
func (c *apiManagementAPIVersionSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAPIVersionSet, err error) {
	result = &v1alpha1.ApiManagementAPIVersionSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementapiversionsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
