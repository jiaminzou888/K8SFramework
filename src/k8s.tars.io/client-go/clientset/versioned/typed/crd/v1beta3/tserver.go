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

package v1beta3

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "k8s.tars.io/client-go/clientset/versioned/scheme"
	v1beta3 "k8s.tars.io/crd/v1beta3"
)

// TServersGetter has a method to return a TServerInterface.
// A group's client should implement this interface.
type TServersGetter interface {
	TServers(namespace string) TServerInterface
}

// TServerInterface has methods to work with TServer resources.
type TServerInterface interface {
	Create(ctx context.Context, tServer *v1beta3.TServer, opts v1.CreateOptions) (*v1beta3.TServer, error)
	Update(ctx context.Context, tServer *v1beta3.TServer, opts v1.UpdateOptions) (*v1beta3.TServer, error)
	UpdateStatus(ctx context.Context, tServer *v1beta3.TServer, opts v1.UpdateOptions) (*v1beta3.TServer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta3.TServer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta3.TServerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TServer, err error)
	TServerExpansion
}

// tServers implements TServerInterface
type tServers struct {
	client rest.Interface
	ns     string
}

// newTServers returns a TServers
func newTServers(c *CrdV1beta3Client, namespace string) *tServers {
	return &tServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tServer, and returns the corresponding tServer object, and an error if there is any.
func (c *tServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.TServer, err error) {
	result = &v1beta3.TServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TServers that match those selectors.
func (c *tServers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.TServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta3.TServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tServers.
func (c *tServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tServer and creates it.  Returns the server's representation of the tServer, and an error, if there is any.
func (c *tServers) Create(ctx context.Context, tServer *v1beta3.TServer, opts v1.CreateOptions) (result *v1beta3.TServer, err error) {
	result = &v1beta3.TServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tServer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tServer and updates it. Returns the server's representation of the tServer, and an error, if there is any.
func (c *tServers) Update(ctx context.Context, tServer *v1beta3.TServer, opts v1.UpdateOptions) (result *v1beta3.TServer, err error) {
	result = &v1beta3.TServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tservers").
		Name(tServer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tServer).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tServers) UpdateStatus(ctx context.Context, tServer *v1beta3.TServer, opts v1.UpdateOptions) (result *v1beta3.TServer, err error) {
	result = &v1beta3.TServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tservers").
		Name(tServer.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tServer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tServer and deletes it. Returns an error if one occurs.
func (c *tServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tservers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tservers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tServer.
func (c *tServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TServer, err error) {
	result = &v1beta3.TServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tservers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}