package contracts

import "parsdevkit.net/application/schemas"

type EnvironmentServiceInterface[T schemas.SchemaInterface] interface {
	List() ([]T, error)
}
