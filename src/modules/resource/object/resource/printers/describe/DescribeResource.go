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

type DescribeResource struct {
	Resource ViewModel
}

func (s *DescribeResource) Print() error {

	printer := &printerx.DetailPrinter{ShowEmpty: true}
	printer.Print(os.Stdout, s.Resource)

	return nil
}
