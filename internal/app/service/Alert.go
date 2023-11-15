package service

import (
	"fmt"
	"net/http"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AlertRepository interface {
	GetAll(user *entity.User) (*[]entity.Alert, error)
	Get(id int, user *entity.User) (*entity.Alert, error)
}
type AlertService struct {
	alertRepo AlertRepository
}

func NewAlertService(alertRepo AlertRepository) *AlertService {
	return &AlertService{alertRepo: alertRepo}
}

func (d *AlertService) GetAll(user *entity.User) (*[]response.Alert, *errorHandler.HttpErr) {
	alertResponses := &[]response.Alert{}

	alert, err := d.alertRepo.GetAll(user)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	alertResponses = mapper.AlertToAlertResponses(alert)

	return alertResponses, nil
}

func (d *AlertService) Get(id int, user *entity.User) (*response.Alert, *errorHandler.HttpErr) {
	alertResponse := &response.Alert{}
	alert, err := d.alertRepo.Get(id, user)
	if alert == nil {
		return nil, errorHandler.New(fmt.Sprintf("Alert with id %d does not exists", id), http.StatusNotFound)
	}
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	alertResponse = mapper.AlertToAlertResponse(alert)

	return alertResponse, nil
}
