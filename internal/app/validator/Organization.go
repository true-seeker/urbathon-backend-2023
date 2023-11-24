package validator

import (
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/pkg/errorHandler"
)

func OrganizationRegistration(organizationRegister *input.OrganizationRegister) *errorHandler.HttpErr {
	if IsStringEmpty(&organizationRegister.Name) {
		return errorHandler.New("name is empty", http.StatusBadRequest)
	}
	if IsStringEmpty(&organizationRegister.Inn) {
		return errorHandler.New("inn is empty", http.StatusBadRequest)
	}
	if IsStringEmpty(&organizationRegister.Address) {
		return errorHandler.New("address is empty", http.StatusBadRequest)
	}
	if IsStringEmpty(&organizationRegister.Phone) {
		return errorHandler.New("phone is empty", http.StatusBadRequest)
	}

	if organizationRegister.Categories == nil || len(*organizationRegister.Categories) == 0 {
		return errorHandler.New("categories are empty", http.StatusBadRequest)
	}
	return nil
}
