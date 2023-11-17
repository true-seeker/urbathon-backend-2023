package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func UserToUserResponse(user *model.Users) *response.User {
	r := &response.User{
		Id:    user.ID,
		Email: user.Email,
		Name:  user.Name,
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

func UserInputToUser(userInput *input.User) *model.Users {
	r := &model.Users{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	}
	return r
}
