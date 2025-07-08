package models

import data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

type FileTemplateIdentifierContext struct {
	Resource data_resource_payload_structs.ResourceBaseStruct
	Section  data_resource_payload_structs.Section
}

func NewFileTemplateIdentifierContext(resource data_resource_payload_structs.ResourceBaseStruct, section data_resource_payload_structs.Section) *FileTemplateIdentifierContext {
	return &FileTemplateIdentifierContext{
		Resource: resource,
		Section:  section,
	}
}
