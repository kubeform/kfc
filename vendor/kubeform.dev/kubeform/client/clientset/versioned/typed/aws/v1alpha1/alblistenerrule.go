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
	"context"
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AlbListenerRulesGetter has a method to return a AlbListenerRuleInterface.
// A group's client should implement this interface.
type AlbListenerRulesGetter interface {
	AlbListenerRules(namespace string) AlbListenerRuleInterface
}

// AlbListenerRuleInterface has methods to work with AlbListenerRule resources.
type AlbListenerRuleInterface interface {
	Create(ctx context.Context, albListenerRule *v1alpha1.AlbListenerRule, opts v1.CreateOptions) (*v1alpha1.AlbListenerRule, error)
	Update(ctx context.Context, albListenerRule *v1alpha1.AlbListenerRule, opts v1.UpdateOptions) (*v1alpha1.AlbListenerRule, error)
	UpdateStatus(ctx context.Context, albListenerRule *v1alpha1.AlbListenerRule, opts v1.UpdateOptions) (*v1alpha1.AlbListenerRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AlbListenerRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AlbListenerRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlbListenerRule, err error)
	AlbListenerRuleExpansion
}

// albListenerRules implements AlbListenerRuleInterface
type albListenerRules struct {
	client rest.Interface
	ns     string
}

// newAlbListenerRules returns a AlbListenerRules
func newAlbListenerRules(c *AwsV1alpha1Client, namespace string) *albListenerRules {
	return &albListenerRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the albListenerRule, and returns the corresponding albListenerRule object, and an error if there is any.
func (c *albListenerRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AlbListenerRules that match those selectors.
func (c *albListenerRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlbListenerRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AlbListenerRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alblistenerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested albListenerRules.
func (c *albListenerRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alblistenerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a albListenerRule and creates it.  Returns the server's representation of the albListenerRule, and an error, if there is any.
func (c *albListenerRules) Create(ctx context.Context, albListenerRule *v1alpha1.AlbListenerRule, opts v1.CreateOptions) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alblistenerrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(albListenerRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a albListenerRule and updates it. Returns the server's representation of the albListenerRule, and an error, if there is any.
func (c *albListenerRules) Update(ctx context.Context, albListenerRule *v1alpha1.AlbListenerRule, opts v1.UpdateOptions) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(albListenerRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(albListenerRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *albListenerRules) UpdateStatus(ctx context.Context, albListenerRule *v1alpha1.AlbListenerRule, opts v1.UpdateOptions) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(albListenerRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(albListenerRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the albListenerRule and deletes it. Returns an error if one occurs.
func (c *albListenerRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *albListenerRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alblistenerrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched albListenerRule.
func (c *albListenerRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlbListenerRule, err error) {
	result = &v1alpha1.AlbListenerRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alblistenerrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
