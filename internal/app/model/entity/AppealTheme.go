package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type AppealTheme struct {
	model.AppealThemes

	AppealCategory *AppealCategory `json:"appeal_category"`
}
