package repository

import (
	"context"
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/storage"
)

type AppealRepository struct {
	db *sql.DB
}

func NewAppealRepository(s storage.Sql) *AppealRepository {
	return &AppealRepository{db: s.GetDb()}
}

var selectAppealStmt = SELECT(Appeals.AllColumns,
	Users.ID.AS("users.id"),
	Users.Name.AS("users.name"),
	Users.Email.AS("users.email"),
	AppealTypes.ID.AS("appealTypes.id"),
	AppealTypes.Title.AS("appealTypes.title"),
	AppealCategories.ID.AS("appealCategories.id"),
	AppealCategories.Title.AS("appealCategories.title"),
	AppealPhotos.ID.AS("appealPhotos.id"),
	AppealPhotos.URL.AS("appealPhotos.url"),
	AppealStatus.ID.AS("appealStatus.id"),
	AppealStatus.Status.AS("appealStatus.status"),
).FROM(Appeals.
	INNER_JOIN(Users, Users.ID.EQ(Appeals.UserID)).
	INNER_JOIN(AppealTypes, AppealTypes.ID.EQ(Appeals.AppealTypeID)).
	INNER_JOIN(AppealCategories, AppealCategories.ID.EQ(AppealTypes.AppealCategoryID)).
	INNER_JOIN(AppealPhotos, AppealPhotos.AppealID.EQ(Appeals.ID)).
	INNER_JOIN(AppealStatus, AppealStatus.ID.EQ(Appeals.StatusID)))

func (a *AppealRepository) Get(id *int32) (*entity.Appeal, error) {
	var u entity.Appeal
	stmt := selectAppealStmt.
		WHERE(Appeals.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealRepository) GetAll(f *input.Filter) (*[]entity.Appeal, error) {
	var u []entity.Appeal
	stmt := selectAppealStmt.
		LIMIT(f.PageSize).
		OFFSET((f.Page - 1) * f.PageSize)
	ORDER_BY(Appeals.ID.DESC())
	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealRepository) GetTotal() (*int, error) {
	var count int
	rawSql, _ := SELECT(Raw("count(*)")).
		FROM(Appeals).Sql()

	if err := a.db.QueryRow(rawSql).Scan(&count); err != nil {
		return nil, err
	}
	return &count, nil
}

func (a *AppealRepository) Create(appeal *model.Appeals, urls *[]string) (*entity.Appeal, error) {
	var u *entity.Appeal
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	stmt := Appeals.INSERT(Appeals.AllColumns.Except(Appeals.ID)).
		MODEL(appeal).
		RETURNING(Appeals.ID)

	if err := stmt.QueryContext(ctx, tx, appeal); err != nil {
		return nil, err
	}

	for _, url := range *urls {
		photosStmt := AppealPhotos.INSERT(AppealPhotos.AppealID, AppealPhotos.URL).
			VALUES(Int32(appeal.ID), String(url))
		if _, err := photosStmt.ExecContext(ctx, tx); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	u, err = a.Get(&appeal.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (a *AppealRepository) Update(appeal *model.Appeals) (*entity.Appeal, error) {
	var u *entity.Appeal
	stmt := Appeals.UPDATE(Appeals.AllColumns.Except(Appeals.ID)).
		MODEL(appeal).
		WHERE(Appeals.ID.EQ(Int32(appeal.ID))).
		RETURNING(Appeals.ID)

	if err := stmt.Query(a.db, appeal); err != nil {
		return nil, err
	}

	u, err := a.Get(&appeal.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (a *AppealRepository) Delete(id int32) error {
	stmt := Appeals.DELETE().
		WHERE(Appeals.ID.EQ(Int32(id)))
	if _, err := stmt.Exec(a.db); err != nil {
		return err
	}
	return nil
}
