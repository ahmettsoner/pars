package models

import objectresource "parsdevkit.net/modules/resource/object_resource_payload"

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
