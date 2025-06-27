package sharedtemplate

type TemplateConfiguration struct {
}

func NewTemplateConfiguration() TemplateConfiguration {
	return TemplateConfiguration{}
}

func (s *TemplateConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

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
