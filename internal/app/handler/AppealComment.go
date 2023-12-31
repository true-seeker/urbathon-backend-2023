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

type AppealCommentService interface {
	GetAllByAppealId(f *filter.Pagination, appealId int32) (*response.AppealCommentPaged, *errorHandler.HttpErr)
	Create(commentInput *input.AppealComment, user *model.Users, appealId *int32) (*response.AppealComment, *errorHandler.HttpErr)
}

type AppealCommentHandler struct {
	appealCommentService AppealCommentService
}

func NewAppealCommentHandler(appealCommentService AppealCommentService) *AppealCommentHandler {
	return &AppealCommentHandler{appealCommentService: appealCommentService}
}

// GetComments Получение комментариев обращения
//
//	@Summary		Получение комментариев обращения
//	@Description	Получение комментариев обращения
//	@Tags			appealComment
//	@Param			id			path	int	true	"appeal id"	default(1)
//	@Param			page		query	int	false	"page"		minimum(1)	default(1)
//	@Param			page_size	query	int	false	"page"		minimum(1)	maximum(20)	default(10)
//	@Produce		json
//	@Success		200	{object}	response.AppealCommentPaged
//	@Failure		400	{object}	errorHandler.HttpErr
//	@Failure		404	{object}	errorHandler.HttpErr
//	@Router			/appeal/{id}/comment [get]
func (d *AppealCommentHandler) GetComments(c *gin.Context) {
	var p filter.Pagination
	_ = c.ShouldBindQuery(&p)

	p2, httpErr := filter.ValidatePagination(&p)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	appealId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	// todo exists validation

	commentsResponse, httpErr := d.appealCommentService.GetAllByAppealId(p2, appealId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, commentsResponse)
}

// CreateComment Создание комментария
//
// @Summary		Создание комментария
// @Description	Создание комментария
// @Tags			appealComment
// @Param			id		path		int					true	"appeal id"	default(1)
// @Param			comment	formData	input.AppealComment	true	"appeal"
// @Param			photos	formData	[]file				false	"photos"
// @Accept			mpfd
// @Produce		json
// @Success		201	{object}	response.AppealComment
// @Failure		401	{object}	errorHandler.HttpErr
// @Failure		404	{object}	errorHandler.HttpErr
// @Router			/appeal/{id}/comment [post]
func (d *AppealCommentHandler) CreateComment(c *gin.Context) {
	appealCommentInput := &input.AppealComment{}

	appealId, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	if err := c.Bind(&appealCommentInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)

	appeal, httpErr := d.appealCommentService.Create(appealCommentInput, user, &appealId)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	c.JSON(http.StatusCreated, appeal)
}
