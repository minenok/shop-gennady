package model

type Warehouse struct {
	ID      uint   `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	Address string `gorm:"column:address"`
}
