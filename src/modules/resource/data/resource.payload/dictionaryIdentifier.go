package data_resource_payload

import (
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"

	"gopkg.in/yaml.v3"
)

type DictionaryIdentifier struct {
	Key string
}

func NewDictionaryIdentifier(key string) DictionaryIdentifier {
	return DictionaryIdentifier{
		Key: key,
	}
}
func (e DictionaryIdentifier) Validate() error {
	if _string.IsEmpty(e.Key) {
		return &errors.ErrFieldRequired{FieldName: "Key"}
	}
	return nil
}

func (s *DictionaryIdentifier) IsKeyExists() bool {
	return !_string.IsEmpty(s.Key)
}

func (s *DictionaryIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Key string `yaml:"Key"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Key = tempObject.Key
		} else {
			return err
		}

	} else {
		s.Key = value
	}

	return nil
}
