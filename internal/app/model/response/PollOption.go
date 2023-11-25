package response

type PollOption struct {
	ID          int32   `json:"id" example:"1"`
	Title       *string `json:"title" example:"Вариант опроса"`
	Votes       *int    `json:"votes" example:"1"`
	IsUserVoted *bool   `json:"is_user_voted,omitempty"`
}
