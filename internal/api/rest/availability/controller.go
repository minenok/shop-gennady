package availability

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/minenok/shop-gennady/internal/model"
	"net/http"
	"strconv"
)

type Controller struct {
	repo Repository
}

type Repository interface {
	AvailabilityOptions(productID uint) ([]model.AvailabilityOption, error)
	WarehouseByID(uint) (model.Warehouse, error)
}

func NewController(repo Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, ok := v["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "missing product id"}`))
		return
	}
	productID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "missing product id"}`))
		return
	}

	options, err := c.repo.AvailabilityOptions(uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "availability fetching error"}`))
		return
	}

	resp := make([]AvailabilityOption, 0, len(options))
	warehouses := make(map[uint]model.Warehouse)
	for _, opt := range options {
		if _, ok := warehouses[opt.WarehouseID]; !ok {
			wh, err := c.repo.WarehouseByID(opt.WarehouseID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "availability fetching error"}`))
				return
			}
			warehouses[opt.WarehouseID] = wh
		}
		resp = append(resp, AvailabilityOption{
			WarehouseID:      opt.WarehouseID,
			WarehouseName:    warehouses[opt.WarehouseID].Name,
			WarehouseAddress: warehouses[opt.WarehouseID].Address,
			Quantity:         opt.Quantity,
		})
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(map[string]interface{}{"products": resp}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "marshalling error"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}
