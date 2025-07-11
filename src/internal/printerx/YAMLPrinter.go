package printerx

import (
	"io"

	"gopkg.in/yaml.v3"
)

type YAMLPrinter struct{}

func (p *YAMLPrinter) Print(out io.Writer, data any) error {
	return yaml.NewEncoder(out).Encode(data)
}
