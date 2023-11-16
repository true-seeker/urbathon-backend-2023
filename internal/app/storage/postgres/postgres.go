package postgres

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"urbathon-backend-2023/pkg/config"
	"urbathon-backend-2023/pkg/projectpath"
)

type Postgres struct {
	db *sql.DB
}

func (s *Postgres) GetDb() *sql.DB {
	return s.db
}

func BuildPostgresConnectionString() string {
	var ConnectionString = fmt.Sprintf("host=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"port=%s "+
		"sslmode=disable "+
		"search_path=public "+
		"TimeZone=Asia/Yekaterinburg ",
		config.GetConfig().Get("database.address"),
		config.GetConfig().Get("database.user"),
		config.GetConfig().Get("database.password"),
		config.GetConfig().Get("database.dbname"),
		config.GetConfig().Get("database.port"))
	return ConnectionString
}
func (s *Postgres) New() *sql.DB {
	connStr := BuildPostgresConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	s.db = db
	return db
}

func (s *Postgres) GetMigrationDriver() *database.Driver {
	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	return &driver
}

func (s *Postgres) GetMigrationPath() string {
	return fmt.Sprintf("file://%s/internal/app/storage/postgres/migrations/", projectpath.Root)
}
