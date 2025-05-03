package objectresource

import (
	"gopkg.in/yaml.v3"
)

type Encapsulation struct {
	Getter EncapsulationGetter
	Setter EncapsulationSetter
}

func NewEncapsulation(getter EncapsulationGetter, setter EncapsulationSetter) Encapsulation {
	return Encapsulation{
		Getter: getter,
		Setter: setter,
	}
}

func (s *Encapsulation) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Getter EncapsulationGetter `yaml:"Getter"`
				Setter EncapsulationSetter `yaml:"Setter"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Getter = tempObject.Getter
				s.Setter = tempObject.Setter
			}

		}
	}

	return nil
}
