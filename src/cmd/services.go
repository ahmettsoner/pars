package cmd

import (
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"

	group "parsdevkit.net/modules/group/group"
	projectApplication "parsdevkit.net/modules/project/application"
	resourceData "parsdevkit.net/modules/resource/data"
	resourceObject "parsdevkit.net/modules/resource/object"
	taskCommon "parsdevkit.net/modules/task/common"
	templateCode "parsdevkit.net/modules/template/code"
	templateFile "parsdevkit.net/modules/template/file"
	templateShared "parsdevkit.net/modules/template/shared"

	groupSchema "parsdevkit.net/modules/group/group"
	applicationProjectSchema "parsdevkit.net/structs/project/application-project"
	dataResourceSchema "parsdevkit.net/structs/resource/data-resource"
	objectResourceSchema "parsdevkit.net/structs/resource/object-resource"
	codeTemplateSchema "parsdevkit.net/structs/template/code-template"
	fileTemplateSchema "parsdevkit.net/structs/template/file-template"
	sharedTemplateSchema "parsdevkit.net/structs/template/shared-template"
)

func RegisterServices() {
	registerEngines()
	registerSchemas()
}
func registerSchemas() {
	schemas.Register(&groupSchema.GroupBaseStruct{})
	schemas.Register(&applicationProjectSchema.ProjectBaseStruct{})
	schemas.Register(&dataResourceSchema.ResourceBaseStruct{})
	schemas.Register(&objectResourceSchema.ResourceBaseStruct{})
	schemas.Register(&codeTemplateSchema.TemplateBaseStruct{})
	schemas.Register(&fileTemplateSchema.TemplateBaseStruct{})
	schemas.Register(&sharedTemplateSchema.TemplateBaseStruct{})
}

func registerEngines() {
	engines.Register(&group.GroupEngine{})
	engines.Register(&projectApplication.ApplicationProjectEngine{})
	engines.Register(&resourceData.DataResourceEngine{})
	engines.Register(&resourceObject.ObjectResourceEngine{})
	engines.Register(&templateCode.CodeTemplateEngine{})
	engines.Register(&templateFile.FileTemplateEngine{})
	engines.Register(&templateShared.SharedTemplateEngine{})
	engines.Register(&taskCommon.CommonTaskEngine{})
}
