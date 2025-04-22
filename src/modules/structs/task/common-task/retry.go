package commontask

import (
	"gopkg.in/yaml.v3"
)

type Retry struct {
	Count    int
	Interval int
	Delay    int
}

func NewRetry(count int, interval int, delay int) Retry {
	return Retry{
		Count:    count,
		Interval: interval,
		Delay:    delay,
	}
}

func (s *Retry) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Count    int `yaml:"Count"`
		Interval int `yaml:"Interval"`
		Delay    int `yaml:"Delay"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Count = tempObject.Count
		s.Interval = tempObject.Interval
		s.Delay = tempObject.Delay
	}

	return nil
}
