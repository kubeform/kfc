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

// DefaultVpcsGetter has a method to return a DefaultVpcInterface.
// A group's client should implement this interface.
type DefaultVpcsGetter interface {
	DefaultVpcs(namespace string) DefaultVpcInterface
}

// DefaultVpcInterface has methods to work with DefaultVpc resources.
type DefaultVpcInterface interface {
	Create(*v1alpha1.DefaultVpc) (*v1alpha1.DefaultVpc, error)
	Update(*v1alpha1.DefaultVpc) (*v1alpha1.DefaultVpc, error)
	UpdateStatus(*v1alpha1.DefaultVpc) (*v1alpha1.DefaultVpc, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DefaultVpc, error)
	List(opts v1.ListOptions) (*v1alpha1.DefaultVpcList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultVpc, err error)
	DefaultVpcExpansion
}

// defaultVpcs implements DefaultVpcInterface
type defaultVpcs struct {
	client rest.Interface
	ns     string
}

// newDefaultVpcs returns a DefaultVpcs
func newDefaultVpcs(c *AwsV1alpha1Client, namespace string) *defaultVpcs {
	return &defaultVpcs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the defaultVpc, and returns the corresponding defaultVpc object, and an error if there is any.
func (c *defaultVpcs) Get(name string, options v1.GetOptions) (result *v1alpha1.DefaultVpc, err error) {
	result = &v1alpha1.DefaultVpc{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultvpcs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DefaultVpcs that match those selectors.
func (c *defaultVpcs) List(opts v1.ListOptions) (result *v1alpha1.DefaultVpcList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DefaultVpcList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultvpcs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested defaultVpcs.
func (c *defaultVpcs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("defaultvpcs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a defaultVpc and creates it.  Returns the server's representation of the defaultVpc, and an error, if there is any.
func (c *defaultVpcs) Create(defaultVpc *v1alpha1.DefaultVpc) (result *v1alpha1.DefaultVpc, err error) {
	result = &v1alpha1.DefaultVpc{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("defaultvpcs").
		Body(defaultVpc).
		Do().
		Into(result)
	return
}

// Update takes the representation of a defaultVpc and updates it. Returns the server's representation of the defaultVpc, and an error, if there is any.
func (c *defaultVpcs) Update(defaultVpc *v1alpha1.DefaultVpc) (result *v1alpha1.DefaultVpc, err error) {
	result = &v1alpha1.DefaultVpc{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultvpcs").
		Name(defaultVpc.Name).
		Body(defaultVpc).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *defaultVpcs) UpdateStatus(defaultVpc *v1alpha1.DefaultVpc) (result *v1alpha1.DefaultVpc, err error) {
	result = &v1alpha1.DefaultVpc{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultvpcs").
		Name(defaultVpc.Name).
		SubResource("status").
		Body(defaultVpc).
		Do().
		Into(result)
	return
}

// Delete takes name of the defaultVpc and deletes it. Returns an error if one occurs.
func (c *defaultVpcs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultvpcs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *defaultVpcs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultvpcs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched defaultVpc.
func (c *defaultVpcs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DefaultVpc, err error) {
	result = &v1alpha1.DefaultVpc{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("defaultvpcs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
