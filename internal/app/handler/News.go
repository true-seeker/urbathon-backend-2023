package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type NewsService interface {
	Get(id *int32) (*response.News, *errorHandler.HttpErr)
	GetAll(filter *filter.Pagination) (*response.NewsPaged, *errorHandler.HttpErr)
	Create(newsInput *input.News, user *model.Users) (*response.News, *errorHandler.HttpErr)
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
	var p filter.Pagination
	_ = c.ShouldBindQuery(&p)

	p2, httpErr := filter.ValidatePagination(&p)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	news, httpErr := d.newsService.GetAll(p2)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, news)
}

// Create
//
// @Summary		create news
// @Description	create news
// @Tags			news
// @Param			news	formData	input.News	true	"news"
// @Param			photo	formData	file		false	"photo"
// @Accept			mpfd
// @Produce		json
// @Success		201	{object}	response.News
// @Failure		400	{object}	errorHandler.HttpErr
// @Router			/news [post]
func (d *NewsHandler) Create(c *gin.Context) {
	newsInput := &input.News{}
	if err := c.Bind(&newsInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)

	appeal, httpErr := d.newsService.Create(newsInput, user)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	c.JSON(http.StatusCreated, appeal)
}
