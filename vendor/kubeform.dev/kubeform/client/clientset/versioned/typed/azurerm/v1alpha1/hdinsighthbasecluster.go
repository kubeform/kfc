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

// HdinsightHbaseClustersGetter has a method to return a HdinsightHbaseClusterInterface.
// A group's client should implement this interface.
type HdinsightHbaseClustersGetter interface {
	HdinsightHbaseClusters(namespace string) HdinsightHbaseClusterInterface
}

// HdinsightHbaseClusterInterface has methods to work with HdinsightHbaseCluster resources.
type HdinsightHbaseClusterInterface interface {
	Create(*v1alpha1.HdinsightHbaseCluster) (*v1alpha1.HdinsightHbaseCluster, error)
	Update(*v1alpha1.HdinsightHbaseCluster) (*v1alpha1.HdinsightHbaseCluster, error)
	UpdateStatus(*v1alpha1.HdinsightHbaseCluster) (*v1alpha1.HdinsightHbaseCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HdinsightHbaseCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.HdinsightHbaseClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HdinsightHbaseCluster, err error)
	HdinsightHbaseClusterExpansion
}

// hdinsightHbaseClusters implements HdinsightHbaseClusterInterface
type hdinsightHbaseClusters struct {
	client rest.Interface
	ns     string
}

// newHdinsightHbaseClusters returns a HdinsightHbaseClusters
func newHdinsightHbaseClusters(c *AzurermV1alpha1Client, namespace string) *hdinsightHbaseClusters {
	return &hdinsightHbaseClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hdinsightHbaseCluster, and returns the corresponding hdinsightHbaseCluster object, and an error if there is any.
func (c *hdinsightHbaseClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.HdinsightHbaseCluster, err error) {
	result = &v1alpha1.HdinsightHbaseCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HdinsightHbaseClusters that match those selectors.
func (c *hdinsightHbaseClusters) List(opts v1.ListOptions) (result *v1alpha1.HdinsightHbaseClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HdinsightHbaseClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hdinsightHbaseClusters.
func (c *hdinsightHbaseClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a hdinsightHbaseCluster and creates it.  Returns the server's representation of the hdinsightHbaseCluster, and an error, if there is any.
func (c *hdinsightHbaseClusters) Create(hdinsightHbaseCluster *v1alpha1.HdinsightHbaseCluster) (result *v1alpha1.HdinsightHbaseCluster, err error) {
	result = &v1alpha1.HdinsightHbaseCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		Body(hdinsightHbaseCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hdinsightHbaseCluster and updates it. Returns the server's representation of the hdinsightHbaseCluster, and an error, if there is any.
func (c *hdinsightHbaseClusters) Update(hdinsightHbaseCluster *v1alpha1.HdinsightHbaseCluster) (result *v1alpha1.HdinsightHbaseCluster, err error) {
	result = &v1alpha1.HdinsightHbaseCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		Name(hdinsightHbaseCluster.Name).
		Body(hdinsightHbaseCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *hdinsightHbaseClusters) UpdateStatus(hdinsightHbaseCluster *v1alpha1.HdinsightHbaseCluster) (result *v1alpha1.HdinsightHbaseCluster, err error) {
	result = &v1alpha1.HdinsightHbaseCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		Name(hdinsightHbaseCluster.Name).
		SubResource("status").
		Body(hdinsightHbaseCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the hdinsightHbaseCluster and deletes it. Returns an error if one occurs.
func (c *hdinsightHbaseClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hdinsightHbaseClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hdinsightHbaseCluster.
func (c *hdinsightHbaseClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HdinsightHbaseCluster, err error) {
	result = &v1alpha1.HdinsightHbaseCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hdinsighthbaseclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
