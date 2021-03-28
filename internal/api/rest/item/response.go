package item

import "github.com/minenok/shop-gennady/internal/model"

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Colour      string `json:"colour"`
	Brand       string `json:"brand"`
	Price       uint   `json:"price"`
}

func productFromModel(p model.Product, price uint) Product {
	return Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Colour:      p.RawProperties["colour"],
		Brand:       p.RawProperties["brand"],
		Price:       price,
	}
}
