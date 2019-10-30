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

// IamRolesGetter has a method to return a IamRoleInterface.
// A group's client should implement this interface.
type IamRolesGetter interface {
	IamRoles(namespace string) IamRoleInterface
}

// IamRoleInterface has methods to work with IamRole resources.
type IamRoleInterface interface {
	Create(*v1alpha1.IamRole) (*v1alpha1.IamRole, error)
	Update(*v1alpha1.IamRole) (*v1alpha1.IamRole, error)
	UpdateStatus(*v1alpha1.IamRole) (*v1alpha1.IamRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamRole, error)
	List(opts v1.ListOptions) (*v1alpha1.IamRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamRole, err error)
	IamRoleExpansion
}

// iamRoles implements IamRoleInterface
type iamRoles struct {
	client rest.Interface
	ns     string
}

// newIamRoles returns a IamRoles
func newIamRoles(c *AwsV1alpha1Client, namespace string) *iamRoles {
	return &iamRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamRole, and returns the corresponding iamRole object, and an error if there is any.
func (c *iamRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.IamRole, err error) {
	result = &v1alpha1.IamRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamRoles that match those selectors.
func (c *iamRoles) List(opts v1.ListOptions) (result *v1alpha1.IamRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamRoles.
func (c *iamRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamRole and creates it.  Returns the server's representation of the iamRole, and an error, if there is any.
func (c *iamRoles) Create(iamRole *v1alpha1.IamRole) (result *v1alpha1.IamRole, err error) {
	result = &v1alpha1.IamRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamroles").
		Body(iamRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamRole and updates it. Returns the server's representation of the iamRole, and an error, if there is any.
func (c *iamRoles) Update(iamRole *v1alpha1.IamRole) (result *v1alpha1.IamRole, err error) {
	result = &v1alpha1.IamRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamroles").
		Name(iamRole.Name).
		Body(iamRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamRoles) UpdateStatus(iamRole *v1alpha1.IamRole) (result *v1alpha1.IamRole, err error) {
	result = &v1alpha1.IamRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamroles").
		Name(iamRole.Name).
		SubResource("status").
		Body(iamRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamRole and deletes it. Returns an error if one occurs.
func (c *iamRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamRole.
func (c *iamRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamRole, err error) {
	result = &v1alpha1.IamRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
