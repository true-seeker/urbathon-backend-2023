package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealToAppealResponse(appeal entity.Appeal) *response.Appeal {
	r := &response.Appeal{
		Id:          appeal.ID,
		Title:       &appeal.Title,
		Description: &appeal.Description,
		Address:     &appeal.Address,
		Latitude:    &appeal.Latitude,
		Longitude:   &appeal.Longitude,
		User: UserToUserResponse(&model.Users{
			ID:    appeal.User.ID,
			Name:  appeal.User.Name,
			Email: appeal.User.Email,
		}),
		AppealType: AppealTypeToAppealTypeResponse(*appeal.AppealType),
	}
	return r
}

func AppealListToAppealResponses(appealList *[]entity.Appeal) *[]response.Appeal {
	rs := make([]response.Appeal, 0)

	for _, appeal := range *appealList {
		rs = append(rs, *AppealToAppealResponse(appeal))
	}

	return &rs
}

func AppealInputToAppeal(appealInput *input.Appeal) *entity.Appeal {
	r := &entity.Appeal{
		Appeals: model.Appeals{ID: appealInput.Id},
	}
	return r
}
