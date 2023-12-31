package response

import (
	"time"
	"urbathon-backend-2023/internal/app/model/filter"
)

type News struct {
	Id        int32         `json:"id" example:"1"`
	Title     *string       `json:"title" example:"Заголовок"`
	Body      *string       `json:"body" example:"Тело новости"`
	Date      *time.Time    `json:"date" example:"2024-02-10T00:00:00+05:00"`
	Category  *NewsCategory `json:"category"`
	PhotoUrl  *string       `json:"photo_url" example:"https://storage.yandexcloud.net/urbathon/test.jpg"`
	Poll      *NewsPoll     `json:"poll"`
	Address   *string       `json:"address" example:"Улица Пушкина"`
	Latitude  *float64      `json:"latitude" example:"54.1234"`
	Longitude *float64      `json:"longitude" example:"122.7656"`
}

type NewsPaged struct {
	Paged
	Items *[]News `json:"items"`
}

func NewNewsPaged(f *filter.Pagination, items *[]News, total *int) *NewsPaged {
	return &NewsPaged{
		Paged: *NewPaged(f.Page, f.PageSize, *total, "news"),
		Items: items,
	}
}
