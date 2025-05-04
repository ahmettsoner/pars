package models

import objectresource "parsdevkit.net/structs/resource/object-resource"

type CodeTemplateIdentifierContext struct {
	Resource objectresource.ResourceBaseStruct
	Section  objectresource.Section
}

func NewCodeTemplateIdentifierContext(resource objectresource.ResourceBaseStruct, section objectresource.Section) *CodeTemplateIdentifierContext {
	return &CodeTemplateIdentifierContext{
		Resource: resource,
		Section:  section,
	}
}
