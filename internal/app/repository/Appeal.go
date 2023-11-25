package repository

import (
	"context"
	"database/sql"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/filter"
	"urbathon-backend-2023/internal/app/storage"
)

type AppealRepository struct {
	db *sql.DB
}

func NewAppealRepository(s storage.Sql) *AppealRepository {
	return &AppealRepository{db: s.GetDb()}
}

func getSelectAppeal(f *filter.AppealFilter) SelectTable {
	stmt := SELECT(Appeals.AllColumns).FROM(Appeals)
	var stmtTable SelectTable
	if f != nil {
		stmtTable = f.GetLimitOffsetStmt(stmt).AsTable("rAppeals")
	} else {
		stmtTable = stmt.AsTable("rAppeals")
	}
	return stmtTable
}

func getSelectJoinAppealStmt(f *filter.AppealFilter) SelectStatement {
	rAppeals := getSelectAppeal(f)
	rAppealsId := Appeals.ID.From(rAppeals)
	rAppealsUserId := Appeals.UserID.From(rAppeals)
	rAppealsAppealTypeID := Appeals.AppealTypeID.From(rAppeals)
	rAppealsStatusID := Appeals.StatusID.From(rAppeals)
	rAppealsLongitude := Appeals.Longitude.From(rAppeals)
	rAppealsLatitude := Appeals.Latitude.From(rAppeals)

	stmt := SELECT(rAppeals.AllColumns(),
		Users.ID.AS("users.id"),
		Users.FirstName.AS("users.firstname"),
		Users.LastName.AS("users.lastname"),
		Users.Patronymic.AS("users.patronymic"),
		Users.Email.AS("users.email"),
		AppealTypes.ID.AS("appealTypes.id"),
		AppealTypes.Title.AS("appealTypes.title"),
		AppealCategories.ID.AS("appealCategories.id"),
		AppealCategories.Title.AS("appealCategories.title"),
		AppealPhotos.ID.AS("appealPhotos.id"),
		AppealPhotos.URL.AS("appealPhotos.url"),
		AppealStatus.ID.AS("appealStatus.id"),
		AppealStatus.Status.AS("appealStatus.status"),
	).FROM(rAppeals.
		INNER_JOIN(Users, Users.ID.EQ(rAppealsUserId)).
		INNER_JOIN(AppealTypes, AppealTypes.ID.EQ(rAppealsAppealTypeID)).
		LEFT_JOIN(AppealPhotos, AppealPhotos.AppealID.EQ(rAppealsId)).
		INNER_JOIN(AppealStatus, AppealStatus.ID.EQ(rAppealsStatusID)).
		INNER_JOIN(AppealCategories, AppealCategories.ID.EQ(AppealTypes.AppealCategoryID)))

	if f != nil {
		if f.UserId != nil {
			stmt = stmt.WHERE(rAppealsUserId.EQ(Int32(*f.UserId)))
		}
		if f.LatUp != nil && f.LatDown != nil && f.LongDown != nil && f.LongUp != nil {
			stmt = stmt.WHERE(rAppealsLongitude.GT(Float(*f.LongUp)).
				AND(rAppealsLongitude.LT(Float(*f.LongDown))).
				AND(rAppealsLatitude.GT(Float(*f.LatDown))).
				AND(rAppealsLatitude.LT(Float(*f.LatUp))))
		}
	}
	return stmt
}

func getSelectAppealStmt() SelectStatement {
	return SELECT(Appeals.AllColumns,
		Users.ID.AS("users.id"),
		Users.FirstName.AS("users.firstname"),
		Users.LastName.AS("users.lastname"),
		Users.Patronymic.AS("users.patronymic"),
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
		LEFT_JOIN(AppealPhotos, AppealPhotos.AppealID.EQ(Appeals.ID)).
		INNER_JOIN(AppealStatus, AppealStatus.ID.EQ(Appeals.StatusID)))

}
func (a *AppealRepository) Get(id *int32) (*entity.Appeal, error) {
	var u entity.Appeal
	stmt := getSelectAppealStmt().
		WHERE(Appeals.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealRepository) GetAll(f *filter.AppealFilter) (*[]entity.Appeal, error) {
	var u []entity.Appeal
	stmt := getSelectJoinAppealStmt(f)

	fmt.Println(stmt.Sql())

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	fmt.Println(111, len(u))
	return &u, nil
}

func (a *AppealRepository) GetTotal(f *filter.AppealFilter) (*int, error) {
	var count int
	stmt := SELECT(Raw("count(*)")).
		FROM(Appeals)

	rawSql, args := stmt.Sql()
	if err := a.db.QueryRow(rawSql, args...).Scan(&count); err != nil {
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
	stmt := Appeals.INSERT(Appeals.AllColumns.Except(Appeals.ID, Appeals.StatusID)).
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
	stmt := Appeals.UPDATE(Appeals.AllColumns.Except(Appeals.ID, Appeals.StatusID)).
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

func (a *AppealRepository) UpdateStatus(appealId int32, statusId int32) error {
	stmt := Appeals.UPDATE(Appeals.StatusID).
		SET(Int32(statusId)).
		WHERE(Appeals.ID.EQ(Int32(appealId)))
	if _, err := stmt.Exec(a.db); err != nil {
		return err
	}

	return nil
}

func makeWhereMap(f *filter.Map, stmt SelectStatement) SelectStatement {
	if f.LatUp != nil && f.LatDown != nil && f.LongDown != nil && f.LongUp != nil {
		stmt = stmt.WHERE(Appeals.Longitude.GT(Float(*f.LongUp)).
			AND(Appeals.Longitude.LT(Float(*f.LongDown))).
			AND(Appeals.Latitude.GT(Float(*f.LatDown))).
			AND(Appeals.Latitude.LT(Float(*f.LatUp))))
	}
	return stmt
}

func (a *AppealRepository) GetForMap(f *filter.Map) (*[]entity.Appeal, error) {
	var u []entity.Appeal
	stmt := Appeals.SELECT(Appeals.ID, Appeals.Latitude, Appeals.Longitude, Appeals.Title).FROM(Appeals)
	stmt = makeWhereMap(f, stmt)

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
