package contracts

type ContextProviderConfig struct {
	Name string
}
type BaseContextProviderInterface interface {
	GetConfig() ContextProviderConfig
}
