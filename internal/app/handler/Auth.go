package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AuthService interface {
	Login(login *input.Login) (*response.User, *errorHandler.HttpErr)
	Create(userInput *input.User) (*response.User, *errorHandler.HttpErr)
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login godoc
// @Summary login
// @Description login
// @Accept json
// @Produce json
// @Param input body input.UserLogin true "login and password"
// @Success 200 {object} response.User
// @Router /auth/login [post]
func (d *AuthHandler) Login(c *gin.Context) {
	loginInput := &input.Login{}
	err := c.BindJSON(&loginInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, httpErr := d.authService.Login(loginInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.Id)
	err = session.Save()
	if err != nil {
		panic(err)
	}
	session.Get("user_id")
	c.JSON(http.StatusOK, user)
}

// Logout godoc
// @Summary logout
// @Description logout
// @Accept json
// @Produce json
// @Success 200
// @Router /auth/logout [post]
func (d *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("user_id")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Save()
	c.Status(http.StatusOK)
}

// Test godoc
// @Summary auth test
// @Description auth test
// @Accept json
// @Produce json
// @Success 200 {object} response.User
// @Router /auth/test [post]
func (d *AuthHandler) Test(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)

	c.JSON(http.StatusOK, user)
}

// Register godoc
// @Summary register
// @Description register
// @Accept json
// @Produce json
// @Param input body input.User true "User"
// @Success 200 {object} response.User
// @Router /auth/register [post]
func (d *AuthHandler) Register(c *gin.Context) {
	userInput := &input.User{}
	err := c.BindJSON(&userInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, httpErr := d.authService.Create(userInput)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.Id)
	err = session.Save()

	c.JSON(http.StatusCreated, user)
}
