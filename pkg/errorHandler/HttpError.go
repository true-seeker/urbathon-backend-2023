package errorHandler

import (
	"fmt"
)

// HttpErr структура для формирования сообщения об ошибке
type HttpErr struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func New(errText string, statusCode int) *HttpErr {
	return &HttpErr{Message: errText, StatusCode: statusCode}
}

func (r *HttpErr) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Message)
}
