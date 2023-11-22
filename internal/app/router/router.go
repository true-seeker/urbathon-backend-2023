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
	newsRepository := repository.NewNewsRepository(storage)
	newsService := service.NewNewsService(newsRepository)
	newsHandler := handler.NewNewsHandler(newsService)
	newsGroup := api.Group("news")
	{
		newsGroup.GET("/:id", newsHandler.Get)
		newsGroup.GET("/", newsHandler.GetAll)
		//newsGroup.POST("/", middleware.Session, newsHandler.Create)
		//newsGroup.PUT("/:id", middleware.Session, newsHandler.Update)
		//newsGroup.DELETE("/:id", middleware.Session, newsHandler.Delete)
	}

	return r
}
