package label

import (
	"strings"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Label struct {
	Key   string
	Value string
}

func NewLabel(key, value string) Label {
	return Label{
		Key:   key,
		Value: value,
	}
}

func NewLabel_KeyOnly(key string) Label {
	return Label{
		Key: key,
	}
}
func (e Label) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Key, v.Required),
	)
}

func (s *Label) IsKeyExists() bool {
	return !utils.IsEmpty(s.Key)
}

func (s *Label) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var value string
	if err := unmarshal(&value); err == nil {
		var parts []string = strings.Split(value, "=")
		if len(parts) == 1 {
			s.Key = value
		} else if len(parts) == 2 {
			key := strings.TrimSpace(strings.ToLower(parts[0]))
			_value := strings.TrimSpace(parts[1])

			s.Key = key

			if utils.IsEmpty(s.Key) {
				return &errors.InvalidLanguageError{Value: key}
			}
			s.Value = _value
		} else {
			return &errors.InvalidFormatForLanguageError{Value: value}
		}

	} else {
		if _, ok := err.(*yaml.TypeError); ok {
			var labels map[string]string
			if err := unmarshal(&labels); err == nil {
				if len(labels) == 1 {
					for key, value := range labels {
						s.Key = key
						s.Value = value
						break
					}
				} else {
					var tempObject struct {
						Key   string `yaml:"Key"`
						Value string `yaml:"Value"`
					}

					if err := unmarshal(&tempObject); err != nil {
						// if _, ok := err.(*yaml.TypeError); !ok {
						// 	return err
						// }
						return err
					} else {
						s.Key = tempObject.Key
						s.Value = tempObject.Value
					}
				}
			}
		} else {
			return err
		}
	}

	return nil
}
