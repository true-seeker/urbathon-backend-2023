package storage

import "database/sql"

type Sql interface {
	GetDb() *sql.DB
	New() *sql.DB
}

var CurrentStorage Sql
