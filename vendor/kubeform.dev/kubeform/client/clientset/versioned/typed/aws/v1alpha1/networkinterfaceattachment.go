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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkInterfaceAttachmentsGetter has a method to return a NetworkInterfaceAttachmentInterface.
// A group's client should implement this interface.
type NetworkInterfaceAttachmentsGetter interface {
	NetworkInterfaceAttachments(namespace string) NetworkInterfaceAttachmentInterface
}

// NetworkInterfaceAttachmentInterface has methods to work with NetworkInterfaceAttachment resources.
type NetworkInterfaceAttachmentInterface interface {
	Create(*v1alpha1.NetworkInterfaceAttachment) (*v1alpha1.NetworkInterfaceAttachment, error)
	Update(*v1alpha1.NetworkInterfaceAttachment) (*v1alpha1.NetworkInterfaceAttachment, error)
	UpdateStatus(*v1alpha1.NetworkInterfaceAttachment) (*v1alpha1.NetworkInterfaceAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkInterfaceAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkInterfaceAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterfaceAttachment, err error)
	NetworkInterfaceAttachmentExpansion
}

// networkInterfaceAttachments implements NetworkInterfaceAttachmentInterface
type networkInterfaceAttachments struct {
	client rest.Interface
	ns     string
}

// newNetworkInterfaceAttachments returns a NetworkInterfaceAttachments
func newNetworkInterfaceAttachments(c *AwsV1alpha1Client, namespace string) *networkInterfaceAttachments {
	return &networkInterfaceAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkInterfaceAttachment, and returns the corresponding networkInterfaceAttachment object, and an error if there is any.
func (c *networkInterfaceAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkInterfaceAttachments that match those selectors.
func (c *networkInterfaceAttachments) List(opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkInterfaceAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkInterfaceAttachments.
func (c *networkInterfaceAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkInterfaceAttachment and creates it.  Returns the server's representation of the networkInterfaceAttachment, and an error, if there is any.
func (c *networkInterfaceAttachments) Create(networkInterfaceAttachment *v1alpha1.NetworkInterfaceAttachment) (result *v1alpha1.NetworkInterfaceAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		Body(networkInterfaceAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkInterfaceAttachment and updates it. Returns the server's representation of the networkInterfaceAttachment, and an error, if there is any.
func (c *networkInterfaceAttachments) Update(networkInterfaceAttachment *v1alpha1.NetworkInterfaceAttachment) (result *v1alpha1.NetworkInterfaceAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		Name(networkInterfaceAttachment.Name).
		Body(networkInterfaceAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkInterfaceAttachments) UpdateStatus(networkInterfaceAttachment *v1alpha1.NetworkInterfaceAttachment) (result *v1alpha1.NetworkInterfaceAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		Name(networkInterfaceAttachment.Name).
		SubResource("status").
		Body(networkInterfaceAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkInterfaceAttachment and deletes it. Returns an error if one occurs.
func (c *networkInterfaceAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkInterfaceAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkInterfaceAttachment.
func (c *networkInterfaceAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterfaceAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkinterfaceattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}