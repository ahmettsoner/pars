package platforms

import (
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs"
	"parsdevkit.net/application/structs/project"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
)

type PlatformInterfaceBase interface {
	GetKey() models.PlatformType
}
type PlatformInterface[T schemas.SchemaInterface] interface {
	GetKey() models.PlatformType
	// ──────────────── Project Lifecycle ────────────────
	CreateProject(project T) error
	RemoveProject(project T) error

	// ──────────────── Group Operations ────────────────
	CreateGroup(project T) error
	DeleteGroup(project T) //error eklenecek
	AddToGroup(project T) error
	RemoveFromGroup(project T) error
	ListProjectsFromGroup(project T) ([]T, error)
	HasProjectOnGroup(project T) (bool, error)
	IsGroupExists(project T, controlFile string) (bool, error)
	IsGroupFileExists(project T) (bool, error)
	IsGroupFolderExists(project T) (bool, error)
	GetGroupFileName(project T) string

	// ──────────────── Build & Execution ────────────────
	BuildProject(project T) error
	CleanProject(project T) error
	InstallProject(project T) error
	// RestoreDependencies(project T) error // öneri
	TestProject(project T) error
	RunProject(project T) error
	PackageProject(project T) error
	// FormatProject(project T) error // öneri
	// LintProject(project T) error   // öneri

	// ──────────────── Project File & Folder Management ────────────────
	GetProjectFileName(project applicationProject.ProjectSpecification) string
	IsProjectFileExists(project T) (bool, error)
	IsProjectFolderExists(project T) (bool, error)
	AddFolderToProjectDefinition(project T, paths ...string) error
	RemoveFolderFromProjectDefinition(project T, paths ...string) error
	ListFoldersFromProjectDefinition(project T) ([]string, error)

	// ──────────────── Layer Management ────────────────
	CreateLayerFolder(project T, layers ...project.Layer) error
	ListLayersFromProject(project T) ([]project.Layer, error)
	HasLayerOnProject(project T, layer string) (bool, error)
	IsLayerFolderExists(project T, layer string) (bool, error)
	IsLayerFoldersExists(project T) (bool, error)

	// ──────────────── Dependency Management ────────────────
	AddDependenciesToProject(project T, dependencies []project.Dependency) error
	RemoveDependenciesFromProject(project T, dependencies []project.Dependency) error
	ListDependenciesFromProject(project T) ([]project.Dependency, error)
	GetDependencyFromProject(project T, _package project.Dependency) error
	HasDependencyOnProject(project T, _package project.Dependency) (bool, error)
	PrintDependencies(dependencies []string) string

	// ──────────────── Reference Management ────────────────
	AddReferenceToProject(project T, references []applicationProject.Reference) error
	RemoveReferenceFromProject(project T, references []applicationProject.Reference) error
	ListReferencesFromProject(project T) ([]applicationProject.Reference, error)
	GetReferenceFromProject(project T, reference applicationProject.Reference) error
	HasReferenceOnProject(project T, reference applicationProject.Reference) (bool, error)

	// ──────────────── Helpers ────────────────
	PrintDataType(dataType structs.DataType) string
	PrintVisibility(visibility structs.VisibilityType) string
	GetDefaultPlatformProjectType(model T) models.ProjectType
	NormalizeText(text string) string

	// // ──────────────── File Cleanup ────────────────
	RemoveDefaultFiles(project T) error
	// RemovePlatformGeneratedFiles(project T) error //obj, bin, node_modules vs silinmesi
	// ListDefaultFiles(project T) ([]string, error) // Default file/folder belirlenmesi
	// ListPlatformGeneratedFiles(project T) ([]string, error) // Default file/folder belirlenmesi

}
