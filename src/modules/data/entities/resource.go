package entities

type Resource struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"unique;not null"`
	Document string `gorm:"type:json"`
	Hash     string `gorm:"not null"`
	Version  int    `gorm:"not null"`
}

func NewResource(id int, name string, document string, hash string, version int) *Resource {
	return &Resource{
		ID:       id,
		Name:     name,
		Document: document,
		Hash:     hash,
		Version:  version,
	}
}
