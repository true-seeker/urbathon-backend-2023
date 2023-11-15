package repository

import (
	"database/sql"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/storage"
)

type IncidentRepository struct {
	db                 *sql.DB
	databaseReposotiry *DatabaseRepository
}

func NewIncidentRepository(s storage.Sql, databaseRepository *DatabaseRepository) *IncidentRepository {
	return &IncidentRepository{db: s.GetDb(), databaseReposotiry: databaseRepository}
}

func (a *IncidentRepository) Get(id int) (*entity.Incident, error) {
	row := a.db.QueryRow("select id, db_id, error, date from incidents WHERE id = $1", id)

	var d entity.Incident
	var dbId int
	switch err := row.Scan(&d.Id, &dbId, &d.Error, &d.Date); err {
	case sql.ErrNoRows:
		return nil, err
	}
	d.Db, _ = a.databaseReposotiry.Get(dbId)
	return &d, nil
}
