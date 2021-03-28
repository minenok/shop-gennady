package availability

type AvailabilityOption struct {
	WarehouseID      uint   `json:"warehouse_id"`
	WarehouseName    string `json:"warehouse_name"`
	WarehouseAddress string `json:"warehouse_address"`
	Quantity         uint   `json:"quantity"`
}
