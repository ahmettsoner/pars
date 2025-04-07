package print

import (
	actionBase "parsdevkit.net/structs/task/actions"
)

const (
	Type = "io/print"
)

type PrintAction struct {
	actionBase.Action
	Value string
}
