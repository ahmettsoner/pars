package printerx

import "io"

type Printer interface {
	Print(out io.Writer, data any) error
}
