package regions

import (
	"context"

	"github.com/acorn-io/mink/pkg/types"
	apiv1 "github.com/acorn-io/runtime/pkg/apis/api.acorn.io/v1"
	v1 "github.com/acorn-io/runtime/pkg/apis/internal.acorn.io/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/storage"
)

type strategy struct {
	startTime metav1.Time
}

func (s *strategy) Get(_ context.Context, _, name string) (types.Object, error) {
	if name != apiv1.LocalRegion {
		return nil, apierrors.NewNotFound(schema.GroupResource{
			Group:    apiv1.SchemeGroupVersion.Group,
			Resource: "regions",
		}, name)
	}

	return &apiv1.Region{
		ObjectMeta: metav1.ObjectMeta{
			Name:              apiv1.LocalRegion,
			CreationTimestamp: s.startTime,
			OwnerReferences: []metav1.OwnerReference{
				{
					Name: apiv1.LocalRegion,
				},
			},
		},
		Spec: apiv1.RegionSpec{
			Description: "Local Region",
			RegionName:  apiv1.LocalRegion,
		},
		Status: apiv1.RegionStatus{
			Conditions: []v1.Condition{
				{
					Type:               apiv1.RegionConditionClusterReady,
					Success:            true,
					Status:             metav1.ConditionTrue,
					LastTransitionTime: s.startTime,
				},
			},
		},
	}, nil
}

func (s *strategy) List(_ context.Context, _ string, _ storage.ListOptions) (types.ObjectList, error) {
	region, _ := s.Get(context.Background(), "", apiv1.LocalRegion)
	return &apiv1.RegionList{
		Items: []apiv1.Region{*(region.(*apiv1.Region))},
	}, nil
}

func (s *strategy) New() types.Object {
	return new(apiv1.Region)
}

func (s *strategy) NewList() types.ObjectList {
	return new(apiv1.RegionList)
}
