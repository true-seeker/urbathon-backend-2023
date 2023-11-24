package filter

import (
	"github.com/gin-gonic/gin"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealFilter struct {
	*Pagination
	*Sort

	UserId   *int32   `form:"user_id" `
	LatUp    *float64 `form:"lat_up" `
	LatDown  *float64 `form:"lat_down" `
	LongUp   *float64 `form:"long_up" `
	LongDown *float64 `form:"long_down" `
}

func NewAppealFilter(c *gin.Context) (*AppealFilter, *errorHandler.HttpErr) {
	var httpErr *errorHandler.HttpErr
	p := AppealFilter{
		Pagination: new(Pagination),
		Sort:       new(Sort),
	}

	_ = c.ShouldBindQuery(&p)

	p.Pagination, httpErr = ValidatePagination(p.Pagination)
	if httpErr != nil {
		return nil, httpErr
	}
	p.Sort, httpErr = ValidateSort(p.Sort, p.GetSortableFields())
	if httpErr != nil {
		return nil, httpErr
	}
	return &p, nil
}

func (a *AppealFilter) GetPagination() *Pagination {
	return a.Pagination
}

func (a *AppealFilter) GetSort() *Sort {
	return a.Sort
}

func (a *AppealFilter) GetSortableFields() []string {
	return []string{"date"}
}
