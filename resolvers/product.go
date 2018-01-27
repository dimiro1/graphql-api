package resolvers

import (
	"strconv"

	"github.com/dimiro1/graphql-api/products"
	"github.com/neelance/graphql-go"
)

type ProductResolver struct {
	product products.Product
}

func (r *ProductResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(r.product.ID, 10))
}

func (r *ProductResolver) Name() string {
	return r.product.Name
}

func (r *ProductResolver) Price() *float64 {
	if r.product.Price == 0 {
		return nil
	}
	return &r.product.Price
}

type ProductConnectionResolver struct {
	pagedProducts products.PagedProducts
}

func (p ProductConnectionResolver) TotalCount() int32 {
	return int32(p.pagedProducts.Total)
}

func (p ProductConnectionResolver) Edges() []*ProductsEdgeResolver {
	var edges []*ProductsEdgeResolver

	for _, p := range p.pagedProducts.Results {
		edges = append(edges, &ProductsEdgeResolver{p})
	}

	return edges
}

func (p ProductConnectionResolver) PageInfo() PageInfoResolver {
	return PageInfoResolver{
		HasPrevious: p.pagedProducts.HasPrevious,
		HasNext:     p.pagedProducts.HasNext,
	}
}

type ProductsEdgeResolver struct {
	product products.Product
}

func (p ProductsEdgeResolver) Cursor() string {
	return strconv.FormatInt(p.product.ID, 10)
}

func (p ProductsEdgeResolver) Node() *ProductResolver {
	return &ProductResolver{p.product}
}

type PageInfoResolver struct {
	HasPrevious bool
	HasNext     bool
}

func (p PageInfoResolver) HasPreviousPage() bool {
	return p.HasPrevious
}

func (p PageInfoResolver) HasNextPage() bool {
	return p.HasNext
}
