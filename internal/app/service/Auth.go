package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/internal/pkg/passwordHash"
	"urbathon-backend-2023/pkg/errorHandler"
)

type UserRepository interface {
	GetByEmail(email *string) (*model.Users, error)
	Get(id *int32) (*model.Users, error)
	Create(userInput *model.Users) (*model.Users, error)
}
type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (d *AuthService) Login(loginInput *input.UserLogin) (*response.User, *errorHandler.HttpErr) {
	userResponse := &response.User{}
	if httpErr := d.validateLogin(loginInput); httpErr != nil {
		return nil, httpErr
	}

	user, err := d.userRepo.GetByEmail(loginInput.Email)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("Wrong email or password", http.StatusUnauthorized)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}

	if !checkPassword(loginInput.Password, user) {
		return nil, errorHandler.New("Wrong email or password", http.StatusUnauthorized)
	}
	userResponse = mapper.UserToUserResponse(user)

	return userResponse, nil
}

func (d *AuthService) Register(userRegister *input.UserRegister) (*response.User, *errorHandler.HttpErr) {
	userResponse := &response.User{}
	if httpErr := d.validateRegister(userRegister); httpErr != nil {
		return nil, httpErr
	}

	user := mapper.UserRegisterInputToUser(userRegister)
	user.Password, user.Salt = generateHashedPass(userRegister.Password)

	user, err := d.userRepo.Create(user)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	userResponse = mapper.UserToUserResponse(user)

	return userResponse, nil
}

func (d *AuthService) validateRegister(userRegister *input.UserRegister) *errorHandler.HttpErr {
	if httpErr := validator.UserRegistration(userRegister); httpErr != nil {
		return httpErr
	}
	user, err := d.userRepo.GetByEmail(userRegister.Email)
	if err != nil && !errors.Is(err, qrm.ErrNoRows) {
		return errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	if user != nil {
		return errorHandler.New("User with email already exists", http.StatusConflict)
	}
	return nil
}

func (d *AuthService) validateLogin(loginInput *input.UserLogin) *errorHandler.HttpErr {
	if httpErr := validator.UserLogin(loginInput); httpErr != nil {
		return httpErr
	}
	return nil
}

func generateHashedPass(userPassword *string) (*[]byte, *[]byte) {
	newSalt := passwordHash.GenerateRandomSalt()
	hashedPassword := passwordHash.HashPassword(*userPassword, newSalt)
	bytePassword := []byte(hashedPassword)
	return &bytePassword, &newSalt
}

func checkPassword(inputPassword *string, user *model.Users) bool {
	stringPass := string(*user.Password)
	return passwordHash.DoPasswordsMatch(stringPass, *inputPassword, *user.Salt)
}
