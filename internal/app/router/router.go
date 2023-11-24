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
	authRepo := repository.NewAuthRepository(storage)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	authGroup := api.Group("auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/logout", authHandler.Logout)
		authGroup.POST("/test", middleware.Session, authHandler.Test)
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/register_organization", authHandler.RegisterOrganization) // todo admin role
	}
	newsRepository := repository.NewNewsRepository(storage)
	newsService := service.NewNewsService(newsRepository)
	newsHandler := handler.NewNewsHandler(newsService)
	newsGroup := api.Group("news")
	{
		newsGroup.GET("/:id", newsHandler.Get)
		newsGroup.GET("/", newsHandler.GetAll)
	}

	appealRepo := repository.NewAppealRepository(storage)
	appealService := service.NewAppealService(appealRepo)
	appealHandler := handler.NewAppealHandler(appealService)
	appealCommentRepo := repository.NewAppealCommentRepository(storage)
	appealCommentService := service.NewAppealCommentService(appealCommentRepo)
	appealCommentHandler := handler.NewAppealCommentHandler(appealCommentService)
	appealGroup := api.Group("appeal")
	{
		appealGroup.GET("/", appealHandler.GetAll)
		appealGroup.GET("/:id", appealHandler.Get)
		appealGroup.POST("/", middleware.Session, appealHandler.Create)
		appealGroup.PUT("/:id", middleware.Session, appealHandler.Update)
		appealGroup.POST("/:id/status/:status_id", middleware.Session, appealHandler.UpdateStatus)
		appealGroup.GET("/:id/comment", middleware.Session, appealCommentHandler.GetComments)
		appealGroup.POST("/:id/comment", middleware.Session, appealCommentHandler.CreateComment)
		appealGroup.DELETE("/:id", middleware.Session, appealHandler.Delete)
	}

	appealCategoryRepo := repository.NewAppealCategoryRepository(storage)
	appealCategoryService := service.NewAppealCategoryService(appealCategoryRepo)
	appealCategoryHandler := handler.NewAppealCategoryHandler(appealCategoryService)
	appealCategoryGroup := api.Group("appeal_category")
	{
		appealCategoryGroup.GET("/", appealCategoryHandler.GetAll)
		appealCategoryGroup.GET("/:id", appealCategoryHandler.Get)
		appealCategoryGroup.GET("/:id/appeal_types", appealCategoryHandler.GetAppealTypes)
	}

	appealTypeRepo := repository.NewAppealTypeRepository(storage)
	appealTypeService := service.NewAppealTypeService(appealTypeRepo)
	appealTypeHandler := handler.NewAppealTypeHandler(appealTypeService)
	appealTypeGroup := api.Group("appeal_type")
	{
		appealTypeGroup.GET("/", appealTypeHandler.GetAll)
		appealTypeGroup.GET("/:id", appealTypeHandler.Get)
	}

	appealStatusRepo := repository.NewAppealStatusRepository(storage)
	appealStatusService := service.NewAppealStatusService(appealStatusRepo)
	appealStatusHandler := handler.NewAppealStatusHandler(appealStatusService)
	appealStatusGroup := api.Group("appeal_status")
	{
		appealStatusGroup.GET("/", appealStatusHandler.GetAll)
		appealStatusGroup.GET("/:id", appealStatusHandler.Get)
	}

	return r
}
