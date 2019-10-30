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

// ComputeDisksGetter has a method to return a ComputeDiskInterface.
// A group's client should implement this interface.
type ComputeDisksGetter interface {
	ComputeDisks(namespace string) ComputeDiskInterface
}

// ComputeDiskInterface has methods to work with ComputeDisk resources.
type ComputeDiskInterface interface {
	Create(*v1alpha1.ComputeDisk) (*v1alpha1.ComputeDisk, error)
	Update(*v1alpha1.ComputeDisk) (*v1alpha1.ComputeDisk, error)
	UpdateStatus(*v1alpha1.ComputeDisk) (*v1alpha1.ComputeDisk, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeDisk, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeDiskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeDisk, err error)
	ComputeDiskExpansion
}

// computeDisks implements ComputeDiskInterface
type computeDisks struct {
	client rest.Interface
	ns     string
}

// newComputeDisks returns a ComputeDisks
func newComputeDisks(c *GoogleV1alpha1Client, namespace string) *computeDisks {
	return &computeDisks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeDisk, and returns the corresponding computeDisk object, and an error if there is any.
func (c *computeDisks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeDisk, err error) {
	result = &v1alpha1.ComputeDisk{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computedisks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeDisks that match those selectors.
func (c *computeDisks) List(opts v1.ListOptions) (result *v1alpha1.ComputeDiskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeDiskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computedisks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeDisks.
func (c *computeDisks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computedisks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeDisk and creates it.  Returns the server's representation of the computeDisk, and an error, if there is any.
func (c *computeDisks) Create(computeDisk *v1alpha1.ComputeDisk) (result *v1alpha1.ComputeDisk, err error) {
	result = &v1alpha1.ComputeDisk{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computedisks").
		Body(computeDisk).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeDisk and updates it. Returns the server's representation of the computeDisk, and an error, if there is any.
func (c *computeDisks) Update(computeDisk *v1alpha1.ComputeDisk) (result *v1alpha1.ComputeDisk, err error) {
	result = &v1alpha1.ComputeDisk{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computedisks").
		Name(computeDisk.Name).
		Body(computeDisk).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeDisks) UpdateStatus(computeDisk *v1alpha1.ComputeDisk) (result *v1alpha1.ComputeDisk, err error) {
	result = &v1alpha1.ComputeDisk{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computedisks").
		Name(computeDisk.Name).
		SubResource("status").
		Body(computeDisk).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeDisk and deletes it. Returns an error if one occurs.
func (c *computeDisks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computedisks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeDisks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computedisks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeDisk.
func (c *computeDisks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeDisk, err error) {
	result = &v1alpha1.ComputeDisk{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computedisks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
