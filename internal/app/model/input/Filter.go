package input

type Filter struct {
	Page     int64 `form:"page"`
	PageSize int64 `form:"page_size"`
}
