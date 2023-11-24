package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type OrganizationService interface {
	Register(organizationInput *input.OrganizationRegister) (*response.Organization, *errorHandler.HttpErr)
	AddUser(organizationId int32, userId int32) *errorHandler.HttpErr
}

type OrganizationHandler struct {
	organizationService OrganizationService
}

func NewOrganizationHandler(organizationService OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{organizationService: organizationService}
}

// Register
// @Summary		register organization
// @Description	register organization
// @Accept			json
// @Tags			organization
// @Produce		json
// @Param			input	body		input.OrganizationRegister	true	"OrganizationRegister"
// @Success		201		{object}	response.User
// @Failure		400		{object}	errorHandler.HttpErr
// @Failure		409		{object}	errorHandler.HttpErr
// @Router			/organization [post]
func (d *OrganizationHandler) Register(c *gin.Context) {
	userInput := &input.OrganizationRegister{}
	err := c.BindJSON(&userInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	organization, httpErr := d.organizationService.Register(userInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	c.JSON(http.StatusCreated, organization)
}

// AddUser
// @Summary		add user to organization
// @Description	add user to organization
// @Tags			organization
// @Param			id	path		int	true	"organization id"
// @Param			user_id	path		int	true	"user id"
// @Success		200		{object}	nil
// @Failure		400		{object}	errorHandler.HttpErr
// @Failure		404		{object}	errorHandler.HttpErr
// @Router			/organization/{id}/add_user/{user_id} [post]
func (d *OrganizationHandler) AddUser(c *gin.Context) {
	organizationId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	userId, httpErr := validator.ValidateAndReturnId(c.Param("user_id"), "user_id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	if httpErr := d.organizationService.AddUser(organizationId, userId); httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	//todo exists validation

	c.Status(http.StatusOK)
}
