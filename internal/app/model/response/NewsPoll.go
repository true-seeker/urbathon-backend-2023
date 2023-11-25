package response

type NewsPoll struct {
	ID      int32         `json:"id" example:"1"`
	Title   *string       `json:"title" example:"Название опроса"`
	Options *[]PollOption `json:"options"`
}
