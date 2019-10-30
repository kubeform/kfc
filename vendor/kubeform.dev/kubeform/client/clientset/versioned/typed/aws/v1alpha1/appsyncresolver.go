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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppsyncResolversGetter has a method to return a AppsyncResolverInterface.
// A group's client should implement this interface.
type AppsyncResolversGetter interface {
	AppsyncResolvers(namespace string) AppsyncResolverInterface
}

// AppsyncResolverInterface has methods to work with AppsyncResolver resources.
type AppsyncResolverInterface interface {
	Create(*v1alpha1.AppsyncResolver) (*v1alpha1.AppsyncResolver, error)
	Update(*v1alpha1.AppsyncResolver) (*v1alpha1.AppsyncResolver, error)
	UpdateStatus(*v1alpha1.AppsyncResolver) (*v1alpha1.AppsyncResolver, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppsyncResolver, error)
	List(opts v1.ListOptions) (*v1alpha1.AppsyncResolverList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppsyncResolver, err error)
	AppsyncResolverExpansion
}

// appsyncResolvers implements AppsyncResolverInterface
type appsyncResolvers struct {
	client rest.Interface
	ns     string
}

// newAppsyncResolvers returns a AppsyncResolvers
func newAppsyncResolvers(c *AwsV1alpha1Client, namespace string) *appsyncResolvers {
	return &appsyncResolvers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appsyncResolver, and returns the corresponding appsyncResolver object, and an error if there is any.
func (c *appsyncResolvers) Get(name string, options v1.GetOptions) (result *v1alpha1.AppsyncResolver, err error) {
	result = &v1alpha1.AppsyncResolver{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppsyncResolvers that match those selectors.
func (c *appsyncResolvers) List(opts v1.ListOptions) (result *v1alpha1.AppsyncResolverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppsyncResolverList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appsyncResolvers.
func (c *appsyncResolvers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appsyncResolver and creates it.  Returns the server's representation of the appsyncResolver, and an error, if there is any.
func (c *appsyncResolvers) Create(appsyncResolver *v1alpha1.AppsyncResolver) (result *v1alpha1.AppsyncResolver, err error) {
	result = &v1alpha1.AppsyncResolver{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		Body(appsyncResolver).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appsyncResolver and updates it. Returns the server's representation of the appsyncResolver, and an error, if there is any.
func (c *appsyncResolvers) Update(appsyncResolver *v1alpha1.AppsyncResolver) (result *v1alpha1.AppsyncResolver, err error) {
	result = &v1alpha1.AppsyncResolver{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		Name(appsyncResolver.Name).
		Body(appsyncResolver).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appsyncResolvers) UpdateStatus(appsyncResolver *v1alpha1.AppsyncResolver) (result *v1alpha1.AppsyncResolver, err error) {
	result = &v1alpha1.AppsyncResolver{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		Name(appsyncResolver.Name).
		SubResource("status").
		Body(appsyncResolver).
		Do().
		Into(result)
	return
}

// Delete takes name of the appsyncResolver and deletes it. Returns an error if one occurs.
func (c *appsyncResolvers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appsyncResolvers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appsyncresolvers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appsyncResolver.
func (c *appsyncResolvers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppsyncResolver, err error) {
	result = &v1alpha1.AppsyncResolver{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appsyncresolvers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
