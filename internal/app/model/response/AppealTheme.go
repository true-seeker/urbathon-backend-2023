package response

type AppealTheme struct {
	Id             int32           `json:"id" example:"1"`
	Title          *string         `json:"title" example:"Тема"`
	AppealCategory *AppealCategory `json:"appeal_category"`
}
