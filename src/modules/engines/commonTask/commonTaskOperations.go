package commonTask

import (
	objectresource "parsdevkit.net/structs/resource/object-resource"
	commontask "parsdevkit.net/structs/task/common-task"

	"parsdevkit.net/persistence/repositories"
)

type CommonTaskOperations struct {
	environment                  string
	generationHistoryRespository *repositories.GenerationHistoryRepository
}

func NewCommonTaskOperations(environment string) CommonTaskOperations {
	return CommonTaskOperations{
		environment:                  environment,
		generationHistoryRespository: repositories.NewGenerationHistoryRepository(environment),
	}
}

func (s CommonTaskOperations) GenerateByResource(model objectresource.ResourceBaseStruct) error {

	return nil
}

func (s CommonTaskOperations) GenerateByTask(model commontask.TaskBaseStruct) error {

	return nil
}
