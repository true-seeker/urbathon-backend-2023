package service

import (
	"net/http"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type UserRepository interface {
	GetByCreds(loginInput *input.Login) (*entity.User, error)
	Get(id *int) (*entity.User, error)
	Create(userInput *input.User) (*entity.User, error)
}
type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (d *AuthService) Login(loginInput *input.Login) (*response.User, *errorHandler.HttpErr) {
	userResponse := &response.User{}

	user, err := d.userRepo.GetByCreds(loginInput)
	if user == nil {
		return nil, errorHandler.New("Wrong email or password", http.StatusForbidden)
	}
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	userResponse = mapper.UserToUserResponse(user)

	return userResponse, nil
}

func (d *AuthService) Create(userInput *input.User) (*response.User, *errorHandler.HttpErr) {
	userResponse := &response.User{}

	user, err := d.userRepo.Create(userInput)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	userResponse = mapper.UserToUserResponse(user)

	return userResponse, nil
}
