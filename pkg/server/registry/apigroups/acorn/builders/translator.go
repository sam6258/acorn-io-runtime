package builders

import (
	mtypes "github.com/acorn-io/mink/pkg/types"
	apiv1 "github.com/acorn-io/runtime/pkg/apis/api.acorn.io/v1"
	v1 "github.com/acorn-io/runtime/pkg/apis/internal.acorn.io/v1"
)

type Translator struct {
}

func (s *Translator) FromPublic(obj mtypes.Object) mtypes.Object {
	return (*v1.BuilderInstance)(obj.(*apiv1.Builder))
}

func (s *Translator) ToPublic(obj mtypes.Object) mtypes.Object {
	return (*apiv1.Builder)(obj.(*v1.BuilderInstance))
}
