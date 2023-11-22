package response

type AppealCategory struct {
	Id    int32   `json:"id" example:"1"`
	Title *string `json:"title" example:"Категория"`
}
