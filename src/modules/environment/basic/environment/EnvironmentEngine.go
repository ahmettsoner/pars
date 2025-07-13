package basic_environment

import (
	"fmt"

	"github.com/sirupsen/logrus"

	basic_environment_payload_structs "parsdevkit.net/modules/environment/basic_environment_payload/structs"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	list_printer "parsdevkit.net/modules/environment/basic_environment/printers/list"
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
	readyToListStructs, err := s.prepareToList(ctx)
	if err != nil {
		return err
	}

	err = s.list(ctx, readyToListStructs)
	if err != nil {
		return err
	}

	return nil
}
func (s EnvironmentEngine) prepareToList(ctx *application.ApplicationContext) ([]basic_environment_payload_structs.EnvironmentBaseStruct, error) {

	service := ioc.Get[basic_environment_contract.EnvironmentInterface]()

	environmentList, err := service.List()
	if err != nil {
		return nil, err
	}

	return environmentList, nil
}
func (s EnvironmentEngine) list(ctx *application.ApplicationContext, models []basic_environment_payload_structs.EnvironmentBaseStruct) error {

	var viewModels []list_printer.ViewModel = make([]list_printer.ViewModel, 0)
	for _, e := range models {
		label := e.Header.Name
		if e.Header.Name == "Default" {
			label = fmt.Sprintf("* %s", label)
		}
		resource := list_printer.ViewModel{
			Name: label,
		}

		viewModels = append(viewModels, resource)
	}

	fmt.Printf("\n🛠️  Environment List (%d):\n\n", len(viewModels))

	printer := list_printer.ListEnvironment{Environments: viewModels}
	printer.Print()

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
		Name:  basic_environment_payload_structs.MODULE_KEY,
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
