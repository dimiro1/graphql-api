package search

import "github.com/dimiro1/graphql-api/pagination"

// Options is used to pass Search and Pagination parameters for repositories
type Options struct {
	pagination.Options
	Q string
}
