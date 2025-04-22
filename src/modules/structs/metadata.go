package structs

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Tags []string
}

func NewMetadata(tags []string) Metadata {
	return Metadata{
		Tags: tags,
	}
}

func (s *Metadata) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Tags interface{} `yaml:"Tags"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			switch tags := tempObject.Tags.(type) {
			case string:
				for _, tag := range strings.Split(tags, ",") {
					s.Tags = append(s.Tags, strings.TrimSpace(tag))
				}
			case []interface{}:
				for _, tag := range tags {
					s.Tags = append(s.Tags, strings.TrimSpace(fmt.Sprint(tag)))
				}
			}
		} else {
			return err
		}

	}

	return nil
}
