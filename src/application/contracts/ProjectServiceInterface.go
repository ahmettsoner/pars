package contracts

import (
	"parsdevkit.net/application/structs/project"
)

type ProjectServiceInterface[T SchemaInterface] interface {
	Create(model T, init bool) (*T, error)
	GenerateProject(model T) (*T, error)
	AddFileToLayer(model T, layer string, paths []string, filename string, content string) (*T, error)
	List() (*([]T), error)
	ListBySet(set string) (*([]T), error)
	ListBySetAndLayers(set string, layers ...string) (*([]T), error)
	ListByWorkspace(workspaceName string) (*([]T), error)
	GetByFullNameWorkspace(name string, workspaceName string) (*T, error)
	ListByFullNameWorkspace(name string, workspaceName string) (*([]T), error)
	CheckIfWorkingOnProject() (*T, error)
	IsDirectoryReserved(path string) (*T, error)
	Remove(name string, workspaceName string, force bool, permanent bool) (*T, error)
	Build(name string, workspaceName string) (*T, error)
	Test(name string, workspaceName string) (*T, error)
	Clean(name string, workspaceName string) (*T, error)
	CleanV2(name string, workspaceName string) (*T, error)
	Release(name string, workspaceName string) (*T, error)
	Run(name string, workspaceName string) (*T, error)

	ValidateProjectStructure(model T) (bool, error)
	ValidateProjectDependencies(model T) (bool, error)
	ValidateProjectReferences(model T) (bool, error)
	IsExists(name string, workspaceName string) (bool, error)
	GetHash(name string, workspaceName string) (string, error)

	GetByName(name string) (*T, error)
	ListByGroupName(group string) (*([]T), error)
	CreateLayerFolder(project T, layers ...project.Layer) error
	DeleteLayerFolder(project T, layers ...project.Layer) error
	AddPackageToProject(model T, packages ...project.Package) error
	RemovePackageToProject(model T, packages ...project.Package) error
	AddReferenceToProject(model T, references ...T) error
	RemoveReferenceFromProject(model T, references ...T) error
	ListIndividualByWorkspace(workspaceName string) (*([]T), error)
}
