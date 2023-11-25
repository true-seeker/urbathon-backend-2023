package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type PollOption struct {
	model.PollOptions

	Votes       *int  `json:"votes"`
	IsUserVoted *bool `json:"is_user_voted"`
}
