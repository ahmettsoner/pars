package entities

type Project struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null;uniqueIndex:idx_workspace_name_group,unique"`
	Document string `gorm:"type:json"`
	Hash     string `gorm:"not null"`
	Version  int    `gorm:"not null"`
}

func NewProjectDocument(id int, name string, document string, hash string, version int) *Project {
	return &Project{
		ID:       id,
		Name:     name,
		Document: document,
		Hash:     hash,
		Version:  version,
	}
}
