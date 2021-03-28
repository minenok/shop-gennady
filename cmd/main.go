package main

import (
	"github.com/gorilla/mux"
	"github.com/minenok/shop-gennady/internal/api/graphql"
	"github.com/minenok/shop-gennady/internal/api/rest"
	"github.com/minenok/shop-gennady/internal/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	productsRepo := repository.NewProducts()
	priceRepo := repository.NewPrices()
	availabilityRepo := repository.NewAvailability()

	r := mux.NewRouter()
	rest.NewAPI(productsRepo, priceRepo, availabilityRepo).Bind(r)

	ga, err := graphql.NewAPI(productsRepo, priceRepo, availabilityRepo)
	if err != nil {
		log.Fatal(err)
	}
	ga.Bind(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
