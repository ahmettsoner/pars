package printerx

import (
	"fmt"
)

type Format string

const (
	Table Format = "table"
	JSON  Format = "json"
	YAML  Format = "yaml"
	Plain Format = "plain"
)

func GetPrinter(format Format) (Printer, error) {
	switch format {
	case Table:
		return &TablePrinter{}, nil
	case JSON:
		return &JSONPrinter{}, nil
	case YAML:
		return &YAMLPrinter{}, nil
	case Plain:
		return &PlainPrinter{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}
