package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealTypeToAppealTypeResponse(appealType entity.AppealType) *response.AppealType {
	r := &response.AppealType{
		Id:             appealType.ID,
		Title:          appealType.Title,
		AppealCategory: AppealCategoryToAppealCategoryResponse(*appealType.AppealCategory),
	}

	return r
}

func AppealTypeModelToAppealTypeResponseByCategory(appealType model.AppealTypes) *response.AppealTypeByCategory {
	r := &response.AppealTypeByCategory{
		Id:    appealType.ID,
		Title: appealType.Title,
	}

	return r
}

func AppealTypeListToAppealTypeResponses(appealTypeList *[]entity.AppealType) *[]response.AppealType {
	rs := make([]response.AppealType, 0)

	for _, appealType := range *appealTypeList {
		rs = append(rs, *AppealTypeToAppealTypeResponse(appealType))
	}

	return &rs
}

func AppealTypeModelListToAppealTypeResponsesByCategory(appealTypeList *[]model.AppealTypes) *[]response.AppealTypeByCategory {
	rs := make([]response.AppealTypeByCategory, 0)

	for _, appealType := range *appealTypeList {
		rs = append(rs, *AppealTypeModelToAppealTypeResponseByCategory(appealType))
	}

	return &rs
}
