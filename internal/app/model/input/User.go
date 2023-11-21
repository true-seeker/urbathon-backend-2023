package input

type UserLogin struct {
	Email    *string `json:"email" example:"test@gmail.com"`
	Password *string `json:"password" example:"123456"`
}

type User struct {
	Email    *string `json:"email" example:"test@gmail.com"`
	Password *string `json:"password" example:"123456"`
	Name     *string `json:"name" example:"Ивано Иван Иванович"`
}
