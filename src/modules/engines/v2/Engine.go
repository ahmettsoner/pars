package v2

import (
	"parsdevkit.net/core"
	"parsdevkit.net/core/schemas"
)

type Engine interface {
	Validate(data []schemas.Schema) bool
	Process(ctx *core.Context, data []schemas.Schema) error
	Destroy(ctx *core.Context, data []schemas.Schema) error
}
