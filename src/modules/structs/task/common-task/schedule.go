package commontask

import (
	"gopkg.in/yaml.v3"
)

type Schedule struct {
	Interval  string
	StartTime string
	Cron      string
}

func NewSchedule(interval, startTime, cron string) Schedule {
	return Schedule{
		Interval:  interval,
		StartTime: startTime,
		Cron:      cron,
	}
}

func (s *Schedule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Interval  string `yaml:"Interval"`
				StartTime string `yaml:"StartTime"`
				Cron      string `yaml:"Cron"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Interval = tempObject.Interval
				s.StartTime = tempObject.StartTime
				s.Cron = tempObject.Cron
			}

		} else {
			return err
		}

	} else {
		s.Interval = value
	}

	return nil
}
