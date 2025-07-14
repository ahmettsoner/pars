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
type WorkspaceInitializerEngineInterface interface {
	Init(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectRunnerEngineInterface interface {
	Run(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectCleanerEngineInterface interface {
	Clean(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectInstallerEngineInterface interface {
	Install(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}

type ApplicationProjectRunnerrEngineInterface interface {
	Run(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}

type ApplicationProjectOpenerEngineInterface interface {
	Open(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectTesterEngineInterface interface {
	Test(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ApplicationProjectReleaserEngineInterface interface {
	Release(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
type ListEngineInterface interface {
	List(ctx *application.ApplicationContext) error
}

type DescribeEngineInterface interface {
	Describe(ctx *application.ApplicationContext, args ...any) error
}
type RemoveEngineInterface interface {
	Remove(ctx *application.ApplicationContext, args ...any) error
}
type ToolsEngineInterface interface {
	Browse(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error
}
