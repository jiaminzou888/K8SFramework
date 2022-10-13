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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta3 "k8s.tars.io/crd/v1beta3"
)

// FakeTEndpoints implements TEndpointInterface
type FakeTEndpoints struct {
	Fake *FakeCrdV1beta3
	ns   string
}

var tendpointsResource = schema.GroupVersionResource{Group: "k8s.tars.io", Version: "v1beta3", Resource: "tendpoints"}

var tendpointsKind = schema.GroupVersionKind{Group: "k8s.tars.io", Version: "v1beta3", Kind: "TEndpoint"}

// Get takes name of the tEndpoint, and returns the corresponding tEndpoint object, and an error if there is any.
func (c *FakeTEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.TEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tendpointsResource, c.ns, name), &v1beta3.TEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TEndpoint), err
}

// List takes label and field selectors, and returns the list of TEndpoints that match those selectors.
func (c *FakeTEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.TEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tendpointsResource, tendpointsKind, c.ns, opts), &v1beta3.TEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta3.TEndpointList{ListMeta: obj.(*v1beta3.TEndpointList).ListMeta}
	for _, item := range obj.(*v1beta3.TEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tEndpoints.
func (c *FakeTEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tendpointsResource, c.ns, opts))

}

// Create takes the representation of a tEndpoint and creates it.  Returns the server's representation of the tEndpoint, and an error, if there is any.
func (c *FakeTEndpoints) Create(ctx context.Context, tEndpoint *v1beta3.TEndpoint, opts v1.CreateOptions) (result *v1beta3.TEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tendpointsResource, c.ns, tEndpoint), &v1beta3.TEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TEndpoint), err
}

// Update takes the representation of a tEndpoint and updates it. Returns the server's representation of the tEndpoint, and an error, if there is any.
func (c *FakeTEndpoints) Update(ctx context.Context, tEndpoint *v1beta3.TEndpoint, opts v1.UpdateOptions) (result *v1beta3.TEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tendpointsResource, c.ns, tEndpoint), &v1beta3.TEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTEndpoints) UpdateStatus(ctx context.Context, tEndpoint *v1beta3.TEndpoint, opts v1.UpdateOptions) (*v1beta3.TEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tendpointsResource, "status", c.ns, tEndpoint), &v1beta3.TEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TEndpoint), err
}

// Delete takes name of the tEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeTEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tendpointsResource, c.ns, name), &v1beta3.TEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta3.TEndpointList{})
	return err
}

// Patch applies the patch and returns the patched tEndpoint.
func (c *FakeTEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tendpointsResource, c.ns, name, pt, data, subresources...), &v1beta3.TEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TEndpoint), err
}
