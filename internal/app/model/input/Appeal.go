package input

import "mime/multipart"

type Appeal struct {
	AppealTypeID *int32                  `form:"appeal_type_id" example:"1"`
	Title        *string                 `form:"title" example:"Обращение"`
	Description  *string                 `form:"description" example:"Текст обращения"`
	Address      *string                 `form:"address" example:"Улица Пушкина"`
	Latitude     *float64                `form:"latitude" example:"54.1234"`
	Longitude    *float64                `form:"longitude" example:"122.7656"`
	Photos       *[]multipart.FileHeader `form:"photos" swaggerignore:"true"`
}

type AppealUpdate struct {
	AppealTypeID *int32   `json:"appeal_type_id" example:"1"`
	Title        *string  `json:"title" example:"Обращение"`
	Description  *string  `json:"description" example:"Текст обращения"`
	Address      *string  `json:"address" example:"Улица Пушкина"`
	Latitude     *float64 `json:"latitude" example:"54.1234"`
	Longitude    *float64 `json:"longitude" example:"122.7656"`
}
