package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func UserToUserResponse(user *model.Users) *response.User {
	r := &response.User{
		Id:             user.ID,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Patronymic:     user.Patronymic,
		OrganizationId: user.OrganizationID,
	}

	return r
}

func UserToUserResponses(users *[]model.Users) *[]response.User {
	rs := make([]response.User, 0)

	for _, user := range *users {
		rs = append(rs, *UserToUserResponse(&user))
	}

	return &rs
}

func UserRegisterInputToUser(userInput *input.UserRegister) *model.Users {
	r := &model.Users{
		FirstName:  userInput.FirstName,
		LastName:   userInput.LastName,
		Patronymic: userInput.Patronymic,
		Email:      userInput.Email,
	}
	return r
}
