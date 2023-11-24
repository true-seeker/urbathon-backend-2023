package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/response"
)

func NewsCategoryModelToNewsCategoryResponse(news model.NewsCategories) *response.NewsCategory {
	r := &response.NewsCategory{
		Id:    news.ID,
		Title: news.Title,
	}

	return r
}

func NewsCategoryModelListToNewsCategoryResponses(newsList *[]model.NewsCategories) *[]response.NewsCategory {
	rs := make([]response.NewsCategory, 0)

	for _, news := range *newsList {
		rs = append(rs, *NewsCategoryModelToNewsCategoryResponse(news))
	}

	return &rs
}
