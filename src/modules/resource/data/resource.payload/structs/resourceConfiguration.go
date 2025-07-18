package structs

type ResourceConfiguration struct {
}

func NewResourceConfiguration() ResourceConfiguration {
	return ResourceConfiguration{}
}

func (s *ResourceConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
	}

	return nil
}
