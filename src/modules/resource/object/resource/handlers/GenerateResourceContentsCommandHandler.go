package handlers

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"
	templateEngine "parsdevkit.net/components/template/engines"
	"parsdevkit.net/modules/resource/object_resource_contract"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/modules/resource/object_resource_payload/commands"
)

type GenerateResourceContentsCommandHandler struct{}

func (s *GenerateResourceContentsCommandHandler) Handle(cmd commands.GenerateResourceContents) error {

	model, ok := cmd.Data.(object_resource_payload_structs.ResourceBaseStruct)
	if !ok {
		return fmt.Errorf("invalid item type: expected object_resource_payload_structs.ResourceBaseStruct, got %T", model)
	}
	if _, err := s.generate(model); err != nil {
		return err
	}

	return nil
}

func (s *GenerateResourceContentsCommandHandler) generate(model object_resource_payload_structs.ResourceBaseStruct) (*object_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()

	result, err := service.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateEngine := templateEngine.NewTemplateOperations(application.GetEnvironment())
	err = templateEngine.Generate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
