package products

import (
	"context"
	"fmt"

	"github.com/dimiro1/graphql-api/pagination"
	"github.com/dimiro1/graphql-api/search"
)

type ProductNotFound struct {
	ID int64
}

func (p ProductNotFound) Error() string {
	return fmt.Sprintf("Could not find product with id: %f", p.ID)
}

// InventoryRepository is an repository of inventory
type InventoryRepository interface {
	// All must return all from the database
	All(context.Context, pagination.ForwardBackwardCursor) (PagedProducts, error)

	// Search, search products
	Search(context.Context, search.Cursor) (PagedProducts, error)

	// FindByID, returns the product by its id
	FindByID(int64) (Product, error)
}
