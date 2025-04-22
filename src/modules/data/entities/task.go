package entities

type Task struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"unique;not null"`
	Document string `gorm:"type:json"`
	Hash     string `gorm:"not null"`
	Version  int    `gorm:"not null"`
}

func NewTask(id int, name string, document string, hash string, version int) *Task {
	return &Task{
		ID:       id,
		Name:     name,
		Document: document,
		Hash:     hash,
		Version:  version,
	}
}
