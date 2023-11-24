package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"mime/multipart"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/s3"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type NewsRepository interface {
	Get(id *int32) (*entity.News, error)
	GetAll(f *input.Filter) (*[]entity.News, error)
	GetTotal() (*int, error)
	Create(news *model.News) (*entity.News, error)
}
type NewsService struct {
	newsRepo NewsRepository
}

func NewNewsService(newsRepository NewsRepository) *NewsService {
	return &NewsService{newsRepo: newsRepository}
}

func (d *NewsService) Get(id *int32) (*response.News, *errorHandler.HttpErr) {
	newsResponse := &response.News{}
	news, err := d.newsRepo.Get(id)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("News with id does not exists", http.StatusNotFound)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	newsResponse = mapper.NewsToNewsResponse(*news)
	return newsResponse, nil
}

func (d *NewsService) GetAll(f *input.Filter) (*response.NewsPaged, *errorHandler.HttpErr) {
	items := &[]response.News{}
	news, err := d.newsRepo.GetAll(f)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	total, err := d.newsRepo.GetTotal()
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	items = mapper.NewsListToNewsResponses(news)
	newsPaged := response.NewNewsPaged(f, items, total)
	return newsPaged, nil
}

func (d *NewsService) Create(newsInput *input.News, user *model.Users) (*response.News, *errorHandler.HttpErr) {
	newsResponse := &response.News{}
	if httpErr := d.validateCreate(newsInput); httpErr != nil {
		return nil, httpErr
	}

	news := mapper.NewsInputToNews(newsInput)
	news.UserID = &user.ID
	news.OrganizationID = user.OrganizationID

	if newsInput.Photo != nil {
		photoUrls, httpErr := s3.UploadPhotos(&[]multipart.FileHeader{*newsInput.Photo})
		if httpErr != nil {
			return nil, httpErr
		}
		news.PhotoURL = &((*photoUrls)[0])
	}

	appealEntity, err := d.newsRepo.Create(news)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	newsResponse = mapper.NewsToNewsResponse(*appealEntity)

	return newsResponse, nil
}

func (d *NewsService) validateCreate(newsInput *input.News) *errorHandler.HttpErr {
	if httpErr := validator.NewsCreate(newsInput); httpErr != nil {
		return httpErr
	}
	// todo appeal_type_id
	return nil
}
