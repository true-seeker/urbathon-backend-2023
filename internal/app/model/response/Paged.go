package response

import "fmt"

type Paged struct {
	Page        int64   `json:"page" example:"1"`
	PageSize    int64   `json:"page_size" example:"10"`
	Total       int     `json:"total" example:"100"`
	NextPageUrl *string `json:"next_page_url" example:"/api/news?page=1&page_size=10"`
}

func NewPaged(page int64, pageSize int64, total int, resource string) *Paged {
	var nextPageUrlP *string
	fmt.Println(resource, page+1, pageSize, nextPageUrlP)
	if pageSize*page < int64(total) {
		temp := fmt.Sprintf("/api/%s?page=%d&page_size=%d",
			resource,
			page+1,
			pageSize)
		nextPageUrlP = &temp
	}

	return &Paged{Page: page,
		PageSize:    pageSize,
		Total:       total,
		NextPageUrl: nextPageUrlP}
}
