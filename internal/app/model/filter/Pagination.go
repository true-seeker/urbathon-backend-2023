package filter

import (
	"github.com/go-jet/jet/v2/postgres"
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
	if p.Page == -1 {
		p.PageSize = -1
		return p, nil
	}
	if p.Page <= 0 {
		p.Page = config.DefaultPage
	}
	if p.PageSize <= 0 || p.PageSize > config.MaxPageSize {
		p.PageSize = config.PageSize
	}
	return p, nil
}

func (p *Pagination) GetLimitOffsetStmt(stmt postgres.SelectStatement) postgres.SelectStatement {
	if p.Page != -1 {
		stmt = stmt.LIMIT(p.PageSize).
			OFFSET((p.Page - 1) * p.PageSize)
	}
	return stmt
}
