package response

import (
	"time"
	"urbathon-backend-2023/internal/app/model/input"
)

type News struct {
	Id    int32      `json:"id"`
	Title *string    `json:"title"`
	Body  *string    `json:"body"`
	Date  *time.Time `json:"date"`
}

type NewsPaged struct {
	Paged
	Items *[]News `json:"items"`
}

func NewNewsPaged(f *input.Filter, items *[]News, total *int) *NewsPaged {
	return &NewsPaged{
		Paged: Paged{
			Page:     f.Page,
			PageSize: f.PageSize,
			Total:    *total,
		},
		Items: items,
	}
}
