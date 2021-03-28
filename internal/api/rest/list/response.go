package list

import "github.com/minenok/shop-gennady/internal/model"

type Product struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Colour string `json:"colour"`
	Price  uint   `json:"price"`
}

func productFromModel(p model.Product, price uint) Product {
	return Product{
		ID:     p.ID,
		Name:   p.Name,
		Colour: p.RawProperties["colour"],
		Price:  price,
	}
}
