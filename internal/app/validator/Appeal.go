package validator

import (
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/pkg/errorHandler"
)

func AppealCreate(appealInput *input.Appeal) *errorHandler.HttpErr {
	if IsStringEmpty(appealInput.Title) {
		return errorHandler.New("title is empty", http.StatusBadRequest)
	}

	if IsStringEmpty(appealInput.Description) {
		return errorHandler.New("description is empty", http.StatusBadRequest)
	}

	if IsStringEmpty(appealInput.Address) {
		return errorHandler.New("address is empty", http.StatusBadRequest)
	}

	if !IsLatitudeCorrect(appealInput.Latitude) {
		return errorHandler.New("latitude must be in [-90;90]", http.StatusBadRequest)
	}
	if !IsLatitudeCorrect(appealInput.Longitude) {
		return errorHandler.New("latitude must be in [-180;180]", http.StatusBadRequest)
	}
	return nil
}
