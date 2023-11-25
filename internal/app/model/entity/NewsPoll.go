package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type NewsPoll struct {
	model.NewsPolls

	Options *[]model.PollOptions `json:"options"`
}
