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

// VirtualMachineExtensionsGetter has a method to return a VirtualMachineExtensionInterface.
// A group's client should implement this interface.
type VirtualMachineExtensionsGetter interface {
	VirtualMachineExtensions(namespace string) VirtualMachineExtensionInterface
}

// VirtualMachineExtensionInterface has methods to work with VirtualMachineExtension resources.
type VirtualMachineExtensionInterface interface {
	Create(*v1alpha1.VirtualMachineExtension) (*v1alpha1.VirtualMachineExtension, error)
	Update(*v1alpha1.VirtualMachineExtension) (*v1alpha1.VirtualMachineExtension, error)
	UpdateStatus(*v1alpha1.VirtualMachineExtension) (*v1alpha1.VirtualMachineExtension, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualMachineExtension, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualMachineExtensionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineExtension, err error)
	VirtualMachineExtensionExpansion
}

// virtualMachineExtensions implements VirtualMachineExtensionInterface
type virtualMachineExtensions struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineExtensions returns a VirtualMachineExtensions
func newVirtualMachineExtensions(c *AzurermV1alpha1Client, namespace string) *virtualMachineExtensions {
	return &virtualMachineExtensions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineExtension, and returns the corresponding virtualMachineExtension object, and an error if there is any.
func (c *virtualMachineExtensions) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineExtension, err error) {
	result = &v1alpha1.VirtualMachineExtension{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineExtensions that match those selectors.
func (c *virtualMachineExtensions) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineExtensionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineExtensionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineExtensions.
func (c *virtualMachineExtensions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineExtension and creates it.  Returns the server's representation of the virtualMachineExtension, and an error, if there is any.
func (c *virtualMachineExtensions) Create(virtualMachineExtension *v1alpha1.VirtualMachineExtension) (result *v1alpha1.VirtualMachineExtension, err error) {
	result = &v1alpha1.VirtualMachineExtension{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		Body(virtualMachineExtension).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineExtension and updates it. Returns the server's representation of the virtualMachineExtension, and an error, if there is any.
func (c *virtualMachineExtensions) Update(virtualMachineExtension *v1alpha1.VirtualMachineExtension) (result *v1alpha1.VirtualMachineExtension, err error) {
	result = &v1alpha1.VirtualMachineExtension{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		Name(virtualMachineExtension.Name).
		Body(virtualMachineExtension).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineExtensions) UpdateStatus(virtualMachineExtension *v1alpha1.VirtualMachineExtension) (result *v1alpha1.VirtualMachineExtension, err error) {
	result = &v1alpha1.VirtualMachineExtension{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		Name(virtualMachineExtension.Name).
		SubResource("status").
		Body(virtualMachineExtension).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineExtension and deletes it. Returns an error if one occurs.
func (c *virtualMachineExtensions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineExtensions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineExtension.
func (c *virtualMachineExtensions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineExtension, err error) {
	result = &v1alpha1.VirtualMachineExtension{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachineextensions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}