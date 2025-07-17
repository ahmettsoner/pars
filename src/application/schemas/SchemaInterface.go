package schemas

type SchemaInterface interface {
	Validate() error
	// PrintInfo()
	GetHeader() SchemaHeader
	// GetSpecification() SchemaSpecification
	GetKey() string
}
