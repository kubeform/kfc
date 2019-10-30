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

// DataLakeAnalyticsAccountsGetter has a method to return a DataLakeAnalyticsAccountInterface.
// A group's client should implement this interface.
type DataLakeAnalyticsAccountsGetter interface {
	DataLakeAnalyticsAccounts(namespace string) DataLakeAnalyticsAccountInterface
}

// DataLakeAnalyticsAccountInterface has methods to work with DataLakeAnalyticsAccount resources.
type DataLakeAnalyticsAccountInterface interface {
	Create(*v1alpha1.DataLakeAnalyticsAccount) (*v1alpha1.DataLakeAnalyticsAccount, error)
	Update(*v1alpha1.DataLakeAnalyticsAccount) (*v1alpha1.DataLakeAnalyticsAccount, error)
	UpdateStatus(*v1alpha1.DataLakeAnalyticsAccount) (*v1alpha1.DataLakeAnalyticsAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DataLakeAnalyticsAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.DataLakeAnalyticsAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeAnalyticsAccount, err error)
	DataLakeAnalyticsAccountExpansion
}

// dataLakeAnalyticsAccounts implements DataLakeAnalyticsAccountInterface
type dataLakeAnalyticsAccounts struct {
	client rest.Interface
	ns     string
}

// newDataLakeAnalyticsAccounts returns a DataLakeAnalyticsAccounts
func newDataLakeAnalyticsAccounts(c *AzurermV1alpha1Client, namespace string) *dataLakeAnalyticsAccounts {
	return &dataLakeAnalyticsAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataLakeAnalyticsAccount, and returns the corresponding dataLakeAnalyticsAccount object, and an error if there is any.
func (c *dataLakeAnalyticsAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	result = &v1alpha1.DataLakeAnalyticsAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataLakeAnalyticsAccounts that match those selectors.
func (c *dataLakeAnalyticsAccounts) List(opts v1.ListOptions) (result *v1alpha1.DataLakeAnalyticsAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DataLakeAnalyticsAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataLakeAnalyticsAccounts.
func (c *dataLakeAnalyticsAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataLakeAnalyticsAccount and creates it.  Returns the server's representation of the dataLakeAnalyticsAccount, and an error, if there is any.
func (c *dataLakeAnalyticsAccounts) Create(dataLakeAnalyticsAccount *v1alpha1.DataLakeAnalyticsAccount) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	result = &v1alpha1.DataLakeAnalyticsAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		Body(dataLakeAnalyticsAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataLakeAnalyticsAccount and updates it. Returns the server's representation of the dataLakeAnalyticsAccount, and an error, if there is any.
func (c *dataLakeAnalyticsAccounts) Update(dataLakeAnalyticsAccount *v1alpha1.DataLakeAnalyticsAccount) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	result = &v1alpha1.DataLakeAnalyticsAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		Name(dataLakeAnalyticsAccount.Name).
		Body(dataLakeAnalyticsAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dataLakeAnalyticsAccounts) UpdateStatus(dataLakeAnalyticsAccount *v1alpha1.DataLakeAnalyticsAccount) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	result = &v1alpha1.DataLakeAnalyticsAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		Name(dataLakeAnalyticsAccount.Name).
		SubResource("status").
		Body(dataLakeAnalyticsAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataLakeAnalyticsAccount and deletes it. Returns an error if one occurs.
func (c *dataLakeAnalyticsAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataLakeAnalyticsAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataLakeAnalyticsAccount.
func (c *dataLakeAnalyticsAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeAnalyticsAccount, err error) {
	result = &v1alpha1.DataLakeAnalyticsAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datalakeanalyticsaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
