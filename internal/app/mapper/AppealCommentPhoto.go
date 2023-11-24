package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealCommentPhotoModelToAppealCommentPhotoResponse(appealCommentPhoto model.AppealCommentPhotos) *response.AppealCommentPhoto {
	r := &response.AppealCommentPhoto{
		Id:  appealCommentPhoto.ID,
		Url: appealCommentPhoto.URL,
	}

	return r
}

func AppealCommentPhotoModelListToAppealCommentPhotoResponses(appealCommentPhotoList *[]model.AppealCommentPhotos) *[]response.AppealCommentPhoto {
	rs := make([]response.AppealCommentPhoto, 0)

	if appealCommentPhotoList == nil {
		return &rs
	}

	for _, appealPhoto := range *appealCommentPhotoList {
		rs = append(rs, *AppealCommentPhotoModelToAppealCommentPhotoResponse(appealPhoto))
	}

	return &rs
}
