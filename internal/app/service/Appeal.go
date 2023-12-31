package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/s3"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealRepository interface {
	Get(id *int32) (*entity.Appeal, error)
	GetAll(f *filter.AppealFilter) (*[]entity.Appeal, error)
	GetTotal(f *filter.AppealFilter) (*int, error)
	Create(appeal *model.Appeals, urls *[]string) (*entity.Appeal, error)
	Update(appeal *model.Appeals) (*entity.Appeal, error)
	Delete(id int32) error
	UpdateStatus(appealId int32, statusId int32) error
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

func (d *AppealService) GetAll(f *filter.AppealFilter) (*response.AppealPaged, *errorHandler.HttpErr) {
	appeal, err := d.appealRepo.GetAll(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	total, err := d.appealRepo.GetTotal(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	items := mapper.AppealListToAppealResponses(appeal)
	appealPaged := response.NewAppealPaged(f.Pagination, items, total)
	return appealPaged, nil
}

func (d *AppealService) Create(appealInput *input.Appeal, user *model.Users) (*response.Appeal, *errorHandler.HttpErr) {
	userResponse := &response.Appeal{}
	if httpErr := d.validateCreate(appealInput); httpErr != nil {
		return nil, httpErr
	}

	appeal := mapper.AppealInputToAppeal(appealInput)

	appeal.UserID = user.ID
	photoUrls, httpErr := s3.UploadPhotos(appealInput.Photos)
	if httpErr != nil {
		return nil, httpErr
	}

	appealEntity, err := d.appealRepo.Create(appeal, photoUrls)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	userResponse = mapper.AppealToAppealResponse(*appealEntity)

	return userResponse, nil
}

func (d *AppealService) Update(appealInput *input.AppealUpdate, user *model.Users, id *int32) (*response.Appeal, *errorHandler.HttpErr) {
	userResponse := &response.Appeal{}
	if httpErr := d.validateUpdate(appealInput); httpErr != nil {
		return nil, httpErr
	}
	//todo exists validation
	appeal := mapper.AppealInputUpdateToAppeal(appealInput)
	appeal.UserID = user.ID
	appeal.ID = *id

	appealEntity, err := d.appealRepo.Update(appeal)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	userResponse = mapper.AppealToAppealResponse(*appealEntity)

	return userResponse, nil
}

func (d *AppealService) Delete(id int32) *errorHandler.HttpErr {
	//todo exists validation
	err := d.appealRepo.Delete(id)
	if err != nil {
		return errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	//todo deletion_date
	return nil
}

func (d *AppealService) UpdateStatus(appealId int32, statusId int32) *errorHandler.HttpErr {
	//todo exists validation
	err := d.appealRepo.UpdateStatus(appealId, statusId)
	if err != nil {
		return errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	//todo deletion_date
	return nil
}

func (d *AppealService) GetMyAppeals(f *filter.AppealFilter) (*response.AppealPaged, *errorHandler.HttpErr) {
	appeal, err := d.appealRepo.GetAll(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	total, err := d.appealRepo.GetTotal(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	items := mapper.AppealListToAppealResponses(appeal)
	appealPaged := response.NewAppealPaged(f.Pagination, items, total)
	return appealPaged, nil
}

func (d *AppealService) validateCreate(appealInput *input.Appeal) *errorHandler.HttpErr {
	if httpErr := validator.AppealCreate(appealInput); httpErr != nil {
		return httpErr
	}
	// todo appeal_type_id
	return nil
}

func (d *AppealService) validateUpdate(appealInput *input.AppealUpdate) *errorHandler.HttpErr {
	if httpErr := validator.AppealUpdate(appealInput); httpErr != nil {
		return httpErr
	}
	// todo appeal_type_id
	return nil
}
