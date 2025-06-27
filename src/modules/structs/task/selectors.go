package task

type Selectors struct {
	Resource Resource
	// Expression Expression
	// - key: domain.resource
	// operator: in
	// values: ["mydomain"]

}

func NewSelectors(resource Resource) Selectors {
	return Selectors{
		Resource: resource,
	}
}

func (s *Selectors) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempObject struct {
		Resource Resource `yaml:"Resource"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Resource = tempObject.Resource
	}

	return nil
}
