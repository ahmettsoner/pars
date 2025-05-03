package services

import (
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"
)

type ApplicationProjectServiceInterface interface {
	Create(model applicationproject.ProjectBaseStruct, init bool) (*applicationproject.ProjectBaseStruct, error)
	GenerateProject(model applicationproject.ProjectBaseStruct) (*applicationproject.ProjectBaseStruct, error)
	AddFileToLayer(model applicationproject.ProjectBaseStruct, layer string, paths []string, filename string, content string) (*applicationproject.ProjectBaseStruct, error)
	List() (*([]applicationproject.ProjectBaseStruct), error)
	ListBySet(set string) (*([]applicationproject.ProjectBaseStruct), error)
	ListBySetAndLayers(set string, layers ...string) (*([]applicationproject.ProjectBaseStruct), error)
	ListByWorkspace(workspaceName string) (*([]applicationproject.ProjectBaseStruct), error)
	GetByFullNameWorkspace(name string, workspaceName string) (*applicationproject.ProjectBaseStruct, error)
	ListByFullNameWorkspace(name string, workspaceName string) (*([]applicationproject.ProjectBaseStruct), error)
	CheckIfWorkingOnProject() (*applicationproject.ProjectBaseStruct, error)
	IsDirectoryReserved(path string) (*applicationproject.ProjectBaseStruct, error)
	Remove(name string, workspaceName string, force bool, permanent bool) (*applicationproject.ProjectSpecification, error)
	Build(name string, workspaceName string) (*applicationproject.ProjectSpecification, error)
	GetProjectWorkspace(workspaceName string) (*workspace.WorkspaceSpecification, error)

	ValidateProjectStructure(model applicationproject.ProjectSpecification) (bool, error)
	ValidateProjectDependencies(model applicationproject.ProjectSpecification) (bool, error)
	ValidateProjectReferences(model applicationproject.ProjectSpecification) (bool, error)
	GetHash(name string, workspaceName string) string
}
