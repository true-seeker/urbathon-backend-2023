package validator

import (
	"net/http"
	"net/mail"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/pkg/errorHandler"
)

func UserRegistration(userRegister *input.UserRegister) *errorHandler.HttpErr {
	if IsStringEmpty(*userRegister.Email) {
		return errorHandler.New("email is empty", http.StatusBadRequest)
	}
	_, err := mail.ParseAddress(*userRegister.Email)
	if err != nil {
		return errorHandler.New("email is invalid", http.StatusBadRequest)
	}

	if IsStringEmpty(*userRegister.Password) {
		return errorHandler.New("password is empty", http.StatusBadRequest)
	}

	if IsStringEmpty(*userRegister.Name) {
		return errorHandler.New("name is empty", http.StatusBadRequest)
	}
	return nil
}

func UserLogin(userLogin *input.UserLogin) *errorHandler.HttpErr {
	if IsStringEmpty(*userLogin.Email) {
		return errorHandler.New("email is empty", http.StatusBadRequest)
	}

	if IsStringEmpty(*userLogin.Password) {
		return errorHandler.New("password is empty", http.StatusBadRequest)
	}
	return nil
}
