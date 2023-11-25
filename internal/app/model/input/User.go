package input

type UserLogin struct {
	Email    *string `json:"email" example:"test@gmail.com"`
	Password *string `json:"password" example:"123456"`
}

type UserRegister struct {
	Email      *string `json:"email" example:"test@gmail.com"`
	Password   *string `json:"password" example:"123456"`
	FirstName  *string `json:"first_name" example:"Иван"`
	LastName   *string `json:"last_name" example:"Иванов"`
	Patronymic *string `json:"patronymic" example:"Иванович"`
}
