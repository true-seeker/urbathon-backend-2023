package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"net/http"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealTypeRepository interface {
	Get(id *int32) (*entity.AppealType, error)
	GetAll() (*[]entity.AppealType, error)
}
type AppealTypeService struct {
	appealTypeRepo AppealTypeRepository
}

func NewAppealTypeService(appealTypeRepository AppealTypeRepository) *AppealTypeService {
	return &AppealTypeService{appealTypeRepo: appealTypeRepository}
}

func (d *AppealTypeService) Get(id *int32) (*response.AppealType, *errorHandler.HttpErr) {
	appealTypeResponse := &response.AppealType{}
	appealType, err := d.appealTypeRepo.Get(id)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("AppealType with id does not exists", http.StatusNotFound)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	appealTypeResponse = mapper.AppealTypeToAppealTypeResponse(*appealType)
	return appealTypeResponse, nil
}

func (d *AppealTypeService) GetAll() (*[]response.AppealType, *errorHandler.HttpErr) {
	appealTypeResponse := &[]response.AppealType{}
	appealType, err := d.appealTypeRepo.GetAll()
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	appealTypeResponse = mapper.AppealTypeListToAppealTypeResponses(appealType)
	return appealTypeResponse, nil
}
