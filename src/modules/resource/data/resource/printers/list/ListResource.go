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

type ListResource struct {
	Resources []ViewModel
}

func (s *ListResource) Print() error {

	format := printerx.Table
	printer, err := printerx.GetPrinter(format)
	if err != nil {
		panic(err)
	}

	err = printer.Print(os.Stdout, s.Resources)
	if err != nil {
		panic(err)
	}

	return nil
}
