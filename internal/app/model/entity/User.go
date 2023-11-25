package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type User struct {
	*model.Users

	Organization *Organization
}
