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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeSubnetworksGetter has a method to return a ComputeSubnetworkInterface.
// A group's client should implement this interface.
type ComputeSubnetworksGetter interface {
	ComputeSubnetworks(namespace string) ComputeSubnetworkInterface
}

// ComputeSubnetworkInterface has methods to work with ComputeSubnetwork resources.
type ComputeSubnetworkInterface interface {
	Create(*v1alpha1.ComputeSubnetwork) (*v1alpha1.ComputeSubnetwork, error)
	Update(*v1alpha1.ComputeSubnetwork) (*v1alpha1.ComputeSubnetwork, error)
	UpdateStatus(*v1alpha1.ComputeSubnetwork) (*v1alpha1.ComputeSubnetwork, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeSubnetwork, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeSubnetworkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSubnetwork, err error)
	ComputeSubnetworkExpansion
}

// computeSubnetworks implements ComputeSubnetworkInterface
type computeSubnetworks struct {
	client rest.Interface
	ns     string
}

// newComputeSubnetworks returns a ComputeSubnetworks
func newComputeSubnetworks(c *GoogleV1alpha1Client, namespace string) *computeSubnetworks {
	return &computeSubnetworks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeSubnetwork, and returns the corresponding computeSubnetwork object, and an error if there is any.
func (c *computeSubnetworks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSubnetwork, err error) {
	result = &v1alpha1.ComputeSubnetwork{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesubnetworks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeSubnetworks that match those selectors.
func (c *computeSubnetworks) List(opts v1.ListOptions) (result *v1alpha1.ComputeSubnetworkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeSubnetworkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesubnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeSubnetworks.
func (c *computeSubnetworks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computesubnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeSubnetwork and creates it.  Returns the server's representation of the computeSubnetwork, and an error, if there is any.
func (c *computeSubnetworks) Create(computeSubnetwork *v1alpha1.ComputeSubnetwork) (result *v1alpha1.ComputeSubnetwork, err error) {
	result = &v1alpha1.ComputeSubnetwork{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computesubnetworks").
		Body(computeSubnetwork).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeSubnetwork and updates it. Returns the server's representation of the computeSubnetwork, and an error, if there is any.
func (c *computeSubnetworks) Update(computeSubnetwork *v1alpha1.ComputeSubnetwork) (result *v1alpha1.ComputeSubnetwork, err error) {
	result = &v1alpha1.ComputeSubnetwork{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesubnetworks").
		Name(computeSubnetwork.Name).
		Body(computeSubnetwork).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeSubnetworks) UpdateStatus(computeSubnetwork *v1alpha1.ComputeSubnetwork) (result *v1alpha1.ComputeSubnetwork, err error) {
	result = &v1alpha1.ComputeSubnetwork{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesubnetworks").
		Name(computeSubnetwork.Name).
		SubResource("status").
		Body(computeSubnetwork).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeSubnetwork and deletes it. Returns an error if one occurs.
func (c *computeSubnetworks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesubnetworks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeSubnetworks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesubnetworks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeSubnetwork.
func (c *computeSubnetworks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSubnetwork, err error) {
	result = &v1alpha1.ComputeSubnetwork{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computesubnetworks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
