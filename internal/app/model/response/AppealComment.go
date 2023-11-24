package response

import (
	"fmt"
	"time"
	"urbathon-backend-2023/internal/app/model/input"
)

type AppealComment struct {
	Id   int32      `json:"id" example:"1"`
	User *User      `json:"user"`
	Date *time.Time `json:"date" example:"2024-02-10T00:00:00+05:00"`
	Text *string    `json:"text" example:"Текст комментария"`
}

type AppealCommentPaged struct {
	Paged
	Items *[]AppealComment `json:"items"`
}

func NewAppealCommentPaged(f *input.Filter, items *[]AppealComment, total *int, appealId int32) *AppealCommentPaged {
	return &AppealCommentPaged{
		Paged: *NewPaged(f.Page, f.PageSize, *total, fmt.Sprintf("appeal/%d/comments", appealId)),
		Items: items,
	}
}
