package main

import (
	"github.com/gorilla/mux"
	"github.com/minenok/shop-gennady/internal/api/rest"
	"github.com/minenok/shop-gennady/internal/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	rest.NewAPI(repository.NewProducts(), repository.NewPrices(), repository.NewAvailability()).Bind(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
