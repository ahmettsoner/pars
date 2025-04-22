package sharedtemplate

import (
	"fmt"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Template struct {
	Source  TemplateSourceType
	Content string
}

func NewTemplate(source TemplateSourceType, content string) Template {
	return Template{
		Source:  source,
		Content: content,
	}
}

func (s *Template) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Source  TemplateSourceType `yaml:"Source"`
		Content string             `yaml:"Content"`
	}

	var rawTemplate map[interface{}]interface{}
	if err := unmarshal(&rawTemplate); err != nil {
		return err
	}

	if len(rawTemplate) == 1 {
		for key, value := range rawTemplate {
			source, err := TemplateSourceTypeEnumFromString(key.(string))
			if err != nil {
				return err
			}

			s.Source = source
			s.Content = value.(string)
			break
		}
	} else if len(rawTemplate) > 1 {
		if err := unmarshal(&tempObject); err != nil {
			if _, ok := err.(*yaml.TypeError); !ok {
				return err
			}
		} else {
			s.Source = tempObject.Source
			s.Content = tempObject.Content
		}
	} else {
		return fmt.Errorf("invalid template format")
	}

	if utils.IsEmpty(s.Source.String()) {
		return &errors.ErrFieldRequired{FieldName: "Source"}
	}

	if utils.IsEmpty(s.Content) {
		return &errors.ErrFieldRequired{FieldName: "Content"}
	}

	return nil
}
