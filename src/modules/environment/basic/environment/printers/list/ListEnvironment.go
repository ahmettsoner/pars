package list

import (
	"fmt"
	"os"

	"parsdevkit.net/internal/printerx"
	basic_environment_payload_structs "parsdevkit.net/modules/environment/basic_environment_payload/structs"
)

type ViewModel struct {
	Name string
}

type ListEnvironment struct {
	Environments []basic_environment_payload_structs.EnvironmentBaseStruct
}

func (s *ListEnvironment) Print() error {

	fmt.Printf("(%d) environment available\n", (len(s.Environments) + 1))
	resources := []ViewModel{
		{"* Default"},
	}
	for _, e := range s.Environments {
		resources = append(resources, ViewModel{
			Name: e.Header.Name,
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
