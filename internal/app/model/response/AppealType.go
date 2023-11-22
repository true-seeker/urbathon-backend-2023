package response

type AppealType struct {
	Id          int32        `json:"id" example:"1"`
	Title       *string      `json:"title" example:"Тип"`
	AppealTheme *AppealTheme `json:"appeal_theme"`
}
