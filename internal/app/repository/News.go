package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/storage"
)

type NewsRepository struct {
	db *sql.DB
}

func NewNewsRepository(s storage.Sql) *NewsRepository {
	return &NewsRepository{db: s.GetDb()}
}

func getSelectNewStmt() SelectStatement {
	return SELECT(News.AllColumns,
		NewsCategories.ID.AS("newsCategories.id"),
		NewsCategories.Title.AS("newsCategories.title")).
		FROM(News.
			LEFT_JOIN(NewsCategories, NewsCategories.ID.EQ(News.CategoryID)))
}

func (a *NewsRepository) Get(id *int32) (*entity.News, error) {
	var u entity.News
	stmt := getSelectNewStmt().
		WHERE(News.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *NewsRepository) GetAll(f *filter.Pagination) (*[]entity.News, error) {
	var u []entity.News
	stmt := getSelectNewStmt().
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

func (a *NewsRepository) Create(news *model.News) (*entity.News, error) {
	var u *entity.News

	stmt := News.INSERT(News.AllColumns.Except(News.ID, News.Date)).
		MODEL(news).
		RETURNING(News.ID)

	if err := stmt.Query(a.db, news); err != nil {
		return nil, err
	}

	u, err := a.Get(&news.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}
