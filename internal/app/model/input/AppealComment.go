package input

import "mime/multipart"

type AppealComment struct {
	Text   *string                 `form:"text" example:"Текст кооментария"`
	Photos *[]multipart.FileHeader `form:"photos" swaggerignore:"true"`
}
