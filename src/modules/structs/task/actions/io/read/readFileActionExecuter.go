package read

import (
	"fmt"
	"os"
	taskEngine "pars/modules/tasks/engines"

	actionBase "parsdevkit.net/structs/task/actions"
)

type ReadActionExecuter struct {
	taskEngine.ActionExecuterInterface
}

func NewReadActionExecuter() ReadActionExecuter {
	return ReadActionExecuter{}
}
func (s ReadActionExecuter) Execute(action actionBase.ActionInterface, input []interface{}) (interface{}, error) {
	readAction, ok := action.(ReadAction)
	if !ok {
		return nil, fmt.Errorf("unexpected action type")
	}

	content, err := os.ReadFile(readAction.FieldName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
