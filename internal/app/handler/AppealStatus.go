package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealStatusService interface {
	Get(id *int32) (*response.AppealStatus, *errorHandler.HttpErr)
	GetAll() (*[]response.AppealStatus, *errorHandler.HttpErr)
}

type AppealStatusHandler struct {
	appealStatusService AppealStatusService
}

func NewAppealStatusHandler(appealStatusService AppealStatusService) *AppealStatusHandler {
	return &AppealStatusHandler{appealStatusService: appealStatusService}
}

// Get Получить статус обращения
//
// @Summary		Получить статус обращения
// @Description	Получить статус обращения
// @Tags			appealStatus
// @Produce		json
// @Param			id	path		int	true	"appealStatus id" default(1)
// @Success		200	{object}	response.AppealStatus
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal_status/{id} [get]
func (d *AppealStatusHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	appealStatus, httpErr := d.appealStatusService.Get(&id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appealStatus)
}

// GetAll Получить все статусы обращений
//
// @Summary		Получить все статусы обращений
// @Description	Получить все статусы обращений
// @Tags			appealStatus
// @Produce		json
// @Success		200	{object}	[]response.AppealStatus
// @Failure		400	{object}	errorHandler.HttpErr
// @Router			/appeal_status [get]
func (d *AppealStatusHandler) GetAll(c *gin.Context) {
	appealCategories, httpErr := d.appealStatusService.GetAll()
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appealCategories)
}
