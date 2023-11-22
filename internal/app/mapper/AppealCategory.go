package mapper

import (
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealCategoryToAppealCategoryResponse(appealCategory entity.AppealCategory) *response.AppealCategory {
	r := &response.AppealCategory{
		Id:    appealCategory.ID,
		Title: appealCategory.Title,
	}

	return r
}

func AppealCategoryListToAppealCategoryResponses(appealCategoryList *[]entity.AppealCategory) *[]response.AppealCategory {
	rs := make([]response.AppealCategory, 0)

	for _, appealCategory := range *appealCategoryList {
		rs = append(rs, *AppealCategoryToAppealCategoryResponse(appealCategory))
	}

	return &rs
}
