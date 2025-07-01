package applicationproject

import (
	"fmt"
	"path/filepath"
	"strings"

	"parsdevkit.net/application/models/label"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/core/utilities"
	"parsdevkit.net/models"

	"parsdevkit.net/core/errors"

	"github.com/sirupsen/logrus"
)

type ProjectSpecification struct {
	applicationProject.ProjectIdentifier
	Platform        Platform
	ProjectType     models.ProjectType
	Set             string
	Package         []string
	Labels          []label.Label
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
	GroupObject     applicationGroup.GroupIdentifier
	Runtime         Runtime
	Language        Language
	Schema          Schema
	Configuration   Configuration
}

func NewProjectSpecification(id int, name, group, workspace string, projectType models.ProjectType, groupObject applicationGroup.GroupIdentifier, set string, _package []string, labels []label.Label, path []string, workspaceObject applicationWorkspace.WorkspaceIdentifier, platform Platform, runtime Runtime, schema Schema, configuration Configuration) ProjectSpecification {
	return ProjectSpecification{
		ProjectIdentifier: applicationProject.NewProjectIdentifier(id, name, path, group, workspace),
		ProjectType:       projectType,
		GroupObject:       groupObject,
		Set:               set,
		Package:           _package,
		Labels:            labels,
		WorkspaceObject:   workspaceObject,
		Platform:          platform,
		Runtime:           runtime,
		Schema:            schema,
		Configuration:     configuration,
	}
}
func (s ProjectSpecification) Validate() error {
	if utilities.IsEmpty(s.ProjectIdentifier.Name) {
		return &errors.ErrFieldRequired{FieldName: "ProjectIdentifier.Name"}
	}
	if utilities.IsEmpty(string(s.ProjectType)) || s.ProjectType.String() == "Unknown" {
		return &errors.ErrFieldRequired{FieldName: "ProjectType"}
	}
	return nil
}
func (s *ProjectSpecification) GetAllPackage() []string {

	projectPackages := []string{}
	projectPackages = append(projectPackages, s.GroupObject.Package...)
	projectPackages = append(projectPackages, s.Package...)

	return projectPackages
}
func (s *ProjectSpecification) GetAllPackageWithLayer(layer string) []string {

	projectPackages := []string{}
	projectPackages = append(projectPackages, s.GroupObject.Package...)
	projectPackages = append(projectPackages, s.Package...)
	for _, layerInProject := range s.Configuration.Layers {
		if layerInProject.Name == layer {
			projectPackages = append(projectPackages, layerInProject.Package...)
			break
		}
	}

	return projectPackages
}
func (s *ProjectSpecification) GetPackageString() string {
	return strings.Join(s.Package, "/")
}

func (s *ProjectSpecification) SetPackageFromString(_package string) {
	s.Package = strings.Split(_package, "/")
}
func (s *ProjectSpecification) AppendPackage(_package ...string) {
	s.Package = append(s.Package, _package...)
}

