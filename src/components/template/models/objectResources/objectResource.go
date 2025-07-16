package objectResources

import (
	"parsdevkit.net/application/models/class"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/modules/resource/data_resource"
	"parsdevkit.net/modules/resource/object_resource"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
)

type ObjectDictionary struct {
	Key        string
	Translates map[string]string
}

type ObjectGroup struct {
	Name    string
	Title   ObjectMessage
	Options []ObjectOption
}

type ObjectMessage struct {
	Text       string
	Dictionary string
}
type ObjectOption struct {
	Key   string
	Value interface{}
}

type FileTemplateComposite struct {
	FileTemplate
	Original file_template_payload_structs.TemplateBaseStruct
}

type FileTemplate struct {
	Name    string
	Package string
	Labels  []object_resource.ObjectLabel
	// Options []ObjectOption
	// Layers     []ObjectLayer
}

type CodeTemplateComposite struct {
	CodeTemplate
	Original code_template_payload_structs.TemplateBaseStruct
}

type CodeTemplate struct {
	Name    string
	Package string
	Labels  []object_resource.ObjectLabel
	// Options []ObjectOption
	// Layers     []ObjectLayer
}

type DataLayerComposite struct {
	data_resource.DataLayer
	Original data_resource_payload_structs.Layer
}

type DataLayer struct {
	Name     string
	Sections []DataSection
}

type DataSectionComposite struct {
	data_resource.DataSection
	Original data_resource_payload_structs.Section
}

type DataSection struct {
	Name    string
	Package string
	Classes []class.Class
	Labels  []object_resource.ObjectLabel
	Options []ObjectOption
}

type ObjectLayerComposite struct {
	object_resource.ObjectLayer
	Original object_resource_payload_structs.Layer
}

type ObjectSectionComposite struct {
	object_resource.ObjectSection
	Original object_resource_payload_structs.Section
}
