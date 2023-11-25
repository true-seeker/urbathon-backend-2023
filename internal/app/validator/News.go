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
		return errorHandler.New("body is empty", http.StatusBadRequest)
	}

	if IsStringEmpty(newsInput.Address) {
		return errorHandler.New("address is empty", http.StatusBadRequest)
	}

	if !IsLatitudeCorrect(newsInput.Latitude) {
		return errorHandler.New("latitude must be in [-90;90]", http.StatusBadRequest)
	}

	if !IsLongitudeCorrect(newsInput.Longitude) {
		return errorHandler.New("latitude must be in [-180;180]", http.StatusBadRequest)
	}

	if newsInput.CategoryId == nil {
		return errorHandler.New("category_id is empty", http.StatusBadRequest)
	}
	return nil
}
