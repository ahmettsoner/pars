package services

import sharedTemplate "parsdevkit.net/structs/template/shared-template"

type SharedTemplateServiceInterface interface {
	GetByName(name string) (*sharedTemplate.TemplateBaseStruct, error)
	Save(model sharedTemplate.TemplateBaseStruct) (*sharedTemplate.TemplateBaseStruct, error)
	List() (*([]sharedTemplate.TemplateBaseStruct), error)
	ListByWorkspace(workspace string) (*([]sharedTemplate.TemplateBaseStruct), error)
	Remove(name, workspace string, permanent bool) (*sharedTemplate.TemplateBaseStruct, error)
	IsExists(name, workspace string) bool
	GetHash(name string) string
}
