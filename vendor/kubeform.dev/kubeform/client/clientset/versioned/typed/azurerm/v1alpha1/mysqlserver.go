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

// MysqlServersGetter has a method to return a MysqlServerInterface.
// A group's client should implement this interface.
type MysqlServersGetter interface {
	MysqlServers(namespace string) MysqlServerInterface
}

// MysqlServerInterface has methods to work with MysqlServer resources.
type MysqlServerInterface interface {
	Create(*v1alpha1.MysqlServer) (*v1alpha1.MysqlServer, error)
	Update(*v1alpha1.MysqlServer) (*v1alpha1.MysqlServer, error)
	UpdateStatus(*v1alpha1.MysqlServer) (*v1alpha1.MysqlServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MysqlServer, error)
	List(opts v1.ListOptions) (*v1alpha1.MysqlServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlServer, err error)
	MysqlServerExpansion
}

// mysqlServers implements MysqlServerInterface
type mysqlServers struct {
	client rest.Interface
	ns     string
}

// newMysqlServers returns a MysqlServers
func newMysqlServers(c *AzurermV1alpha1Client, namespace string) *mysqlServers {
	return &mysqlServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mysqlServer, and returns the corresponding mysqlServer object, and an error if there is any.
func (c *mysqlServers) Get(name string, options v1.GetOptions) (result *v1alpha1.MysqlServer, err error) {
	result = &v1alpha1.MysqlServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MysqlServers that match those selectors.
func (c *mysqlServers) List(opts v1.ListOptions) (result *v1alpha1.MysqlServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mysqlServers.
func (c *mysqlServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mysqlServer and creates it.  Returns the server's representation of the mysqlServer, and an error, if there is any.
func (c *mysqlServers) Create(mysqlServer *v1alpha1.MysqlServer) (result *v1alpha1.MysqlServer, err error) {
	result = &v1alpha1.MysqlServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlservers").
		Body(mysqlServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mysqlServer and updates it. Returns the server's representation of the mysqlServer, and an error, if there is any.
func (c *mysqlServers) Update(mysqlServer *v1alpha1.MysqlServer) (result *v1alpha1.MysqlServer, err error) {
	result = &v1alpha1.MysqlServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlservers").
		Name(mysqlServer.Name).
		Body(mysqlServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mysqlServers) UpdateStatus(mysqlServer *v1alpha1.MysqlServer) (result *v1alpha1.MysqlServer, err error) {
	result = &v1alpha1.MysqlServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlservers").
		Name(mysqlServer.Name).
		SubResource("status").
		Body(mysqlServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the mysqlServer and deletes it. Returns an error if one occurs.
func (c *mysqlServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mysqlServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mysqlServer.
func (c *mysqlServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlServer, err error) {
	result = &v1alpha1.MysqlServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
