package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealStatusModelToAppealStatusResponse(appealStatus model.AppealStatus) *response.AppealStatus {
	r := &response.AppealStatus{
		Id:     appealStatus.ID,
		Status: appealStatus.Status,
	}

	return r
}

func AppealStatusModelListToAppealStatusResponses(appealStatusList *[]model.AppealStatus) *[]response.AppealStatus {
	rs := make([]response.AppealStatus, 0)

	for _, appealStatus := range *appealStatusList {
		rs = append(rs, *AppealStatusModelToAppealStatusResponse(appealStatus))
	}

	return &rs
}
