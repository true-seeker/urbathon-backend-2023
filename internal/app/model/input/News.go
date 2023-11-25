package input

import "mime/multipart"

type News struct {
	Title      *string               `form:"title"`
	Body       *string               `form:"body"`
	CategoryId *int32                `form:"category_id"`
	Photo      *multipart.FileHeader `form:"photo" swaggerignore:"true"`
	Poll       *string               `form:"poll" example:"тут json голосования"`
	Address    *string               `form:"address" example:"Улица Пушкина"`
	Latitude   *float64              `form:"latitude" example:"54.1234"`
	Longitude  *float64              `form:"longitude" example:"122.7656"`
}
