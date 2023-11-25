package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type MapService interface {
	GetMapElements(f *filter.Map) (*[]response.MapElement, *errorHandler.HttpErr)
}

type MapHandler struct {
	mapService MapService
}

func NewMapHandler(mapService MapService) *MapHandler {
	return &MapHandler{mapService: mapService}
}

// GetMapElements get map elements
//
// @Summary		get map elements
// @Description	get map elements
// @Tags			map
// @Param			page		query	filter.Map	false	"page"
// @Produce		json
// @Success		200	{object}	[]response.MapElement
// @Failure		400	{object}	errorHandler.HttpErr
// @Router			/map/get_map_elements [get]
func (d *MapHandler) GetMapElements(c *gin.Context) {
	f, httpErr := filter.NewMapFilter(c)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	appeal, httpErr := d.mapService.GetMapElements(f)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.JSON(http.StatusOK, appeal)
}
