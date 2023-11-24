package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealPhotoModelToAppealPhotoResponse(appealPhoto model.AppealPhotos) *response.AppealPhoto {
	r := &response.AppealPhoto{
		Id:  appealPhoto.ID,
		Url: appealPhoto.URL,
	}

	return r
}

func AppealPhotoModelListToAppealPhotoResponses(appealPhotoList *[]model.AppealPhotos) *[]response.AppealPhoto {
	rs := make([]response.AppealPhoto, 0)

	if appealPhotoList == nil {
		return &rs
	}

	for _, appealPhoto := range *appealPhotoList {
		rs = append(rs, *AppealPhotoModelToAppealPhotoResponse(appealPhoto))
	}

	return &rs
}
