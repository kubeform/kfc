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

// ConfigAggregateAuthorizationsGetter has a method to return a ConfigAggregateAuthorizationInterface.
// A group's client should implement this interface.
type ConfigAggregateAuthorizationsGetter interface {
	ConfigAggregateAuthorizations(namespace string) ConfigAggregateAuthorizationInterface
}

// ConfigAggregateAuthorizationInterface has methods to work with ConfigAggregateAuthorization resources.
type ConfigAggregateAuthorizationInterface interface {
	Create(*v1alpha1.ConfigAggregateAuthorization) (*v1alpha1.ConfigAggregateAuthorization, error)
	Update(*v1alpha1.ConfigAggregateAuthorization) (*v1alpha1.ConfigAggregateAuthorization, error)
	UpdateStatus(*v1alpha1.ConfigAggregateAuthorization) (*v1alpha1.ConfigAggregateAuthorization, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ConfigAggregateAuthorization, error)
	List(opts v1.ListOptions) (*v1alpha1.ConfigAggregateAuthorizationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigAggregateAuthorization, err error)
	ConfigAggregateAuthorizationExpansion
}

// configAggregateAuthorizations implements ConfigAggregateAuthorizationInterface
type configAggregateAuthorizations struct {
	client rest.Interface
	ns     string
}

// newConfigAggregateAuthorizations returns a ConfigAggregateAuthorizations
func newConfigAggregateAuthorizations(c *AwsV1alpha1Client, namespace string) *configAggregateAuthorizations {
	return &configAggregateAuthorizations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configAggregateAuthorization, and returns the corresponding configAggregateAuthorization object, and an error if there is any.
func (c *configAggregateAuthorizations) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	result = &v1alpha1.ConfigAggregateAuthorization{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigAggregateAuthorizations that match those selectors.
func (c *configAggregateAuthorizations) List(opts v1.ListOptions) (result *v1alpha1.ConfigAggregateAuthorizationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigAggregateAuthorizationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configAggregateAuthorizations.
func (c *configAggregateAuthorizations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a configAggregateAuthorization and creates it.  Returns the server's representation of the configAggregateAuthorization, and an error, if there is any.
func (c *configAggregateAuthorizations) Create(configAggregateAuthorization *v1alpha1.ConfigAggregateAuthorization) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	result = &v1alpha1.ConfigAggregateAuthorization{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		Body(configAggregateAuthorization).
		Do().
		Into(result)
	return
}

// Update takes the representation of a configAggregateAuthorization and updates it. Returns the server's representation of the configAggregateAuthorization, and an error, if there is any.
func (c *configAggregateAuthorizations) Update(configAggregateAuthorization *v1alpha1.ConfigAggregateAuthorization) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	result = &v1alpha1.ConfigAggregateAuthorization{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		Name(configAggregateAuthorization.Name).
		Body(configAggregateAuthorization).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *configAggregateAuthorizations) UpdateStatus(configAggregateAuthorization *v1alpha1.ConfigAggregateAuthorization) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	result = &v1alpha1.ConfigAggregateAuthorization{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		Name(configAggregateAuthorization.Name).
		SubResource("status").
		Body(configAggregateAuthorization).
		Do().
		Into(result)
	return
}

// Delete takes name of the configAggregateAuthorization and deletes it. Returns an error if one occurs.
func (c *configAggregateAuthorizations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configAggregateAuthorizations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched configAggregateAuthorization.
func (c *configAggregateAuthorizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	result = &v1alpha1.ConfigAggregateAuthorization{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configaggregateauthorizations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
