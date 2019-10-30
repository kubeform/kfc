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

// LambdaFunctionsGetter has a method to return a LambdaFunctionInterface.
// A group's client should implement this interface.
type LambdaFunctionsGetter interface {
	LambdaFunctions(namespace string) LambdaFunctionInterface
}

// LambdaFunctionInterface has methods to work with LambdaFunction resources.
type LambdaFunctionInterface interface {
	Create(*v1alpha1.LambdaFunction) (*v1alpha1.LambdaFunction, error)
	Update(*v1alpha1.LambdaFunction) (*v1alpha1.LambdaFunction, error)
	UpdateStatus(*v1alpha1.LambdaFunction) (*v1alpha1.LambdaFunction, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LambdaFunction, error)
	List(opts v1.ListOptions) (*v1alpha1.LambdaFunctionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LambdaFunction, err error)
	LambdaFunctionExpansion
}

// lambdaFunctions implements LambdaFunctionInterface
type lambdaFunctions struct {
	client rest.Interface
	ns     string
}

// newLambdaFunctions returns a LambdaFunctions
func newLambdaFunctions(c *AwsV1alpha1Client, namespace string) *lambdaFunctions {
	return &lambdaFunctions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lambdaFunction, and returns the corresponding lambdaFunction object, and an error if there is any.
func (c *lambdaFunctions) Get(name string, options v1.GetOptions) (result *v1alpha1.LambdaFunction, err error) {
	result = &v1alpha1.LambdaFunction{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lambdafunctions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LambdaFunctions that match those selectors.
func (c *lambdaFunctions) List(opts v1.ListOptions) (result *v1alpha1.LambdaFunctionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LambdaFunctionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lambdafunctions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lambdaFunctions.
func (c *lambdaFunctions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lambdafunctions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lambdaFunction and creates it.  Returns the server's representation of the lambdaFunction, and an error, if there is any.
func (c *lambdaFunctions) Create(lambdaFunction *v1alpha1.LambdaFunction) (result *v1alpha1.LambdaFunction, err error) {
	result = &v1alpha1.LambdaFunction{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lambdafunctions").
		Body(lambdaFunction).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lambdaFunction and updates it. Returns the server's representation of the lambdaFunction, and an error, if there is any.
func (c *lambdaFunctions) Update(lambdaFunction *v1alpha1.LambdaFunction) (result *v1alpha1.LambdaFunction, err error) {
	result = &v1alpha1.LambdaFunction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lambdafunctions").
		Name(lambdaFunction.Name).
		Body(lambdaFunction).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lambdaFunctions) UpdateStatus(lambdaFunction *v1alpha1.LambdaFunction) (result *v1alpha1.LambdaFunction, err error) {
	result = &v1alpha1.LambdaFunction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lambdafunctions").
		Name(lambdaFunction.Name).
		SubResource("status").
		Body(lambdaFunction).
		Do().
		Into(result)
	return
}

// Delete takes name of the lambdaFunction and deletes it. Returns an error if one occurs.
func (c *lambdaFunctions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lambdafunctions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lambdaFunctions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lambdafunctions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lambdaFunction.
func (c *lambdaFunctions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LambdaFunction, err error) {
	result = &v1alpha1.LambdaFunction{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lambdafunctions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
