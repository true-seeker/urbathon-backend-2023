package response

type AppealStatus struct {
	Id     int32   `json:"id" example:"1"`
	Status *string `json:"status" example:"Решено"`
}
