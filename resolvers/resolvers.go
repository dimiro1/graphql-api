package resolvers

import (
	"context"
	"strconv"

	"github.com/dimiro1/graphql-api/pagination"
	"github.com/dimiro1/graphql-api/products"
	"github.com/dimiro1/graphql-api/search"
	"github.com/neelance/graphql-go"
)

type Resolver struct {
	InventoryRepository products.InventoryRepository
}

type ProductsQueryArgs struct {
	First *int32
	After *int32
}

func (r *Resolver) Products(ctx context.Context, args ProductsQueryArgs) []*productResolver {
	var resolverList []*productResolver
	var first int32 = 10
	var after int32

	if args.First != nil {
		first = *args.First
	}
	if args.After != nil {
		after = *args.After
	}

	productList, err := r.InventoryRepository.All(pagination.Options{Limit: int(first), Offset: int(after)})
	if err != nil {
		return resolverList
	}

	for _, p := range productList {
		resolverList = append(resolverList, &productResolver{p})
	}
	return resolverList
}

type SearchQueryArgs struct {
	Q     string
	First *int32
	After *int32
}

func (r *Resolver) Search(args SearchQueryArgs) []*productResolver {
	var resolverList []*productResolver
	var first int32 = 10
	var after int32

	if args.First != nil {
		first = *args.First
	}
	if args.After != nil {
		after = *args.After
	}

	productList, err := r.InventoryRepository.Search(search.Options{
		Options: pagination.Options{Limit: int(first), Offset: int(after)},
		Q:       args.Q,
	})
	if err != nil {
		return resolverList
	}

	for _, p := range productList {
		resolverList = append(resolverList, &productResolver{p})
	}
	return resolverList
}

type ProductQueryArgs struct {
	ID graphql.ID
}

func (r *Resolver) Product(args ProductQueryArgs) *productResolver {
	id, err := strconv.ParseInt(string(args.ID), 10, 2)
	if err != nil {
		return nil
	}

	product, err := r.InventoryRepository.FindByID(id)
	if err != nil {
		return nil
	}

	return &productResolver{product}
}
