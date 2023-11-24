package mapper

import (
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
)

func OrganizationToOrganizationResponse(organization *model.Organizations) *response.Organization {
	r := &response.Organization{
		Id:      organization.ID,
		Name:    *organization.Name,
		Inn:     *organization.Address,
		Address: *organization.Address,
		Phone:   *organization.Phone,
	}

	return r
}

func OrganizationsToOrganizationResponses(organizations *[]model.Organizations) *[]response.Organization {
	rs := make([]response.Organization, 0)

	for _, organization := range *organizations {
		rs = append(rs, *OrganizationToOrganizationResponse(&organization))
	}

	return &rs
}

func OrganizationRegisterInputToOrganization(organizationInput *input.OrganizationRegister) *model.Organizations {
	r := &model.Organizations{
		Name:    &organizationInput.Name,
		Inn:     &organizationInput.Address,
		Address: &organizationInput.Address,
		Phone:   &organizationInput.Phone,
	}
	return r
}
