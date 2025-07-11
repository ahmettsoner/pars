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

type DescribeTemplate struct {
	Template ViewModel
}

func (s *DescribeTemplate) Print() error {

	printer := &printerx.DetailPrinter{ShowEmpty: true}
	printer.Print(os.Stdout, s.Template)

	return nil
}
