package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type AppealType struct {
	model.AppealTypes

	AppealTheme *AppealTheme `json:"appeal_theme"`
}
