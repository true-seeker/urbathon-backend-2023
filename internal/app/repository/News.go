package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/storage"
)

type NewsRepository struct {
	db *sql.DB
}

func NewNewsRepository(s storage.Sql) *NewsRepository {
	return &NewsRepository{db: s.GetDb()}
}

func (a *NewsRepository) Get(id *int32) (*model.News, error) {
	var u model.News
	stmt := SELECT(News.AllColumns).
		FROM(News).
		WHERE(News.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *NewsRepository) GetAll(f *input.Filter) (*[]model.News, error) {
	var u []model.News
	stmt := SELECT(News.AllColumns).
		FROM(News).
		LIMIT(f.PageSize).
		OFFSET((f.Page - 1) * f.PageSize).
		ORDER_BY(News.Date.DESC())
	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *NewsRepository) GetTotal() (*int, error) {
	var count int
	rawSql, _ := SELECT(Raw("count(*)")).
		FROM(News).Sql()

	if err := a.db.QueryRow(rawSql).Scan(&count); err != nil {
		return nil, err
	}
	return &count, nil
}
