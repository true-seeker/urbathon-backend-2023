package entity

import "urbathon-backend-2023/.gen/urbathon/public/model"

type Organization struct {
	*model.Organizations

	OrganizationUsers *[]model.Users
}
