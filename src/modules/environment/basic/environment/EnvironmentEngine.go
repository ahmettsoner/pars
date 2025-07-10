package basic_environment

import (
	"fmt"

	"github.com/sirupsen/logrus"

	basic_environment_payload_structs "parsdevkit.net/modules/environment/basic_environment_payload/structs"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/modules/environment/basic_environment_contract"
	_string "parsdevkit.net/pkg/utilities/string"
)

type EnvironmentEngine struct{}

func (s EnvironmentEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*basic_environment_payload_structs.EnvironmentBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s EnvironmentEngine) List(ctx *application.ApplicationContext) error {
	err := s.list(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (s EnvironmentEngine) prepareToList(ctx *application.ApplicationContext) ([]basic_environment_payload_structs.EnvironmentBaseStruct, error) {

	service := ioc.Get[basic_environment_contract.EnvironmentInterface]()

	readyToListStructs, err := service.List()
	if err != nil {
		return nil, err
	}

	return readyToListStructs, nil
}
func (s EnvironmentEngine) list(ctx *application.ApplicationContext) error {

	readyToListStructs, err := s.prepareToList(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("(%d) environment available\n", (len(readyToListStructs) + 1))
	fmt.Println("* Default")
	for _, e := range readyToListStructs {
		fmt.Printf("- %v\n", e.Header.Name)
	}

	return nil
}

func (s EnvironmentEngine) completeInformation(ctx *application.ApplicationContext, model *basic_environment_payload_structs.EnvironmentBaseStruct) error {

	logrus.Debugf("filling environment (%v) information", model.Header.Name)

	if _string.IsEmpty(model.Specifications.Name) {
		model.Specifications.Name = model.Header.Name
	}

	return nil
}

func (s EnvironmentEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Environment",
		Order: 1000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]basic_environment_payload_structs.EnvironmentBaseStruct, error) {
	r := make([]basic_environment_payload_structs.EnvironmentBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*basic_environment_payload_structs.EnvironmentBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected basic_environment_payload_structs.EnvironmentBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
