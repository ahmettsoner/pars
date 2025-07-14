package engines

import (
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/template/shared_template_contract"
)

type EngineFuncs struct{}

func (t EngineFuncs) RenderContent(templateName string, data any) string {
	content, err := RenderTemplate(templateName, data)
	if err != nil {
		return ""
	}

	return content
}

func (t EngineFuncs) GetContent(templateName string) string {
	sharedTemplateService := ioc.Get[shared_template_contract.TemplateInterface]()

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
	content, err := RenderTemplate(t.GetContent(templateName), data)
	if err != nil {
		return ""
	}

	return content
}
