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

// RdsClusterParameterGroupsGetter has a method to return a RdsClusterParameterGroupInterface.
// A group's client should implement this interface.
type RdsClusterParameterGroupsGetter interface {
	RdsClusterParameterGroups(namespace string) RdsClusterParameterGroupInterface
}

// RdsClusterParameterGroupInterface has methods to work with RdsClusterParameterGroup resources.
type RdsClusterParameterGroupInterface interface {
	Create(*v1alpha1.RdsClusterParameterGroup) (*v1alpha1.RdsClusterParameterGroup, error)
	Update(*v1alpha1.RdsClusterParameterGroup) (*v1alpha1.RdsClusterParameterGroup, error)
	UpdateStatus(*v1alpha1.RdsClusterParameterGroup) (*v1alpha1.RdsClusterParameterGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RdsClusterParameterGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.RdsClusterParameterGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsClusterParameterGroup, err error)
	RdsClusterParameterGroupExpansion
}

// rdsClusterParameterGroups implements RdsClusterParameterGroupInterface
type rdsClusterParameterGroups struct {
	client rest.Interface
	ns     string
}

// newRdsClusterParameterGroups returns a RdsClusterParameterGroups
func newRdsClusterParameterGroups(c *AwsV1alpha1Client, namespace string) *rdsClusterParameterGroups {
	return &rdsClusterParameterGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rdsClusterParameterGroup, and returns the corresponding rdsClusterParameterGroup object, and an error if there is any.
func (c *rdsClusterParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	result = &v1alpha1.RdsClusterParameterGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RdsClusterParameterGroups that match those selectors.
func (c *rdsClusterParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.RdsClusterParameterGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RdsClusterParameterGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rdsClusterParameterGroups.
func (c *rdsClusterParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a rdsClusterParameterGroup and creates it.  Returns the server's representation of the rdsClusterParameterGroup, and an error, if there is any.
func (c *rdsClusterParameterGroups) Create(rdsClusterParameterGroup *v1alpha1.RdsClusterParameterGroup) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	result = &v1alpha1.RdsClusterParameterGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		Body(rdsClusterParameterGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rdsClusterParameterGroup and updates it. Returns the server's representation of the rdsClusterParameterGroup, and an error, if there is any.
func (c *rdsClusterParameterGroups) Update(rdsClusterParameterGroup *v1alpha1.RdsClusterParameterGroup) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	result = &v1alpha1.RdsClusterParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		Name(rdsClusterParameterGroup.Name).
		Body(rdsClusterParameterGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *rdsClusterParameterGroups) UpdateStatus(rdsClusterParameterGroup *v1alpha1.RdsClusterParameterGroup) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	result = &v1alpha1.RdsClusterParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		Name(rdsClusterParameterGroup.Name).
		SubResource("status").
		Body(rdsClusterParameterGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the rdsClusterParameterGroup and deletes it. Returns an error if one occurs.
func (c *rdsClusterParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rdsClusterParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rdsClusterParameterGroup.
func (c *rdsClusterParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsClusterParameterGroup, err error) {
	result = &v1alpha1.RdsClusterParameterGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rdsclusterparametergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
