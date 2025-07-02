package platforms

type PlatformInterfaceE interface {
	Run(args []string) (string error)
	GetKey() string
}
