package schemas

type Schema interface {
	Validate() error
	// PrintInfo()
	GetHeader() SchemaHeader
}
