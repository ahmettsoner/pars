package schemas

import (
	"reflect"
)

type SchemaSpecification struct {
}

func (l SchemaSpecification) IsEqual(other SchemaSpecification) bool {
	return reflect.DeepEqual(l, other)
}
