package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"urbathon-backend-2023/internal/app/router"
	"urbathon-backend-2023/internal/app/storage"
	"urbathon-backend-2023/pkg/config"
)

type App struct {
	router *gin.Engine
	srv    *http.Server
}

// New Конструктор приложения
func New() *App {
	return &App{}
}

func (a *App) InitServer() {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", config.GetConfig().GetString("server.address"),
			config.GetConfig().GetString("server.port")),
		Handler: a.router,
	}
	a.srv = srv
}

func (a *App) Run(r *gin.Engine, storage storage.Sql) error {
	a.router = router.InitRoutes(r, storage)
	a.InitServer()
	if err := a.srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
