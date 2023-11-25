package response

type User struct {
	Id             int32   `json:"id" example:"1"`
	Email          *string `json:"email,omitempty" example:"test@gmail.com"`
	FirstName      *string `json:"first_name" example:"Иван"`
	LastName       *string `json:"last_name" example:"Иванов"`
	Patronymic     *string `json:"patronymic" example:"Иванович"`
	OrganizationId *int32  `json:"organization_id,omitempty" example:"1"`
}
