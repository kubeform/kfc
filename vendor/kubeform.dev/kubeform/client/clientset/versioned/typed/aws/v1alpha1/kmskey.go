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

// KmsKeysGetter has a method to return a KmsKeyInterface.
// A group's client should implement this interface.
type KmsKeysGetter interface {
	KmsKeys(namespace string) KmsKeyInterface
}

// KmsKeyInterface has methods to work with KmsKey resources.
type KmsKeyInterface interface {
	Create(*v1alpha1.KmsKey) (*v1alpha1.KmsKey, error)
	Update(*v1alpha1.KmsKey) (*v1alpha1.KmsKey, error)
	UpdateStatus(*v1alpha1.KmsKey) (*v1alpha1.KmsKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KmsKey, error)
	List(opts v1.ListOptions) (*v1alpha1.KmsKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsKey, err error)
	KmsKeyExpansion
}

// kmsKeys implements KmsKeyInterface
type kmsKeys struct {
	client rest.Interface
	ns     string
}

// newKmsKeys returns a KmsKeys
func newKmsKeys(c *AwsV1alpha1Client, namespace string) *kmsKeys {
	return &kmsKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kmsKey, and returns the corresponding kmsKey object, and an error if there is any.
func (c *kmsKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsKey, err error) {
	result = &v1alpha1.KmsKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmskeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KmsKeys that match those selectors.
func (c *kmsKeys) List(opts v1.ListOptions) (result *v1alpha1.KmsKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KmsKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmskeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kmsKeys.
func (c *kmsKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kmskeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kmsKey and creates it.  Returns the server's representation of the kmsKey, and an error, if there is any.
func (c *kmsKeys) Create(kmsKey *v1alpha1.KmsKey) (result *v1alpha1.KmsKey, err error) {
	result = &v1alpha1.KmsKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kmskeys").
		Body(kmsKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kmsKey and updates it. Returns the server's representation of the kmsKey, and an error, if there is any.
func (c *kmsKeys) Update(kmsKey *v1alpha1.KmsKey) (result *v1alpha1.KmsKey, err error) {
	result = &v1alpha1.KmsKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmskeys").
		Name(kmsKey.Name).
		Body(kmsKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kmsKeys) UpdateStatus(kmsKey *v1alpha1.KmsKey) (result *v1alpha1.KmsKey, err error) {
	result = &v1alpha1.KmsKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmskeys").
		Name(kmsKey.Name).
		SubResource("status").
		Body(kmsKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the kmsKey and deletes it. Returns an error if one occurs.
func (c *kmsKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmskeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kmsKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmskeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kmsKey.
func (c *kmsKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsKey, err error) {
	result = &v1alpha1.KmsKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kmskeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
