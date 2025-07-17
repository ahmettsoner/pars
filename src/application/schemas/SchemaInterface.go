package schemas

type SchemaInterface interface {
	Validate() error
	// PrintInfo()
	GetHeader() SchemaHeader
	GetSpecification() any
	GetKey() string
}
