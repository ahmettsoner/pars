package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
)

type ViewModel struct {
	Name string
	Tags []string
}

type ListWorkspace struct {
	Workspaces []ViewModel
}

func (s *ListWorkspace) Print() error {

	format := printerx.Table
	printer, err := printerx.GetPrinter(format)
	if err != nil {
		panic(err)
	}

	err = printer.Print(os.Stdout, s.Workspaces)
	if err != nil {
		panic(err)
	}

	return nil
}
