package core

import (
	applicationproject "parsdevkit.net/structs/project/application-project"
	objectresource "parsdevkit.net/structs/resource/object-resource"
)

type ManagerInterface interface {
	CreateGroup(project applicationproject.ProjectSpecification) error
	DeleteGroup(project applicationproject.ProjectSpecification)
	AddToGroup(project applicationproject.ProjectSpecification) error
	RemoveFromGroup(project applicationproject.ProjectSpecification) error
	RemoveDefaultFiles(project applicationproject.ProjectSpecification) error

	GetGroupFileName(project applicationproject.ProjectSpecification) string
	IsGroupFileExists(project applicationproject.ProjectSpecification) (bool, error)
	IsGroupFolderExists(project applicationproject.ProjectSpecification) (bool, error)
	IsGroupExists(project applicationproject.ProjectSpecification, controlFile string) (bool, error)
	ListProjectsFromGroup(project applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error)
	HasProjectOnGroup(project applicationproject.ProjectSpecification) (bool, error)

	CreateProject(project applicationproject.ProjectSpecification) error
	RemoveProject(project applicationproject.ProjectSpecification) error
	BuildProject(project applicationproject.ProjectSpecification) error
	CleanProject(project applicationproject.ProjectSpecification) error
	InstallProject(project applicationproject.ProjectSpecification) error
	TestProject(project applicationproject.ProjectSpecification) error
	RunProject(project applicationproject.ProjectSpecification) error
	PackageProject(project applicationproject.ProjectSpecification) error

	PrintPackage(packages []string) string
	PrintDataType(dataType objectresource.DataType) string
	PrintVisibility(visibility objectresource.VisibilityType) string

	IsProjectFolderExists(project applicationproject.ProjectSpecification) (bool, error)
	GetProjectFileName(project applicationproject.ProjectSpecification) string
	IsProjectFileExists(project applicationproject.ProjectSpecification) (bool, error)
	AddFolderToProjectDefinition(project applicationproject.ProjectSpecification, paths ...string)
	RemoveFolderFromProjectDefinition(project applicationproject.ProjectSpecification, paths ...string)
	ListFoldersFromProjectDefinition(project applicationproject.ProjectSpecification) ([]string, error)

	CreateLayerFolder(project applicationproject.ProjectSpecification, layers ...applicationproject.Layer) error
	ListLayersFromProject(projectSpecification applicationproject.ProjectSpecification) ([]applicationproject.Layer, error)
	HasLayerOnProject(project applicationproject.ProjectSpecification, layer string) (bool, error)
	IsLayerFolderExists(project applicationproject.ProjectSpecification, layer string) (bool, error)
	IsLayerFoldersExists(project applicationproject.ProjectSpecification) (bool, error)

	AddPackageToProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error
	RemovePackageFromProject(project applicationproject.ProjectSpecification, packages []applicationproject.Package) error
	ListPackagesFromProject(project applicationproject.ProjectSpecification) ([]applicationproject.Package, error)
	GetPackageFromProject(project applicationproject.ProjectSpecification, _package applicationproject.Package) error
	HasPackageOnProject(project applicationproject.ProjectSpecification, _package applicationproject.Package) (bool, error)

	AddReferenceToProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error
	RemoveReferenceFromProject(project applicationproject.ProjectSpecification, references []applicationproject.ProjectSpecification) error
	ListReferencesFromProject(project applicationproject.ProjectSpecification) ([]applicationproject.ProjectSpecification, error)
	GetReferenceFromProject(project applicationproject.ProjectSpecification, reference applicationproject.ProjectSpecification) error
	HasReferenceOnProject(project applicationproject.ProjectSpecification, reference applicationproject.ProjectSpecification) (bool, error)

	NormalizeText(text string) string
}
