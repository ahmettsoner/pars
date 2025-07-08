package handlers

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/template/file_template_contract"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	engineOperations "parsdevkit.net/engines"
	"parsdevkit.net/modules/template/file_template_payload/commands"
)

type GenerateTemplateContentsCommandHandler struct{}

func (s *GenerateTemplateContentsCommandHandler) Handle(cmd commands.GenerateTemplateContents) error {

	model, ok := cmd.Data.(file_template_payload_structs.TemplateBaseStruct)
	if !ok {
		return fmt.Errorf("invalid item type: expected file_template_payload_structs.TemplateBaseStruct, got %T", model)
	}
	if _, err := s.generate(model); err != nil {
		return err
	}
	fmt.Printf("[CommandHandler] Template '%s' generated\n", cmd.Data.GetHeader().Name)

	return nil
}

func (s *GenerateTemplateContentsCommandHandler) generate(model file_template_payload_structs.TemplateBaseStruct) (*file_template_payload_structs.TemplateBaseStruct, error) {

	service := ioc.Get[file_template_contract.TemplateInterface]()

	result, err := service.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateOperations := engineOperations.NewFileTemplateOperations(application.GetEnvironment())
	err = templateOperations.GenerateByTemplate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
