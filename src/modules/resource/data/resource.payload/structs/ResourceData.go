package structs

type ResourceData struct {
	Dictionary []Dictionary
	Data       any
}

func NewResourceData(data any) ResourceData {
	return ResourceData{
		Data: data,
	}
}
func (e ResourceData) Validate() error {
	return nil
}

func (s *ResourceData) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Data any `yaml:"Data"`
	}

	if err := unmarshal(&tempObject); err != nil {
		return err

	} else {
		s.Data = tempObject.Data
	}

	return nil
}
