package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealService interface {
	Get(id *int32) (*response.Appeal, *errorHandler.HttpErr)
	GetAll(filter *input.Filter) (*response.AppealPaged, *errorHandler.HttpErr)
}

type AppealHandler struct {
	appealService AppealService
}

func NewAppealHandler(appealService AppealService) *AppealHandler {
	return &AppealHandler{appealService: appealService}
}

// Get get appeal by id
//
// @Summary		get appeal by id
// @Description	get appeal by id
// @Tags			appeal
// @Produce		json
// @Param			id	path		int	true	"appeal id"
// @Success		200	{object}	response.Appeal
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal/{id} [get]
func (d *AppealHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	appeal, httpErr := d.appealService.Get(&id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appeal)
}

//	GetAll get all appeal
//
// @Summary		get all appeal
// @Description	get all appeal
// @Tags			appeal
// @Param			page		query	int	false	"page"	minimum(1)	default(1)
// @Param			page_size	query	int	false	"page"	minimum(1)	maximum(20)	default(10)
// @Produce		json
// @Success		200	{object}	response.AppealPaged
// @Failure		400	{object}	errorHandler.HttpErr
// @Router			/appeal [get]
func (d *AppealHandler) GetAll(c *gin.Context) {
	f, httpErr := validator.ValidateQueryFilter(c)

	appeal, httpErr := d.appealService.GetAll(f)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appeal)
}

func (d *AppealHandler) Create(c *gin.Context) {
	panic("not implemented")
}

func (d *AppealHandler) Update(c *gin.Context) {
	panic("not implemented")
}

func (d *AppealHandler) Delete(c *gin.Context) {
	panic("not implemented")
}
