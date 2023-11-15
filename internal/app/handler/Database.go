package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/model/response"
	"urbathon-backend-2023/internal/app/validator"
	"urbathon-backend-2023/pkg/errorHandler"
)

type DatabaseService interface {
	GetAll() (*[]response.Database, *errorHandler.HttpErr)
	Get(id int) (*response.Database, *errorHandler.HttpErr)
	Create(databaseInput *input.Database) (*response.Database, *errorHandler.HttpErr)
	Edit(databaseInput *input.Database, id int) (*response.Database, *errorHandler.HttpErr)
	Delete(id int) (*bool, *errorHandler.HttpErr)
}

type DatabaseHandler struct {
	databaseService DatabaseService
}

func NewDatabaseHandler(databaseService DatabaseService) *DatabaseHandler {
	return &DatabaseHandler{databaseService: databaseService}
}

func (d *DatabaseHandler) GetAll(c *gin.Context) {
	//id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	//if httpErr != nil {
	//	c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
	//	return
	//}
	database, httpErr := d.databaseService.GetAll()
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, database)
}

func (d *DatabaseHandler) Get(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	database, httpErr := d.databaseService.Get(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	c.JSON(http.StatusOK, database)
}

func (d *DatabaseHandler) Create(c *gin.Context) {
	databaseInput := &input.Database{}
	err := c.BindJSON(&databaseInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	database, httpErr := d.databaseService.Create(databaseInput)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusCreated, database)
}

func (d *DatabaseHandler) Delete(c *gin.Context) {
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}
	ok, httpErr := d.databaseService.Delete(id)
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	if *ok {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusBadRequest)
	}
}

func (d *DatabaseHandler) Edit(c *gin.Context) {
	databaseInput := &input.Database{}
	err := c.BindJSON(&databaseInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	id, httpErr := validator.ValidateAndReturnId(c.Param("id"), "id")
	if httpErr != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	database, httpErr := d.databaseService.Edit(databaseInput, id)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.StatusCode, httpErr.Err.Error())
		return
	}

	c.JSON(http.StatusCreated, database)
}
