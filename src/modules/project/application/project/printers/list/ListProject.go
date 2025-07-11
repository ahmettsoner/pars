package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
	"parsdevkit.net/models"
)

type ViewModel struct {
	Name        string
	Set         string
	Group       string
	Platform    string
	ProjectType models.ProjectType
	Tags        []string
	Labels      []string
}

type ListProject struct {
	Projects []ViewModel
}

func (s *ListProject) Print() error {

	format := printerx.Table
	printer, err := printerx.GetPrinter(format)
	if err != nil {
		panic(err)
	}

	err = printer.Print(os.Stdout, s.Projects)
	if err != nil {
		panic(err)
	}

	return nil
}
