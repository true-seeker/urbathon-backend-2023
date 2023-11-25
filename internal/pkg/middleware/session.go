package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/repository"
	"urbathon-backend-2023/internal/app/storage"
	"urbathon-backend-2023/pkg/errorHandler"
)

func Session(c *gin.Context) {
	session := sessions.Default(c)
	userIdAny := session.Get("user_id")
	if userIdAny == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Next()
		return
	}
	userId := userIdAny.(int32)
	db := storage.CurrentStorage
	userRepo := repository.NewAuthRepository(db)
	user, err := userRepo.Get(&userId)
	if err != nil {
		httpErr := errorHandler.New("Something went wrong", http.StatusBadRequest)
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.Set("user", user)
	c.Next()
}

func OptionalSession(c *gin.Context) {
	session := sessions.Default(c)
	userIdAny := session.Get("user_id")
	if userIdAny == nil {
		c.Next()
		return
	}
	userId := userIdAny.(int32)
	db := storage.CurrentStorage
	userRepo := repository.NewAuthRepository(db)
	user, err := userRepo.Get(&userId)
	if err != nil {
		httpErr := errorHandler.New("Something went wrong", http.StatusBadRequest)
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}
	c.Set("user", user)
	c.Next()
}
