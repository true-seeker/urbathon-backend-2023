package response

type User struct {
	Id    int32   `json:"id" example:"1"`
	Email *string `json:"email" example:"test@gmail.com"`
	Name  *string `json:"name" example:"Ивано Иван Иванович"`
}
