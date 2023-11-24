package service

import (
	"fmt"
	"net/http"
	"time"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/s3"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/config"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealCommentRepository interface {
	GetAllComments(f *input.Filter, appealId int32) (*[]entity.AppealComment, error)
	GetTotalComments(appealId int32) (*int, error)
	Create(appealComments *model.AppealComments, urls *[]string) (*entity.AppealComment, error)
}
type AppealCommentService struct {
	appealCommentRepo AppealCommentRepository
}

func NewAppealCommentService(appealCommentRepository AppealCommentRepository) *AppealCommentService {
	return &AppealCommentService{appealCommentRepo: appealCommentRepository}
}

func (d *AppealCommentService) GetAllByAppealId(f *input.Filter, appealId int32) (*response.AppealCommentPaged, *errorHandler.HttpErr) {
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

	photoUrls, httpErr := uploadAppealCommentPhotos(appealCommentInput)
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

func uploadAppealCommentPhotos(appealInput *input.AppealComment) (*[]string, *errorHandler.HttpErr) {
	var urls []string
	if appealInput.Photos == nil {
		return &urls, nil
	}
	for _, photo := range *appealInput.Photos {
		filename := fmt.Sprintf("%s_%s", time.Now().Format(config.DateTimeLayout), photo.Filename)
		openedFile, _ := photo.Open()
		url, err := s3.BucketBase.UploadFile("urbathon", filename, openedFile)
		if err != nil {
			return nil, errorHandler.New("Yandex S3 not available", http.StatusBadRequest)
		}
		urls = append(urls, url)
	}

	return &urls, nil
}
