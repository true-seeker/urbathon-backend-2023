package filter

import (
	"urbathon-backend-2023/pkg/config"
	"urbathon-backend-2023/pkg/errorHandler"
)

type Pagination struct {
	Page     int64 `form:"page"`
	PageSize int64 `form:"page_size"`
}

type PaginationInterface interface {
	GetPagination() *Pagination
}

func ValidatePagination(p *Pagination) (*Pagination, *errorHandler.HttpErr) {
	if p.Page <= 0 {
		p.Page = config.DefaultPage
	}
	if p.PageSize <= 0 || p.PageSize > config.MaxPageSize {
		p.PageSize = config.PageSize
	}
	return p, nil
}
