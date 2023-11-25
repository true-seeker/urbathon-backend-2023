package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func NewsModelToNewsResponse(news model.News) *response.News {
	r := &response.News{
		Id:        news.ID,
		Title:     &news.Title,
		Body:      &news.Body,
		Date:      news.Date,
		PhotoUrl:  news.PhotoURL,
		Address:   news.Address,
		Latitude:  news.Latitude,
		Longitude: news.Longitude,
	}

	return r
}

func NewsToNewsResponse(news entity.News) *response.News {
	r := &response.News{
		Id:        news.ID,
		Title:     &news.Title,
		Body:      &news.Body,
		Date:      news.Date,
		Category:  NewsCategoryModelToNewsCategoryResponse(*news.NewsCategory),
		PhotoUrl:  news.PhotoURL,
		Poll:      NewsPollToNewsPollResponse(news.Poll),
		Address:   news.Address,
		Latitude:  news.Latitude,
		Longitude: news.Longitude,
	}

	return r
}

func NewsModelListToNewsResponses(newsList *[]model.News) *[]response.News {
	rs := make([]response.News, 0)

	for _, news := range *newsList {
		rs = append(rs, *NewsModelToNewsResponse(news))
	}

	return &rs
}

func NewsListToNewsResponses(newsList *[]entity.News) *[]response.News {
	rs := make([]response.News, 0)

	for _, news := range *newsList {
		rs = append(rs, *NewsToNewsResponse(news))
	}

	return &rs
}

func NewsInputToNews(newsInput *input.News) *model.News {
	r := &model.News{
		Title:      *newsInput.Title,
		Body:       *newsInput.Body,
		CategoryID: newsInput.CategoryId,
	}
	return r
}
