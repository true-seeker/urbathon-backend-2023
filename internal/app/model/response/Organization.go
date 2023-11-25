package response

type Organization struct {
	Id      int32   `json:"id" example:"1"`
	Name    *string `json:"name" example:"ЖКХ"`
	Inn     *string `json:"inn,omitempty"`
	Address *string `json:"address,omitempty"`
	Phone   *string `json:"phone,omitempty"`
}
