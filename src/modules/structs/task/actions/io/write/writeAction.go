package write

import (
	actionBase "parsdevkit.net/structs/task/actions"
)

const (
	Type = "io/write"
)

type WriteAction struct {
	actionBase.Action
	FieldName string
}
