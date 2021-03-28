package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/minenok/shop-gennady/internal/model"
)

func (a *API) init() error {
	warehouseOptionType := graphql.NewObject(graphql.ObjectConfig{
		Name: "warehouse",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "ID склада",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if w, ok := p.Source.(model.Warehouse); ok {
						return w.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Название склада",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if w, ok := p.Source.(model.Warehouse); ok {
						return w.Name, nil
					}
					return nil, nil
				},
			},
			"address": &graphql.Field{
				Type:        graphql.String,
				Description: "Адрес склада",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if w, ok := p.Source.(model.Warehouse); ok {
						return w.Address, nil
					}
					return nil, nil
				},
			},
		},
	})

	availabilityOptionType := graphql.NewObject(graphql.ObjectConfig{
		Name: "availability_option",
		Fields: graphql.Fields{
			"warehouse": &graphql.Field{
				Type:        warehouseOptionType,
				Description: "Склад",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ao, ok := p.Source.(model.AvailabilityOption); ok {
						wh, err := a.availabilityRepo.WarehouseByID(ao.WarehouseID)
						if err != nil {
							return nil, err
						}
						return wh, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.Int,
				Description: "Количество товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ao, ok := p.Source.(model.AvailabilityOption); ok {
						return ao.Quantity, nil
					}
					return nil, nil
				},
			},
		},
	})

	productType := graphql.NewObject(graphql.ObjectConfig{
		Name: "product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "ID товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if pr, ok := p.Source.(model.Product); ok {
						return pr.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Название товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if pr, ok := p.Source.(model.Product); ok {
						return pr.Name, nil
					}
					return nil, nil
				},
			},
			"description": &graphql.Field{
				Type:        graphql.String,
				Description: "Описание товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if pr, ok := p.Source.(model.Product); ok {
						return pr.Description, nil
					}
					return nil, nil
				},
			},
			"colour": &graphql.Field{
				Type:        graphql.String,
				Description: "Цвет товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if pr, ok := p.Source.(model.Product); ok {
						return pr.RawProperties["colour"], nil
					}
					return nil, nil
				},
			},
			"brand": &graphql.Field{
				Type:        graphql.String,
				Description: "Бренд товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if pr, ok := p.Source.(model.Product); ok {
						return pr.RawProperties["brand"], nil
					}
					return nil, nil
				},
			},
			"price": &graphql.Field{
				Type:        graphql.Int,
				Description: "цена товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if pr, ok := p.Source.(model.Product); ok {
						prices, err := a.priceRepo.CurrentPrice([]model.Product{pr})
						if err != nil {
							return nil, err
						}
						return prices[pr.ID], nil
					}
					return nil, nil
				},
			},
			"availability": &graphql.Field{
				Type:        graphql.NewList(availabilityOptionType),
				Description: "наличие товара",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if pr, ok := p.Source.(model.Product); ok {
						av, err := a.availabilityRepo.AvailabilityOptions(pr.ID)
						if err != nil {
							return nil, err
						}
						return av, nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"products": &graphql.Field{
				Type: graphql.NewList(productType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id товара",
						Type:        graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return a.productRepo.FindProducts()
					}
					pr, err := a.productRepo.FindProduct(uint(id))
					if err != nil {
						return nil, err
					}
					return []model.Product{pr}, nil
				},
			},
		},
	})
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		return err
	}
	a.schema = s
	return nil
}
