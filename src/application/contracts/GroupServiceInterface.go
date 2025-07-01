package contracts

type GroupServiceInterface[T SchemaInterface] interface {
	GetByName(name string) (*T, error)
	Save(model T) (*T, error)
	List() (*([]T), error)
	Remove(name string, permanent bool) (*T, error)
	IsExists(name string) (bool, error)
	GetHash(name string) (string, error)
}
