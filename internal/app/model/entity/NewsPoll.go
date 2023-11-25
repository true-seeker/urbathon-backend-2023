package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type NewsPoll struct {
	model.NewsPolls

	Options *[]PollOption `json:"options"`
}
