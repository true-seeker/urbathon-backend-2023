package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func AppealToMapElementResponse(appeal entity.Appeal) *response.MapElement {
	t := "appeal"
	r := &response.MapElement{
		Id:        appeal.ID,
		Latitude:  &appeal.Latitude,
		Longitude: &appeal.Longitude,
		Title:     &appeal.Title,
		Type:      &t,
	}

	return r
}

func AppealsToMapElementResponses(organizations *[]entity.Appeal) *[]response.MapElement {
	rs := make([]response.MapElement, 0)

	for _, organization := range *organizations {
		rs = append(rs, *AppealToMapElementResponse(organization))
	}

	return &rs
}

func TkoToMapElementResponse(tko model.Tko) *response.MapElement {
	t := "tko"
	r := &response.MapElement{
		Id:        tko.ID,
		Latitude:  tko.Latitude,
		Longitude: tko.Longitude,
		Title:     tko.Type,
		Type:      &t,
	}

	return r
}

func TkosToMapElementResponses(organizations *[]model.Tko) *[]response.MapElement {
	rs := make([]response.MapElement, 0)

	for _, organization := range *organizations {
		rs = append(rs, *TkoToMapElementResponse(organization))
	}

	return &rs
}
