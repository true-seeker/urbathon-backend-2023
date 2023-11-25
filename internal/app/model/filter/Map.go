package filter

import (
	"github.com/gin-gonic/gin"
	"urbathon-backend-2023/pkg/errorHandler"
)

type Map struct {
	LatUp    *float64 `form:"lat_up" `
	LatDown  *float64 `form:"lat_down" `
	LongUp   *float64 `form:"long_up" `
	LongDown *float64 `form:"long_down" `
}

func NewMapFilter(c *gin.Context) (*Map, *errorHandler.HttpErr) {
	p := Map{}

	_ = c.ShouldBindQuery(&p)

	return &p, nil
}
