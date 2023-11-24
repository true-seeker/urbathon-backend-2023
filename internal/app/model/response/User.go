package response

type User struct {
	Id    int32   `json:"id" example:"1"`
	Email *string `json:"email,omitempty" example:"test@gmail.com"`
	Name  *string `json:"name" example:"Иванов Иван Иванович"`
}
