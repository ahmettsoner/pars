package handlers

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/resource/data_resource_contract"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	engineOperations "parsdevkit.net/engines"
	"parsdevkit.net/modules/resource/data_resource_payload/commands"
)

type GenerateResourceContentsCommandHandler struct{}

func (s *GenerateResourceContentsCommandHandler) Handle(cmd commands.GenerateResourceContents) error {

	model, ok := cmd.Data.(data_resource_payload_structs.ResourceBaseStruct)
	if !ok {
		return fmt.Errorf("invalid item type: expected data_resource_payload_structs.ResourceBaseStruct, got %T", model)
	}
	if _, err := s.generate(model); err != nil {
		return err
	}

	return nil
}

func (s *GenerateResourceContentsCommandHandler) generate(model data_resource_payload_structs.ResourceBaseStruct) (*data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()

	result, err := service.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateOperations := engineOperations.NewFileTemplateOperations(application.GetEnvironment())
	err = templateOperations.GenerateByResource(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
