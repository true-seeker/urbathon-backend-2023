package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/pkg/passwordHash"
	"urbathon-backend-2023/pkg/errorHandler"
)

type UserRepository interface {
	GetByEmail(loginInput *input.Login) (*model.Users, error)
	Get(id *int32) (*model.Users, error)
	Create(userInput *model.Users) (*model.Users, error)
}
type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (d *AuthService) Login(loginInput *input.Login) (*response.User, *errorHandler.HttpErr) {
	userResponse := &response.User{}
	user, err := d.userRepo.GetByEmail(loginInput)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, errorHandler.New("Wrong email or password", http.StatusForbidden)
	case err != nil:
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	if !checkPassword(loginInput.Password, user) {
		return nil, errorHandler.New("Wrong email or password", http.StatusForbidden)
	}
	userResponse = mapper.UserToUserResponse(user)

	return userResponse, nil
}

func (d *AuthService) Create(userInput *input.User) (*response.User, *errorHandler.HttpErr) {
	userResponse := &response.User{}
	user := mapper.UserInputToUser(userInput)
	user.Password, user.Salt = generateHashedPass(userInput.Password)

	user, err := d.userRepo.Create(user)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	userResponse = mapper.UserToUserResponse(user)

	return userResponse, nil
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
