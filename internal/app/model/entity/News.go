package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type News struct {
	model.News

	NewsCategory *model.NewsCategories `json:"news_category"`
	Poll         *NewsPoll             `json:"poll"`
}
