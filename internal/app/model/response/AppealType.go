package response

type AppealType struct {
	Id          int32        `json:"id"`
	Title       *string      `json:"title"`
	AppealTheme *AppealTheme `json:"appeal_theme"`
}
