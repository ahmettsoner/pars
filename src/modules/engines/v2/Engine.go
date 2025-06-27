package v2

import "parsdevkit.net/structs"

type Engine interface {
	Validate(data []structs.Schema) bool
	Process(data []structs.Schema) error
	Destroy(data []structs.Schema) error
}
