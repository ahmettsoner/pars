package events

import (
	"parsdevkit.net/application/schemas"
)

type TemplateCreated struct {
	Data schemas.SchemaInterface
}
