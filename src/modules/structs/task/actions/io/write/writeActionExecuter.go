package write

import (
	"fmt"
	taskEngine "pars/modules/tasks/engines"

	actionBase "parsdevkit.net/structs/task/actions"
)

type WriteActionExecuter struct {
	taskEngine.ActionExecuterInterface
}

func NewWriteActionExecuter() WriteActionExecuter {
	return WriteActionExecuter{}
}
func (s WriteActionExecuter) Execute(action actionBase.ActionInterface, input []interface{}) (interface{}, error) {
	_, ok := action.(WriteAction)
	if !ok {
		return nil, fmt.Errorf("unexpected action type")
	}

	return "", nil
}
