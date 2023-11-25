package repository

import (
	"context"
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
		NewsCategories.Title.AS("newsCategories.title"),
		NewsPolls.ID.AS("newsPolls.id"),
		NewsPolls.Title.AS("newsPolls.title"),
		PollOptions.ID.AS("pollOptions.id"),
		PollOptions.Title.AS("pollOptions.title")).
		FROM(News.
			LEFT_JOIN(NewsCategories, NewsCategories.ID.EQ(News.CategoryID)).
			LEFT_JOIN(NewsPolls, NewsPolls.ID.EQ(News.PollID)).
			LEFT_JOIN(PollOptions, PollOptions.PollID.EQ(NewsPolls.ID)))
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
	stmt := getSelectNewStmt()
	stmt = f.GetLimitOffsetStmt(stmt).
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

func (a *NewsRepository) Create(news *model.News, poll entity.NewsPoll) (*entity.News, error) {
	var u *entity.News
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	stmt := NewsPolls.
		INSERT(NewsPolls.Title).
		MODEL(poll.NewsPolls).
		RETURNING(NewsPolls.ID)
	if err := stmt.QueryContext(ctx, tx, &poll.NewsPolls); err != nil {
		return nil, err
	}

	for _, option := range *poll.Options {
		option.PollID = &poll.NewsPolls.ID
		stmt = PollOptions.
			INSERT(PollOptions.Title, PollOptions.PollID).
			MODEL(option).
			RETURNING(PollOptions.ID)
		if err := stmt.QueryContext(ctx, tx, &option); err != nil {
			return nil, err
		}
	}

	news.PollID = &poll.NewsPolls.ID
	stmt = News.INSERT(News.AllColumns.Except(News.ID, News.Date)).
		MODEL(news).
		RETURNING(News.ID)

	if err := stmt.QueryContext(ctx, tx, news); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	u, err = a.Get(&news.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}
