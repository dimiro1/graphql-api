package products

import (
	"context"
	"fmt"

	"github.com/dimiro1/graphql-api/pagination"
	"github.com/dimiro1/graphql-api/search"
	"github.com/jmoiron/sqlx"
)

// DatabaseInventoryRepository is a InventoryRepository baked by a SQL database
type DatabaseInventoryRepository struct {
	DB *sqlx.DB
}

// All must return all from the database using a SQL database
func (d *DatabaseInventoryRepository) All(ctx context.Context, options pagination.ForwardBackwardCursor) (PagedProducts, error) {
	var products []Product
	var paged PagedProducts

	var direction = ">"
	var limit = options.First
	var offset = options.After

	if options.Before > 0 {
		direction = "<"
		limit = options.Last
		offset = options.Before
	}

	err := d.DB.SelectContext(ctx,
		&products,
		fmt.Sprintf("SELECT * FROM Products WHERE id %s ? LIMIT ?", direction), offset, limit+1)
	if err != nil {
		return paged, err
	}
	if len(products) > limit {
		paged.HasNext = true
	}

	if options.Before > 0 {
		paged.HasPrevious = true
	}
	paged.Results = products[:len(products)-1]
	paged.Total = uint64(len(paged.Results))

	return paged, err
}

// Search products using a SQL Database
func (d *DatabaseInventoryRepository) Search(ctx context.Context, options search.Cursor) (PagedProducts, error) {
	var products []Product
	var paged PagedProducts

	err := d.DB.SelectContext(ctx, &products, "SELECT * FROM Products WHERE id > ? AND name LIKE  '%' || ? || '%' LIMIT ?",
		options.After,
		options.Q,
		options.First+1, // Hack to check if we have the next page
	)
	if len(products) > options.First {
		paged.HasNext = true
	}

	paged.Results = products[:len(products)-1]
	paged.Total = uint64(len(paged.Results))

	return paged, err
}

// FindByID, returns the product by its id
func (d *DatabaseInventoryRepository) FindByID(id int64) (Product, error) {
	var product Product
	err := d.DB.Get(&product, "SELECT * FROM Products WHERE id = ?", id)
	if err != nil {
		return product, ProductNotFound{id}
	}

	return product, nil
}
