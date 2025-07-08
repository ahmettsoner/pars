package events

import (
	"parsdevkit.net/application/schemas"
)

type ResourceCreated struct {
	Data schemas.SchemaInterface
}
