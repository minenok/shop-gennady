package graphql

import (
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/minenok/shop-gennady/internal/model"
)

type API struct {
	productRepo      ProductRepository
	priceRepo        PriceRepository
	availabilityRepo AvailabilityRepository
	schema           graphql.Schema
}

type ProductRepository interface {
	FindProducts() ([]model.Product, error)
	FindProduct(id uint) (model.Product, error)
}

type PriceRepository interface {
	CurrentPrice([]model.Product) (map[uint]uint, error)
}

type AvailabilityRepository interface {
	AvailabilityOptions(productID uint) ([]model.AvailabilityOption, error)
	WarehouseByID(uint) (model.Warehouse, error)
}

func NewAPI(productRepo ProductRepository, priceRepo PriceRepository, availabilityRepo AvailabilityRepository) (*API, error) {
	a := &API{productRepo: productRepo, priceRepo: priceRepo, availabilityRepo: availabilityRepo}
	if err := a.init(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *API) Bind(r *mux.Router) {
	h := handler.New(&handler.Config{
		Schema:   &a.schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Handle("/graphql", h)
}
