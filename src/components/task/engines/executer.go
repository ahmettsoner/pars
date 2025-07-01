package engines

import (
	"errors"
	"fmt"

	actionBase "parsdevkit.net/structs/task/actions"
)

type ExecuterContext struct {
	Executers map[string]ActionExecuterInterface
}

func NewExecuterContext() ExecuterContext {
	return ExecuterContext{
		Executers: make(map[string]ActionExecuterInterface),
	}
}

func (m *ExecuterContext) RegisterExecuter(_type string, executer ActionExecuterInterface) {
	m.Executers[_type] = executer
}

type ActionExecuterInterface interface {
	Execute(action actionBase.ActionInterface, input []interface{}) (interface{}, error)
}

type TaskExecuter struct {
	executerContext ExecuterContext
	context         map[string]interface{}
	Actions         []actionBase.ActionInterface
}

func NewTaskExecuter(ctx ExecuterContext, actions ...actionBase.ActionInterface) TaskExecuter {
	return TaskExecuter{
		executerContext: ctx,
		context:         make(map[string]interface{}),
		Actions:         actions,
	}
}

func (s TaskExecuter) Execute() error {
	if len(s.Actions) == 0 {
		return errors.New("No actions in the pipeline")
	}

	var inputData []interface{} = nil
	for _, action := range s.Actions {
		if len(action.GetInputs()) > 0 {
			inputVariables := make([]interface{}, 0)
			for _, v := range action.GetInputs() {
				inputVariable, ok := s.context[v]
				if !ok {
					return fmt.Errorf("%v input is not exists in context", v)
				} else {
					inputVariables = append(inputVariables, inputVariable)
				}
				inputData = inputVariables
			}
		} else {
			for _, value := range s.context {
				inputData = append(inputData, value)
			}
		}

		executer, ok := s.executerContext.Executers[action.GetType()]
		if !ok {
			return fmt.Errorf("Executor for %v type is not found", action.GetType())
		}

		outputData, err := executer.Execute(action, inputData)
		if err != nil {
			return err
		}

		if action.GetOutput() != "" {
			s.context[fmt.Sprintf("%v.%v", action.GetName(), action.GetOutput())] = outputData
		} else {
			s.context[fmt.Sprintf("%v.result", action.GetName())] = outputData
		}
	}
	return nil
}