func (s *ProjectSpecification) GetCodeBasePath() string {
	return filepath.Join(s.WorkspaceObject.GetCodeBaseFolder())
}
func (s *ProjectSpecification) GetRelativeGroupPath() string {

	folders := utilities.CombinePaths([]string{s.GroupObject.GetRelativeGroupPath()})

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetRelativeBaseGroupPath() string {
	if !utilities.IsEmpty(s.GroupObject.GetRelativeGroupPath()) {

		paths := utilities.PathToArray(s.GroupObject.GetRelativeGroupPath())[:1]
		folders := utilities.CombinePaths(paths)

		relativeFullPath := filepath.Join(folders...)

		return relativeFullPath
	}
	return ""
}
func (s *ProjectSpecification) GetAbsoluteGroupPath() string {
	return filepath.Join(s.WorkspaceObject.GetCodeBaseFolder(), s.GroupObject.GetRelativeGroupPath())
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetAbsoluteBaseGroupPath() string {
	if !utilities.IsEmpty(s.GroupObject.GetRelativeGroupPath()) {
		paths := utilities.PathToArray(s.GroupObject.GetRelativeGroupPath())[:1]
		absolutePath := filepath.Join(s.WorkspaceObject.GetCodeBaseFolder(), strings.Join(paths, "/"))

		return absolutePath
	}
	return ""
}

func (s *ProjectSpecification) GetRelativeProjectPath() string {

	folders := utilities.CombinePaths([]string{s.GetRelativeGroupPath()}, s.Path)

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetRelativeBaseProjectPath() string {

	folders := utilities.CombinePaths([]string{s.GetRelativeGroupPath()}, s.Path[:1])

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}
func (s *ProjectSpecification) GetAbsoluteProjectPath() string {
	folders := utilities.CombinePaths([]string{s.GetAbsoluteGroupPath()}, s.Path)

	absoluteFullPath := filepath.Join(folders...)

	logrus.Debugf("Project Absolute Path: %v", absoluteFullPath)

	return absoluteFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetAbsoluteBaseProjectPath() string {
	folders := utilities.CombinePaths([]string{s.GetAbsoluteGroupPath()}, s.Path[:1])

	absoluteFullPath := filepath.Join(folders...)

	logrus.Debugf("Project Absolute Path: %v", absoluteFullPath)

	return absoluteFullPath
}
func (s *ProjectSpecification) GetRelativeProjectLayerPath(layer string) string {
	var existingLayer *applicationProject.Layer = nil
	for _, value := range s.Configuration.Layers {
		if value.Name == layer {
			existingLayer = &value
			break
		}
	}

	if existingLayer == nil {
		return ""
	}

	return filepath.Join(s.GetRelativeProjectPath(), existingLayer.Path)
}
func (s *ProjectSpecification) GetAbsoluteProjectLayerPath(layer string) string {
	var existingLayer *applicationProject.Layer = nil
	for _, value := range s.Configuration.Layers {
		if value.Name == layer {
			existingLayer = &value
			break
		}
	}

	if existingLayer == nil {
		return ""
	}

	return filepath.Join(s.GetAbsoluteProjectPath(), existingLayer.Path)
}

func (s *ProjectSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		applicationProject.ProjectIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {
		s.ProjectIdentifier = tempIdentifierObject.ProjectIdentifier
	}

	var tempObject struct {
		Name          string             `yaml:"Name"`
		Platform      Platform           `yaml:"Platform"`
		ProjectType   models.ProjectType `yaml:"ProjectType"`
		Set           string             `yaml:"Set"`
		Path          string             `yaml:"Path"`
		Package       interface{}        `yaml:"Package"`
		Labels        []label.Label      `yaml:"Labels"`
		Runtime       Runtime            `yaml:"Runtime"`
		Language      Language           `yaml:"Language"`
		Schema        Schema             `yaml:"Schema"`
		Configuration Configuration      `yaml:"Configuration"`
	}

	err := unmarshal(&tempObject)
	if err != nil {
		return err
	}

	s.Platform = tempObject.Platform
	s.ProjectType = tempObject.ProjectType
	s.Set = tempObject.Set
	s.Path = utilities.PathToArray(tempObject.Path)

	switch packages := tempObject.Package.(type) {
	case string:
		s.SetPackageFromString(packages)
	case []interface{}:
		for _, _package := range packages {
			s.AppendPackage(fmt.Sprint(_package))
		}
	}

	s.Labels = tempObject.Labels
	s.Runtime = tempObject.Runtime
	s.Language = tempObject.Language
	s.Schema = tempObject.Schema
	s.Configuration = tempObject.Configuration

	if len(s.Package) == 0 && !utilities.IsEmpty(s.Name) {
		s.AppendPackage(s.Name)
	}

	if len(s.Path) == 0 {
		s.Path = utilities.PathToArray(tempObject.Name)
	}

	return nil
}
