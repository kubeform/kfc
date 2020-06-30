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
	"context"
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeInstanceGroupsGetter has a method to return a ComputeInstanceGroupInterface.
// A group's client should implement this interface.
type ComputeInstanceGroupsGetter interface {
	ComputeInstanceGroups(namespace string) ComputeInstanceGroupInterface
}

// ComputeInstanceGroupInterface has methods to work with ComputeInstanceGroup resources.
type ComputeInstanceGroupInterface interface {
	Create(ctx context.Context, computeInstanceGroup *v1alpha1.ComputeInstanceGroup, opts v1.CreateOptions) (*v1alpha1.ComputeInstanceGroup, error)
	Update(ctx context.Context, computeInstanceGroup *v1alpha1.ComputeInstanceGroup, opts v1.UpdateOptions) (*v1alpha1.ComputeInstanceGroup, error)
	UpdateStatus(ctx context.Context, computeInstanceGroup *v1alpha1.ComputeInstanceGroup, opts v1.UpdateOptions) (*v1alpha1.ComputeInstanceGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ComputeInstanceGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ComputeInstanceGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeInstanceGroup, err error)
	ComputeInstanceGroupExpansion
}

// computeInstanceGroups implements ComputeInstanceGroupInterface
type computeInstanceGroups struct {
	client rest.Interface
	ns     string
}

// newComputeInstanceGroups returns a ComputeInstanceGroups
func newComputeInstanceGroups(c *GoogleV1alpha1Client, namespace string) *computeInstanceGroups {
	return &computeInstanceGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeInstanceGroup, and returns the corresponding computeInstanceGroup object, and an error if there is any.
func (c *computeInstanceGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeInstanceGroup, err error) {
	result = &v1alpha1.ComputeInstanceGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeInstanceGroups that match those selectors.
func (c *computeInstanceGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeInstanceGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeInstanceGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeInstanceGroups.
func (c *computeInstanceGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeInstanceGroup and creates it.  Returns the server's representation of the computeInstanceGroup, and an error, if there is any.
func (c *computeInstanceGroups) Create(ctx context.Context, computeInstanceGroup *v1alpha1.ComputeInstanceGroup, opts v1.CreateOptions) (result *v1alpha1.ComputeInstanceGroup, err error) {
	result = &v1alpha1.ComputeInstanceGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeInstanceGroup and updates it. Returns the server's representation of the computeInstanceGroup, and an error, if there is any.
func (c *computeInstanceGroups) Update(ctx context.Context, computeInstanceGroup *v1alpha1.ComputeInstanceGroup, opts v1.UpdateOptions) (result *v1alpha1.ComputeInstanceGroup, err error) {
	result = &v1alpha1.ComputeInstanceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		Name(computeInstanceGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *computeInstanceGroups) UpdateStatus(ctx context.Context, computeInstanceGroup *v1alpha1.ComputeInstanceGroup, opts v1.UpdateOptions) (result *v1alpha1.ComputeInstanceGroup, err error) {
	result = &v1alpha1.ComputeInstanceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		Name(computeInstanceGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeInstanceGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeInstanceGroup and deletes it. Returns an error if one occurs.
func (c *computeInstanceGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeInstanceGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinstancegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeInstanceGroup.
func (c *computeInstanceGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeInstanceGroup, err error) {
	result = &v1alpha1.ComputeInstanceGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeinstancegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
