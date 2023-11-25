package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/response"
)

func PollOptionModelToPollOptionResponse(pollOption model.PollOptions) *response.PollOption {
	r := &response.PollOption{
		ID:    pollOption.ID,
		Title: pollOption.Title,
	}

	return r
}

func PollOptionModelListToPollOptionResponses(pollOptionList *[]model.PollOptions) *[]response.PollOption {
	rs := make([]response.PollOption, 0)
	if pollOptionList == nil {
		return &rs
	}

	for _, pollOption := range *pollOptionList {
		rs = append(rs, *PollOptionModelToPollOptionResponse(pollOption))
	}

	return &rs
}
