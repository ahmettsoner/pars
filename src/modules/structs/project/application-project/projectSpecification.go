package applicationproject

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"parsdevkit.net/models"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"github.com/sirupsen/logrus"
)

type ProjectSpecification struct {
	ProjectIdentifier
	Platform        Platform
	ProjectType     models.ProjectType
	Set             string
	Package         []string
	Path            []string
	Labels          []label.Label
	WorkspaceObject workspace.WorkspaceSpecification
	GroupObject     group.GroupSpecification
	Runtime         Runtime
	Language        Language
	Schema          Schema
	Configuration   Configuration
}

func NewProjectSpecification(id int, name, group, workspace string, projectType models.ProjectType, groupObject group.GroupSpecification, set string, _package []string, labels []label.Label, path []string, workspaceObject workspace.WorkspaceSpecification, platform Platform, runtime Runtime, schema Schema, configuration Configuration) ProjectSpecification {
	return ProjectSpecification{
		ProjectIdentifier: NewProjectIdentifier(id, name, group, workspace),
		ProjectType:       projectType,
		GroupObject:       groupObject,
		Set:               set,
		Package:           _package,
		Labels:            labels,
		Path:              path,
		WorkspaceObject:   workspaceObject,
		Platform:          platform,
		Runtime:           runtime,
		Schema:            schema,
		Configuration:     configuration,
	}
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

	folders := utils.CombinePaths([]string{s.GroupObject.GetRelativeGroupPath()})

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetRelativeBaseGroupPath() string {
	if !utils.IsEmpty(s.GroupObject.GetRelativeGroupPath()) {

		paths := utils.PathToArray(s.GroupObject.GetRelativeGroupPath())[:1]
		folders := utils.CombinePaths(paths)

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
	if !utils.IsEmpty(s.GroupObject.GetRelativeGroupPath()) {
		paths := utils.PathToArray(s.GroupObject.GetRelativeGroupPath())[:1]
		absolutePath := filepath.Join(s.WorkspaceObject.GetCodeBaseFolder(), strings.Join(paths, "/"))

		return absolutePath
	}
	return ""
}
func (s *ProjectSpecification) GetProjectPath() string {

	folders := utils.CombinePaths(s.Path)

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}
func (s *ProjectSpecification) GetRelativeProjectPath() string {

	folders := utils.CombinePaths([]string{s.GetRelativeGroupPath()}, s.Path)

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetRelativeBaseProjectPath() string {

	folders := utils.CombinePaths([]string{s.GetRelativeGroupPath()}, s.Path[:1])

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}
func (s *ProjectSpecification) GetAbsoluteProjectPath() string {
	folders := utils.CombinePaths([]string{s.GetAbsoluteGroupPath()}, s.Path)

	absoluteFullPath := filepath.Join(folders...)

	logrus.Debugf("Project Absolute Path: %v", absoluteFullPath)

	return absoluteFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetAbsoluteBaseProjectPath() string {
	folders := utils.CombinePaths([]string{s.GetAbsoluteGroupPath()}, s.Path[:1])

	absoluteFullPath := filepath.Join(folders...)

	logrus.Debugf("Project Absolute Path: %v", absoluteFullPath)

	return absoluteFullPath
}
func (s *ProjectSpecification) GetRelativeProjectLayerPath(layer string) string {
	var existingLayer *Layer = nil
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
	var existingLayer *Layer = nil
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
		ProjectIdentifier
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
	s.Path = utils.PathToArray(tempObject.Path)

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

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	if len(s.Package) == 0 {
		s.AppendPackage(s.Name)
	}

	if len(s.Path) == 0 {
		s.Path = utils.PathToArray(tempObject.Name)
	}

	if s.ProjectType.String() == "Unknown" {
		return &errors.ErrFieldRequired{FieldName: "ProjectType"}
	}

	return nil
}

func (s *ProjectSpecification) UnmarshalJSON(data []byte) error {

	var tempIdentifierObject struct {
		ProjectIdentifier
	}

	if err := json.Unmarshal(data, &tempIdentifierObject); err != nil {
		return err
	} else {

		s.ProjectIdentifier = tempIdentifierObject.ProjectIdentifier
	}

	var tempObject struct {
		Platform        Platform
		ProjectType     models.ProjectType
		Set             string
		Package         []string
		Labels          []label.Label
		Path            []string
		WorkspaceObject workspace.WorkspaceSpecification
		GroupObject     group.GroupSpecification
		Runtime         Runtime
		Language        Language
		Schema          Schema
		Configuration   Configuration
	}

	err := json.Unmarshal(data, &tempObject)
	if err != nil {
		return err
	}

	s.Platform = tempObject.Platform
	s.ProjectType = tempObject.ProjectType
	s.Set = tempObject.Set
	s.Path = tempObject.Path
	s.WorkspaceObject = tempObject.WorkspaceObject
	s.GroupObject = tempObject.GroupObject
	s.Package = tempObject.Package
	s.Labels = tempObject.Labels
	s.Runtime = tempObject.Runtime
	s.Language = tempObject.Language
	s.Schema = tempObject.Schema
	s.Configuration = tempObject.Configuration

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	if s.ProjectType.String() == "Unknown" {
		return &errors.ErrFieldRequired{FieldName: "ProjectType"}
	}

	return nil
}
