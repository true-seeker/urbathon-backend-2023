package validator

import (
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/pkg/errorHandler"
)

func AppealCommentCreate(appealInput *input.AppealComment) *errorHandler.HttpErr {
	if IsStringEmpty(appealInput.Text) {
		return errorHandler.New("text is empty", http.StatusBadRequest)
	}
	return nil
}
