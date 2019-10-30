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

// NeptuneParameterGroupsGetter has a method to return a NeptuneParameterGroupInterface.
// A group's client should implement this interface.
type NeptuneParameterGroupsGetter interface {
	NeptuneParameterGroups(namespace string) NeptuneParameterGroupInterface
}

// NeptuneParameterGroupInterface has methods to work with NeptuneParameterGroup resources.
type NeptuneParameterGroupInterface interface {
	Create(*v1alpha1.NeptuneParameterGroup) (*v1alpha1.NeptuneParameterGroup, error)
	Update(*v1alpha1.NeptuneParameterGroup) (*v1alpha1.NeptuneParameterGroup, error)
	UpdateStatus(*v1alpha1.NeptuneParameterGroup) (*v1alpha1.NeptuneParameterGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NeptuneParameterGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.NeptuneParameterGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NeptuneParameterGroup, err error)
	NeptuneParameterGroupExpansion
}

// neptuneParameterGroups implements NeptuneParameterGroupInterface
type neptuneParameterGroups struct {
	client rest.Interface
	ns     string
}

// newNeptuneParameterGroups returns a NeptuneParameterGroups
func newNeptuneParameterGroups(c *AwsV1alpha1Client, namespace string) *neptuneParameterGroups {
	return &neptuneParameterGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the neptuneParameterGroup, and returns the corresponding neptuneParameterGroup object, and an error if there is any.
func (c *neptuneParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.NeptuneParameterGroup, err error) {
	result = &v1alpha1.NeptuneParameterGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NeptuneParameterGroups that match those selectors.
func (c *neptuneParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.NeptuneParameterGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NeptuneParameterGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested neptuneParameterGroups.
func (c *neptuneParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a neptuneParameterGroup and creates it.  Returns the server's representation of the neptuneParameterGroup, and an error, if there is any.
func (c *neptuneParameterGroups) Create(neptuneParameterGroup *v1alpha1.NeptuneParameterGroup) (result *v1alpha1.NeptuneParameterGroup, err error) {
	result = &v1alpha1.NeptuneParameterGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		Body(neptuneParameterGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a neptuneParameterGroup and updates it. Returns the server's representation of the neptuneParameterGroup, and an error, if there is any.
func (c *neptuneParameterGroups) Update(neptuneParameterGroup *v1alpha1.NeptuneParameterGroup) (result *v1alpha1.NeptuneParameterGroup, err error) {
	result = &v1alpha1.NeptuneParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		Name(neptuneParameterGroup.Name).
		Body(neptuneParameterGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *neptuneParameterGroups) UpdateStatus(neptuneParameterGroup *v1alpha1.NeptuneParameterGroup) (result *v1alpha1.NeptuneParameterGroup, err error) {
	result = &v1alpha1.NeptuneParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		Name(neptuneParameterGroup.Name).
		SubResource("status").
		Body(neptuneParameterGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the neptuneParameterGroup and deletes it. Returns an error if one occurs.
func (c *neptuneParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *neptuneParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched neptuneParameterGroup.
func (c *neptuneParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NeptuneParameterGroup, err error) {
	result = &v1alpha1.NeptuneParameterGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("neptuneparametergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
