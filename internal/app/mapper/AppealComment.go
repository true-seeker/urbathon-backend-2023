package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealCommentToAppealCommentResponse(appeal entity.AppealComment) *response.AppealComment {
	r := &response.AppealComment{
		Id:   appeal.ID,
		Text: appeal.Text,
		Date: appeal.Date,
		User: UserToUserResponse(&model.Users{
			ID:    appeal.User.ID,
			Name:  appeal.User.Name,
			Email: appeal.User.Email,
		}),
		AppealCommentPhotos: AppealCommentPhotoModelListToAppealCommentPhotoResponses(appeal.AppealCommentPhotos),
	}
	return r
}

func AppealCommentsToAppealCommentResponses(appealList *[]entity.AppealComment) *[]response.AppealComment {
	rs := make([]response.AppealComment, 0)

	for _, appeal := range *appealList {
		rs = append(rs, *AppealCommentToAppealCommentResponse(appeal))
	}

	return &rs
}

func AppealCommentInputToAppealComment(appealInput *input.AppealComment) *model.AppealComments {
	r := &model.AppealComments{
		Text: appealInput.Text,
	}

	return r
}
