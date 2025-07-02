package contracts

import "parsdevkit.net/application/schemas"

type GroupServiceInterface[T schemas.SchemaInterface] interface {
	GetByName(name string) (*T, error)
	Save(model T) (*T, error)
	List() (*([]T), error)
	Remove(name string, permanent bool) (*T, error)
	IsExists(name string) (bool, error)
	GetHash(name string) (string, error)
}
