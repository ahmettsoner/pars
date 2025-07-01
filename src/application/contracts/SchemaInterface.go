package contracts

import "parsdevkit.net/application/schemas"

type SchemaInterface interface {
	Validate() error
	// PrintInfo()
	GetHeader() schemas.SchemaHeader
}
