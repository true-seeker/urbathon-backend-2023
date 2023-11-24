package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AppealService interface {
	Get(id *int32) (*response.Appeal, *errorHandler.HttpErr)
	GetAll(filter *input.Filter) (*response.AppealPaged, *errorHandler.HttpErr)
	Create(appealInput *input.Appeal, user *model.Users) (*response.Appeal, *errorHandler.HttpErr)
	Update(appealInput *input.AppealUpdate, user *model.Users, id *int32) (*response.Appeal, *errorHandler.HttpErr)
	Delete(id int32) *errorHandler.HttpErr
	UpdateStatus(appealId int32, statusId int32) *errorHandler.HttpErr
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

// GetAll get all appeals
//
// @Summary		get all appeals
// @Description	get all appeals
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

// Create
//
// @Summary		create appeal
// @Description	create appeal
// @Tags			appeal
// @Param			appeal	formData	input.Appeal	true	"appeal"
// @Param			photos	formData	[]file			true	"photos"
// @Accept			mpfd
// @Produce		json
// @Success		201	{object}	response.Appeal
// @Failure		400	{object}	errorHandler.HttpErr
// @Router			/appeal [post]
func (d *AppealHandler) Create(c *gin.Context) {
	appealInput := &input.Appeal{}
	if err := c.Bind(&appealInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)

	appeal, httpErr := d.appealService.Create(appealInput, user)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	c.JSON(http.StatusCreated, appeal)
}

// Update
//
// @Summary		update appeal
// @Description	update appeal
// @Tags			appeal
// @Param			appeal		body	input.AppealUpdate	true	"appeal"
// @Param			id	path	int				true	"appeal id"	default(1)
// @Produce		json
// @Success		200	{object}	response.Appeal
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal/{id} [put]
func (d *AppealHandler) Update(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	appealInput := &input.AppealUpdate{}
	if err := c.BindJSON(&appealInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)
	// todo валидация изменений только своих обращений

	appeal, httpErr := d.appealService.Update(appealInput, user, &id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	c.JSON(http.StatusOK, appeal)
}

// Delete
//
// @Summary		delete appeal
// @Description	delete appeal
// @Tags			appeal
// @Param			id	path	int	true	"appeal id"	default(1)
// @Produce		json
// @Success		200	{object}	nil
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal/{id} [delete]
func (d *AppealHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	if httpErr = d.appealService.Delete(id); httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	c.Status(http.StatusOK)
}

// UpdateStatus
//
// @Summary		update appeal status
// @Description	update appeal status
// @Tags			appeal
// @Param			id	path	int	true	"appeal id"	default(1)
// @Param			status_id	path	int	true	"new status id"	default(1)
// @Produce		json
// @Success		200	{object}	nil
// @Failure		400	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal/{id}/status/{status_id} [post]
func (d *AppealHandler) UpdateStatus(c *gin.Context) {
	appealId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	statusId, httpErr := validator.ValidateAndReturnId(c.Param("status_id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	// todo exists validation
	if httpErr = d.appealService.UpdateStatus(appealId, statusId); httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	c.Status(http.StatusOK)
}
