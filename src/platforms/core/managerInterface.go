package core

import (
	"parsdevkit.net/application/structs"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"
)

type ManagerInterface interface {

	// ──────────────── Project Lifecycle ────────────────
	CreateProject(project applicationproject.ProjectSpecification) error
	RemoveProject(project applicationproject.ProjectSpecification) error

	// ──────────────── Group Operations ────────────────
	CreateGroup(project applicationproject.ProjectSpecification) error
	DeleteGroup(project applicationproject.ProjectSpecification) //error eklenecek
	AddToGroup(project applicationproject.ProjectSpecification) error
	RemoveFromGroup(project applicationproject.ProjectSpecification) error
	ListProjectsFromGroup(project applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error)
	HasProjectOnGroup(project applicationproject.ProjectSpecification) (bool, error)
	IsGroupExists(project applicationproject.ProjectSpecification, controlFile string) (bool, error)
	IsGroupFileExists(project applicationproject.ProjectSpecification) (bool, error)
	IsGroupFolderExists(project applicationproject.ProjectSpecification) (bool, error)
	GetGroupFileName(project applicationproject.ProjectSpecification) string

	// ──────────────── Build & Execution ────────────────
	BuildProject(project applicationproject.ProjectSpecification) error
	CleanProject(project applicationproject.ProjectSpecification) error
	InstallProject(project applicationproject.ProjectSpecification) error
	// RestoreDependencies(project applicationproject.ProjectSpecification) error // öneri
	TestProject(project applicationproject.ProjectSpecification) error
	RunProject(project applicationproject.ProjectSpecification) error
	PackageProject(project applicationproject.ProjectSpecification) error
	// FormatProject(project applicationproject.ProjectSpecification) error // öneri
	// LintProject(project applicationproject.ProjectSpecification) error   // öneri

	// ──────────────── Project File & Folder Management ────────────────
	GetProjectFileName(project applicationproject.ProjectSpecification) string
	IsProjectFileExists(project applicationproject.ProjectSpecification) (bool, error)
	IsProjectFolderExists(project applicationproject.ProjectSpecification) (bool, error)
	AddFolderToProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) error
	RemoveFolderFromProjectDefinition(project applicationproject.ProjectSpecification, paths ...string) error
	ListFoldersFromProjectDefinition(project applicationproject.ProjectSpecification) ([]string, error)

	// ──────────────── Layer Management ────────────────
	CreateLayerFolder(project applicationproject.ProjectSpecification, layers ...applicationProject.Layer) error
	ListLayersFromProject(project applicationproject.ProjectSpecification) ([]applicationProject.Layer, error)
	HasLayerOnProject(project applicationproject.ProjectSpecification, layer string) (bool, error)
	IsLayerFolderExists(project applicationproject.ProjectSpecification, layer string) (bool, error)
	IsLayerFoldersExists(project applicationproject.ProjectSpecification) (bool, error)

	// ──────────────── Package Management ────────────────
	// Package değil dependency olacak
	AddDependenciesToProject(project applicationproject.ProjectSpecification, dependencies []applicationProject.Package) error
	RemoveDependenciesFromProject(project applicationproject.ProjectSpecification, dependencies []applicationProject.Package) error
	ListDependenciesFromProject(project applicationproject.ProjectSpecification) ([]applicationProject.Package, error)
	GetDependencyFromProject(project applicationproject.ProjectSpecification, _package applicationProject.Package) error
	HasDependencyOnProject(project applicationproject.ProjectSpecification, _package applicationProject.Package) (bool, error)
	PrintDependencies(dependencies []string) string

	// ──────────────── Reference Management ────────────────
	AddReferenceToProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error
	RemoveReferenceFromProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error
	ListReferencesFromProject(project applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error)
	GetReferenceFromProject(project applicationproject.ProjectSpecification, reference applicationproject.ProjectSpecification) error
	HasReferenceOnProject(project applicationproject.ProjectSpecification, reference applicationproject.ProjectSpecification) (bool, error)

	// ──────────────── Helpers ────────────────
	PrintDataType(dataType structs.DataType) string
	PrintVisibility(visibility structs.VisibilityType) string
	GetDefaultPlatformProjectType(model applicationproject.ProjectSpecification) models.ProjectType
	NormalizeText(text string) string

	// // ──────────────── File Cleanup ────────────────
	RemoveDefaultFiles(project applicationproject.ProjectSpecification) error
	// RemovePlatformGeneratedFiles(project applicationproject.ProjectSpecification) error //obj, bin, node_modules vs silinmesi
	// ListDefaultFiles(project applicationproject.ProjectSpecification) ([]string, error) // Default file/folder belirlenmesi
	// ListPlatformGeneratedFiles(project applicationproject.ProjectSpecification) ([]string, error) // Default file/folder belirlenmesi

}
