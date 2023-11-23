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
}

type AppealCategoryHandler struct {
	appealCategoryService AppealCategoryService
}

func NewAppealCategoryHandler(appealCategoryService AppealCategoryService) *AppealCategoryHandler {
	return &AppealCategoryHandler{appealCategoryService: appealCategoryService}
}

// Get get appealCategory by id
//
// @Summary		get appealCategory by id
// @Description	get appealCategory by id
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

// GetAll get all appealCategories
//
// @Summary		get all appealCategories
// @Description	get all appealCategories
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
