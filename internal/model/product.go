package model

type Product struct {
	ID            uint              `gorm:"column:id"`
	Name          string            `gorm:"column:name"`
	Description   string            `gorm:"column:description"`
	RawProperties map[string]string `gorm:"column:raw_properties"`
}
