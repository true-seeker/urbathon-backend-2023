package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"urbathon-backend-2023/docs"
	"urbathon-backend-2023/internal/app/storage"
	"urbathon-backend-2023/internal/app/storage/postgres"
	"urbathon-backend-2023/internal/pkg/app"
	"urbathon-backend-2023/pkg/config"
)

// Init Инициализация сервиса
func Init() {
	config.Init("development")
	docs.SwaggerInfo.BasePath = "/api"
}

func InitDB() *postgres.Postgres {
	st := &postgres.Postgres{}
	st.New()
	storage.CurrentStorage = st
	return st
}

func main() {
	Init()
	st := InitDB()

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.Use(cors.Default())

	a := app.New()

	go func() {
		err := a.Run(r, st)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	log.Println("Closing DB connection")
	st.GetDb().Close()

	select {
	case <-ctx.Done():
	}
}
