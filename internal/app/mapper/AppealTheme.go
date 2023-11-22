package mapper

import (
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealThemeToAppealThemeResponse(appealTheme entity.AppealTheme) *response.AppealTheme {
	r := &response.AppealTheme{
		Id:             appealTheme.ID,
		Title:          appealTheme.Title,
		AppealCategory: AppealCategoryToAppealCategoryResponse(*appealTheme.AppealCategory),
	}

	return r
}

func AppealThemeListToAppealThemeResponses(appealThemeList *[]entity.AppealTheme) *[]response.AppealTheme {
	rs := make([]response.AppealTheme, 0)

	for _, appealTheme := range *appealThemeList {
		rs = append(rs, *AppealThemeToAppealThemeResponse(appealTheme))
	}

	return &rs
}
