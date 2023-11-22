package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type NewsRepository interface {
	Get(id *int32) (*model.News, error)
	GetAll(f *input.Filter) (*[]model.News, error)
	GetTotal() (*int, error)
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
func (d *NewsService) validateLogin(loginInput *input.UserLogin) *errorHandler.HttpErr {
	if httpErr := validator.UserLogin(loginInput); httpErr != nil {
		return httpErr
	}
	return nil
}
