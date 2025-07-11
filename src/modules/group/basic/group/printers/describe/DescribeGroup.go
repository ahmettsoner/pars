package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
)

type ViewModel struct {
	Name     string
	Path     string
	Package  []string
	Tags     []string
	Projects []ProjectViewModel
}

type ProjectViewModel struct {
	Name string
}

type DescribeGroup struct {
	Group ViewModel
}

func (s *DescribeGroup) Print() error {

	printer := &printerx.DetailPrinter{ShowEmpty: true}
	printer.Print(os.Stdout, s.Group)

	return nil
}
