package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type AppealComment struct {
	model.AppealComments

	User                *User                        `json:"user"`
	AppealCommentPhotos *[]model.AppealCommentPhotos `json:"appeal_comment_photos"`
}
