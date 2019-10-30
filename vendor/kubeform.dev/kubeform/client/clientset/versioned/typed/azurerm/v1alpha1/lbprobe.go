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

// LbProbesGetter has a method to return a LbProbeInterface.
// A group's client should implement this interface.
type LbProbesGetter interface {
	LbProbes(namespace string) LbProbeInterface
}

// LbProbeInterface has methods to work with LbProbe resources.
type LbProbeInterface interface {
	Create(*v1alpha1.LbProbe) (*v1alpha1.LbProbe, error)
	Update(*v1alpha1.LbProbe) (*v1alpha1.LbProbe, error)
	UpdateStatus(*v1alpha1.LbProbe) (*v1alpha1.LbProbe, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LbProbe, error)
	List(opts v1.ListOptions) (*v1alpha1.LbProbeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbProbe, err error)
	LbProbeExpansion
}

// lbProbes implements LbProbeInterface
type lbProbes struct {
	client rest.Interface
	ns     string
}

// newLbProbes returns a LbProbes
func newLbProbes(c *AzurermV1alpha1Client, namespace string) *lbProbes {
	return &lbProbes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lbProbe, and returns the corresponding lbProbe object, and an error if there is any.
func (c *lbProbes) Get(name string, options v1.GetOptions) (result *v1alpha1.LbProbe, err error) {
	result = &v1alpha1.LbProbe{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbprobes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LbProbes that match those selectors.
func (c *lbProbes) List(opts v1.ListOptions) (result *v1alpha1.LbProbeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LbProbeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbprobes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lbProbes.
func (c *lbProbes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lbprobes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lbProbe and creates it.  Returns the server's representation of the lbProbe, and an error, if there is any.
func (c *lbProbes) Create(lbProbe *v1alpha1.LbProbe) (result *v1alpha1.LbProbe, err error) {
	result = &v1alpha1.LbProbe{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lbprobes").
		Body(lbProbe).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lbProbe and updates it. Returns the server's representation of the lbProbe, and an error, if there is any.
func (c *lbProbes) Update(lbProbe *v1alpha1.LbProbe) (result *v1alpha1.LbProbe, err error) {
	result = &v1alpha1.LbProbe{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbprobes").
		Name(lbProbe.Name).
		Body(lbProbe).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lbProbes) UpdateStatus(lbProbe *v1alpha1.LbProbe) (result *v1alpha1.LbProbe, err error) {
	result = &v1alpha1.LbProbe{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbprobes").
		Name(lbProbe.Name).
		SubResource("status").
		Body(lbProbe).
		Do().
		Into(result)
	return
}

// Delete takes name of the lbProbe and deletes it. Returns an error if one occurs.
func (c *lbProbes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbprobes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lbProbes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbprobes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lbProbe.
func (c *lbProbes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbProbe, err error) {
	result = &v1alpha1.LbProbe{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lbprobes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
