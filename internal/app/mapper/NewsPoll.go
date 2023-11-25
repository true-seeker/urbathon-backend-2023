package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func NewsPollModelToNewsPollResponse(newsPoll model.NewsPolls) *response.NewsPoll {
	r := &response.NewsPoll{
		ID:    newsPoll.ID,
		Title: newsPoll.Title,
	}

	return r
}

func NewsPollToNewsPollResponse(newsPoll *entity.NewsPoll) *response.NewsPoll {
	if newsPoll == nil {
		return nil
	}
	r := &response.NewsPoll{
		ID:      newsPoll.ID,
		Title:   newsPoll.Title,
		Options: PollOptionListToPollOptionResponses(newsPoll.Options),
	}

	return r
}

func NewsPollModelListToNewsPollResponses(newsPollList *[]model.NewsPolls) *[]response.NewsPoll {
	rs := make([]response.NewsPoll, 0)

	for _, newsPoll := range *newsPollList {
		rs = append(rs, *NewsPollModelToNewsPollResponse(newsPoll))
	}

	return &rs
}

func NewsPollListToNewsPollResponses(newsPollList *[]entity.NewsPoll) *[]response.NewsPoll {
	rs := make([]response.NewsPoll, 0)

	for _, newsPoll := range *newsPollList {
		rs = append(rs, *NewsPollToNewsPollResponse(&newsPoll))
	}

	return &rs
}
