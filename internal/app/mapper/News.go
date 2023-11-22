package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func NewsToNewsResponse(news *model.News) *response.News {
	r := &response.News{
		Id:    news.ID,
		Title: &news.Title,
		Body:  &news.Body,
		Date:  news.Date,
	}

	return r
}

func NewsListToNewsResponses(newss *[]model.News) *[]response.News {
	rs := make([]response.News, 0)

	for _, news := range *newss {
		rs = append(rs, *NewsToNewsResponse(&news))
	}

	return &rs
}

func NewsInputToNews(newsInput *input.News) *model.News {
	r := &model.News{
		ID:    newsInput.Id,
		Title: *newsInput.Title,
		Body:  *newsInput.Body,
		Date:  newsInput.Date,
	}
	return r
}
