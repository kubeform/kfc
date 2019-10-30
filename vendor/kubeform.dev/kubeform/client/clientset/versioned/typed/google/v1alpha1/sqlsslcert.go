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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SqlSSLCertsGetter has a method to return a SqlSSLCertInterface.
// A group's client should implement this interface.
type SqlSSLCertsGetter interface {
	SqlSSLCerts(namespace string) SqlSSLCertInterface
}

// SqlSSLCertInterface has methods to work with SqlSSLCert resources.
type SqlSSLCertInterface interface {
	Create(*v1alpha1.SqlSSLCert) (*v1alpha1.SqlSSLCert, error)
	Update(*v1alpha1.SqlSSLCert) (*v1alpha1.SqlSSLCert, error)
	UpdateStatus(*v1alpha1.SqlSSLCert) (*v1alpha1.SqlSSLCert, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SqlSSLCert, error)
	List(opts v1.ListOptions) (*v1alpha1.SqlSSLCertList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlSSLCert, err error)
	SqlSSLCertExpansion
}

// sqlSSLCerts implements SqlSSLCertInterface
type sqlSSLCerts struct {
	client rest.Interface
	ns     string
}

// newSqlSSLCerts returns a SqlSSLCerts
func newSqlSSLCerts(c *GoogleV1alpha1Client, namespace string) *sqlSSLCerts {
	return &sqlSSLCerts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sqlSSLCert, and returns the corresponding sqlSSLCert object, and an error if there is any.
func (c *sqlSSLCerts) Get(name string, options v1.GetOptions) (result *v1alpha1.SqlSSLCert, err error) {
	result = &v1alpha1.SqlSSLCert{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SqlSSLCerts that match those selectors.
func (c *sqlSSLCerts) List(opts v1.ListOptions) (result *v1alpha1.SqlSSLCertList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SqlSSLCertList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sqlSSLCerts.
func (c *sqlSSLCerts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sqlSSLCert and creates it.  Returns the server's representation of the sqlSSLCert, and an error, if there is any.
func (c *sqlSSLCerts) Create(sqlSSLCert *v1alpha1.SqlSSLCert) (result *v1alpha1.SqlSSLCert, err error) {
	result = &v1alpha1.SqlSSLCert{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		Body(sqlSSLCert).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sqlSSLCert and updates it. Returns the server's representation of the sqlSSLCert, and an error, if there is any.
func (c *sqlSSLCerts) Update(sqlSSLCert *v1alpha1.SqlSSLCert) (result *v1alpha1.SqlSSLCert, err error) {
	result = &v1alpha1.SqlSSLCert{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		Name(sqlSSLCert.Name).
		Body(sqlSSLCert).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sqlSSLCerts) UpdateStatus(sqlSSLCert *v1alpha1.SqlSSLCert) (result *v1alpha1.SqlSSLCert, err error) {
	result = &v1alpha1.SqlSSLCert{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		Name(sqlSSLCert.Name).
		SubResource("status").
		Body(sqlSSLCert).
		Do().
		Into(result)
	return
}

// Delete takes name of the sqlSSLCert and deletes it. Returns an error if one occurs.
func (c *sqlSSLCerts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sqlSSLCerts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqlsslcerts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sqlSSLCert.
func (c *sqlSSLCerts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlSSLCert, err error) {
	result = &v1alpha1.SqlSSLCert{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sqlsslcerts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
