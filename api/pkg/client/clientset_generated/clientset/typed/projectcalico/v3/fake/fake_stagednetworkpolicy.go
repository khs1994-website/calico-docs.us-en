// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStagedNetworkPolicies implements StagedNetworkPolicyInterface
type FakeStagedNetworkPolicies struct {
	Fake *FakeProjectcalicoV3
	ns   string
}

var stagednetworkpoliciesResource = v3.SchemeGroupVersion.WithResource("stagednetworkpolicies")

var stagednetworkpoliciesKind = v3.SchemeGroupVersion.WithKind("StagedNetworkPolicy")

// Get takes name of the stagedNetworkPolicy, and returns the corresponding stagedNetworkPolicy object, and an error if there is any.
func (c *FakeStagedNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.StagedNetworkPolicy, err error) {
	emptyResult := &v3.StagedNetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(stagednetworkpoliciesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of StagedNetworkPolicies that match those selectors.
func (c *FakeStagedNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.StagedNetworkPolicyList, err error) {
	emptyResult := &v3.StagedNetworkPolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(stagednetworkpoliciesResource, stagednetworkpoliciesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.StagedNetworkPolicyList{ListMeta: obj.(*v3.StagedNetworkPolicyList).ListMeta}
	for _, item := range obj.(*v3.StagedNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stagedNetworkPolicies.
func (c *FakeStagedNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(stagednetworkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a stagedNetworkPolicy and creates it.  Returns the server's representation of the stagedNetworkPolicy, and an error, if there is any.
func (c *FakeStagedNetworkPolicies) Create(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.CreateOptions) (result *v3.StagedNetworkPolicy, err error) {
	emptyResult := &v3.StagedNetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(stagednetworkpoliciesResource, c.ns, stagedNetworkPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}

// Update takes the representation of a stagedNetworkPolicy and updates it. Returns the server's representation of the stagedNetworkPolicy, and an error, if there is any.
func (c *FakeStagedNetworkPolicies) Update(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.UpdateOptions) (result *v3.StagedNetworkPolicy, err error) {
	emptyResult := &v3.StagedNetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(stagednetworkpoliciesResource, c.ns, stagedNetworkPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}

// Delete takes name of the stagedNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeStagedNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(stagednetworkpoliciesResource, c.ns, name, opts), &v3.StagedNetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStagedNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(stagednetworkpoliciesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.StagedNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched stagedNetworkPolicy.
func (c *FakeStagedNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedNetworkPolicy, err error) {
	emptyResult := &v3.StagedNetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(stagednetworkpoliciesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}
