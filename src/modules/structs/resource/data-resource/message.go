package dataresource

import (
	"reflect"

	"parsdevkit.net/pkg/errors"

	"gopkg.in/yaml.v3"
	_string "parsdevkit.net/pkg/utilities/string"
)

type Message struct {
	Text       string
	Dictionary DictionaryIdentifier
}

func NewMessage(text string, dictionary DictionaryIdentifier) Message {
	return Message{
		Text:       text,
		Dictionary: dictionary,
	}
}

func (e Message) Validate() error {
	if _string.IsEmpty(e.Text) {
		return &errors.ErrFieldRequired{FieldName: "Text"}
	}
	return nil
}
func (s *Message) IsTextExists() bool {
	return !_string.IsEmpty(s.Text)
}

func (s *Message) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Text       string               `yaml:"Text"`
				Dictionary DictionaryIdentifier `yaml:"RefMessage"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

			} else {
				s.Text = tempObject.Text
				s.Dictionary = tempObject.Dictionary
			}

		} else {
			return err
		}
	} else {
		s.Text = value
	}

	if reflect.DeepEqual(s.Dictionary, Dictionary{}) {
		return &errors.ErrFieldRequired{FieldName: "Text|Dictionary"}
	}

	return nil
}
