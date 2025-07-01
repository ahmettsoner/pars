package contracts

type ResourceServiceInterface[T SchemaInterface] interface {
	GetByName(name string) (*T, error)
	Save(model T) (*T, error)
	List() (*([]T), error)
	ListByWorkspace(workspace string) (*([]T), error)
	ListBySet(set string) (*([]T), error)
	ListByWorkspaceAndSet(workspace, set string) (*([]T), error)
	ListBySetAndLayers(set string, layers ...string) (*([]T), error)
	ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]T), error)
	Remove(name, workspace string, force, permanent bool) (*T, error)
	IsExists(name, workspace string) (bool, error)
	GetHash(name string) (string, error)
}
