package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type NewsService interface {
	Get(id *int32) (*response.News, *errorHandler.HttpErr)
	GetAll(filter *input.Filter) (*response.NewsPaged, *errorHandler.HttpErr)
}

type NewsHandler struct {
	newsService NewsService
}

func NewNewsHandler(newsService NewsService) *NewsHandler {
	return &NewsHandler{newsService: newsService}
}

// Get get news by id
//
//	@Summary		get news by id
//	@Description	get news by id
//	@Tags			news
//	@Produce		json
//	@Param			id	path		int	true	"news id"
//	@Success		200	{object}	response.News
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Router			/news/{id} [get]
func (d *NewsHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	news, httpErr := d.newsService.Get(&id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, news)
}

// GetAll get all news
//
//	@Summary		get all news
//	@Description	get all news
//	@Tags			news
//	@Param			page		query	int	false	"page"	minimum(1)	default(1)
//	@Param			page_size	query	int	false	"page"	minimum(1)	maximum(20)	default(10)
//	@Produce		json
//	@Success		200	{object}	response.NewsPaged
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Router			/news [get]
func (d *NewsHandler) GetAll(c *gin.Context) {
	f, httpErr := validator.ValidateQueryFilter(c)

	news, httpErr := d.newsService.GetAll(f)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, news)
}

func (d *NewsHandler) Create(c *gin.Context) {
	panic("not implemented")
}

func (d *NewsHandler) Update(c *gin.Context) {
	panic("not implemented")
}

func (d *NewsHandler) Delete(c *gin.Context) {
	panic("not implemented")
}
