package schemas

type SchemaHeader struct {
	Type StructType `yaml:"Type"`
	Kind string     `yaml:"Kind"`
	Name string     `yaml:"Name"`
}
