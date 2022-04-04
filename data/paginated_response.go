package data

import "net/url"

type PaginatedResponse[T any] struct {
	Metadata *PaginationMetaData
	Links    *PaginationLinks
	Items    []*T
}

type PaginationMetaData struct {
	Total     *int
	FirstPage *int
	PrevPage  *int
	LastPage  *int
	NextPage  *int
}

type PaginationLinks struct {
	First *url.URL
	Prev  *url.URL
	Next  *url.URL
	Last  *url.URL
}
