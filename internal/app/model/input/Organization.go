package input

type OrganizationRegister struct {
	Name       string   `json:"name" example:"ЖКХ"`
	Inn        string   `json:"inn"`
	Address    string   `json:"address"`
	Phone      string   `json:"phone"`
	Categories *[]int32 `json:"category_ids"`
}

type OrganizationAddUser struct {
	UserId int32   `json:"user_id" example:"1"`
	Job    *string `json:"job" example:"Главный по главным"`
}
