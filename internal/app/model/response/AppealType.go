package response

type AppealType struct {
	Id             int32           `json:"id" example:"1"`
	Title          *string         `json:"title" example:"Тип"`
	AppealCategory *AppealCategory `json:"appeal_category"`
}

type AppealTypeByCategory struct {
	Id    int32   `json:"id" example:"1"`
	Title *string `json:"title" example:"Тип"`
}
