package main

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"os/signal"
	"syscall"
	"urbathon-backend-2023/docs"
	"urbathon-backend-2023/internal/app/s3"
	"urbathon-backend-2023/internal/app/storage"
	"urbathon-backend-2023/internal/app/storage/postgres"
	rd "urbathon-backend-2023/internal/app/store/redis"
	"urbathon-backend-2023/internal/pkg/app"
	"urbathon-backend-2023/pkg/config"
)

func Init() {
	config.Init("development")
	docs.SwaggerInfo.BasePath = "/api"
	s3.Init()
}

func initDB() *postgres.Postgres {
	st := &postgres.Postgres{}
	st.New()
	storage.CurrentStorage = st
	return st
}

func initCache(a *app.App) *cookie.Store {
	var store cookie.Store
	if a.IsRunningInDockerContainer() {
		r := rd.New()
		store, _ = redis.NewStore(10, "tcp", fmt.Sprintf("%s:%s", r.Address, r.Port), "", []byte("secret"))
	} else {
		store = cookie.NewStore([]byte("secret"))
	}
	return &store
}

func initGin(store *cookie.Store) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers"},
	}))
	r.Use(sessions.Sessions("session", *store))
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

// @title			Urbathon-2023
// @description	Спецификация приложения команды подCRUDули
func main() {
	Init()
	st := initDB()
	a := app.New()
	store := initCache(a)
	r := initGin(store)

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
