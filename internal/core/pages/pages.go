package pages

type PagedResult[PageItem any] struct {
	Items   []PageItem
	Total   int
	HasNext bool
}

type PageOptions struct {
	Page, PageSize int
}
