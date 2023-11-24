package repository

import (
	"context"
	"database/sql"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/storage"
)

type AppealCommentRepository struct {
	db *sql.DB
}

func NewAppealCommentRepository(s storage.Sql) *AppealCommentRepository {
	return &AppealCommentRepository{db: s.GetDb()}
}

var selectAppealCommentsStmt = SELECT(AppealComments.AllColumns,
	Users.ID.AS("users.id"),
	Users.Name.AS("users.name"),
	AppealCommentPhotos.ID.AS("appealCommentPhotos.id"),
	AppealCommentPhotos.URL.AS("appealCommentPhotos.url"),
).FROM(AppealComments.
	INNER_JOIN(Users, Users.ID.EQ(AppealComments.UserID)).
	LEFT_JOIN(AppealCommentPhotos, AppealCommentPhotos.AppealCommentID.EQ(AppealComments.ID)))

func (a *AppealCommentRepository) Get(id *int32) (*entity.AppealComment, error) {
	var u entity.AppealComment
	stmt := selectAppealCommentsStmt.
		WHERE(AppealComments.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealCommentRepository) Create(appealComments *model.AppealComments, urls *[]string) (*entity.AppealComment, error) {
	var u *entity.AppealComment
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	stmt := AppealComments.
		INSERT(AppealComments.AllColumns.Except(AppealComments.ID, AppealComments.Date)).
		MODEL(appealComments).
		RETURNING(AppealComments.ID)

	if err := stmt.QueryContext(ctx, tx, appealComments); err != nil {
		return nil, err
	}

	for _, url := range *urls {
		photosStmt := AppealCommentPhotos.
			INSERT(AppealCommentPhotos.AppealCommentID, AppealCommentPhotos.URL).
			VALUES(Int32(appealComments.ID), String(url))
		if _, err := photosStmt.ExecContext(ctx, tx); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	u, err = a.Get(&appealComments.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (a *AppealCommentRepository) GetAllComments(f *input.Filter, appealId int32) (*[]entity.AppealComment, error) {
	var u []entity.AppealComment
	stmt := selectAppealCommentsStmt.
		WHERE(AppealComments.AppealID.EQ(Int32(appealId))).
		LIMIT(f.PageSize).
		OFFSET((f.Page - 1) * f.PageSize).
		ORDER_BY(AppealComments.Date.DESC())
	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealCommentRepository) GetTotalComments(appealId int32) (*int, error) {
	var count int
	rawSql, args := SELECT(Raw("count(*)")).
		FROM(AppealComments).
		WHERE(AppealComments.AppealID.EQ(Int32(appealId))).
		Sql()

	if err := a.db.QueryRow(rawSql, args...).Scan(&count); err != nil {
		fmt.Println(rawSql, args)
		return nil, err
	}
	return &count, nil
}