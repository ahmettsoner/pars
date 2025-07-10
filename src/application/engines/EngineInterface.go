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
	GetConfig() EngineConfig
}

type ProcessorInterface interface {
	Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type DestroyerInterface interface {
	Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectRunnerEngineInterface interface {
	Run(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectCleanerEngineInterface interface {
	Clean(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}

type ApplicationProjectExecuterEngineInterface interface {
	Execute(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectReleaserEngineInterface interface {
	Release(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ListEngineInterface interface {
	List(ctx *application.ApplicationContext) error
}
type ToolsEngineInterface interface {
	Browse(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
