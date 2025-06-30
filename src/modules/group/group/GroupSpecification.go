package group

import (
	"fmt"

	applicationGroup "parsdevkit.net/application/structs/group"
	"parsdevkit.net/core/utils"

	"gopkg.in/yaml.v3"
)

type GroupSpecification struct {
	applicationGroup.GroupIdentifier
}

func NewGroupSpecification(id int, name, path string, _package []string) GroupSpecification {
	return GroupSpecification{
		GroupIdentifier: applicationGroup.NewGroupIdentifier(id, name, path, _package),
	}
}

func NewGroupSpecification_Empty(name string) GroupSpecification {
	return GroupSpecification{
		GroupIdentifier: applicationGroup.NewGroupIdentifier(0, name, name, []string{}),
	}
}

func (s *GroupSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		applicationGroup.GroupIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return fmt.Errorf("xxx: Group Specification Çözümlenemedi\n%w", err)
	} else {
		s.GroupIdentifier = tempIdentifierObject.GroupIdentifier
	}

	var tempObject struct {
		Path    string      `yaml:"Path"`
		Package interface{} `yaml:"Package"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return fmt.Errorf("xxx: Group Specification Path ve Package dönüştürme hatası oluştu\n%w", err)
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
	if utils.IsEmpty(s.Path) {
		s.Path = s.Name
	}

	return nil
}
