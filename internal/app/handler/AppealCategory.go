package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealCategoryService interface {
	Get(id *int32) (*response.AppealCategory, *errorHandler.HttpErr)
	GetAll() (*[]response.AppealCategory, *errorHandler.HttpErr)
	GetAppealTypes(id *int32) (*[]response.AppealTypeByCategory, *errorHandler.HttpErr)
}

type AppealCategoryHandler struct {
	appealCategoryService AppealCategoryService
}

func NewAppealCategoryHandler(appealCategoryService AppealCategoryService) *AppealCategoryHandler {
	return &AppealCategoryHandler{appealCategoryService: appealCategoryService}
}

// Get Получение категории обращения по айди
//
// @Summary		Получение категории обращения по айди
// @Description	Получение категории обращения по айди
// @Tags			appealCategory
// @Produce		json
// @Param			id	path		int	true	"appealCategory id"
// @Success		200	{object}	response.AppealCategory
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal_category/{id} [get]
func (d *AppealCategoryHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	appealCategory, httpErr := d.appealCategoryService.Get(&id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appealCategory)
}

// GetAll Получение всех категорий обращений
//
// @Summary		Получение всех категорий обращений
// @Description	Получение всех категорий обращений
// @Tags			appealCategory
// @Produce		json
// @Success		200	{object}	[]response.AppealCategory
// @Failure		400	{object}	errorHandler.HttpErr
// @Router			/appeal_category [get]
func (d *AppealCategoryHandler) GetAll(c *gin.Context) {
	appealCategories, httpErr := d.appealCategoryService.GetAll()
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appealCategories)
}

// GetAppealTypes Получение всех типов обращений по айди категории
//
// @Summary		Получение всех типов обращений по айди категории
// @Description	Получение всех типов обращений по айди категории
// @Tags			appealCategory
// @Produce		json
// @Param			id	path		int	true	"appealCategory id"
// @Success		200	{object}	[]response.AppealTypeByCategory
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal_category/{id}/appeal_types [get]
func (d *AppealCategoryHandler) GetAppealTypes(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	appealTypes, httpErr := d.appealCategoryService.GetAppealTypes(&id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appealTypes)
}
