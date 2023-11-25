package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	"urbathon-backend-2023/internal/app/mapper"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/pkg/errorHandler"
)

type AuthService interface {
	Login(login *input.UserLogin) (*response.User, *errorHandler.HttpErr)
	Register(userInput *input.UserRegister) (*response.User, *errorHandler.HttpErr)
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login Вход
// @Summary		Вход
// @Description	Вход
// @Accept			json
// @Tags			auth
// @Produce		json
// @Param			input	body		input.UserLogin	true	"login and password"
// @Success		200		{object}	response.User
// @Failure		400		{object}	errorHandler.HttpErr
// @Failure		401		{object}	errorHandler.HttpErr
// @Router			/auth/login [post]
func (d *AuthHandler) Login(c *gin.Context) {
	loginInput := &input.UserLogin{}
	err := c.BindJSON(&loginInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	user, httpErr := d.authService.Login(loginInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
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

// Logout Выход
// @Summary		Выход
// @Description	Выход
// @Tags			auth
// @Success		200
// @Router			/auth/logout [post]
func (d *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Get("user_id")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Save()
	c.Status(http.StatusOK)
}

// Test
// @Summary		auth test
// @Description	auth test
// @Accept			json
// @Tags			auth
// @Produce		json
// @Success		200	{object}	response.User
// @Router			/auth/test [post]
func (d *AuthHandler) Test(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(*model.Users)
	userResponse := mapper.UserModelToUserResponse(user)

	c.JSON(http.StatusOK, userResponse)
}

// Register Регистрация
// @Summary		Регистрация
// @Description	Регистрация
// @Accept			json
// @Tags			auth
// @Produce		json
// @Param			input	body		input.UserRegister	true	"UserRegister"
// @Success		201		{object}	response.User
// @Failure		400		{object}	errorHandler.HttpErr
// @Failure		409		{object}	errorHandler.HttpErr
// @Router			/auth/register [post]
func (d *AuthHandler) Register(c *gin.Context) {
	userInput := &input.UserRegister{}
	err := c.BindJSON(&userInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, httpErr := d.authService.Register(userInput)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr)
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.Id)
	err = session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}
