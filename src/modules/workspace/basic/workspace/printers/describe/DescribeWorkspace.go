package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
)

type ViewModel struct {
	Name string
	Tags []string
}

type ProjectViewModel struct {
	Name string
}

type DescribeWorkspace struct {
	Workspace ViewModel
}

func (s *DescribeWorkspace) Print() error {

	printer := &printerx.DetailPrinter{ShowEmpty: true}
	printer.Print(os.Stdout, s.Workspace)

	return nil
}
