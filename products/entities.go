package products

import "github.com/dimiro1/graphql-api/pagination"

// Product represents a Product
type Product struct {
	ID    int64
	Name  string
	Price float64
}

// PagedProducts represents a single page, it contains Metadata metadata and the Products results
type PagedProducts struct {
	pagination.Metadata
	Results []Product
}
