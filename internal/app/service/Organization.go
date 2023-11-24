package service

import (
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type OrganizationRepository interface {
	Register(organization *model.Organizations, organizationInputCategories *[]int32) (*model.Organizations, error)
}
type OrganizationService struct {
	organizationRepo OrganizationRepository
}

func NewOrganizationService(organizationRepository OrganizationRepository) *OrganizationService {
	return &OrganizationService{organizationRepo: organizationRepository}
}

func (d *OrganizationService) Register(organizationInput *input.OrganizationRegister) (*response.Organization, *errorHandler.HttpErr) {
	organizationResponse := &response.Organization{}
	if httpErr := d.validateRegisterOrganization(organizationInput); httpErr != nil {
		return nil, httpErr
	}

	organization := mapper.OrganizationRegisterInputToOrganization(organizationInput)

	organization, err := d.organizationRepo.Register(organization, organizationInput.Categories)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	organizationResponse = mapper.OrganizationToOrganizationResponse(organization)

	return organizationResponse, nil
}

func (d *OrganizationService) validateRegisterOrganization(organizationRegister *input.OrganizationRegister) *errorHandler.HttpErr {
	if httpErr := validator.OrganizationRegistration(organizationRegister); httpErr != nil {
		return httpErr
	}
	return nil
}
