package engines

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"
	sharedtemplate "parsdevkit.net/structs/template/shared-template"
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
	sharedTemplateService := ioc.Get[contracts.TemplateServiceInterface[sharedtemplate.TemplateBaseStruct]]()

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
