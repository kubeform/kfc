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

// RedisFirewallRulesGetter has a method to return a RedisFirewallRuleInterface.
// A group's client should implement this interface.
type RedisFirewallRulesGetter interface {
	RedisFirewallRules(namespace string) RedisFirewallRuleInterface
}

// RedisFirewallRuleInterface has methods to work with RedisFirewallRule resources.
type RedisFirewallRuleInterface interface {
	Create(*v1alpha1.RedisFirewallRule) (*v1alpha1.RedisFirewallRule, error)
	Update(*v1alpha1.RedisFirewallRule) (*v1alpha1.RedisFirewallRule, error)
	UpdateStatus(*v1alpha1.RedisFirewallRule) (*v1alpha1.RedisFirewallRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RedisFirewallRule, error)
	List(opts v1.ListOptions) (*v1alpha1.RedisFirewallRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisFirewallRule, err error)
	RedisFirewallRuleExpansion
}

// redisFirewallRules implements RedisFirewallRuleInterface
type redisFirewallRules struct {
	client rest.Interface
	ns     string
}

// newRedisFirewallRules returns a RedisFirewallRules
func newRedisFirewallRules(c *AzurermV1alpha1Client, namespace string) *redisFirewallRules {
	return &redisFirewallRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisFirewallRule, and returns the corresponding redisFirewallRule object, and an error if there is any.
func (c *redisFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisFirewallRule, err error) {
	result = &v1alpha1.RedisFirewallRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisFirewallRules that match those selectors.
func (c *redisFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.RedisFirewallRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RedisFirewallRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisFirewallRules.
func (c *redisFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a redisFirewallRule and creates it.  Returns the server's representation of the redisFirewallRule, and an error, if there is any.
func (c *redisFirewallRules) Create(redisFirewallRule *v1alpha1.RedisFirewallRule) (result *v1alpha1.RedisFirewallRule, err error) {
	result = &v1alpha1.RedisFirewallRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		Body(redisFirewallRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a redisFirewallRule and updates it. Returns the server's representation of the redisFirewallRule, and an error, if there is any.
func (c *redisFirewallRules) Update(redisFirewallRule *v1alpha1.RedisFirewallRule) (result *v1alpha1.RedisFirewallRule, err error) {
	result = &v1alpha1.RedisFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		Name(redisFirewallRule.Name).
		Body(redisFirewallRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *redisFirewallRules) UpdateStatus(redisFirewallRule *v1alpha1.RedisFirewallRule) (result *v1alpha1.RedisFirewallRule, err error) {
	result = &v1alpha1.RedisFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		Name(redisFirewallRule.Name).
		SubResource("status").
		Body(redisFirewallRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the redisFirewallRule and deletes it. Returns an error if one occurs.
func (c *redisFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisfirewallrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched redisFirewallRule.
func (c *redisFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisFirewallRule, err error) {
	result = &v1alpha1.RedisFirewallRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redisfirewallrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
