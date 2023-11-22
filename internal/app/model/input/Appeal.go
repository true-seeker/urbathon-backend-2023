package input

type Appeal struct {
	AppealTypeID *int32   `json:"appeal_type_id" example:"1"`
	Title        *string  `json:"title" example:"Обращение"`
	Description  *string  `json:"description" example:"Текст обращения"`
	Address      *string  `json:"address" example:"Улица Пушкина"`
	Latitude     *float64 `json:"latitude" example:"54.1234"`
	Longitude    *float64 `json:"longitude" example:"122.7656"`
}
