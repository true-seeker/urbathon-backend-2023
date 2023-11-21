package input

type Login struct {
	Email    *string `json:"email" example:"test@gmail.com"`
	Password *string `json:"password" example:"123456"`
}
