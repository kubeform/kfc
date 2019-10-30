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

// NatGatewaysGetter has a method to return a NatGatewayInterface.
// A group's client should implement this interface.
type NatGatewaysGetter interface {
	NatGateways(namespace string) NatGatewayInterface
}

// NatGatewayInterface has methods to work with NatGateway resources.
type NatGatewayInterface interface {
	Create(*v1alpha1.NatGateway) (*v1alpha1.NatGateway, error)
	Update(*v1alpha1.NatGateway) (*v1alpha1.NatGateway, error)
	UpdateStatus(*v1alpha1.NatGateway) (*v1alpha1.NatGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NatGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.NatGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NatGateway, err error)
	NatGatewayExpansion
}

// natGateways implements NatGatewayInterface
type natGateways struct {
	client rest.Interface
	ns     string
}

// newNatGateways returns a NatGateways
func newNatGateways(c *AwsV1alpha1Client, namespace string) *natGateways {
	return &natGateways{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the natGateway, and returns the corresponding natGateway object, and an error if there is any.
func (c *natGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.NatGateway, err error) {
	result = &v1alpha1.NatGateway{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natgateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NatGateways that match those selectors.
func (c *natGateways) List(opts v1.ListOptions) (result *v1alpha1.NatGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NatGatewayList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested natGateways.
func (c *natGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("natgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a natGateway and creates it.  Returns the server's representation of the natGateway, and an error, if there is any.
func (c *natGateways) Create(natGateway *v1alpha1.NatGateway) (result *v1alpha1.NatGateway, err error) {
	result = &v1alpha1.NatGateway{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("natgateways").
		Body(natGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a natGateway and updates it. Returns the server's representation of the natGateway, and an error, if there is any.
func (c *natGateways) Update(natGateway *v1alpha1.NatGateway) (result *v1alpha1.NatGateway, err error) {
	result = &v1alpha1.NatGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("natgateways").
		Name(natGateway.Name).
		Body(natGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *natGateways) UpdateStatus(natGateway *v1alpha1.NatGateway) (result *v1alpha1.NatGateway, err error) {
	result = &v1alpha1.NatGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("natgateways").
		Name(natGateway.Name).
		SubResource("status").
		Body(natGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the natGateway and deletes it. Returns an error if one occurs.
func (c *natGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natgateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *natGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natgateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched natGateway.
func (c *natGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NatGateway, err error) {
	result = &v1alpha1.NatGateway{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("natgateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
