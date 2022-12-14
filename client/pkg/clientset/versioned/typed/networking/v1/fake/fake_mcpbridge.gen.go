// Copyright (c) 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	networkingv1 "github.com/alibaba/higress/client/pkg/apis/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMcpBridges implements McpBridgeInterface
type FakeMcpBridges struct {
	Fake *FakeNetworkingV1
	ns   string
}

var mcpbridgesResource = schema.GroupVersionResource{Group: "networking.higress.io", Version: "v1", Resource: "mcpbridges"}

var mcpbridgesKind = schema.GroupVersionKind{Group: "networking.higress.io", Version: "v1", Kind: "McpBridge"}

// Get takes name of the mcpBridge, and returns the corresponding mcpBridge object, and an error if there is any.
func (c *FakeMcpBridges) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkingv1.McpBridge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mcpbridgesResource, c.ns, name), &networkingv1.McpBridge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.McpBridge), err
}

// List takes label and field selectors, and returns the list of McpBridges that match those selectors.
func (c *FakeMcpBridges) List(ctx context.Context, opts v1.ListOptions) (result *networkingv1.McpBridgeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mcpbridgesResource, mcpbridgesKind, c.ns, opts), &networkingv1.McpBridgeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1.McpBridgeList{ListMeta: obj.(*networkingv1.McpBridgeList).ListMeta}
	for _, item := range obj.(*networkingv1.McpBridgeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mcpBridges.
func (c *FakeMcpBridges) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mcpbridgesResource, c.ns, opts))

}

// Create takes the representation of a mcpBridge and creates it.  Returns the server's representation of the mcpBridge, and an error, if there is any.
func (c *FakeMcpBridges) Create(ctx context.Context, mcpBridge *networkingv1.McpBridge, opts v1.CreateOptions) (result *networkingv1.McpBridge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mcpbridgesResource, c.ns, mcpBridge), &networkingv1.McpBridge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.McpBridge), err
}

// Update takes the representation of a mcpBridge and updates it. Returns the server's representation of the mcpBridge, and an error, if there is any.
func (c *FakeMcpBridges) Update(ctx context.Context, mcpBridge *networkingv1.McpBridge, opts v1.UpdateOptions) (result *networkingv1.McpBridge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mcpbridgesResource, c.ns, mcpBridge), &networkingv1.McpBridge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.McpBridge), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMcpBridges) UpdateStatus(ctx context.Context, mcpBridge *networkingv1.McpBridge, opts v1.UpdateOptions) (*networkingv1.McpBridge, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mcpbridgesResource, "status", c.ns, mcpBridge), &networkingv1.McpBridge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.McpBridge), err
}

// Delete takes name of the mcpBridge and deletes it. Returns an error if one occurs.
func (c *FakeMcpBridges) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mcpbridgesResource, c.ns, name), &networkingv1.McpBridge{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMcpBridges) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mcpbridgesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1.McpBridgeList{})
	return err
}

// Patch applies the patch and returns the patched mcpBridge.
func (c *FakeMcpBridges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkingv1.McpBridge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mcpbridgesResource, c.ns, name, pt, data, subresources...), &networkingv1.McpBridge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.McpBridge), err
}
