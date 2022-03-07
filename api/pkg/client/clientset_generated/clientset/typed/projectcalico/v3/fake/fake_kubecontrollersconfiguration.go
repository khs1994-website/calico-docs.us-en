// Copyright (c) 2022 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeControllersConfigurations implements KubeControllersConfigurationInterface
type FakeKubeControllersConfigurations struct {
	Fake *FakeProjectcalicoV3
}

var kubecontrollersconfigurationsResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "kubecontrollersconfigurations"}

var kubecontrollersconfigurationsKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "KubeControllersConfiguration"}

// Get takes name of the kubeControllersConfiguration, and returns the corresponding kubeControllersConfiguration object, and an error if there is any.
func (c *FakeKubeControllersConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.KubeControllersConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubecontrollersconfigurationsResource, name), &v3.KubeControllersConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.KubeControllersConfiguration), err
}

// List takes label and field selectors, and returns the list of KubeControllersConfigurations that match those selectors.
func (c *FakeKubeControllersConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v3.KubeControllersConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubecontrollersconfigurationsResource, kubecontrollersconfigurationsKind, opts), &v3.KubeControllersConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.KubeControllersConfigurationList{ListMeta: obj.(*v3.KubeControllersConfigurationList).ListMeta}
	for _, item := range obj.(*v3.KubeControllersConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeControllersConfigurations.
func (c *FakeKubeControllersConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubecontrollersconfigurationsResource, opts))
}

// Create takes the representation of a kubeControllersConfiguration and creates it.  Returns the server's representation of the kubeControllersConfiguration, and an error, if there is any.
func (c *FakeKubeControllersConfigurations) Create(ctx context.Context, kubeControllersConfiguration *v3.KubeControllersConfiguration, opts v1.CreateOptions) (result *v3.KubeControllersConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubecontrollersconfigurationsResource, kubeControllersConfiguration), &v3.KubeControllersConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.KubeControllersConfiguration), err
}

// Update takes the representation of a kubeControllersConfiguration and updates it. Returns the server's representation of the kubeControllersConfiguration, and an error, if there is any.
func (c *FakeKubeControllersConfigurations) Update(ctx context.Context, kubeControllersConfiguration *v3.KubeControllersConfiguration, opts v1.UpdateOptions) (result *v3.KubeControllersConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubecontrollersconfigurationsResource, kubeControllersConfiguration), &v3.KubeControllersConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.KubeControllersConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeControllersConfigurations) UpdateStatus(ctx context.Context, kubeControllersConfiguration *v3.KubeControllersConfiguration, opts v1.UpdateOptions) (*v3.KubeControllersConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubecontrollersconfigurationsResource, "status", kubeControllersConfiguration), &v3.KubeControllersConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.KubeControllersConfiguration), err
}

// Delete takes name of the kubeControllersConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeKubeControllersConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubecontrollersconfigurationsResource, name, opts), &v3.KubeControllersConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeControllersConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubecontrollersconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.KubeControllersConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched kubeControllersConfiguration.
func (c *FakeKubeControllersConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.KubeControllersConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubecontrollersconfigurationsResource, name, pt, data, subresources...), &v3.KubeControllersConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.KubeControllersConfiguration), err
}
