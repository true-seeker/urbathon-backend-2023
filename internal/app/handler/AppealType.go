package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealTypeService interface {
	Get(id *int32) (*response.AppealType, *errorHandler.HttpErr)
	GetAll() (*[]response.AppealType, *errorHandler.HttpErr)
}

type AppealTypeHandler struct {
	appealTypeService AppealTypeService
}

func NewAppealTypeHandler(appealTypeService AppealTypeService) *AppealTypeHandler {
	return &AppealTypeHandler{appealTypeService: appealTypeService}
}

// Get get appealType by id
//
// @Summary		get appealType by id
// @Description	get appealType by id
// @Tags			appealType
// @Produce		json
// @Param			id	path		int	true	"appealType id"
// @Success		200	{object}	response.AppealType
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal_type/{id} [get]
func (d *AppealTypeHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	appealType, httpErr := d.appealTypeService.Get(&id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appealType)
}

// GetAll get all appealTypes
//
// @Summary		get all appealTypes
// @Description	get all appealTypes
// @Tags			appealType
// @Produce		json
// @Success		200	{object}	[]response.AppealType
// @Failure		400	{object}	errorHandler.HttpErr
// @Router			/appeal_type [get]
func (d *AppealTypeHandler) GetAll(c *gin.Context) {
	appealTypes, httpErr := d.appealTypeService.GetAll()
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appealTypes)
}
