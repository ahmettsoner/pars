package dataresource

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

// Dictionary global bi struct olarak ta tanımlanabilmeli, hem global hem  resource bağımlı şekilde tanımlanabilmeli
type Dictionary struct {
	DictionaryIdentifier
	Translates map[string]string
}

func NewDictionary(key string, translates map[string]string) Dictionary {
	return Dictionary{
		DictionaryIdentifier: NewDictionaryIdentifier(key),
		Translates:           translates,
	}
}

func (s *Dictionary) IsKeyExists() bool {
	return !utils.IsEmpty(s.Key)
}

func (s *Dictionary) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempIdentifierObject struct {
		DictionaryIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {
		s.DictionaryIdentifier = tempIdentifierObject.DictionaryIdentifier
	}

	var tempObject struct {
		Translates map[string]string `yaml:"Translates"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Translates = tempObject.Translates

	}

	if utils.IsEmpty(s.Key) {
		return &errors.ErrFieldRequired{FieldName: "Dictionary.Key"}
	}

	return nil
}
