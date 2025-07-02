package engines

import (
	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"
)

type EngineConfig struct {
	Name  string
	Order int
}
type EngineInterface interface {
	// Init(ctx *application.ApplicationContext) error
	Validate(data []schemas.SchemaInterface) bool
	Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
	Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
	GetConfig() EngineConfig
}
