package response

import (
	"time"
	"urbathon-backend-2023/internal/app/model/input"
)

type News struct {
	Id    int32      `json:"id" example:"1"`
	Title *string    `json:"title" example:"Заголовок"`
	Body  *string    `json:"body" example:"Тело новости"`
	Date  *time.Time `json:"date" example:"2024-02-10T00:00:00+05:00"`
}

type NewsPaged struct {
	Paged
	Items *[]News `json:"items"`
}

func NewNewsPaged(f *input.Filter, items *[]News, total *int) *NewsPaged {
	return &NewsPaged{
		Paged: *NewPaged(f.Page, f.PageSize, *total, "news"),
		Items: items,
	}
}
