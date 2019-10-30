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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MysqlVirtualNetworkRulesGetter has a method to return a MysqlVirtualNetworkRuleInterface.
// A group's client should implement this interface.
type MysqlVirtualNetworkRulesGetter interface {
	MysqlVirtualNetworkRules(namespace string) MysqlVirtualNetworkRuleInterface
}

// MysqlVirtualNetworkRuleInterface has methods to work with MysqlVirtualNetworkRule resources.
type MysqlVirtualNetworkRuleInterface interface {
	Create(*v1alpha1.MysqlVirtualNetworkRule) (*v1alpha1.MysqlVirtualNetworkRule, error)
	Update(*v1alpha1.MysqlVirtualNetworkRule) (*v1alpha1.MysqlVirtualNetworkRule, error)
	UpdateStatus(*v1alpha1.MysqlVirtualNetworkRule) (*v1alpha1.MysqlVirtualNetworkRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MysqlVirtualNetworkRule, error)
	List(opts v1.ListOptions) (*v1alpha1.MysqlVirtualNetworkRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlVirtualNetworkRule, err error)
	MysqlVirtualNetworkRuleExpansion
}

// mysqlVirtualNetworkRules implements MysqlVirtualNetworkRuleInterface
type mysqlVirtualNetworkRules struct {
	client rest.Interface
	ns     string
}

// newMysqlVirtualNetworkRules returns a MysqlVirtualNetworkRules
func newMysqlVirtualNetworkRules(c *AzurermV1alpha1Client, namespace string) *mysqlVirtualNetworkRules {
	return &mysqlVirtualNetworkRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mysqlVirtualNetworkRule, and returns the corresponding mysqlVirtualNetworkRule object, and an error if there is any.
func (c *mysqlVirtualNetworkRules) Get(name string, options v1.GetOptions) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	result = &v1alpha1.MysqlVirtualNetworkRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MysqlVirtualNetworkRules that match those selectors.
func (c *mysqlVirtualNetworkRules) List(opts v1.ListOptions) (result *v1alpha1.MysqlVirtualNetworkRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlVirtualNetworkRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mysqlVirtualNetworkRules.
func (c *mysqlVirtualNetworkRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mysqlVirtualNetworkRule and creates it.  Returns the server's representation of the mysqlVirtualNetworkRule, and an error, if there is any.
func (c *mysqlVirtualNetworkRules) Create(mysqlVirtualNetworkRule *v1alpha1.MysqlVirtualNetworkRule) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	result = &v1alpha1.MysqlVirtualNetworkRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		Body(mysqlVirtualNetworkRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mysqlVirtualNetworkRule and updates it. Returns the server's representation of the mysqlVirtualNetworkRule, and an error, if there is any.
func (c *mysqlVirtualNetworkRules) Update(mysqlVirtualNetworkRule *v1alpha1.MysqlVirtualNetworkRule) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	result = &v1alpha1.MysqlVirtualNetworkRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		Name(mysqlVirtualNetworkRule.Name).
		Body(mysqlVirtualNetworkRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mysqlVirtualNetworkRules) UpdateStatus(mysqlVirtualNetworkRule *v1alpha1.MysqlVirtualNetworkRule) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	result = &v1alpha1.MysqlVirtualNetworkRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		Name(mysqlVirtualNetworkRule.Name).
		SubResource("status").
		Body(mysqlVirtualNetworkRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the mysqlVirtualNetworkRule and deletes it. Returns an error if one occurs.
func (c *mysqlVirtualNetworkRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mysqlVirtualNetworkRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mysqlVirtualNetworkRule.
func (c *mysqlVirtualNetworkRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	result = &v1alpha1.MysqlVirtualNetworkRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlvirtualnetworkrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
