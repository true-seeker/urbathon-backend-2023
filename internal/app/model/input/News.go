package input

import "time"

type News struct {
	Id    int32      `json:"id"`
	Title *string    `json:"title"`
	Body  *string    `json:"body"`
	Date  *time.Time `json:"date"`
}
