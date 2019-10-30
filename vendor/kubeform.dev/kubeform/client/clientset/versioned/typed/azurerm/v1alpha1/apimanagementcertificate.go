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

// ApiManagementCertificatesGetter has a method to return a ApiManagementCertificateInterface.
// A group's client should implement this interface.
type ApiManagementCertificatesGetter interface {
	ApiManagementCertificates(namespace string) ApiManagementCertificateInterface
}

// ApiManagementCertificateInterface has methods to work with ApiManagementCertificate resources.
type ApiManagementCertificateInterface interface {
	Create(*v1alpha1.ApiManagementCertificate) (*v1alpha1.ApiManagementCertificate, error)
	Update(*v1alpha1.ApiManagementCertificate) (*v1alpha1.ApiManagementCertificate, error)
	UpdateStatus(*v1alpha1.ApiManagementCertificate) (*v1alpha1.ApiManagementCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementCertificate, err error)
	ApiManagementCertificateExpansion
}

// apiManagementCertificates implements ApiManagementCertificateInterface
type apiManagementCertificates struct {
	client rest.Interface
	ns     string
}

// newApiManagementCertificates returns a ApiManagementCertificates
func newApiManagementCertificates(c *AzurermV1alpha1Client, namespace string) *apiManagementCertificates {
	return &apiManagementCertificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementCertificate, and returns the corresponding apiManagementCertificate object, and an error if there is any.
func (c *apiManagementCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementCertificate, err error) {
	result = &v1alpha1.ApiManagementCertificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementCertificates that match those selectors.
func (c *apiManagementCertificates) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementCertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementCertificates.
func (c *apiManagementCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementCertificate and creates it.  Returns the server's representation of the apiManagementCertificate, and an error, if there is any.
func (c *apiManagementCertificates) Create(apiManagementCertificate *v1alpha1.ApiManagementCertificate) (result *v1alpha1.ApiManagementCertificate, err error) {
	result = &v1alpha1.ApiManagementCertificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		Body(apiManagementCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementCertificate and updates it. Returns the server's representation of the apiManagementCertificate, and an error, if there is any.
func (c *apiManagementCertificates) Update(apiManagementCertificate *v1alpha1.ApiManagementCertificate) (result *v1alpha1.ApiManagementCertificate, err error) {
	result = &v1alpha1.ApiManagementCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		Name(apiManagementCertificate.Name).
		Body(apiManagementCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementCertificates) UpdateStatus(apiManagementCertificate *v1alpha1.ApiManagementCertificate) (result *v1alpha1.ApiManagementCertificate, err error) {
	result = &v1alpha1.ApiManagementCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		Name(apiManagementCertificate.Name).
		SubResource("status").
		Body(apiManagementCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementCertificate and deletes it. Returns an error if one occurs.
func (c *apiManagementCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementCertificate.
func (c *apiManagementCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementCertificate, err error) {
	result = &v1alpha1.ApiManagementCertificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
