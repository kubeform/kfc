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

// CloudfrontPublicKeysGetter has a method to return a CloudfrontPublicKeyInterface.
// A group's client should implement this interface.
type CloudfrontPublicKeysGetter interface {
	CloudfrontPublicKeys(namespace string) CloudfrontPublicKeyInterface
}

// CloudfrontPublicKeyInterface has methods to work with CloudfrontPublicKey resources.
type CloudfrontPublicKeyInterface interface {
	Create(*v1alpha1.CloudfrontPublicKey) (*v1alpha1.CloudfrontPublicKey, error)
	Update(*v1alpha1.CloudfrontPublicKey) (*v1alpha1.CloudfrontPublicKey, error)
	UpdateStatus(*v1alpha1.CloudfrontPublicKey) (*v1alpha1.CloudfrontPublicKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudfrontPublicKey, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudfrontPublicKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfrontPublicKey, err error)
	CloudfrontPublicKeyExpansion
}

// cloudfrontPublicKeys implements CloudfrontPublicKeyInterface
type cloudfrontPublicKeys struct {
	client rest.Interface
	ns     string
}

// newCloudfrontPublicKeys returns a CloudfrontPublicKeys
func newCloudfrontPublicKeys(c *AwsV1alpha1Client, namespace string) *cloudfrontPublicKeys {
	return &cloudfrontPublicKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudfrontPublicKey, and returns the corresponding cloudfrontPublicKey object, and an error if there is any.
func (c *cloudfrontPublicKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudfrontPublicKey, err error) {
	result = &v1alpha1.CloudfrontPublicKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudfrontPublicKeys that match those selectors.
func (c *cloudfrontPublicKeys) List(opts v1.ListOptions) (result *v1alpha1.CloudfrontPublicKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudfrontPublicKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudfrontPublicKeys.
func (c *cloudfrontPublicKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudfrontPublicKey and creates it.  Returns the server's representation of the cloudfrontPublicKey, and an error, if there is any.
func (c *cloudfrontPublicKeys) Create(cloudfrontPublicKey *v1alpha1.CloudfrontPublicKey) (result *v1alpha1.CloudfrontPublicKey, err error) {
	result = &v1alpha1.CloudfrontPublicKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		Body(cloudfrontPublicKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudfrontPublicKey and updates it. Returns the server's representation of the cloudfrontPublicKey, and an error, if there is any.
func (c *cloudfrontPublicKeys) Update(cloudfrontPublicKey *v1alpha1.CloudfrontPublicKey) (result *v1alpha1.CloudfrontPublicKey, err error) {
	result = &v1alpha1.CloudfrontPublicKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		Name(cloudfrontPublicKey.Name).
		Body(cloudfrontPublicKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudfrontPublicKeys) UpdateStatus(cloudfrontPublicKey *v1alpha1.CloudfrontPublicKey) (result *v1alpha1.CloudfrontPublicKey, err error) {
	result = &v1alpha1.CloudfrontPublicKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		Name(cloudfrontPublicKey.Name).
		SubResource("status").
		Body(cloudfrontPublicKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudfrontPublicKey and deletes it. Returns an error if one occurs.
func (c *cloudfrontPublicKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudfrontPublicKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudfrontPublicKey.
func (c *cloudfrontPublicKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfrontPublicKey, err error) {
	result = &v1alpha1.CloudfrontPublicKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudfrontpublickeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
