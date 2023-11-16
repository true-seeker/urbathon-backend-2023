package main

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"os/signal"
	"syscall"
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

func initGin() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.Use(cors.Default())
	return r
}

func runMigrations(sql storage.Sql) {
	m, err := migrate.NewWithDatabaseInstance(
		sql.GetMigrationPath(),
		"postgres",
		*sql.GetMigrationDriver())
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}

func main() {
	Init()
	st := InitDB()
	r := initGin()
	a := app.New()

	runMigrations(st)

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

	log.Println("Closing DB connection")
	st.GetDb().Close()

	log.Println("Stopping http server")
}
