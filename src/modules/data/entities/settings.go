package entities

type Settings struct {
	Key   string `gorm:"primaryKey;unique"`
	Value string
}

func NewSettings(key, value string) *Settings {
	return &Settings{
		Key:   key,
		Value: value,
	}
}
