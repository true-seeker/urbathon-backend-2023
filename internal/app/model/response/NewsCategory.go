package response

type NewsCategory struct {
	Id    int32   `json:"id" example:"1"`
	Title *string `json:"title" example:"Отходы"`
}
