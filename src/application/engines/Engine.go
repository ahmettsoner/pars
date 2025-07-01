package engines

import (
	"parsdevkit.net/application"
	"parsdevkit.net/application/contracts"
)

type Engine interface {
	// Init(ctx *application.ApplicationContext) error
	Validate(data []contracts.SchemaInterface) bool
	Process(ctx *application.ApplicationContext, data []contracts.SchemaInterface) error
	Destroy(ctx *application.ApplicationContext, data []contracts.SchemaInterface) error
}
