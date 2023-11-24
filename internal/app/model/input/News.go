package input

import "mime/multipart"

type News struct {
	Title      *string               `form:"title"`
	Body       *string               `form:"body"`
	CategoryId *int32                `form:"category_id"`
	Photo      *multipart.FileHeader `form:"photo" swaggerignore:"true"`
}
