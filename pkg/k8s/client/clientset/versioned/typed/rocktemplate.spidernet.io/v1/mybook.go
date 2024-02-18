// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1"
	scheme "github.com/koffloader-io/koffloader/pkg/k8s/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MybooksGetter has a method to return a MybookInterface.
// A group's client should implement this interface.
type MybooksGetter interface {
	Mybooks() MybookInterface
}

// MybookInterface has methods to work with Mybook resources.
type MybookInterface interface {
	Create(ctx context.Context, mybook *v1.Mybook, opts metav1.CreateOptions) (*v1.Mybook, error)
	Update(ctx context.Context, mybook *v1.Mybook, opts metav1.UpdateOptions) (*v1.Mybook, error)
	UpdateStatus(ctx context.Context, mybook *v1.Mybook, opts metav1.UpdateOptions) (*v1.Mybook, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Mybook, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MybookList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Mybook, err error)
	MybookExpansion
}

// mybooks implements MybookInterface
type mybooks struct {
	client rest.Interface
}

// newMybooks returns a Mybooks
func newMybooks(c *koffloaderV1Client) *mybooks {
	return &mybooks{
		client: c.RESTClient(),
	}
}

// Get takes name of the mybook, and returns the corresponding mybook object, and an error if there is any.
func (c *mybooks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Mybook, err error) {
	result = &v1.Mybook{}
	err = c.client.Get().
		Resource("mybooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Mybooks that match those selectors.
func (c *mybooks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MybookList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MybookList{}
	err = c.client.Get().
		Resource("mybooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mybooks.
func (c *mybooks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("mybooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mybook and creates it.  Returns the server's representation of the mybook, and an error, if there is any.
func (c *mybooks) Create(ctx context.Context, mybook *v1.Mybook, opts metav1.CreateOptions) (result *v1.Mybook, err error) {
	result = &v1.Mybook{}
	err = c.client.Post().
		Resource("mybooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mybook).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mybook and updates it. Returns the server's representation of the mybook, and an error, if there is any.
func (c *mybooks) Update(ctx context.Context, mybook *v1.Mybook, opts metav1.UpdateOptions) (result *v1.Mybook, err error) {
	result = &v1.Mybook{}
	err = c.client.Put().
		Resource("mybooks").
		Name(mybook.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mybook).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *mybooks) UpdateStatus(ctx context.Context, mybook *v1.Mybook, opts metav1.UpdateOptions) (result *v1.Mybook, err error) {
	result = &v1.Mybook{}
	err = c.client.Put().
		Resource("mybooks").
		Name(mybook.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mybook).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mybook and deletes it. Returns an error if one occurs.
func (c *mybooks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("mybooks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mybooks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("mybooks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mybook.
func (c *mybooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Mybook, err error) {
	result = &v1.Mybook{}
	err = c.client.Patch(pt).
		Resource("mybooks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
