package services

import codetemplate "parsdevkit.net/structs/template/code-template"

type CodeTemplateServiceInterface interface {
	GetByName(name string) (*codetemplate.TemplateBaseStruct, error)
	Save(model codetemplate.TemplateBaseStruct) (*codetemplate.TemplateBaseStruct, error)
	List() (*([]codetemplate.TemplateBaseStruct), error)
	ListBySetAndLayers(set string, layers ...string) (*([]codetemplate.TemplateBaseStruct), error)
	Remove(name, workspace string, permanent bool) (*codetemplate.TemplateBaseStruct, error)
	IsExists(name, workspace string) bool
	GetHash(name string) string
}
