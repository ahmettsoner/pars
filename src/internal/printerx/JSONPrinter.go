package printerx

import (
	"encoding/json"
	"io"
)

type JSONPrinter struct{}

func (p *JSONPrinter) Print(out io.Writer, data any) error {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
