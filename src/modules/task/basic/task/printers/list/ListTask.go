package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
)

type ViewModel struct {
	Name string
	Tags []string
}

type ListTask struct {
	Tasks []ViewModel
}

func (s *ListTask) Print() error {

	format := printerx.Table
	printer, err := printerx.GetPrinter(format)
	if err != nil {
		panic(err)
	}

	err = printer.Print(os.Stdout, s.Tasks)
	if err != nil {
		panic(err)
	}

	return nil
}
