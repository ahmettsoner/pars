package contracts

type WorkspaceServiceInterface[T SchemaInterface] interface {
	GetByName(name string) (*T, error)
	Save(model T) (*T, error)
	List() (*([]T), error)
	Remove(name string, force, permanent bool) (*T, error)
	IsExists(name string) (bool, error)
	GetHash(name string) (string, error)
	GetSelectedWorkspace() (*T, error)
	GetActiveWorkspace() (*T, error)
	ChangeCurrentWorkspace(name string) (*T, error)
	ListByNameStartWith(name string) (*([]T), error)
}
