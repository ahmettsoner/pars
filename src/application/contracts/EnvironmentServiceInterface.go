package contracts

type EnvironmentServiceInterface interface {
	List() ([]string, error)
}
