package pagination

// Metadata is a struct that represents a Metadata in both database and HTTP APIs
type Metadata struct {
	HasNext     bool   `json:"has_next"`
	HasPrevious bool   `json:"has_previous"`
	Total       uint64 `json:"total"`
}

// Options is used to pass Pagination parameters for repositories
type Options struct {
	Limit  int
	Offset int
}

type ForwardCursor struct {
	First int
	After int64
}

type BackwardCursor struct {
	Last   int
	Before int64
}

type ForwardBackwardCursor struct {
	ForwardCursor
	BackwardCursor
}
