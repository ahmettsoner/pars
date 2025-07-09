package contracts

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
)

type ProjectServiceInterface[T schemas.SchemaInterface] interface {
	//General
	Create(model T, init bool) (*T, error) //GenerateProject, CreateProjectFolder, SaveProject
	// Destroy(model T, init bool) (*T, error) //Remove, RemoveProjectFolder, DeleteProject
	GenerateProject(model T) (*T, error)     //with provider
	UndoGenerateProject(model T) (*T, error) //with provider
	// RemoveProject(model T) (*T, error) //with provider
	SaveProject(model T) (*T, error)     //Todb
	UndoSaveProject(model T) (*T, error) //from db
	// DeleteProject(model T) (*T, error) //from db
	GetByName(name string) (*T, error)
	Remove(name string, workspaceName string, force bool, permanent bool) (*T, error) //Object üzerinden yapılablir Create gibi, buda RemoveByName olabilir
	List() (*([]T), error)                                                            //ListAll olarak değişecek
	IsExists(name string, workspaceName string) (bool, error)

	//Query
	// GetByName(name string, workspaceName string) (*T, error)
	ListByFilter(set, workspace string, layers []string, tags []string, labels []label.Label) (*([]T), error)

	GetByFullNameWorkspace(name string, workspaceName string) (*T, error) //GetByName olarak değişecek
	ListBySet(set string) (*([]T), error)
	ListBySetAndLayers(set string, layers ...string) (*([]T), error)
	ListByWorkspace(workspaceName string) (*([]T), error)
	ListByFullNameWorkspace(name string, workspaceName string) (*([]T), error)
	ListByGroupName(group string) (*([]T), error)
	ListIndividualByWorkspace(workspaceName string) (*([]T), error)

	//Helper
	GetHash(name string, workspaceName string) (string, error)
	GetDefaultPlatformProjectType(model T) (models.ProjectType, error)

	//Structure (File/Folder)
	CheckIfWorkingOnProject() (*T, error)
	IsDirectoryReserved(path string) (*T, error)
	ValidateProjectStructure(model T) (bool, error)
	CreateProjectFolder(model T, paths ...string) (string, error)
	DeleteProjectFolder(model T, paths ...string) (string, error)
	RemoveUnnecessaryFiles(model T) (bool, error)

	//Layer
	CreateProjectLayers(project T) error
	AddProjectLayer(project T, layers ...project.Layer) error
	CreateLayerFolder(project T, layers ...project.Layer) error
	DeleteLayerFolder(project T, layers ...project.Layer) error
	AddFileToLayer(model T, layer string, paths []string, filename string, content string) (*T, error)

	//Reference
	SetProjectReferences(model T) (bool, error)
	ValidateProjectReferences(model T) (bool, error)
	AddReferenceToProject(model T, references ...T) error
	RemoveReferenceFromProject(model T, references ...T) error
	// ListReferences

	//Dependency
	SetProjectDependencies(model T) (bool, error)
	ValidateProjectDependencies(model T) (bool, error)
	AddDependenciesToProject(model T, packages ...project.Dependency) error
	RemoveDependencyFromProject(model T, packages ...project.Dependency) error
	//ListPackages

	// Actions
	Build(name string, workspaceName string) (*T, error)
	Test(name string, workspaceName string) (*T, error)
	Clean(name string, workspaceName string) (*T, error)
	CleanV2(name string, workspaceName string) (*T, error)
	Release(name string, workspaceName string) (*T, error)
	Run(name string, workspaceName string) (*T, error)
}
