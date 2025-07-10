package contracts

import (
	"parsdevkit.net/application/schemas"
)

type ToolServiceInterface[T schemas.SchemaInterface] interface {
	Browse(model T) (*T, error)
}
