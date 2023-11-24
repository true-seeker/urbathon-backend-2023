package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type OrganizationService interface {
	Register(organizationInput *input.OrganizationRegister) (*response.Organization, *errorHandler.HttpErr)
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
