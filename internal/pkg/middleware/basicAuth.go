package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/repository"
	"urbathon-backend-2023/internal/app/storage"
)

func DecodeCredentials(c *gin.Context) (string, string, bool) {
	r := c.Request
	return r.BasicAuth()
}

func GetAccountByCreds(c *gin.Context) (*entity.User, error) {
	login, password, ok := DecodeCredentials(c)
	if !ok {
		return nil, errors.New("")
	}

	loginInput := &input.Login{
		Email:    &login,
		Password: &password,
	}
	st := storage.CurrentStorage
	userRepo := repository.NewUserRepository(st)
	user, _ := userRepo.GetByCreds(loginInput)
	return user, nil
}

// BasicAuth middleware для basic auth
func BasicAuth(c *gin.Context) {
	user, err := GetAccountByCreds(c)

	if err != nil || user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Next()
		return
	}

	c.Set("user", user)
	c.Next()
}
