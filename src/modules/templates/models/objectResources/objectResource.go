package objectResources

import (
	"parsdevkit.net/structs/class"
	applicationproject "parsdevkit.net/structs/project/application-project"
	dataresource "parsdevkit.net/structs/resource/data-resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	codetemplate "parsdevkit.net/structs/template/code-template"
	filetemplate "parsdevkit.net/structs/template/file-template"
	"parsdevkit.net/structs/workspace"
)

type WorkspaceComposite struct {
	Workspace
	Original workspace.WorkspaceBaseStruct
}

type Workspace struct {
	Name string
}

type ApplicationProjectComposite struct {
	ApplicationProject
	Original applicationproject.ProjectBaseStruct
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
	Original objectresource.ResourceBaseStruct
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
	Original dataresource.ResourceBaseStruct
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
	Original filetemplate.TemplateBaseStruct
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
	Original codetemplate.TemplateBaseStruct
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
	Original dataresource.Layer
}

type DataLayer struct {
	Name     string
	Sections []DataSection
}

type DataSectionComposite struct {
	DataSection
	Original dataresource.Section
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
	Original objectresource.Layer
}

type ObjectSectionComposite struct {
	ObjectSection
	Original objectresource.Section
}
