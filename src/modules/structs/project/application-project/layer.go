package applicationproject

import (
	"fmt"
	"path/filepath"
	"strings"

	layerPkg "parsdevkit.net/structs/layer"

	"parsdevkit.net/core/utils"

	"gopkg.in/yaml.v3"
)

type Layer struct {
	layerPkg.LayerIdentifier
	Path    string
	Package []string
	BasedOn string
}

func NewLayer(id int, name, path string, _package []string, based_on string) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(id, name),
		Path:            path,
		Package:         _package,
		BasedOn:         based_on,
	}
}

func NewLayer_Basic(id int, name, path string, _package []string) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(id, name),
		Path:            path,
		Package:         _package,
	}
}

func NewLayer_NameOnly(name string) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(0, name),
	}
}

func NewLayer_NameOnlyDefault(name string) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(0, name),
		Package:         strings.Split(name, ":"),
		Path:            filepath.Join(strings.Split(name, ":")...),
	}
}
func NewLayer_NameOnlyDefaultPath(name string) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(0, name),
		Path:            filepath.Join(strings.Split(name, ":")...),
	}
}

func NewLayer_NameOnlyDefaultPackage(name string) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(0, name),
		Package:         strings.Split(name, ":"),
	}
}

func NewLayer_Empty_Basic(name, path string) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(0, name),
		Path:            path,
	}
}

func (s *Layer) GetPathFromName() string {
	return strings.Join(strings.Split(s.Name, ":"), "/")
}

func (s *Layer) IsPathExists() bool {
	return !utils.IsEmpty(s.Path)
}
func (s *Layer) GetPathAsArray() []string {
	return utils.PathToArray(s.Path)
}

func (s *Layer) GetPackageString() string {
	return strings.Join(s.Package, "/")
}
func (s *Layer) SetPackageFromString(_package string) {
	s.Package = strings.Split(_package, "/")
}
func (s *Layer) AppendPackage(_package ...string) {
	s.Package = append(s.Package, _package...)
}
func (s *Layer) IsPackageExists() bool {
	return len(s.Package) > 0
}

func (s *Layer) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		layerPkg.LayerIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.LayerIdentifier = tempIdentifierObject.LayerIdentifier
	}

	var tempObject struct {
		Path    string      `yaml:"Path"`
		Package interface{} `yaml:"Package"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Path = tempObject.Path

		switch packages := tempObject.Package.(type) {
		case string:
			s.SetPackageFromString(packages)
		case []interface{}:
			for _, _package := range packages {
				s.AppendPackage(fmt.Sprint(_package))
			}
		}
	}

	if len(s.Package) == 0 {
		s.Package = []string{s.Name}
	}
	if !s.IsPathExists() {
		s.Path = s.GetPathFromName()
	}

	return nil
}
