package response

import (
	"urbathon-backend-2023/internal/app/model/input"
)

type Appeal struct {
	Id          int32       `json:"id"`
	User        *User       `json:"user"`
	AppealType  *AppealType `json:"appeal_type"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Address     *string     `json:"address"`
	Latitude    *float64    `json:"latitude"`
	Longitude   *float64    `json:"longitude"`
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
