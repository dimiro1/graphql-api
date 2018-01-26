package products

import (
	"fmt"

	"github.com/dimiro1/graphql-api/pagination"
	"github.com/dimiro1/graphql-api/search"
	"github.com/jmoiron/sqlx"
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
	All(pagination.Options) ([]Product, error)

	// Search, search products
	Search(search.Options) ([]Product, error)

	// FindByID, returns the product by its id
	FindByID(int64) (Product, error)
}

// DatabaseInventoryRepository is a InventoryRepository baked by a SQL database
type DatabaseInventoryRepository struct {
	DB *sqlx.DB
}

// All must return all from the database using a SQL database
func (d *DatabaseInventoryRepository) All(options pagination.Options) ([]Product, error) {
	var products []Product

	err := d.DB.Select(&products, "SELECT * FROM Products WHERE id > ? LIMIT ?", options.Offset, options.Limit)

	return products, err
}

// Search products using a SQL Database
func (d *DatabaseInventoryRepository) Search(options search.Options) ([]Product, error) {
	var products []Product

	err := d.DB.Select(&products, "SELECT * FROM Products WHERE id > ? AND name LIKE  '%' || ? || '%' LIMIT ?",
		options.Offset,
		options.Q,
		options.Limit,
	)

	return products, err
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
