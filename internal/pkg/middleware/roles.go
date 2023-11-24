package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
)

func AdminRequired(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)
	if *user.Role != int32(3) {
		c.AbortWithStatusJSON(http.StatusForbidden, "Only admin can access this endpoint")
		return
	}

	c.Next()
}

func AdminOrChipperRequired(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)
	if *user.Role != int32(2) {
		c.AbortWithStatusJSON(http.StatusForbidden, "Only municipal service can access this endpoint")
		return
	}

	c.Next()
}
