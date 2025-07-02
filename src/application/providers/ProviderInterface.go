package providers

type ProviderInterface interface {
	Run(args []string) (string error)
	GetKey() string
}
