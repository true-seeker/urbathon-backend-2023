package mapper

import (
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/response"
)

func UserToUserResponse(user *entity.User) *response.User {
	r := &response.User{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
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
