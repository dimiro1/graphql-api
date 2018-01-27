package search

import "github.com/dimiro1/graphql-api/pagination"

// ForwardBackwardCursor is used to pass Search and Pagination parameters for repositories
type Cursor struct {
	pagination.ForwardCursor
	Q string
}
