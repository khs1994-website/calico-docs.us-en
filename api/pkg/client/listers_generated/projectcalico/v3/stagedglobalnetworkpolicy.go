// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// StagedGlobalNetworkPolicyLister helps list StagedGlobalNetworkPolicies.
// All objects returned here must be treated as read-only.
type StagedGlobalNetworkPolicyLister interface {
	// List lists all StagedGlobalNetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.StagedGlobalNetworkPolicy, err error)
	// Get retrieves the StagedGlobalNetworkPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.StagedGlobalNetworkPolicy, error)
	StagedGlobalNetworkPolicyListerExpansion
}

// stagedGlobalNetworkPolicyLister implements the StagedGlobalNetworkPolicyLister interface.
type stagedGlobalNetworkPolicyLister struct {
	listers.ResourceIndexer[*v3.StagedGlobalNetworkPolicy]
}

// NewStagedGlobalNetworkPolicyLister returns a new StagedGlobalNetworkPolicyLister.
func NewStagedGlobalNetworkPolicyLister(indexer cache.Indexer) StagedGlobalNetworkPolicyLister {
	return &stagedGlobalNetworkPolicyLister{listers.New[*v3.StagedGlobalNetworkPolicy](indexer, v3.Resource("stagedglobalnetworkpolicy"))}
}
