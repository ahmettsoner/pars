package functions

import (
	"parsdevkit.net/core/utils/data"
)

type DataFuncs struct{}

func (d DataFuncs) Default(value any, defaultValue any) any {
	return data.Default(value, defaultValue)
}

func (d DataFuncs) IsDefault(value any) bool {
	return data.IsDefault(value)
}
