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
	AddUser(organizationId int32, orgUserInout *input.OrganizationAddUser) error
}
type OrganizationService struct {
	organizationRepo OrganizationRepository
}

func NewOrganizationService(organizationRepository OrganizationRepository) *OrganizationService {
	return &OrganizationService{organizationRepo: organizationRepository}
}

func (d *OrganizationService) Register(organizationInput *input.OrganizationRegister) (*response.Organization, *errorHandler.HttpErr) {
	if httpErr := d.validateRegisterOrganization(organizationInput); httpErr != nil {
		return nil, httpErr
	}

	organization := mapper.OrganizationRegisterInputToOrganization(organizationInput)

	organization, err := d.organizationRepo.Register(organization, organizationInput.Categories)
	if err != nil {
		return nil, errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	organizationResponse := mapper.OrganizationModelToOrganizationResponse(organization)

	return organizationResponse, nil
}

func (d *OrganizationService) AddUser(organizationId int32, orgUserInout *input.OrganizationAddUser) *errorHandler.HttpErr {
	if err := d.organizationRepo.AddUser(organizationId, orgUserInout); err != nil {
		return errorHandler.New(err.Error(), http.StatusBadRequest)
	}
	return nil
}

func (d *OrganizationService) validateRegisterOrganization(organizationRegister *input.OrganizationRegister) *errorHandler.HttpErr {
	if httpErr := validator.OrganizationRegistration(organizationRegister); httpErr != nil {
		return httpErr
	}
	return nil
}
