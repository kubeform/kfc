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

// IamRolePoliciesGetter has a method to return a IamRolePolicyInterface.
// A group's client should implement this interface.
type IamRolePoliciesGetter interface {
	IamRolePolicies(namespace string) IamRolePolicyInterface
}

// IamRolePolicyInterface has methods to work with IamRolePolicy resources.
type IamRolePolicyInterface interface {
	Create(*v1alpha1.IamRolePolicy) (*v1alpha1.IamRolePolicy, error)
	Update(*v1alpha1.IamRolePolicy) (*v1alpha1.IamRolePolicy, error)
	UpdateStatus(*v1alpha1.IamRolePolicy) (*v1alpha1.IamRolePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamRolePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.IamRolePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamRolePolicy, err error)
	IamRolePolicyExpansion
}

// iamRolePolicies implements IamRolePolicyInterface
type iamRolePolicies struct {
	client rest.Interface
	ns     string
}

// newIamRolePolicies returns a IamRolePolicies
func newIamRolePolicies(c *AwsV1alpha1Client, namespace string) *iamRolePolicies {
	return &iamRolePolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamRolePolicy, and returns the corresponding iamRolePolicy object, and an error if there is any.
func (c *iamRolePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IamRolePolicy, err error) {
	result = &v1alpha1.IamRolePolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamRolePolicies that match those selectors.
func (c *iamRolePolicies) List(opts v1.ListOptions) (result *v1alpha1.IamRolePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamRolePolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamRolePolicies.
func (c *iamRolePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamRolePolicy and creates it.  Returns the server's representation of the iamRolePolicy, and an error, if there is any.
func (c *iamRolePolicies) Create(iamRolePolicy *v1alpha1.IamRolePolicy) (result *v1alpha1.IamRolePolicy, err error) {
	result = &v1alpha1.IamRolePolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		Body(iamRolePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamRolePolicy and updates it. Returns the server's representation of the iamRolePolicy, and an error, if there is any.
func (c *iamRolePolicies) Update(iamRolePolicy *v1alpha1.IamRolePolicy) (result *v1alpha1.IamRolePolicy, err error) {
	result = &v1alpha1.IamRolePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		Name(iamRolePolicy.Name).
		Body(iamRolePolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamRolePolicies) UpdateStatus(iamRolePolicy *v1alpha1.IamRolePolicy) (result *v1alpha1.IamRolePolicy, err error) {
	result = &v1alpha1.IamRolePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		Name(iamRolePolicy.Name).
		SubResource("status").
		Body(iamRolePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamRolePolicy and deletes it. Returns an error if one occurs.
func (c *iamRolePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamRolePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamrolepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamRolePolicy.
func (c *iamRolePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamRolePolicy, err error) {
	result = &v1alpha1.IamRolePolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamrolepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
