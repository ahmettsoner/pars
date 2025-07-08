package objectResources

import (
	"parsdevkit.net/application/models/class"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
)

type WorkspaceComposite struct {
	Workspace
	Original basic_workspace_payload_structs.WorkspaceBaseStruct
}

type Workspace struct {
	Name string
}

type ApplicationProjectComposite struct {
	ApplicationProject
	Original application_project_payload_structs.ProjectBaseStruct
}

type ApplicationProject struct {
	Name    string
	Package string
	Labels  []ObjectLabel
	// Options []ObjectOption
	// Layers     []ObjectLayer
}

type ObjectResourceComposite struct {
	ObjectResource
	Original object_resource_payload_structs.ResourceBaseStruct
}

type ObjectResource struct {
	Name       string
	Package    string
	Labels     []ObjectLabel
	Layers     []ObjectLayer
	Dictionary []ObjectDictionary
	Groups     []ObjectGroup
	Attributes []ObjectResourceAttribute
	Methods    []ObjectResourceMethod
	Imports    []ObjectResourceImport
}

type DataResourceComposite struct {
	DataResource
	Original data_resource_payload_structs.ResourceBaseStruct
}

type DataResource struct {
	Name       string
	Package    string
	Labels     []ObjectLabel
	Layers     []DataLayer
	Dictionary []ObjectDictionary
	Groups     []ObjectGroup
	Data       any
}

type ObjectResourceAttribute struct {
	Name         string
	TypePackage  string
	Type         string
	TypeCategory string
	Visibility   string
	Labels       []ObjectLabel
	Options      []ObjectOption
	Common       bool
}

type ObjectResourceMethod struct {
	Name        string
	Visibility  string
	Parameters  []ObjectResourceMethodParameter
	ReturnTypes []string
	Labels      []ObjectLabel
	Options     []ObjectOption
	Code        string
	Common      bool
}
type ObjectResourceMethodParameter struct {
	Name string
	Type string
}
type ObjectResourceImport struct {
	Aliases []string
	Package string
}
type ObjectSection struct {
	Name       string
	Package    string
	Classes    []class.Class
	Labels     []ObjectLabel
	Options    []ObjectOption
	Attributes []ObjectResourceAttribute
	Methods    []ObjectResourceMethod
	Imports    []ObjectResourceImport
}

type ObjectLabel struct {
	Key   string
	Value string
}

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

type ObjectLayer struct {
	Name     string
	Sections []ObjectSection
}

type FileTemplateComposite struct {
	FileTemplate
	Original file_template_payload_structs.TemplateBaseStruct
}

type FileTemplate struct {
	Name    string
	Package string
	Labels  []ObjectLabel
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
	Labels  []ObjectLabel
	// Options []ObjectOption
	// Layers     []ObjectLayer
}

type DataLayerComposite struct {
	DataLayer
	Original data_resource_payload_structs.Layer
}

type DataLayer struct {
	Name     string
	Sections []DataSection
}

type DataSectionComposite struct {
	DataSection
	Original data_resource_payload_structs.Section
}

type DataSection struct {
	Name    string
	Package string
	Classes []class.Class
	Labels  []ObjectLabel
	Options []ObjectOption
}

type ObjectLayerComposite struct {
	ObjectLayer
	Original object_resource_payload_structs.Layer
}

type ObjectSectionComposite struct {
	ObjectSection
	Original object_resource_payload_structs.Section
}
