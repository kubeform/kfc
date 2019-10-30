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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeRoutesGetter has a method to return a ComputeRouteInterface.
// A group's client should implement this interface.
type ComputeRoutesGetter interface {
	ComputeRoutes(namespace string) ComputeRouteInterface
}

// ComputeRouteInterface has methods to work with ComputeRoute resources.
type ComputeRouteInterface interface {
	Create(*v1alpha1.ComputeRoute) (*v1alpha1.ComputeRoute, error)
	Update(*v1alpha1.ComputeRoute) (*v1alpha1.ComputeRoute, error)
	UpdateStatus(*v1alpha1.ComputeRoute) (*v1alpha1.ComputeRoute, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeRoute, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeRouteList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRoute, err error)
	ComputeRouteExpansion
}

// computeRoutes implements ComputeRouteInterface
type computeRoutes struct {
	client rest.Interface
	ns     string
}

// newComputeRoutes returns a ComputeRoutes
func newComputeRoutes(c *GoogleV1alpha1Client, namespace string) *computeRoutes {
	return &computeRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeRoute, and returns the corresponding computeRoute object, and an error if there is any.
func (c *computeRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRoute, err error) {
	result = &v1alpha1.ComputeRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeRoutes that match those selectors.
func (c *computeRoutes) List(opts v1.ListOptions) (result *v1alpha1.ComputeRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeRoutes.
func (c *computeRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeRoute and creates it.  Returns the server's representation of the computeRoute, and an error, if there is any.
func (c *computeRoutes) Create(computeRoute *v1alpha1.ComputeRoute) (result *v1alpha1.ComputeRoute, err error) {
	result = &v1alpha1.ComputeRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeroutes").
		Body(computeRoute).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeRoute and updates it. Returns the server's representation of the computeRoute, and an error, if there is any.
func (c *computeRoutes) Update(computeRoute *v1alpha1.ComputeRoute) (result *v1alpha1.ComputeRoute, err error) {
	result = &v1alpha1.ComputeRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeroutes").
		Name(computeRoute.Name).
		Body(computeRoute).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeRoutes) UpdateStatus(computeRoute *v1alpha1.ComputeRoute) (result *v1alpha1.ComputeRoute, err error) {
	result = &v1alpha1.ComputeRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeroutes").
		Name(computeRoute.Name).
		SubResource("status").
		Body(computeRoute).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeRoute and deletes it. Returns an error if one occurs.
func (c *computeRoutes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeroutes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeroutes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeRoute.
func (c *computeRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRoute, err error) {
	result = &v1alpha1.ComputeRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeroutes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
