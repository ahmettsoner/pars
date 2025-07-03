package codetemplate

import (
	"fmt"

	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
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
func (s Template) Validate() error {
	if _string.IsEmpty(s.Source.String()) {
		return &errors.ErrFieldRequired{FieldName: "Source"}
	}

	if _string.IsEmpty(s.Content) {
		return &errors.ErrFieldRequired{FieldName: "Content"}
	}
	return nil
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
			// if _, ok := err.(*yaml.TypeError); !ok {
			// 	return err
			// }
			return err

		} else {
			s.Source = tempObject.Source
			s.Content = tempObject.Content
		}
	} else {
		return fmt.Errorf("invalid template format")
	}

	return nil
}
