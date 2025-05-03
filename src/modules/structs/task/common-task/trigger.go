package commontask

import (
	"gopkg.in/yaml.v3"
)

type Trigger struct {
	Event    string
	Schedule Schedule
	OnDemand bool
}

func NewTrigger(event string, schedule Schedule, onDemand bool) Trigger {
	return Trigger{
		Event:    event,
		Schedule: schedule,
		OnDemand: onDemand,
	}
}

func (s *Trigger) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Event    string   `yaml:"Event"`
				Schedule Schedule `yaml:"Schedule"`
				OnDemand bool     `yaml:"OnDemand"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Event = tempObject.Event
				s.Schedule = tempObject.Schedule
				s.OnDemand = tempObject.OnDemand
			}

		} else {
			return err
		}

	} else {
		s.Event = value
	}

	return nil
}
