package graphql

import (
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type API struct {
	schema graphql.Schema
}

func NewAPI(productRepo ProductRepository, priceRepo PriceRepository, availabilityRepo AvailabilityRepository) (*API, error) {
	bld := newSchemaBuilder(productRepo, priceRepo, availabilityRepo)
	schema, err := bld.Build()
	if err != nil {
		return nil, err
	}
	return &API{schema: schema}, nil
}

func (a *API) Bind(r *mux.Router) {
	h := handler.New(&handler.Config{
		Schema:   &a.schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Handle("/graphql", h)
}
