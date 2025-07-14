package structs

type ResourceObject struct {
	Dictionary []Dictionary
	Groups     []Group
	Attributes []Attribute
	Methods    []Method
}

func NewResourceObject(attributes []Attribute, methods []Method) ResourceObject {
	return ResourceObject{
		Attributes: attributes,
		Methods:    methods,
	}
}

func (e ResourceObject) Validate() error {
	return nil
}

func (s *ResourceObject) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempObject struct {
		Attributes []Attribute `yaml:"Attributes"`
		Methods    []Method    `yaml:"Methods"`
	}

	if err := unmarshal(&tempObject); err != nil {
		return err

	} else {
		s.Attributes = tempObject.Attributes
		s.Methods = tempObject.Methods
	}

	return nil
}
