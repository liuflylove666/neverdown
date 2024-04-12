/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/kzz45/neverdown/pkg/apis/openx/v1"
	scheme "github.com/kzz45/neverdown/pkg/client-go/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TolerationsGetter has a method to return a TolerationInterface.
// A group's client should implement this interface.
type TolerationsGetter interface {
	Tolerations(namespace string) TolerationInterface
}

// TolerationInterface has methods to work with Toleration resources.
type TolerationInterface interface {
	Create(ctx context.Context, toleration *v1.Toleration, opts metav1.CreateOptions) (*v1.Toleration, error)
	Update(ctx context.Context, toleration *v1.Toleration, opts metav1.UpdateOptions) (*v1.Toleration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Toleration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TolerationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Toleration, err error)
	TolerationExpansion
}

// tolerations implements TolerationInterface
type tolerations struct {
	client rest.Interface
	ns     string
}

// newTolerations returns a Tolerations
func newTolerations(c *OpenxV1Client, namespace string) *tolerations {
	return &tolerations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the toleration, and returns the corresponding toleration object, and an error if there is any.
func (c *tolerations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Toleration, err error) {
	result = &v1.Toleration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tolerations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tolerations that match those selectors.
func (c *tolerations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TolerationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TolerationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tolerations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tolerations.
func (c *tolerations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tolerations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a toleration and creates it.  Returns the server's representation of the toleration, and an error, if there is any.
func (c *tolerations) Create(ctx context.Context, toleration *v1.Toleration, opts metav1.CreateOptions) (result *v1.Toleration, err error) {
	result = &v1.Toleration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tolerations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(toleration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a toleration and updates it. Returns the server's representation of the toleration, and an error, if there is any.
func (c *tolerations) Update(ctx context.Context, toleration *v1.Toleration, opts metav1.UpdateOptions) (result *v1.Toleration, err error) {
	result = &v1.Toleration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tolerations").
		Name(toleration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(toleration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the toleration and deletes it. Returns an error if one occurs.
func (c *tolerations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tolerations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tolerations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tolerations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched toleration.
func (c *tolerations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Toleration, err error) {
	result = &v1.Toleration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tolerations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}