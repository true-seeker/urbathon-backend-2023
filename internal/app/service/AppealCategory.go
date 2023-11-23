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

type AppealCategoryRepository interface {
	Get(id *int32) (*model.AppealCategories, error)
	GetAll() (*[]model.AppealCategories, error)
	GetAppealTypes(id *int32) (*[]model.AppealTypes, error)
}
type AppealCategoryService struct {
	appealCategoryRepo AppealCategoryRepository
}

func NewAppealCategoryService(appealCategoryRepository AppealCategoryRepository) *AppealCategoryService {
	return &AppealCategoryService{appealCategoryRepo: appealCategoryRepository}
}

func (d *AppealCategoryService) Get(id *int32) (*response.AppealCategory, *errorHandler.HttpErr) {
	appealCategoryResponse := &response.AppealCategory{}
	appealCategory, err := d.appealCategoryRepo.Get(id)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("AppealCategory with id does not exists", http.StatusNotFound)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	appealCategoryResponse = mapper.AppealCategoryModelToAppealCategoryResponse(*appealCategory)
	return appealCategoryResponse, nil
}

func (d *AppealCategoryService) GetAll() (*[]response.AppealCategory, *errorHandler.HttpErr) {
	appealCategoryResponse := &[]response.AppealCategory{}
	appealCategory, err := d.appealCategoryRepo.GetAll()
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	appealCategoryResponse = mapper.AppealCategoryModelListToAppealCategoryResponses(appealCategory)
	return appealCategoryResponse, nil
}

func (d *AppealCategoryService) GetAppealTypes(id *int32) (*[]response.AppealTypeByCategory, *errorHandler.HttpErr) {
	appealTypeResponses := &[]response.AppealTypeByCategory{}
	_, err := d.appealCategoryRepo.Get(id)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("AppealCategory with id does not exists", http.StatusNotFound)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	appealType, err := d.appealCategoryRepo.GetAppealTypes(id)

	appealTypeResponses = mapper.AppealTypeModelListToAppealTypeResponsesByCategory(appealType)
	return appealTypeResponses, nil
}
