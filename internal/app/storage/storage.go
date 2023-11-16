package storage

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database"
)

type Sql interface {
	GetDb() *sql.DB
	New() *sql.DB
	GetMigrationDriver() *database.Driver
	GetMigrationPath() string
}

var CurrentStorage Sql
