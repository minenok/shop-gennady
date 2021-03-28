package rest

import (
	"github.com/gorilla/mux"
	"github.com/minenok/shop-gennady/internal/api/rest/availability"
	"github.com/minenok/shop-gennady/internal/api/rest/item"
	"github.com/minenok/shop-gennady/internal/api/rest/list"
	"github.com/minenok/shop-gennady/internal/repository"
)

type API struct {
	repoProducts     *repository.Products
	repoPrices       *repository.Prices
	repoAvailability *repository.Availability
}

func NewAPI(repoProducts *repository.Products, repoPrices *repository.Prices, repoAv *repository.Availability) *API {
	return &API{repoProducts: repoProducts, repoPrices: repoPrices, repoAvailability: repoAv}
}

func (a *API) Bind(r *mux.Router) {
	r.Handle("/items", list.NewController(a.repoProducts, a.repoPrices))
	r.Handle("/items/{id}", item.NewController(a.repoProducts, a.repoPrices))
	r.Handle("/items/{id}/availability", availability.NewController(a.repoAvailability))
}
