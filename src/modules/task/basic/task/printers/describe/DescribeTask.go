package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
)

type ViewModel struct {
	Name string
	Tags []string
}

type DescribeTask struct {
	Task ViewModel
}

func (s *DescribeTask) Print() error {

	printer := &printerx.DetailPrinter{ShowEmpty: true}
	printer.Print(os.Stdout, s.Task)

	return nil
}
