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

// CognitoUserGroupsGetter has a method to return a CognitoUserGroupInterface.
// A group's client should implement this interface.
type CognitoUserGroupsGetter interface {
	CognitoUserGroups(namespace string) CognitoUserGroupInterface
}

// CognitoUserGroupInterface has methods to work with CognitoUserGroup resources.
type CognitoUserGroupInterface interface {
	Create(*v1alpha1.CognitoUserGroup) (*v1alpha1.CognitoUserGroup, error)
	Update(*v1alpha1.CognitoUserGroup) (*v1alpha1.CognitoUserGroup, error)
	UpdateStatus(*v1alpha1.CognitoUserGroup) (*v1alpha1.CognitoUserGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CognitoUserGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.CognitoUserGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoUserGroup, err error)
	CognitoUserGroupExpansion
}

// cognitoUserGroups implements CognitoUserGroupInterface
type cognitoUserGroups struct {
	client rest.Interface
	ns     string
}

// newCognitoUserGroups returns a CognitoUserGroups
func newCognitoUserGroups(c *AwsV1alpha1Client, namespace string) *cognitoUserGroups {
	return &cognitoUserGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cognitoUserGroup, and returns the corresponding cognitoUserGroup object, and an error if there is any.
func (c *cognitoUserGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitoUserGroup, err error) {
	result = &v1alpha1.CognitoUserGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitousergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CognitoUserGroups that match those selectors.
func (c *cognitoUserGroups) List(opts v1.ListOptions) (result *v1alpha1.CognitoUserGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CognitoUserGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitousergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cognitoUserGroups.
func (c *cognitoUserGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cognitousergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cognitoUserGroup and creates it.  Returns the server's representation of the cognitoUserGroup, and an error, if there is any.
func (c *cognitoUserGroups) Create(cognitoUserGroup *v1alpha1.CognitoUserGroup) (result *v1alpha1.CognitoUserGroup, err error) {
	result = &v1alpha1.CognitoUserGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cognitousergroups").
		Body(cognitoUserGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cognitoUserGroup and updates it. Returns the server's representation of the cognitoUserGroup, and an error, if there is any.
func (c *cognitoUserGroups) Update(cognitoUserGroup *v1alpha1.CognitoUserGroup) (result *v1alpha1.CognitoUserGroup, err error) {
	result = &v1alpha1.CognitoUserGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitousergroups").
		Name(cognitoUserGroup.Name).
		Body(cognitoUserGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cognitoUserGroups) UpdateStatus(cognitoUserGroup *v1alpha1.CognitoUserGroup) (result *v1alpha1.CognitoUserGroup, err error) {
	result = &v1alpha1.CognitoUserGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitousergroups").
		Name(cognitoUserGroup.Name).
		SubResource("status").
		Body(cognitoUserGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the cognitoUserGroup and deletes it. Returns an error if one occurs.
func (c *cognitoUserGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitousergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cognitoUserGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitousergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cognitoUserGroup.
func (c *cognitoUserGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoUserGroup, err error) {
	result = &v1alpha1.CognitoUserGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cognitousergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
