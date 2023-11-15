package errorHandler

import (
	"errors"
	"fmt"
)

// HttpErr структура для формирования сообщения об ошибке
type HttpErr struct {
	Err        error
	StatusCode int
}

func New(errText string, statusCode int) *HttpErr {
	return &HttpErr{Err: errors.New(errText), StatusCode: statusCode}
}

func (r *HttpErr) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
