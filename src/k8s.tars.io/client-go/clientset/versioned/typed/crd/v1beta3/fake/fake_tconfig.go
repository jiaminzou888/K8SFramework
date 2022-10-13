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

// FakeTConfigs implements TConfigInterface
type FakeTConfigs struct {
	Fake *FakeCrdV1beta3
	ns   string
}

var tconfigsResource = schema.GroupVersionResource{Group: "k8s.tars.io", Version: "v1beta3", Resource: "tconfigs"}

var tconfigsKind = schema.GroupVersionKind{Group: "k8s.tars.io", Version: "v1beta3", Kind: "TConfig"}

// Get takes name of the tConfig, and returns the corresponding tConfig object, and an error if there is any.
func (c *FakeTConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.TConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tconfigsResource, c.ns, name), &v1beta3.TConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TConfig), err
}

// List takes label and field selectors, and returns the list of TConfigs that match those selectors.
func (c *FakeTConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.TConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tconfigsResource, tconfigsKind, c.ns, opts), &v1beta3.TConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta3.TConfigList{ListMeta: obj.(*v1beta3.TConfigList).ListMeta}
	for _, item := range obj.(*v1beta3.TConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tConfigs.
func (c *FakeTConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tconfigsResource, c.ns, opts))

}

// Create takes the representation of a tConfig and creates it.  Returns the server's representation of the tConfig, and an error, if there is any.
func (c *FakeTConfigs) Create(ctx context.Context, tConfig *v1beta3.TConfig, opts v1.CreateOptions) (result *v1beta3.TConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tconfigsResource, c.ns, tConfig), &v1beta3.TConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TConfig), err
}

// Update takes the representation of a tConfig and updates it. Returns the server's representation of the tConfig, and an error, if there is any.
func (c *FakeTConfigs) Update(ctx context.Context, tConfig *v1beta3.TConfig, opts v1.UpdateOptions) (result *v1beta3.TConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tconfigsResource, c.ns, tConfig), &v1beta3.TConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TConfig), err
}

// Delete takes name of the tConfig and deletes it. Returns an error if one occurs.
func (c *FakeTConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tconfigsResource, c.ns, name), &v1beta3.TConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta3.TConfigList{})
	return err
}

// Patch applies the patch and returns the patched tConfig.
func (c *FakeTConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.TConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tconfigsResource, c.ns, name, pt, data, subresources...), &v1beta3.TConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.TConfig), err
}
