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

// ConfigConfigurationRecordersGetter has a method to return a ConfigConfigurationRecorderInterface.
// A group's client should implement this interface.
type ConfigConfigurationRecordersGetter interface {
	ConfigConfigurationRecorders(namespace string) ConfigConfigurationRecorderInterface
}

// ConfigConfigurationRecorderInterface has methods to work with ConfigConfigurationRecorder resources.
type ConfigConfigurationRecorderInterface interface {
	Create(*v1alpha1.ConfigConfigurationRecorder) (*v1alpha1.ConfigConfigurationRecorder, error)
	Update(*v1alpha1.ConfigConfigurationRecorder) (*v1alpha1.ConfigConfigurationRecorder, error)
	UpdateStatus(*v1alpha1.ConfigConfigurationRecorder) (*v1alpha1.ConfigConfigurationRecorder, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ConfigConfigurationRecorder, error)
	List(opts v1.ListOptions) (*v1alpha1.ConfigConfigurationRecorderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigConfigurationRecorder, err error)
	ConfigConfigurationRecorderExpansion
}

// configConfigurationRecorders implements ConfigConfigurationRecorderInterface
type configConfigurationRecorders struct {
	client rest.Interface
	ns     string
}

// newConfigConfigurationRecorders returns a ConfigConfigurationRecorders
func newConfigConfigurationRecorders(c *AwsV1alpha1Client, namespace string) *configConfigurationRecorders {
	return &configConfigurationRecorders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configConfigurationRecorder, and returns the corresponding configConfigurationRecorder object, and an error if there is any.
func (c *configConfigurationRecorders) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigConfigurationRecorder, err error) {
	result = &v1alpha1.ConfigConfigurationRecorder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigConfigurationRecorders that match those selectors.
func (c *configConfigurationRecorders) List(opts v1.ListOptions) (result *v1alpha1.ConfigConfigurationRecorderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigConfigurationRecorderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configConfigurationRecorders.
func (c *configConfigurationRecorders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a configConfigurationRecorder and creates it.  Returns the server's representation of the configConfigurationRecorder, and an error, if there is any.
func (c *configConfigurationRecorders) Create(configConfigurationRecorder *v1alpha1.ConfigConfigurationRecorder) (result *v1alpha1.ConfigConfigurationRecorder, err error) {
	result = &v1alpha1.ConfigConfigurationRecorder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		Body(configConfigurationRecorder).
		Do().
		Into(result)
	return
}

// Update takes the representation of a configConfigurationRecorder and updates it. Returns the server's representation of the configConfigurationRecorder, and an error, if there is any.
func (c *configConfigurationRecorders) Update(configConfigurationRecorder *v1alpha1.ConfigConfigurationRecorder) (result *v1alpha1.ConfigConfigurationRecorder, err error) {
	result = &v1alpha1.ConfigConfigurationRecorder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		Name(configConfigurationRecorder.Name).
		Body(configConfigurationRecorder).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *configConfigurationRecorders) UpdateStatus(configConfigurationRecorder *v1alpha1.ConfigConfigurationRecorder) (result *v1alpha1.ConfigConfigurationRecorder, err error) {
	result = &v1alpha1.ConfigConfigurationRecorder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		Name(configConfigurationRecorder.Name).
		SubResource("status").
		Body(configConfigurationRecorder).
		Do().
		Into(result)
	return
}

// Delete takes name of the configConfigurationRecorder and deletes it. Returns an error if one occurs.
func (c *configConfigurationRecorders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configConfigurationRecorders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched configConfigurationRecorder.
func (c *configConfigurationRecorders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigConfigurationRecorder, err error) {
	result = &v1alpha1.ConfigConfigurationRecorder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configconfigurationrecorders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
