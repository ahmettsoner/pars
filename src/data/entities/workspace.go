package entities

type Workspace struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null;uniqueIndex:idx_workspace_name,unique"`
	Document string `gorm:"type:json"`
	Hash     string `gorm:"not null"`
	Version  int    `gorm:"not null"`
}

func NewWorkspace(id int, name string, document string, hash string, version int) Workspace {
	return Workspace{
		ID:       id,
		Name:     name,
		Document: document,
		Hash:     hash,
		Version:  version,
	}
}
