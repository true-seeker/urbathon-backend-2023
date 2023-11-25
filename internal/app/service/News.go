package service

import (
	"encoding/json"
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"mime/multipart"
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

type NewsRepository interface {
	Get(id *int32) (*entity.News, error)
	GetAll(f *filter.Pagination) (*[]entity.News, error)
	GetTotal() (*int, error)
	Create(news *model.News, poll entity.NewsPoll) (*entity.News, error)
	Vote(newsId int32, OptionId int32) error
	GetPollOptionVotesCount(optionId int32) (int, error)
	GetUserVoteOptionId(id *int32) (int, error)
}
type NewsService struct {
	newsRepo NewsRepository
}

func NewNewsService(newsRepository NewsRepository) *NewsService {
	return &NewsService{newsRepo: newsRepository}
}

func (d *NewsService) Get(id *int32, userId *int32) (*response.News, *errorHandler.HttpErr) {
	newsResponse := &response.News{}
	news, err := d.newsRepo.Get(id)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("News with id does not exists", http.StatusNotFound)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	newsResponse = mapper.NewsToNewsResponse(*news)

	if newsResponse.Poll != nil {
		var selectedOptionId int
		if userId != nil {
			selectedOptionId, err = d.newsRepo.GetUserVoteOptionId(userId)
			if err != nil {
				return nil, errorHandler.New("News with id does not exists", http.StatusNotFound)
			}
		}
		for i, option := range *newsResponse.Poll.Options {
			count, err := d.newsRepo.GetPollOptionVotesCount(option.ID)
			if err != nil {
				return nil, errorHandler.New("News with id does not exists", http.StatusNotFound)
			}
			if option.ID == int32(selectedOptionId) {
				t := true
				(*newsResponse.Poll.Options)[i].IsUserVoted = &t
			}
			(*newsResponse.Poll.Options)[i].Votes = &count
		}

	}

	return newsResponse, nil
}

func (d *NewsService) GetAll(f *filter.Pagination) (*response.NewsPaged, *errorHandler.HttpErr) {
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
	if httpErr := d.validateCreate(newsInput); httpErr != nil {
		return nil, httpErr
	}

	news := mapper.NewsInputToNews(newsInput)
	news.UserID = &user.ID
	news.OrganizationID = user.OrganizationID

	var poll entity.NewsPoll
	if newsInput.Poll != nil {
		if err := json.Unmarshal([]byte(*newsInput.Poll), &poll); err != nil {
			return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
		}
	}

	if newsInput.Photo != nil {
		photoUrls, httpErr := s3.UploadPhotos(&[]multipart.FileHeader{*newsInput.Photo})
		if httpErr != nil {
			return nil, httpErr
		}
		news.PhotoURL = &((*photoUrls)[0])
	}

	appealEntity, err := d.newsRepo.Create(news, poll)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	newsResponse, httpErr := d.Get(&appealEntity.ID, &user.ID)
	if httpErr != nil {
		return nil, httpErr
	}
	return newsResponse, nil
}

func (d *NewsService) Vote(userId int32, OptionId int32, newsId int32) (*response.News, *errorHandler.HttpErr) {
	if err := d.newsRepo.Vote(userId, OptionId); err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	newsResponse, err := d.Get(&newsId, &userId)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	return newsResponse, nil
}

func (d *NewsService) validateCreate(newsInput *input.News) *errorHandler.HttpErr {
	if httpErr := validator.NewsCreate(newsInput); httpErr != nil {
		return httpErr
	}
	// todo appeal_type_id
	return nil
}
