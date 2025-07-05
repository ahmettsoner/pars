package core

import (
	"parsdevkit.net/application/structs"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"
)

type ManagerInterface interface {

	// ──────────────── Project Lifecycle ────────────────
	CreateProject(project applicationproject.ProjectBaseStruct) error
	RemoveProject(project applicationproject.ProjectBaseStruct) error

	// ──────────────── Group Operations ────────────────
	CreateGroup(project applicationproject.ProjectBaseStruct) error
	DeleteGroup(project applicationproject.ProjectBaseStruct) //error eklenecek
	AddToGroup(project applicationproject.ProjectBaseStruct) error
	RemoveFromGroup(project applicationproject.ProjectBaseStruct) error
	ListProjectsFromGroup(project applicationproject.ProjectBaseStruct) ([]applicationproject.ProjectBaseStruct, error)
	HasProjectOnGroup(project applicationproject.ProjectBaseStruct) (bool, error)
	IsGroupExists(project applicationproject.ProjectBaseStruct, controlFile string) (bool, error)
	IsGroupFileExists(project applicationproject.ProjectBaseStruct) (bool, error)
	IsGroupFolderExists(project applicationproject.ProjectBaseStruct) (bool, error)
	GetGroupFileName(project applicationproject.ProjectBaseStruct) string

	// ──────────────── Build & Execution ────────────────
	BuildProject(project applicationproject.ProjectBaseStruct) error
	CleanProject(project applicationproject.ProjectBaseStruct) error
	InstallProject(project applicationproject.ProjectBaseStruct) error
	// RestoreDependencies(project applicationproject.ProjectBaseStruct) error // öneri
	TestProject(project applicationproject.ProjectBaseStruct) error
	RunProject(project applicationproject.ProjectBaseStruct) error
	PackageProject(project applicationproject.ProjectBaseStruct) error
	// FormatProject(project applicationproject.ProjectBaseStruct) error // öneri
	// LintProject(project applicationproject.ProjectBaseStruct) error   // öneri

	// ──────────────── Project File & Folder Management ────────────────
	GetProjectFileName(project applicationproject.ProjectBaseStruct) string
	IsProjectFileExists(project applicationproject.ProjectBaseStruct) (bool, error)
	IsProjectFolderExists(project applicationproject.ProjectBaseStruct) (bool, error)
	AddFolderToProjectDefinition(project applicationproject.ProjectBaseStruct, paths ...string) error
	RemoveFolderFromProjectDefinition(project applicationproject.ProjectBaseStruct, paths ...string) error
	ListFoldersFromProjectDefinition(project applicationproject.ProjectBaseStruct) ([]string, error)

	// ──────────────── Layer Management ────────────────
	CreateLayerFolder(project applicationproject.ProjectBaseStruct, layers ...applicationProject.Layer) error
	ListLayersFromProject(project applicationproject.ProjectBaseStruct) ([]applicationProject.Layer, error)
	HasLayerOnProject(project applicationproject.ProjectBaseStruct, layer string) (bool, error)
	IsLayerFolderExists(project applicationproject.ProjectBaseStruct, layer string) (bool, error)
	IsLayerFoldersExists(project applicationproject.ProjectBaseStruct) (bool, error)

	// ──────────────── Package Management ────────────────
	// Package değil dependency olacak
	AddDependenciesToProject(project applicationproject.ProjectBaseStruct, dependencies []applicationProject.Package) error
	RemoveDependenciesFromProject(project applicationproject.ProjectBaseStruct, dependencies []applicationProject.Package) error
	ListDependenciesFromProject(project applicationproject.ProjectBaseStruct) ([]applicationProject.Package, error)
	GetDependencyFromProject(project applicationproject.ProjectBaseStruct, _package applicationProject.Package) error
	HasDependencyOnProject(project applicationproject.ProjectBaseStruct, _package applicationProject.Package) (bool, error)
	PrintDependencies(dependencies []string) string

	// ──────────────── Reference Management ────────────────
	AddReferenceToProject(project applicationproject.ProjectBaseStruct, references []applicationproject.ProjectBaseStruct) error
	RemoveReferenceFromProject(project applicationproject.ProjectBaseStruct, references []applicationproject.ProjectBaseStruct) error
	ListReferencesFromProject(project applicationproject.ProjectBaseStruct) ([]applicationproject.ProjectBaseStruct, error)
	GetReferenceFromProject(project applicationproject.ProjectBaseStruct, reference applicationproject.ProjectBaseStruct) error
	HasReferenceOnProject(project applicationproject.ProjectBaseStruct, reference applicationproject.ProjectBaseStruct) (bool, error)

	// ──────────────── Helpers ────────────────
	PrintDataType(dataType structs.DataType) string
	PrintVisibility(visibility structs.VisibilityType) string
	GetDefaultPlatformProjectType(model applicationproject.ProjectBaseStruct) models.ProjectType
	NormalizeText(text string) string

	// // ──────────────── File Cleanup ────────────────
	RemoveDefaultFiles(project applicationproject.ProjectBaseStruct) error
	// RemovePlatformGeneratedFiles(project applicationproject.ProjectBaseStruct) error //obj, bin, node_modules vs silinmesi
	// ListDefaultFiles(project applicationproject.ProjectBaseStruct) ([]string, error) // Default file/folder belirlenmesi
	// ListPlatformGeneratedFiles(project applicationproject.ProjectBaseStruct) ([]string, error) // Default file/folder belirlenmesi

}
