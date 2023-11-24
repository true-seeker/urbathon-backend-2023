package response

type AppealCommentPhoto struct {
	Id  int32   `json:"id" example:"1"`
	Url *string `json:"url" example:"https://storage.yandexcloud.net/urbathon/test.jpg"`
}
