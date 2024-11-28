package pagination

type Page[T any] struct {
	Items    []T      `json:"items,omitempty"`
	PageInfo PageInfo `json:"pageInfo,omitempty"`
}

type PageInfo struct {
	HasNext    bool `json:"hasNext"`
	TotalItems int  `json:"totalItems"`
}
