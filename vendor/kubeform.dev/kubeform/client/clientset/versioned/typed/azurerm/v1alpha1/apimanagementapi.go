/*
Copyright 2019 The Kubeform Authors.

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

// ApiManagementAPIsGetter has a method to return a ApiManagementAPIInterface.
// A group's client should implement this interface.
type ApiManagementAPIsGetter interface {
	ApiManagementAPIs(namespace string) ApiManagementAPIInterface
}

// ApiManagementAPIInterface has methods to work with ApiManagementAPI resources.
type ApiManagementAPIInterface interface {
	Create(*v1alpha1.ApiManagementAPI) (*v1alpha1.ApiManagementAPI, error)
	Update(*v1alpha1.ApiManagementAPI) (*v1alpha1.ApiManagementAPI, error)
	UpdateStatus(*v1alpha1.ApiManagementAPI) (*v1alpha1.ApiManagementAPI, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementAPI, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementAPIList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAPI, err error)
	ApiManagementAPIExpansion
}

// apiManagementAPIs implements ApiManagementAPIInterface
type apiManagementAPIs struct {
	client rest.Interface
	ns     string
}

// newApiManagementAPIs returns a ApiManagementAPIs
func newApiManagementAPIs(c *AzurermV1alpha1Client, namespace string) *apiManagementAPIs {
	return &apiManagementAPIs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementAPI, and returns the corresponding apiManagementAPI object, and an error if there is any.
func (c *apiManagementAPIs) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementAPI, err error) {
	result = &v1alpha1.ApiManagementAPI{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementapis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementAPIs that match those selectors.
func (c *apiManagementAPIs) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementAPIList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementAPIList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementAPIs.
func (c *apiManagementAPIs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementAPI and creates it.  Returns the server's representation of the apiManagementAPI, and an error, if there is any.
func (c *apiManagementAPIs) Create(apiManagementAPI *v1alpha1.ApiManagementAPI) (result *v1alpha1.ApiManagementAPI, err error) {
	result = &v1alpha1.ApiManagementAPI{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementapis").
		Body(apiManagementAPI).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementAPI and updates it. Returns the server's representation of the apiManagementAPI, and an error, if there is any.
func (c *apiManagementAPIs) Update(apiManagementAPI *v1alpha1.ApiManagementAPI) (result *v1alpha1.ApiManagementAPI, err error) {
	result = &v1alpha1.ApiManagementAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementapis").
		Name(apiManagementAPI.Name).
		Body(apiManagementAPI).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementAPIs) UpdateStatus(apiManagementAPI *v1alpha1.ApiManagementAPI) (result *v1alpha1.ApiManagementAPI, err error) {
	result = &v1alpha1.ApiManagementAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementapis").
		Name(apiManagementAPI.Name).
		SubResource("status").
		Body(apiManagementAPI).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementAPI and deletes it. Returns an error if one occurs.
func (c *apiManagementAPIs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementapis").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementAPIs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementapis").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementAPI.
func (c *apiManagementAPIs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAPI, err error) {
	result = &v1alpha1.ApiManagementAPI{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementapis").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
