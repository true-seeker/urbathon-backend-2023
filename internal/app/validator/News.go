package validator

import (
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/pkg/errorHandler"
)

func NewsCreate(newsInput *input.News) *errorHandler.HttpErr {
	if IsStringEmpty(newsInput.Title) {
		return errorHandler.New("title is empty", http.StatusBadRequest)
	}

	if IsStringEmpty(newsInput.Body) {
		return errorHandler.New("description is empty", http.StatusBadRequest)
	}

	if newsInput.CategoryId == nil {
		return errorHandler.New("category_id is empty", http.StatusBadRequest)
	}
	return nil
}
