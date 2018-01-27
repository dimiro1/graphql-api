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
	First  *int32
	After  *graphql.ID
	Last   *int32
	Before *graphql.ID
}

func (r *Resolver) Products(ctx context.Context, args ProductsQueryArgs) (ProductConnectionResolver, error) {
	productConnectionResolver := ProductConnectionResolver{
		pagedProducts: products.PagedProducts{},
	}

	var (
		first  int32 = 10
		after  int64 = 0
		last   int32 = 10
		before int64 = 0
	)

	if args.After != nil {
		after, _ = strconv.ParseInt(string(*args.After), 10, 64)
	}
	if args.First != nil {
		first = *args.First
	}
	if args.Before != nil {
		after, _ = strconv.ParseInt(string(*args.Before), 10, 64)
	}
	if args.Last != nil {
		last = *args.Last
	}

	var err error
	productConnectionResolver.pagedProducts, err = r.InventoryRepository.All(ctx, pagination.ForwardBackwardCursor{
		ForwardCursor: pagination.ForwardCursor{
			First: int(first),
			After: after,
		},
		BackwardCursor: pagination.BackwardCursor{
			Last:   int(last),
			Before: before,
		},
	})
	if err != nil {
		return productConnectionResolver, err
	}

	return productConnectionResolver, nil
}

type SearchQueryArgs struct {
	Q     string
	First *int32
	After *graphql.ID
}

func (r *Resolver) Search(ctx context.Context, args SearchQueryArgs) (ProductConnectionResolver, error) {
	productConnectionResolver := ProductConnectionResolver{
		pagedProducts: products.PagedProducts{},
	}

	var (
		first int32 = 10
		after int64 = 0
	)

	if args.After != nil {
		after, _ = strconv.ParseInt(string(*args.After), 10, 64)
	}
	if args.First != nil {
		first = *args.First
	}

	var err error
	productConnectionResolver.pagedProducts, err = r.InventoryRepository.Search(ctx, search.Cursor{
		ForwardCursor: pagination.ForwardCursor{First: int(first), After: after},
		Q:             args.Q,
	})
	if err != nil {
		return productConnectionResolver, err
	}

	return productConnectionResolver, nil
}

type ProductQueryArgs struct {
	ID graphql.ID
}

func (r *Resolver) Product(ctx context.Context, args ProductQueryArgs) *ProductResolver {
	id, err := strconv.ParseInt(string(args.ID), 10, 2)
	if err != nil {
		return nil
	}

	product, err := r.InventoryRepository.FindByID(id)
	if err != nil {
		return nil
	}

	return &ProductResolver{product}
}
