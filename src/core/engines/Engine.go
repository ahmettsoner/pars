package engines

import (
	"parsdevkit.net/core"
	"parsdevkit.net/core/schemas"
)

type Engine interface {
	Validate(data []schemas.Schema) bool
	Process(ctx *core.ApplicationContext, data []schemas.Schema) error
	Destroy(ctx *core.ApplicationContext, data []schemas.Schema) error
}
