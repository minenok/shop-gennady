package repository

import (
	"errors"
	"fmt"
	"github.com/minenok/shop-gennady/internal/model"
	"sync"
)

type Prices struct {
	prices map[uint]uint
	mu     sync.RWMutex
}

func NewPrices() *Prices {
	return &Prices{
		prices: map[uint]uint{
			1: 100, 2: 200, 3: 300, 4: 400, 5: 450,
		},
	}
}

func (r *Prices) CurrentPrice(products []model.Product) (map[uint]uint, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ret := make(map[uint]uint)
	for _, pr := range products {
		price, ok := r.prices[pr.ID]
		if !ok {
			return nil, errors.New(fmt.Sprintf("price not found for product %d", pr.ID))
		}
		ret[pr.ID] = price
	}
	return ret, nil
}
