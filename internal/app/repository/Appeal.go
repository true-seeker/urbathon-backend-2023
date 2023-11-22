package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
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

var selectStmt = SELECT(Appeals.AllColumns,
	Users.ID.AS("users.id"),
	Users.Name.AS("users.name"),
	Users.Email.AS("users.email"),
	AppealTypes.ID.AS("appealTypes.id"),
	AppealTypes.Title.AS("appealTypes.title"),
	AppealThemes.ID.AS("appealThemes.id"),
	AppealThemes.Title.AS("appealThemes.title"),
	AppealCategories.ID.AS("appealCategories.id"),
	AppealCategories.Title.AS("appealCategories.title"),
).FROM(Appeals.
	INNER_JOIN(Users, Users.ID.EQ(Appeals.UserID)).
	INNER_JOIN(AppealTypes, AppealTypes.ID.EQ(Appeals.AppealTypeID)).
	INNER_JOIN(AppealThemes, AppealThemes.ID.EQ(AppealTypes.AppealThemeID)).
	INNER_JOIN(AppealCategories, AppealCategories.ID.EQ(AppealThemes.AppealCategoryID)))

func (a *AppealRepository) Get(id *int32) (*entity.Appeal, error) {
	var u entity.Appeal
	stmt := selectStmt.
		WHERE(Appeals.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealRepository) GetAll(f *input.Filter) (*[]entity.Appeal, error) {
	var u []entity.Appeal
	stmt := selectStmt.
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
