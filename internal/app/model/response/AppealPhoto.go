package response

type AppealPhoto struct {
	Id  int32   `json:"id" example:"1"`
	Url *string `json:"url example:"https://storage.yandexcloud.net/urbathon/test.jpg"`
}
