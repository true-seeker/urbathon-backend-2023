package response

type Paged struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
	Total    int   `json:"total"`
}
