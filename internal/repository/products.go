package repository

import (
	"errors"
	"github.com/minenok/shop-gennady/internal/model"
	"sync"
)

type Products struct {
	prs map[uint]model.Product
	mu  sync.RWMutex
}

func NewProducts() *Products {
	return &Products{
		prs: map[uint]model.Product{
			1: {
				ID:            1,
				Name:          "Мочалка",
				Description:   "Трёт",
				RawProperties: map[string]string{"colour": "бежевый", "brand": "чисто-чисто"},
			},
			2: {
				ID:            1,
				Name:          "Колбаса",
				Description:   "Вкусная",
				RawProperties: map[string]string{"colour": "розовый", "brand": "докторская"},
			},
			3: {
				ID:            1,
				Name:          "Чапельник",
				Description:   "Вставляется",
				RawProperties: map[string]string{"colour": "черный", "brand": "noname"},
			},
			4: {
				ID:            1,
				Name:          "Машинка",
				Description:   "Ездит",
				RawProperties: map[string]string{"colour": "желтый", "brand": "малыш"},
			},
			5: {
				ID:            1,
				Name:          "Шкаф",
				Description:   "Стоит",
				RawProperties: map[string]string{"colour": "коричневый", "brand": "известный"},
			},
		},
	}
}

func (r *Products) FindProduct(id uint) (model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.prs[id]
	if !ok {
		return model.Product{}, errors.New("not found")
	}
	return p, nil
}

func (r *Products) FindProducts() ([]model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ret := make([]model.Product, 0, len(r.prs))
	for _, p := range r.prs {
		ret = append(ret, p)
	}
	return ret, nil
}
