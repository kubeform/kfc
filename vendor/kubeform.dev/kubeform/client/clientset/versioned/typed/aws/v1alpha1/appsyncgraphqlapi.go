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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppsyncGraphqlAPIsGetter has a method to return a AppsyncGraphqlAPIInterface.
// A group's client should implement this interface.
type AppsyncGraphqlAPIsGetter interface {
	AppsyncGraphqlAPIs(namespace string) AppsyncGraphqlAPIInterface
}

// AppsyncGraphqlAPIInterface has methods to work with AppsyncGraphqlAPI resources.
type AppsyncGraphqlAPIInterface interface {
	Create(*v1alpha1.AppsyncGraphqlAPI) (*v1alpha1.AppsyncGraphqlAPI, error)
	Update(*v1alpha1.AppsyncGraphqlAPI) (*v1alpha1.AppsyncGraphqlAPI, error)
	UpdateStatus(*v1alpha1.AppsyncGraphqlAPI) (*v1alpha1.AppsyncGraphqlAPI, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppsyncGraphqlAPI, error)
	List(opts v1.ListOptions) (*v1alpha1.AppsyncGraphqlAPIList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppsyncGraphqlAPI, err error)
	AppsyncGraphqlAPIExpansion
}

// appsyncGraphqlAPIs implements AppsyncGraphqlAPIInterface
type appsyncGraphqlAPIs struct {
	client rest.Interface
	ns     string
}

// newAppsyncGraphqlAPIs returns a AppsyncGraphqlAPIs
func newAppsyncGraphqlAPIs(c *AwsV1alpha1Client, namespace string) *appsyncGraphqlAPIs {
	return &appsyncGraphqlAPIs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appsyncGraphqlAPI, and returns the corresponding appsyncGraphqlAPI object, and an error if there is any.
func (c *appsyncGraphqlAPIs) Get(name string, options v1.GetOptions) (result *v1alpha1.AppsyncGraphqlAPI, err error) {
	result = &v1alpha1.AppsyncGraphqlAPI{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppsyncGraphqlAPIs that match those selectors.
func (c *appsyncGraphqlAPIs) List(opts v1.ListOptions) (result *v1alpha1.AppsyncGraphqlAPIList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppsyncGraphqlAPIList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appsyncGraphqlAPIs.
func (c *appsyncGraphqlAPIs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appsyncGraphqlAPI and creates it.  Returns the server's representation of the appsyncGraphqlAPI, and an error, if there is any.
func (c *appsyncGraphqlAPIs) Create(appsyncGraphqlAPI *v1alpha1.AppsyncGraphqlAPI) (result *v1alpha1.AppsyncGraphqlAPI, err error) {
	result = &v1alpha1.AppsyncGraphqlAPI{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		Body(appsyncGraphqlAPI).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appsyncGraphqlAPI and updates it. Returns the server's representation of the appsyncGraphqlAPI, and an error, if there is any.
func (c *appsyncGraphqlAPIs) Update(appsyncGraphqlAPI *v1alpha1.AppsyncGraphqlAPI) (result *v1alpha1.AppsyncGraphqlAPI, err error) {
	result = &v1alpha1.AppsyncGraphqlAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		Name(appsyncGraphqlAPI.Name).
		Body(appsyncGraphqlAPI).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appsyncGraphqlAPIs) UpdateStatus(appsyncGraphqlAPI *v1alpha1.AppsyncGraphqlAPI) (result *v1alpha1.AppsyncGraphqlAPI, err error) {
	result = &v1alpha1.AppsyncGraphqlAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		Name(appsyncGraphqlAPI.Name).
		SubResource("status").
		Body(appsyncGraphqlAPI).
		Do().
		Into(result)
	return
}

// Delete takes name of the appsyncGraphqlAPI and deletes it. Returns an error if one occurs.
func (c *appsyncGraphqlAPIs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appsyncGraphqlAPIs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appsyncGraphqlAPI.
func (c *appsyncGraphqlAPIs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppsyncGraphqlAPI, err error) {
	result = &v1alpha1.AppsyncGraphqlAPI{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appsyncgraphqlapis").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
