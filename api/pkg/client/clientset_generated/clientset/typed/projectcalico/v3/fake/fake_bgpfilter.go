// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	projectcalicov3 "github.com/projectcalico/api/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
	gentype "k8s.io/client-go/gentype"
)

// fakeBGPFilters implements BGPFilterInterface
type fakeBGPFilters struct {
	*gentype.FakeClientWithList[*v3.BGPFilter, *v3.BGPFilterList]
	Fake *FakeProjectcalicoV3
}

func newFakeBGPFilters(fake *FakeProjectcalicoV3) projectcalicov3.BGPFilterInterface {
	return &fakeBGPFilters{
		gentype.NewFakeClientWithList[*v3.BGPFilter, *v3.BGPFilterList](
			fake.Fake,
			"",
			v3.SchemeGroupVersion.WithResource("bgpfilters"),
			v3.SchemeGroupVersion.WithKind("BGPFilter"),
			func() *v3.BGPFilter { return &v3.BGPFilter{} },
			func() *v3.BGPFilterList { return &v3.BGPFilterList{} },
			func(dst, src *v3.BGPFilterList) { dst.ListMeta = src.ListMeta },
			func(list *v3.BGPFilterList) []*v3.BGPFilter { return gentype.ToPointerSlice(list.Items) },
			func(list *v3.BGPFilterList, items []*v3.BGPFilter) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
