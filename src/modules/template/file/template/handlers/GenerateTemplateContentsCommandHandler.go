package handlers

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/template/file_template_contract"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	templateEngine "parsdevkit.net/components/template/engines"
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

	templateOperations := templateEngine.NewTemplateOperations(application.GetEnvironment())
	err = templateOperations.Generate(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
