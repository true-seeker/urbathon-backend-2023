package mapper

import (
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

func AppealTypeListToAppealTypeResponses(appealTypeList *[]entity.AppealType) *[]response.AppealType {
	rs := make([]response.AppealType, 0)

	for _, appealType := range *appealTypeList {
		rs = append(rs, *AppealTypeToAppealTypeResponse(appealType))
	}

	return &rs
}
