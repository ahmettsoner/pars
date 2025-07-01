package engines

import (
	"parsdevkit.net/application"
	"parsdevkit.net/core/schemas"
)

type Engine interface {
	// Init(ctx *application.ApplicationContext) error
	Validate(data []schemas.Schema) bool
	Process(ctx *application.ApplicationContext, data []schemas.Schema) error
	Destroy(ctx *application.ApplicationContext, data []schemas.Schema) error
}
