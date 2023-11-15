package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"urbathon-backend-2023/pkg/config"
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
		"sslmode=disable TimeZone=Asia/Yekaterinburg",
		config.GetConfig().Get("database.address"),
		config.GetConfig().Get("database.user"),
		config.GetConfig().Get("database.password"),
		config.GetConfig().Get("database.dbname"),
		config.GetConfig().Get("database.port"))
	return ConnectionString
}

func (s *Postgres) New() *sql.DB {
	connStr := BuildPostgresConnectionString()
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	s.db = db
	s.dbInit()
	return db
}

func (s *Postgres) dbInit() {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS users(
							id             serial primary key,
							name           varchar(255),
							email          varchar(128),
							password       varchar(128)
												);`)
	if err != nil {
		panic(err)
	}
}
