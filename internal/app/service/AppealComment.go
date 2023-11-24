package service

import (
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

type AppealCommentRepository interface {
	GetAllComments(f *filter.Pagination, appealId int32) (*[]entity.AppealComment, error)
	GetTotalComments(appealId int32) (*int, error)
	Create(appealComments *model.AppealComments, urls *[]string) (*entity.AppealComment, error)
}
type AppealCommentService struct {
	appealCommentRepo AppealCommentRepository
}

func NewAppealCommentService(appealCommentRepository AppealCommentRepository) *AppealCommentService {
	return &AppealCommentService{appealCommentRepo: appealCommentRepository}
}

func (d *AppealCommentService) GetAllByAppealId(f *filter.Pagination, appealId int32) (*response.AppealCommentPaged, *errorHandler.HttpErr) {
	items := &[]response.AppealComment{}
	comments, err := d.appealCommentRepo.GetAllComments(f, appealId)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	total, err := d.appealCommentRepo.GetTotalComments(appealId)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	items = mapper.AppealCommentsToAppealCommentResponses(comments)
	appealPaged := response.NewAppealCommentPaged(f, items, total, appealId)
	return appealPaged, nil
}

func (d *AppealCommentService) validateCreate(appealComment *input.AppealComment) *errorHandler.HttpErr {
	if httpErr := validator.AppealCommentCreate(appealComment); httpErr != nil {
		return httpErr
	}
	// todo appeal_type_id
	return nil
}

func (d *AppealCommentService) Create(appealCommentInput *input.AppealComment, user *model.Users, appealId *int32) (*response.AppealComment, *errorHandler.HttpErr) {
	appealCommentResponse := &response.AppealComment{}
	if httpErr := d.validateCreate(appealCommentInput); httpErr != nil {
		return nil, httpErr
	}

	appealComment := mapper.AppealCommentInputToAppealComment(appealCommentInput)
	appealComment.UserID = &user.ID
	appealComment.AppealID = appealId

	photoUrls, httpErr := s3.UploadPhotos(appealCommentInput.Photos)
	if httpErr != nil {
		return nil, httpErr
	}

	appeal, err := d.appealCommentRepo.Create(appealComment, photoUrls)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	appealCommentResponse = mapper.AppealCommentToAppealCommentResponse(*appeal)

	return appealCommentResponse, nil
}
