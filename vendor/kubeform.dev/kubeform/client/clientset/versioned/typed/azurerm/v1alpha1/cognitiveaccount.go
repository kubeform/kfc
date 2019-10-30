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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CognitiveAccountsGetter has a method to return a CognitiveAccountInterface.
// A group's client should implement this interface.
type CognitiveAccountsGetter interface {
	CognitiveAccounts(namespace string) CognitiveAccountInterface
}

// CognitiveAccountInterface has methods to work with CognitiveAccount resources.
type CognitiveAccountInterface interface {
	Create(*v1alpha1.CognitiveAccount) (*v1alpha1.CognitiveAccount, error)
	Update(*v1alpha1.CognitiveAccount) (*v1alpha1.CognitiveAccount, error)
	UpdateStatus(*v1alpha1.CognitiveAccount) (*v1alpha1.CognitiveAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CognitiveAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.CognitiveAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitiveAccount, err error)
	CognitiveAccountExpansion
}

// cognitiveAccounts implements CognitiveAccountInterface
type cognitiveAccounts struct {
	client rest.Interface
	ns     string
}

// newCognitiveAccounts returns a CognitiveAccounts
func newCognitiveAccounts(c *AzurermV1alpha1Client, namespace string) *cognitiveAccounts {
	return &cognitiveAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cognitiveAccount, and returns the corresponding cognitiveAccount object, and an error if there is any.
func (c *cognitiveAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitiveAccount, err error) {
	result = &v1alpha1.CognitiveAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CognitiveAccounts that match those selectors.
func (c *cognitiveAccounts) List(opts v1.ListOptions) (result *v1alpha1.CognitiveAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CognitiveAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cognitiveAccounts.
func (c *cognitiveAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cognitiveAccount and creates it.  Returns the server's representation of the cognitiveAccount, and an error, if there is any.
func (c *cognitiveAccounts) Create(cognitiveAccount *v1alpha1.CognitiveAccount) (result *v1alpha1.CognitiveAccount, err error) {
	result = &v1alpha1.CognitiveAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		Body(cognitiveAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cognitiveAccount and updates it. Returns the server's representation of the cognitiveAccount, and an error, if there is any.
func (c *cognitiveAccounts) Update(cognitiveAccount *v1alpha1.CognitiveAccount) (result *v1alpha1.CognitiveAccount, err error) {
	result = &v1alpha1.CognitiveAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		Name(cognitiveAccount.Name).
		Body(cognitiveAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cognitiveAccounts) UpdateStatus(cognitiveAccount *v1alpha1.CognitiveAccount) (result *v1alpha1.CognitiveAccount, err error) {
	result = &v1alpha1.CognitiveAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		Name(cognitiveAccount.Name).
		SubResource("status").
		Body(cognitiveAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the cognitiveAccount and deletes it. Returns an error if one occurs.
func (c *cognitiveAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cognitiveAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cognitiveAccount.
func (c *cognitiveAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitiveAccount, err error) {
	result = &v1alpha1.CognitiveAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cognitiveaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
