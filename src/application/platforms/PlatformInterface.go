package platforms

import (
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs"
	applicationProject "parsdevkit.net/application/structs/project"
)

type PlatformInterface[T schemas.SchemaInterface] interface {
	GetKey() string
	CreateGroup(project T) error
	DeleteGroup(project T)
	AddToGroup(project T) error
	RemoveFromGroup(project T) error
	RemoveDefaultFiles(project T) error

	GetGroupFileName(project T) string
	IsGroupFileExists(project T) (bool, error)
	IsGroupFolderExists(project T) (bool, error)
	IsGroupExists(project T, controlFile string) (bool, error)
	ListProjectsFromGroup(project T) ([]T, error)
	HasProjectOnGroup(project T) (bool, error)

	CreateProject(project T) error
	RemoveProject(project T) error
	BuildProject(project T) error
	CleanProject(project T) error
	InstallProject(project T) error
	TestProject(project T) error
	RunProject(project T) error
	PackageProject(project T) error

	PrintPackage(packages []string) string
	PrintDataType(dataType structs.DataType) string
	PrintVisibility(visibility structs.VisibilityType) string

	IsProjectFolderExists(project T) (bool, error)
	GetProjectFileName(project T) string
	IsProjectFileExists(project T) (bool, error)
	AddFolderToProjectDefinition(project T, paths ...string) error
	RemoveFolderFromProjectDefinition(project T, paths ...string) error
	ListFoldersFromProjectDefinition(project T) ([]string, error)

	CreateLayerFolder(project T, layers ...applicationProject.Layer) error
	ListLayersFromProject(projectSpecification T) ([]applicationProject.Layer, error)
	HasLayerOnProject(project T, layer string) (bool, error)
	IsLayerFolderExists(project T, layer string) (bool, error)
	IsLayerFoldersExists(project T) (bool, error)

	AddPackageToProject(project T, packages []applicationProject.Package) error
	RemovePackageFromProject(project T, packages []applicationProject.Package) error
	ListPackagesFromProject(project T) ([]applicationProject.Package, error)
	GetPackageFromProject(project T, _package applicationProject.Package) error
	HasPackageOnProject(project T, _package applicationProject.Package) (bool, error)

	AddReferenceToProject(project T, references []T) error
	RemoveReferenceFromProject(project T, references []T) error
	ListReferencesFromProject(project T) ([]T, error)
	GetReferenceFromProject(project T, reference T) error
	HasReferenceOnProject(project T, reference T) (bool, error)

	NormalizeText(text string) string
}
