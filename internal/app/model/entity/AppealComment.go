package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type AppealComment struct {
	model.AppealComments

	User *model.Users `json:"user"`
}
