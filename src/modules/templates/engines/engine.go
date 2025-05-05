package engines

import (
	"parsdevkit.net/core/utils"
	"parsdevkit.net/operation/services"
)

type EngineFuncs struct{}

func (t EngineFuncs) RenderContent(templateName string, data any) string {
	content, err := TemplateEngine(templateName, data)
	if err != nil {
		return ""
	}

	return content
}

func (t EngineFuncs) GetContent(templateName string) string {
	sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())
	sharedTemplate, err := sharedTemplateService.GetByName(templateName)
	if err != nil {
		return ""
	}
	if sharedTemplate == nil {
		return ""
	}

	return sharedTemplate.Specifications.Template.Content
}
func (t EngineFuncs) RenderTemplate(templateName string, data any) string {
	content, err := TemplateEngine(t.GetContent(templateName), data)
	if err != nil {
		return ""
	}

	return content
}
