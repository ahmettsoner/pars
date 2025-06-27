package entities

type Group struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null;uniqueIndex:idx_name,unique"`
	Document string `gorm:"type:json"`
	Hash     string `gorm:"not null"`
	Version  int    `gorm:"not null"`
}

func NewGroup(id int, name string, document string, hash string, version int) Group {
	return Group{
		ID:       id,
		Name:     name,
		Document: document,
		Hash:     hash,
		Version:  version,
	}
}
