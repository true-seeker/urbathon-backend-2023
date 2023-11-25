package service

import (
	"net/http"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealMapRepository interface {
	GetMapElements(f *filter.Map) (*[]entity.Appeal, error)
}

type MapService struct {
	appealRepo AppealMapRepository
}

func NewMapService(appealRepo AppealMapRepository) *MapService {
	return &MapService{appealRepo: appealRepo}
}

func (d *MapService) GetMapElements(f *filter.Map) (*[]response.MapElement, *errorHandler.HttpErr) {
	mapElements, err := d.appealRepo.GetMapElements(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	mapElementResponse := mapper.AppealsToMapElementResponses(mapElements)
	return mapElementResponse, nil
}
