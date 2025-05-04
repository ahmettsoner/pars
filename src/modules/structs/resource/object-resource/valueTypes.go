package objectresource

import (
	"fmt"
	"strings"
)

type ValueTypeDefinition string

var ValueTypes = struct {
	Object    ValueTypeDefinition
	ShortInt  ValueTypeDefinition
	Int       ValueTypeDefinition
	LongInt   ValueTypeDefinition
	Float     ValueTypeDefinition
	Float64   ValueTypeDefinition
	Double    ValueTypeDefinition
	Decimal   ValueTypeDefinition
	Char      ValueTypeDefinition
	String    ValueTypeDefinition
	Date      ValueTypeDefinition
	DateTime  ValueTypeDefinition
	Time      ValueTypeDefinition
	Boolean   ValueTypeDefinition
	Byte      ValueTypeDefinition
	ShortBlob ValueTypeDefinition
	Blob      ValueTypeDefinition
	LongBlob  ValueTypeDefinition
}{
	Object:    "Object",
	ShortInt:  "ShortInt",
	Int:       "Int",
	LongInt:   "LongInt",
	Float:     "Float",
	Float64:   "Float64",
	Double:    "Double",
	Decimal:   "Decimal",
	Char:      "Char",
	String:    "String",
	Date:      "Date",
	DateTime:  "DateTime",
	Time:      "Time",
	Boolean:   "Boolean",
	Byte:      "Byte",
	ShortBlob: "ShortBlob",
	Blob:      "Blob",
	LongBlob:  "LongBlob",
}

func (c ValueTypeDefinition) String() string {
	switch c {

	case ValueTypes.Object:
		return "Object"
	case ValueTypes.ShortInt:
		return "ShortInt"
	case ValueTypes.Int:
		return "Int"
	case ValueTypes.LongInt:
		return "LongInt"
	case ValueTypes.Float:
		return "Float"
	case ValueTypes.Float64:
		return "Float64"
	case ValueTypes.Double:
		return "Double"
	case ValueTypes.Decimal:
		return "Decimal"
	case ValueTypes.Char:
		return "Char"
	case ValueTypes.String:
		return "String"
	case ValueTypes.Date:
		return "Date"
	case ValueTypes.DateTime:
		return "DateTime"
	case ValueTypes.Time:
		return "Time"
	case ValueTypes.Boolean:
		return "Boolean"
	case ValueTypes.Byte:
		return "Byte"
	case ValueTypes.ShortBlob:
		return "ShortBlob"
	case ValueTypes.Blob:
		return "Blob"
	case ValueTypes.LongBlob:
		return "LongBlob"
	default:
		return "Unknown"
	}
}

func DataTypeEnumFromString(enum string) (ValueTypeDefinition, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Object"):
		return ValueTypes.Object, nil
	case strings.ToLower("ShortInt"):
		return ValueTypes.ShortInt, nil
	case strings.ToLower("Int"):
		return ValueTypes.Int, nil
	case strings.ToLower("LongInt"):
		return ValueTypes.LongInt, nil
	case strings.ToLower("Float"):
		return ValueTypes.Float, nil
	case strings.ToLower("Float64"):
		return ValueTypes.Float64, nil
	case strings.ToLower("Double"):
		return ValueTypes.Double, nil
	case strings.ToLower("Decimal"):
		return ValueTypes.Decimal, nil
	case strings.ToLower("Char"):
		return ValueTypes.Char, nil
	case strings.ToLower("String"):
		return ValueTypes.String, nil
	case strings.ToLower("Date"):
		return ValueTypes.Date, nil
	case strings.ToLower("DateTime"):
		return ValueTypes.DateTime, nil
	case strings.ToLower("Time"):
		return ValueTypes.Time, nil
	case strings.ToLower("Boolean"):
		return ValueTypes.Boolean, nil
	case strings.ToLower("Byte"):
		return ValueTypes.Byte, nil
	case strings.ToLower("ShortBlob"):
		return ValueTypes.ShortBlob, nil
	case strings.ToLower("Blob"):
		return ValueTypes.Blob, nil
	case strings.ToLower("LongBlob"):
		return ValueTypes.LongBlob, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func IsDataTypeEnum(enum string) bool {
	switch strings.ToLower(enum) {
	case strings.ToLower("Object"):
		return true
	case strings.ToLower("ShortInt"):
		return true
	case strings.ToLower("Int"):
		return true
	case strings.ToLower("LongInt"):
		return true
	case strings.ToLower("Float"):
		return true
	case strings.ToLower("Float64"):
		return true
	case strings.ToLower("Double"):
		return true
	case strings.ToLower("Decimal"):
		return true
	case strings.ToLower("Char"):
		return true
	case strings.ToLower("String"):
		return true
	case strings.ToLower("Date"):
		return true
	case strings.ToLower("DateTime"):
		return true
	case strings.ToLower("Time"):
		return true
	case strings.ToLower("Boolean"):
		return true
	case strings.ToLower("Byte"):
		return true
	case strings.ToLower("ShortBlob"):
		return true
	case strings.ToLower("Blob"):
		return true
	case strings.ToLower("LongBlob"):
		return true
	default:
		return false
	}
}

func (s *ValueTypeDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := DataTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func DataTypeToArray() []ValueTypeDefinition {
	return []ValueTypeDefinition{
		ValueTypes.Object,
		ValueTypes.ShortInt,
		ValueTypes.Int,
		ValueTypes.LongInt,
		ValueTypes.Float,
		ValueTypes.Float64,
		ValueTypes.Double,
		ValueTypes.Decimal,
		ValueTypes.Char,
		ValueTypes.String,
		ValueTypes.Date,
		ValueTypes.DateTime,
		ValueTypes.Time,
		ValueTypes.Boolean,
		ValueTypes.Byte,
		ValueTypes.ShortBlob,
		ValueTypes.Blob,
		ValueTypes.LongBlob,
	}
}

type DataTypeEnumFlag struct {
	Value ValueTypeDefinition
}

func (e *DataTypeEnumFlag) Type() string {
	return "ValueTypeDefinition"
}

func (e *DataTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *DataTypeEnumFlag) Set(value string) error {
	validEnumValues := DataTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
