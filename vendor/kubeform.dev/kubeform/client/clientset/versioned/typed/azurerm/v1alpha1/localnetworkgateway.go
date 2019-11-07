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

// LocalNetworkGatewaysGetter has a method to return a LocalNetworkGatewayInterface.
// A group's client should implement this interface.
type LocalNetworkGatewaysGetter interface {
	LocalNetworkGateways(namespace string) LocalNetworkGatewayInterface
}

// LocalNetworkGatewayInterface has methods to work with LocalNetworkGateway resources.
type LocalNetworkGatewayInterface interface {
	Create(*v1alpha1.LocalNetworkGateway) (*v1alpha1.LocalNetworkGateway, error)
	Update(*v1alpha1.LocalNetworkGateway) (*v1alpha1.LocalNetworkGateway, error)
	UpdateStatus(*v1alpha1.LocalNetworkGateway) (*v1alpha1.LocalNetworkGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LocalNetworkGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.LocalNetworkGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LocalNetworkGateway, err error)
	LocalNetworkGatewayExpansion
}

// localNetworkGateways implements LocalNetworkGatewayInterface
type localNetworkGateways struct {
	client rest.Interface
	ns     string
}

// newLocalNetworkGateways returns a LocalNetworkGateways
func newLocalNetworkGateways(c *AzurermV1alpha1Client, namespace string) *localNetworkGateways {
	return &localNetworkGateways{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the localNetworkGateway, and returns the corresponding localNetworkGateway object, and an error if there is any.
func (c *localNetworkGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.LocalNetworkGateway, err error) {
	result = &v1alpha1.LocalNetworkGateway{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalNetworkGateways that match those selectors.
func (c *localNetworkGateways) List(opts v1.ListOptions) (result *v1alpha1.LocalNetworkGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LocalNetworkGatewayList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localNetworkGateways.
func (c *localNetworkGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a localNetworkGateway and creates it.  Returns the server's representation of the localNetworkGateway, and an error, if there is any.
func (c *localNetworkGateways) Create(localNetworkGateway *v1alpha1.LocalNetworkGateway) (result *v1alpha1.LocalNetworkGateway, err error) {
	result = &v1alpha1.LocalNetworkGateway{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		Body(localNetworkGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a localNetworkGateway and updates it. Returns the server's representation of the localNetworkGateway, and an error, if there is any.
func (c *localNetworkGateways) Update(localNetworkGateway *v1alpha1.LocalNetworkGateway) (result *v1alpha1.LocalNetworkGateway, err error) {
	result = &v1alpha1.LocalNetworkGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		Name(localNetworkGateway.Name).
		Body(localNetworkGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *localNetworkGateways) UpdateStatus(localNetworkGateway *v1alpha1.LocalNetworkGateway) (result *v1alpha1.LocalNetworkGateway, err error) {
	result = &v1alpha1.LocalNetworkGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		Name(localNetworkGateway.Name).
		SubResource("status").
		Body(localNetworkGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the localNetworkGateway and deletes it. Returns an error if one occurs.
func (c *localNetworkGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localNetworkGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localnetworkgateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched localNetworkGateway.
func (c *localNetworkGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LocalNetworkGateway, err error) {
	result = &v1alpha1.LocalNetworkGateway{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("localnetworkgateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}