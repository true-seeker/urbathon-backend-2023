package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/storage"
)

type AppealTypeRepository struct {
	db *sql.DB
}

func NewAppealTypeRepository(s storage.Sql) *AppealTypeRepository {
	return &AppealTypeRepository{db: s.GetDb()}
}

var selectAppealTypeStmt = SELECT(AppealTypes.AllColumns,
	AppealCategories.ID.AS("appealCategories.id"),
	AppealCategories.Title.AS("appealCategories.title"),
).FROM(AppealTypes.
	INNER_JOIN(AppealCategories, AppealCategories.ID.EQ(AppealTypes.AppealCategoryID)))

func (a *AppealTypeRepository) Get(id *int32) (*entity.AppealType, error) {
	var u entity.AppealType
	stmt := selectAppealTypeStmt.
		WHERE(AppealTypes.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealTypeRepository) GetAll() (*[]entity.AppealType, error) {
	var u []entity.AppealType
	stmt := selectAppealTypeStmt.
		ORDER_BY(AppealTypes.Title.ASC())
	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
