package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
)

type ViewModel struct {
	Name    string
	Path    string
	Package []string
	Tags    []string
}

type ListGroup struct {
	Groups []ViewModel
}

func (s *ListGroup) Print() error {

	format := printerx.Table
	printer, err := printerx.GetPrinter(format)
	if err != nil {
		panic(err)
	}

	err = printer.Print(os.Stdout, s.Groups)
	if err != nil {
		panic(err)
	}

	return nil
}
