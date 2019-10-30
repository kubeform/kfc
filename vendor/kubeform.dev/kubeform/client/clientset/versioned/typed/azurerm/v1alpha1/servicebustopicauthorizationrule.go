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

// ServicebusTopicAuthorizationRulesGetter has a method to return a ServicebusTopicAuthorizationRuleInterface.
// A group's client should implement this interface.
type ServicebusTopicAuthorizationRulesGetter interface {
	ServicebusTopicAuthorizationRules(namespace string) ServicebusTopicAuthorizationRuleInterface
}

// ServicebusTopicAuthorizationRuleInterface has methods to work with ServicebusTopicAuthorizationRule resources.
type ServicebusTopicAuthorizationRuleInterface interface {
	Create(*v1alpha1.ServicebusTopicAuthorizationRule) (*v1alpha1.ServicebusTopicAuthorizationRule, error)
	Update(*v1alpha1.ServicebusTopicAuthorizationRule) (*v1alpha1.ServicebusTopicAuthorizationRule, error)
	UpdateStatus(*v1alpha1.ServicebusTopicAuthorizationRule) (*v1alpha1.ServicebusTopicAuthorizationRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServicebusTopicAuthorizationRule, error)
	List(opts v1.ListOptions) (*v1alpha1.ServicebusTopicAuthorizationRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusTopicAuthorizationRule, err error)
	ServicebusTopicAuthorizationRuleExpansion
}

// servicebusTopicAuthorizationRules implements ServicebusTopicAuthorizationRuleInterface
type servicebusTopicAuthorizationRules struct {
	client rest.Interface
	ns     string
}

// newServicebusTopicAuthorizationRules returns a ServicebusTopicAuthorizationRules
func newServicebusTopicAuthorizationRules(c *AzurermV1alpha1Client, namespace string) *servicebusTopicAuthorizationRules {
	return &servicebusTopicAuthorizationRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the servicebusTopicAuthorizationRule, and returns the corresponding servicebusTopicAuthorizationRule object, and an error if there is any.
func (c *servicebusTopicAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.ServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusTopicAuthorizationRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServicebusTopicAuthorizationRules that match those selectors.
func (c *servicebusTopicAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.ServicebusTopicAuthorizationRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServicebusTopicAuthorizationRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested servicebusTopicAuthorizationRules.
func (c *servicebusTopicAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a servicebusTopicAuthorizationRule and creates it.  Returns the server's representation of the servicebusTopicAuthorizationRule, and an error, if there is any.
func (c *servicebusTopicAuthorizationRules) Create(servicebusTopicAuthorizationRule *v1alpha1.ServicebusTopicAuthorizationRule) (result *v1alpha1.ServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusTopicAuthorizationRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		Body(servicebusTopicAuthorizationRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a servicebusTopicAuthorizationRule and updates it. Returns the server's representation of the servicebusTopicAuthorizationRule, and an error, if there is any.
func (c *servicebusTopicAuthorizationRules) Update(servicebusTopicAuthorizationRule *v1alpha1.ServicebusTopicAuthorizationRule) (result *v1alpha1.ServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusTopicAuthorizationRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		Name(servicebusTopicAuthorizationRule.Name).
		Body(servicebusTopicAuthorizationRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *servicebusTopicAuthorizationRules) UpdateStatus(servicebusTopicAuthorizationRule *v1alpha1.ServicebusTopicAuthorizationRule) (result *v1alpha1.ServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusTopicAuthorizationRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		Name(servicebusTopicAuthorizationRule.Name).
		SubResource("status").
		Body(servicebusTopicAuthorizationRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the servicebusTopicAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *servicebusTopicAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *servicebusTopicAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched servicebusTopicAuthorizationRule.
func (c *servicebusTopicAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusTopicAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusTopicAuthorizationRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicebustopicauthorizationrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
