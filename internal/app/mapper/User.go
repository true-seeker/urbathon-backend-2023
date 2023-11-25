package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func UserModelToUserResponse(user *model.Users) *response.User {
	r := &response.User{
		Id:             user.ID,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Patronymic:     user.Patronymic,
		OrganizationId: user.OrganizationID,
		Phone:          user.PhoneNumber,
		Job:            user.Job,
	}

	return r
}

func UserModelToUserResponses(users *[]model.Users) *[]response.User {
	rs := make([]response.User, 0)

	for _, user := range *users {
		rs = append(rs, *UserModelToUserResponse(&user))
	}

	return &rs
}

func UserToUserResponse(user *entity.User) *response.User {
	r := &response.User{
		Id:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Patronymic:   user.Patronymic,
		Phone:        user.PhoneNumber,
		Job:          user.Job,
		Organization: OrganizationToOrganizationResponse(user.Organization),
	}

	return r
}

func UserToUserResponses(users *[]entity.User) *[]response.User {
	rs := make([]response.User, 0)

	for _, user := range *users {
		rs = append(rs, *UserToUserResponse(&user))
	}

	return &rs
}

func UserRegisterInputToUser(userInput *input.UserRegister) *model.Users {
	r := &model.Users{
		FirstName:   userInput.FirstName,
		LastName:    userInput.LastName,
		Patronymic:  userInput.Patronymic,
		Email:       userInput.Email,
		PhoneNumber: userInput.Phone,
	}
	return r
}
