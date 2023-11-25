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

// Get Получить тип обращения по айди
//
// @Summary		Получить тип обращения по айди
// @Description	Получить тип обращения по айди
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

// GetAll Получить все типы обращений
//
// @Summary		Получить все типы обращений
// @Description	Получить все типы обращений
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
