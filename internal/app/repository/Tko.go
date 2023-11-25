package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/storage"
)

type TkoRepository struct {
	db *sql.DB
}

func NewTkoRepository(s storage.Sql) *TkoRepository {
	return &TkoRepository{db: s.GetDb()}
}

func (t *TkoRepository) GetForMap(f *filter.Map) (*[]model.Tko, error) {
	var tko []model.Tko
	stmt :=
		Tko.SELECT(Tko.AllColumns)
	stmt = makeWhere(f, stmt).LIMIT(100)
	if err := stmt.Query(t.db, &tko); err != nil {
		return nil, err
	}
	return &tko, nil
}

func makeWhere(f *filter.Map, stmt SelectStatement) SelectStatement {
	if f.LatUp != nil && f.LatDown != nil && f.LongDown != nil && f.LongUp != nil {
		stmt = stmt.WHERE(Tko.Longitude.GT(Float(*f.LongUp)).
			AND(Tko.Longitude.LT(Float(*f.LongDown))).
			AND(Tko.Latitude.GT(Float(*f.LatDown))).
			AND(Tko.Latitude.LT(Float(*f.LatUp))))
	}
	return stmt
}
