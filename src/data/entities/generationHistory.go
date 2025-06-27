package entities

import "time"

type GenerationHistory struct {
	ID           int       `gorm:"primaryKey"`
	Set          string    `gorm:"not null;"`
	Resource     string    `gorm:"not null"`
	ResourceHash string    `gorm:"not null"`
	Template     string    `gorm:"not null"`
	TemplateHash string    `gorm:"not null"`
	Section      string    `gorm:"not null"`
	SectionHash  string    `gorm:"not null"`
	Layer        string    `gorm:"not null"`
	Timestamp    time.Time `gorm:"not null"`
}

func NewGenerationHistory(set, resource, resourceHash, template, templateHash, section, sectionHash, layer string) *GenerationHistory {
	return &GenerationHistory{
		Set:          set,
		Resource:     resource,
		ResourceHash: resourceHash,
		Template:     template,
		TemplateHash: templateHash,
		Section:      section,
		SectionHash:  sectionHash,
		Layer:        layer,
	}
}
