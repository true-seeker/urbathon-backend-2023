package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealStatusRepository interface {
	Get(id *int32) (*model.AppealStatus, error)
	GetAll() (*[]model.AppealStatus, error)
}
type AppealStatusService struct {
	appealStatusRepo AppealStatusRepository
}

func NewAppealStatusService(appealStatusRepository AppealStatusRepository) *AppealStatusService {
	return &AppealStatusService{appealStatusRepo: appealStatusRepository}
}

func (d *AppealStatusService) Get(id *int32) (*response.AppealStatus, *errorHandler.HttpErr) {
	appealStatusResponse := &response.AppealStatus{}
	appealStatus, err := d.appealStatusRepo.Get(id)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("AppealStatus with id does not exists", http.StatusNotFound)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	appealStatusResponse = mapper.AppealStatusModelToAppealStatusResponse(*appealStatus)
	return appealStatusResponse, nil
}

func (d *AppealStatusService) GetAll() (*[]response.AppealStatus, *errorHandler.HttpErr) {
	appealStatusResponse := &[]response.AppealStatus{}
	appealStatus, err := d.appealStatusRepo.GetAll()
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	appealStatusResponse = mapper.AppealStatusModelListToAppealStatusResponses(appealStatus)
	return appealStatusResponse, nil
}
