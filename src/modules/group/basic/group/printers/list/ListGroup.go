package list

import (
	"fmt"
	"os"

	"parsdevkit.net/internal/printerx"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

type ViewModel struct {
	Name    string
	Path    string
	Package []string
	Tags    []string
}

type ListGroup struct {
	Groups []basic_group_payload_structs.GroupBaseStruct
}

func (s *ListGroup) Print() error {

	fmt.Printf("(%d) group available\n", (len(s.Groups) + 1))
	resources := []ViewModel{}
	for _, e := range s.Groups {
		resources = append(resources, ViewModel{
			Name:    e.Header.Name,
			Tags:    e.Header.Metadata.Tags,
			Path:    e.Specifications.Path,
			Package: e.Specifications.Package,
		})
	}

	format := printerx.Table
	printer, err := printerx.GetPrinter(format)
	if err != nil {
		panic(err)
	}

	err = printer.Print(os.Stdout, resources)
	if err != nil {
		panic(err)
	}

	return nil
}
