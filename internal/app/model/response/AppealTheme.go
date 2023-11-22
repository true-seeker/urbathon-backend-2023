package response

type AppealTheme struct {
	Id             int32           `json:"id"`
	Title          *string         `json:"title"`
	AppealCategory *AppealCategory `json:"appeal_category"`
}
