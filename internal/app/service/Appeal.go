package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"net/http"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealRepository interface {
	Get(id *int32) (*entity.Appeal, error)
	GetAll(f *input.Filter) (*[]entity.Appeal, error)
	GetTotal() (*int, error)
}
type AppealService struct {
	appealRepo AppealRepository
}

func NewAppealService(appealRepository AppealRepository) *AppealService {
	return &AppealService{appealRepo: appealRepository}
}

func (d *AppealService) Get(id *int32) (*response.Appeal, *errorHandler.HttpErr) {
	appealResponse := &response.Appeal{}
	appeal, err := d.appealRepo.Get(id)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("Appeal with id does not exists", http.StatusNotFound)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	appealResponse = mapper.AppealToAppealResponse(*appeal)
	return appealResponse, nil
}

func (d *AppealService) GetAll(f *input.Filter) (*response.AppealPaged, *errorHandler.HttpErr) {
	items := &[]response.Appeal{}
	appeal, err := d.appealRepo.GetAll(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	total, err := d.appealRepo.GetTotal()
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	items = mapper.AppealListToAppealResponses(appeal)
	appealPaged := response.NewAppealPaged(f, items, total)
	return appealPaged, nil
}
func (d *AppealService) validateLogin(loginInput *input.UserLogin) *errorHandler.HttpErr {
	if httpErr := validator.UserLogin(loginInput); httpErr != nil {
		return httpErr
	}
	return nil
}
