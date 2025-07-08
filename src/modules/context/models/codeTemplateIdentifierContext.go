package models

import object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

type CodeTemplateIdentifierContext struct {
	Resource object_resource_payload_structs.ResourceBaseStruct
	Section  object_resource_payload_structs.Section
}

func NewCodeTemplateIdentifierContext(resource object_resource_payload_structs.ResourceBaseStruct, section object_resource_payload_structs.Section) *CodeTemplateIdentifierContext {
	return &CodeTemplateIdentifierContext{
		Resource: resource,
		Section:  section,
	}
}
