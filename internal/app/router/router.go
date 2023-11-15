package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"urbathon-backend-2023/internal/app/handler"
	"urbathon-backend-2023/internal/app/repository"
	"urbathon-backend-2023/internal/app/service"
	"urbathon-backend-2023/internal/app/storage"
	"urbathon-backend-2023/internal/pkg/middleware"
)

// InitRoutes Инициализация путей эндпоинтов, сервисов и репозиториев
func InitRoutes(r *gin.Engine, storage storage.Sql) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api")
	databaseRepo := repository.NewDatabaseRepository(storage)
	databaseService := service.NewDatabaseService(databaseRepo)
	databaseHandler := handler.NewDatabaseHandler(databaseService)
	databasesGroup := api.Group("database")
	{
		databasesGroup.GET("/", middleware.BasicAuth, databaseHandler.GetAll)
		databasesGroup.GET("/:id", middleware.BasicAuth, databaseHandler.Get)
		databasesGroup.POST("/", middleware.BasicAuth, databaseHandler.Create)
		databasesGroup.DELETE("/:id", middleware.BasicAuth, databaseHandler.Delete)
		databasesGroup.PATCH("/:id", middleware.BasicAuth, databaseHandler.Edit)
	}
	incidentRepo := repository.NewIncidentRepository(storage, databaseRepo)

	alertRepo := repository.NewAlertRepository(storage, incidentRepo)
	alertService := service.NewAlertService(alertRepo)
	alertHandler := handler.NewAlertHandler(alertService)
	alertsGroup := api.Group("alert")
	{
		alertsGroup.GET("/", middleware.BasicAuth, alertHandler.GetAll)
		alertsGroup.GET("/:id", middleware.BasicAuth, alertHandler.Get)
	}

	userRepo := repository.NewUserRepository(storage)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	authGroup := api.Group("auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/logout", authHandler.Logout)
		authGroup.POST("/test", middleware.Session, authHandler.Test)
		authGroup.POST("/register", authHandler.Register)
	}
	return r
}
