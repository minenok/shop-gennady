package model

type AvailabilityOption struct {
	WarehouseID uint `gorm:"column:warehouse_id"`
	ProductID   uint `gorm:"column:product_id"`
	Quantity    uint `gorm:"column:quantity"`
}
