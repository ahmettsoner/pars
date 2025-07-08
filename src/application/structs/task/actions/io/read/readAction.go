package read

import (
	actionBase "parsdevkit.net/application/structs/task/actions"
)

const (
	Type = "io/read"
)

type ReadAction struct {
	actionBase.Action
	FieldName string
}
