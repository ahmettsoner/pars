package objectresource

import (
	"strings"

	"parsdevkit.net/core/utils"

	parsErrors "parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type DataType struct {
	Name     string
	Package  TypePackage
	Category DataTypeCategory
	Modifier ModifierType
	Generics []DataType
}

func NewDataType(name string, _package TypePackage, category DataTypeCategory, modifier ModifierType, generics []DataType) DataType {
	return DataType{
		Name:     name,
		Package:  _package,
		Category: category,
		Modifier: modifier,
		Generics: generics,
	}
}

func New_Int() DataType {
	return NewDataType(string(ValueTypes.Int), TypePackage{}, DataTypeCategories.Value, ModifierTypes.Object, []DataType(nil))
}
func New_String() DataType {
	return NewDataType(string(ValueTypes.String), TypePackage{}, DataTypeCategories.Value, ModifierTypes.Object, []DataType(nil))
}
func New_Reference(name string, _package TypePackage) DataType {
	return NewDataType(name, _package, DataTypeCategories.Reference, ModifierTypes.Object, []DataType(nil))
}
func New_Generic_Reference(name string, _package TypePackage, generic_parameters ...DataType) DataType {
	return NewDataType(name, _package, DataTypeCategories.Reference, ModifierTypes.Object, generic_parameters)
}

func (s *DataType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name     string           `yaml:"Name"`
				Package  TypePackage      `yaml:"Package"`
				Category DataTypeCategory `yaml:"Category"`
				Modifier ModifierType     `yaml:"Modifier"`
				Generics []DataType       `yaml:"Generics"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Package = tempObject.Package
				s.Category = tempObject.Category
				if tempObject.Category == DataTypeCategories.Value {
					_type, err := DataTypeEnumFromString(tempObject.Name)
					if err != nil {

					} else {
						s.Name = string(_type)
					}
				} else {
					s.Name = tempObject.Name
				}
				s.Modifier = tempObject.Modifier
				s.Generics = tempObject.Generics
			}
		} else {
			return err
		}
	} else {
		if IsDataTypeEnum(value) {
			_type, err := DataTypeEnumFromString(value)
			if err != nil {
				return err
			} else {
				s.Name = string(_type)
				s.Category = DataTypeCategories.Value
			}
		} else {
			s.Name = value
			s.Category = DataTypeCategories.Resource
		}
	}

	if utils.IsEmpty(string(s.Name)) {
		return &parsErrors.ErrFieldRequired{FieldName: "DataType.Name"}
	}

	if utils.IsEmpty(string(s.Category)) {
		s.Category = DetectDataTypeCategory(s.Name)
	}

	if utils.IsEmpty(string(s.Modifier)) {
		s.Modifier = ModifierTypes.Object
	}

	return nil
}

func DetectDataTypeCategory(name string) DataTypeCategory {
	if IsDataTypeEnum(name) {
		return DataTypeCategories.Value
	} else {
		return DataTypeCategories.Resource
	}
}
func DetectDataTypeModifier(name string) (string, ModifierType) {

	if strings.HasSuffix(name, "[]") {
		return strings.TrimSuffix(name, "[]"), ModifierTypes.Array
	} else {
		return name, ModifierTypes.Object
	}
}
