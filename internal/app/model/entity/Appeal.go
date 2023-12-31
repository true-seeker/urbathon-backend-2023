package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type Appeal struct {
	model.Appeals

	User         *model.Users          `json:"user"`
	AppealType   *AppealType           `json:"appeal_type"`
	AppealPhotos *[]model.AppealPhotos `json:"appeal_photos"`
	AppealStatus *model.AppealStatus   `json:"appeal_status"`
}
