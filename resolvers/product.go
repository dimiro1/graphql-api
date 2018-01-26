package resolvers

import (
	"strconv"

	"github.com/dimiro1/graphql-api/products"
	"github.com/neelance/graphql-go"
)

type productResolver struct {
	product products.Product
}

func (r *productResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(r.product.ID, 10))
}

func (r *productResolver) Name() string {
	return r.product.Name
}

func (r *productResolver) Price() *float64 {
	if r.product.Price == 0 {
		return nil
	}
	return &r.product.Price
}
