package handlers

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	engineOperations "parsdevkit.net/engines"
	"parsdevkit.net/modules/template/code_template_payload/commands"
)

type GenerateTemplateContentsCommandHandler struct{}

func (s *GenerateTemplateContentsCommandHandler) Handle(cmd commands.GenerateTemplateContents) error {

	model, ok := cmd.Data.(code_template_payload_structs.TemplateBaseStruct)
	if !ok {
		return fmt.Errorf("invalid item type: expected code_template_payload_structs.TemplateBaseStruct, got %T", model)
	}
	if _, err := s.generate(model); err != nil {
		return err
	}

	return nil
}

func (s *GenerateTemplateContentsCommandHandler) generate(model code_template_payload_structs.TemplateBaseStruct) (*code_template_payload_structs.TemplateBaseStruct, error) {

	templateService := ioc.Get[code_template_contract.TemplateInterface]()

	result, err := templateService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateEngine := engineOperations.NewCodeTemplateOperations(application.GetEnvironment())
	err = templateEngine.GenerateByTemplate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
