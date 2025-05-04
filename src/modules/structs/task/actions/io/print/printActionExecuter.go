package print

import (
	"fmt"
	taskEngine "pars/modules/tasks/engines"

	actionBase "parsdevkit.net/structs/task/actions"
)

type PrintActionExecuter struct {
	taskEngine.ActionExecuterInterface
}

func NewPrintActionExecuter() PrintActionExecuter {
	return PrintActionExecuter{}
}
func (s PrintActionExecuter) Execute(action actionBase.ActionInterface, input []interface{}) (interface{}, error) {
	printAction, ok := action.(PrintAction)
	if !ok {
		return nil, fmt.Errorf("unexpected action type")
	}

	for _, v := range input {
		fmt.Printf("%v", v)
	}

	fmt.Printf("%v", printAction.Value)

	return input, nil
}
