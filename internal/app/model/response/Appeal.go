package response

import (
	"urbathon-backend-2023/internal/app/model/input"
)

type Appeal struct {
	Id           int32          `json:"id" example:"1"`
	User         *User          `json:"user"`
	AppealType   *AppealType    `json:"appeal_type"`
	Title        *string        `json:"title" example:"Обращение"`
	Description  *string        `json:"description" example:"Текст обращения"`
	Address      *string        `json:"address" example:"Улица Пушкина"`
	Latitude     *float64       `json:"latitude" example:"54.1234"`
	Longitude    *float64       `json:"longitude" example:"122.7656"`
	AppealPhotos *[]AppealPhoto `json:"appeal_photos"`
}

type AppealPaged struct {
	Paged
	Items *[]Appeal `json:"items"`
}

func NewAppealPaged(f *input.Filter, items *[]Appeal, total *int) *AppealPaged {
	return &AppealPaged{
		Paged: Paged{
			Page:     f.Page,
			PageSize: f.PageSize,
			Total:    *total,
		},
		Items: items,
	}
}
