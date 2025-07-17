package structs

import (
	"fmt"
	"path/filepath"
	"strings"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/pkg/utilities/file"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/pkg/errors"

	"github.com/sirupsen/logrus"
)

type ProjectSpecification struct {
	applicationProject.ProjectIdentifier
	// schemas.SchemaSpecification
	Platform        Platform
	Set             string
	Package         []string
	Labels          []label.Label
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
	GroupObject     applicationGroup.GroupIdentifier
	Layers          []applicationProject.Layer
	Dependencies    []applicationProject.Dependency
	References      []ProjectBaseStruct //schemas.SchemaHeader olmalı?
}

func NewProjectSpecification(
	id int,
	name,
	group,
	workspace string,
	groupObject applicationGroup.GroupIdentifier,
	set string,
	_package []string,
	labels []label.Label,
	path []string,
	workspaceObject applicationWorkspace.WorkspaceIdentifier,
	platform Platform,
	layers []applicationProject.Layer,
	dependencies []applicationProject.Dependency,
	references []ProjectBaseStruct,
) ProjectSpecification {
	return ProjectSpecification{
		ProjectIdentifier: applicationProject.NewProjectIdentifier(id, name, path, group, workspace),
		GroupObject:       groupObject,
		Set:               set,
		Package:           _package,
		Labels:            labels,
		WorkspaceObject:   workspaceObject,
		Platform:          platform,
		Layers:            layers,
		Dependencies:      dependencies,
		References:        references,
	}
}
func (s ProjectSpecification) Validate() error {
	if _string.IsEmpty(s.ProjectIdentifier.Name) {
		return &errors.ErrFieldRequired{FieldName: "ProjectIdentifier.Name"}
	}
	return nil
}
func (s *ProjectSpecification) AppendReferences(references ...ProjectBaseStruct) {
	s.References = append(s.References, references...)
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
	for _, layerInProject := range s.Layers {
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

	folders := file.CombinePaths([]string{s.GroupObject.GetRelativeGroupPath()})

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetRelativeBaseGroupPath() string {
	if !_string.IsEmpty(s.GroupObject.GetRelativeGroupPath()) {

		paths := file.PathToArray(s.GroupObject.GetRelativeGroupPath())[:1]
		folders := file.CombinePaths(paths)

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
	if !_string.IsEmpty(s.GroupObject.GetRelativeGroupPath()) {
		paths := file.PathToArray(s.GroupObject.GetRelativeGroupPath())[:1]
		absolutePath := filepath.Join(s.WorkspaceObject.GetCodeBaseFolder(), strings.Join(paths, "/"))

		return absolutePath
	}
	return ""
}

func (s *ProjectSpecification) GetRelativeProjectPath() string {

	folders := file.CombinePaths([]string{s.GetRelativeGroupPath()}, s.Path)

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetRelativeBaseProjectPath() string {

	folders := file.CombinePaths([]string{s.GetRelativeGroupPath()}, s.Path[:1])

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}
func (s *ProjectSpecification) GetAbsoluteProjectPath() string {
	folders := file.CombinePaths([]string{s.GetAbsoluteGroupPath()}, s.Path)

	absoluteFullPath := filepath.Join(folders...)

	logrus.Debugf("Project Absolute Path: %v", absoluteFullPath)

	return absoluteFullPath
}

// TODO: Gerekli testler tamamlanmalı
func (s *ProjectSpecification) GetAbsoluteBaseProjectPath() string {
	folders := file.CombinePaths([]string{s.GetAbsoluteGroupPath()}, s.Path[:1])

	absoluteFullPath := filepath.Join(folders...)

	logrus.Debugf("Project Absolute Path: %v", absoluteFullPath)

	return absoluteFullPath
}
func (s *ProjectSpecification) GetRelativeProjectLayerPath(layer string) string {
	var existingLayer *applicationProject.Layer = nil
	for _, value := range s.Layers {
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
	for _, value := range s.Layers {
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
		Name         string                                 `yaml:"Name"`
		Platform     Platform                               `yaml:"Platform"`
		Set          string                                 `yaml:"Set"`
		Path         string                                 `yaml:"Path"`
		Package      interface{}                            `yaml:"Package"`
		Labels       []label.Label                          `yaml:"Labels"`
		Layers       []applicationProject.Layer             `yaml:"Layers"`       //Burda inline defination eklenmeli, "Persistence:Data:Repository, Persistence:Data:Entity, Persistence:Data:Migration" gibi
		Dependencies []applicationProject.Dependency        `yaml:"Dependencies"` //Burda inline defination eklenmeli, "gopkg.in/yaml.v3@v3.0.1, gopkg.in/gorm" gibi
		References   []applicationProject.ProjectIdentifier `yaml:"References"`   //Burda inline defination eklenmeli, Workspace::Group/Name formatında "pars::core/utils, pars::service/project" gibi
	}

	err := unmarshal(&tempObject)
	if err != nil {
		return err
	}

	s.Platform = tempObject.Platform
	s.Set = tempObject.Set
	s.Path = file.PathToArray(tempObject.Path)

	switch packages := tempObject.Package.(type) {
	case string:
		s.SetPackageFromString(packages)
	case []interface{}:
		for _, _package := range packages {
			s.AppendPackage(fmt.Sprint(_package))
		}
	}
	s.Labels = tempObject.Labels
	s.Layers = tempObject.Layers
	s.Dependencies = tempObject.Dependencies
	for _, ref := range tempObject.References {
		reference := NewProjectBaseStruct(
			schemas.NewSchemaHeader(schemas.StructTypes.Project, PROJECT_KIND, ref.Name, schemas.Metadata{}),
			NewProjectSpecification(
				0,
				"",
				ref.Group,
				ref.Workspace,
				applicationGroup.GroupIdentifier{},
				"",
				[]string(nil),
				[]label.Label(nil),
				[]string(nil),
				applicationWorkspace.WorkspaceIdentifier{},
				Platform{},
				[]applicationProject.Layer(nil),
				[]applicationProject.Dependency(nil),
				[]ProjectBaseStruct(nil),
			),
			Application{},
		)
		s.AppendReferences(reference)
	}

	if len(s.Package) == 0 && !_string.IsEmpty(s.Name) {
		s.AppendPackage(s.Name)
	}

	if len(s.Path) == 0 {
		s.Path = file.PathToArray(tempObject.Name)
	}

	return nil
}
