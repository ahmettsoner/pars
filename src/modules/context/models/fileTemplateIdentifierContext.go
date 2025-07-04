package models

import dataresource "parsdevkit.net/modules/resource/data_resource_payload"

type FileTemplateIdentifierContext struct {
	Resource dataresource.ResourceBaseStruct
	Section  dataresource.Section
}

func NewFileTemplateIdentifierContext(resource dataresource.ResourceBaseStruct, section dataresource.Section) *FileTemplateIdentifierContext {
	return &FileTemplateIdentifierContext{
		Resource: resource,
		Section:  section,
	}
}
