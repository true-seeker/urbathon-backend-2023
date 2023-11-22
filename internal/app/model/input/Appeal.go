package input

type Appeal struct {
	Id           int32    `json:"id"`
	UserID       *int32   `json:"user_id"`
	AppealTypeID *int32   `json:"appeal_type_id"`
	Title        *string  `json:"title"`
	Description  *string  `json:"description"`
	Address      *string  `json:"address"`
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
}
