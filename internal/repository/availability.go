package repository

import (
	"errors"
	"github.com/minenok/shop-gennady/internal/model"
	"sync"
)

type Availability struct {
	whs map[uint]model.Warehouse
	av  map[uint][]model.AvailabilityOption
	mu  sync.RWMutex
}

func NewAvailability() *Availability {
	return &Availability{
		whs: map[uint]model.Warehouse{
			1: {ID: 1, Name: "Маленький склад", Address: "Огородная, дом 8"},
			2: {ID: 2, Name: "Большой склад", Address: "Коровинское шоссе, дом 8712"},
		},
		av: map[uint][]model.AvailabilityOption{
			1: {{WarehouseID: 2, ProductID: 1, Quantity: 2}, {WarehouseID: 1, ProductID: 1, Quantity: 1}},
			2: {{WarehouseID: 2, ProductID: 2, Quantity: 3}, {WarehouseID: 1, ProductID: 2, Quantity: 2}},
			3: {{WarehouseID: 2, ProductID: 3, Quantity: 4}},
			4: {{WarehouseID: 2, ProductID: 4, Quantity: 5}},
			5: {},
		},
	}
}

func (r *Availability) AvailabilityOptions(productID uint) ([]model.AvailabilityOption, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ret, ok := r.av[productID]
	if !ok {
		return nil, errors.New("not found")
	}
	return ret, nil
}

func (r *Availability) WarehouseByID(whID uint) (model.Warehouse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ret, ok := r.whs[whID]
	if !ok {
		return model.Warehouse{}, errors.New("not found")
	}
	return ret, nil
}
