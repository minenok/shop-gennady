package item

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/minenok/shop-gennady/internal/model"
	"net/http"
	"strconv"
)

type Controller struct {
	repo      Repository
	priceRepo PriceRepository
}

type Repository interface {
	FindProduct(id uint) (model.Product, error)
}

type PriceRepository interface {
	CurrentPrice([]model.Product) (map[uint]uint, error)
}

func NewController(repo Repository, priceRepo PriceRepository) *Controller {
	return &Controller{repo: repo, priceRepo: priceRepo}
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

	p, err := c.repo.FindProduct(uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "product not found"}`))
		return
	}
	prices, err := c.priceRepo.CurrentPrice([]model.Product{p})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "price fetching error"}`))
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(map[string]interface{}{"product": productFromModel(p, prices[p.ID])}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "marshalling error"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}
