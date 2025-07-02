package schemas

type SchemaInterface interface {
	Validate() error
	// PrintInfo()
	GetHeader() SchemaHeader
	GetKey() string
}
