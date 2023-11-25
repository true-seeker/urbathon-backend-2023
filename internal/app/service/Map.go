package service

import (
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealMapRepository interface {
	GetForMap(f *filter.Map) (*[]entity.Appeal, error)
}
type TkoMapRepository interface {
	GetForMap(f *filter.Map) (*[]model.Tko, error)
}
type NewsMapRepository interface {
	GetForMap(f *filter.Map) (*[]model.News, error)
}

type MapService struct {
	appealRepo AppealMapRepository
	tkoRepo    TkoMapRepository
	newsRepo   NewsMapRepository
}

func NewMapService(appealRepo AppealMapRepository, tkoRepo TkoMapRepository, newsRepo NewsMapRepository) *MapService {
	return &MapService{appealRepo: appealRepo, tkoRepo: tkoRepo, newsRepo: newsRepo}
}

func (d *MapService) GetMapElements(f *filter.Map) (*[]response.MapElement, *errorHandler.HttpErr) {
	response := []response.MapElement{}
	appealEntity, err := d.appealRepo.GetForMap(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	mapElementResponse := mapper.AppealsToMapElementResponses(appealEntity)

	tkoModel, err := d.tkoRepo.GetForMap(f)
	tkomapElementResponse := mapper.TkosToMapElementResponses(tkoModel)

	newsModel, err := d.newsRepo.GetForMap(f)
	newsMapElementResponse := mapper.NewsToMapElementResponses(newsModel)

	response = append(response, *mapElementResponse...)
	response = append(response, *newsMapElementResponse...)
	response = append(response, *tkomapElementResponse...)
	return &response, nil
}
