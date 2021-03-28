package list

import (
	"encoding/json"
	"github.com/minenok/shop-gennady/internal/model"
	"net/http"
)

type Controller struct {
	repo      Repository
	priceRepo PriceRepository
}

type Repository interface {
	FindProducts() ([]model.Product, error)
}

type PriceRepository interface {
	CurrentPrice([]model.Product) (map[uint]uint, error)
}

func NewController(r Repository, pr PriceRepository) *Controller {
	return &Controller{repo: r, priceRepo: pr}
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	products, err := c.repo.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "storage error"}`))
		return
	}
	prices, err := c.priceRepo.CurrentPrice(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "storage error"}`))
		return
	}

	resp := make([]Product, 0, len(products))
	for _, p := range products {
		resp = append(resp, productFromModel(p, prices[p.ID]))
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(map[string]interface{}{"products": resp}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "marshalling error"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}
