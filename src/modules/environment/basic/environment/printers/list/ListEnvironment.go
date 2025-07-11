package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
)

type ViewModel struct {
	Name string
}

type ListEnvironment struct {
	Environments []ViewModel
}

func (s *ListEnvironment) Print() error {

	format := printerx.Table
	printer, err := printerx.GetPrinter(format)
	if err != nil {
		panic(err)
	}

	err = printer.Print(os.Stdout, s.Environments)
	if err != nil {
		panic(err)
	}

	return nil
}
