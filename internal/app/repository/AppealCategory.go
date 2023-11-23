package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/storage"
)

type AppealCategoryRepository struct {
	db *sql.DB
}

func NewAppealCategoryRepository(s storage.Sql) *AppealCategoryRepository {
	return &AppealCategoryRepository{db: s.GetDb()}
}

func (a *AppealCategoryRepository) Get(id *int32) (*model.AppealCategories, error) {
	var u model.AppealCategories
	stmt := SELECT(AppealCategories.AllColumns).
		FROM(AppealCategories).
		WHERE(AppealCategories.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealCategoryRepository) GetAll() (*[]model.AppealCategories, error) {
	var u []model.AppealCategories
	stmt := SELECT(AppealCategories.AllColumns).
		FROM(AppealCategories).
		ORDER_BY(AppealCategories.Title.ASC())
	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealCategoryRepository) GetAppealTypes(id *int32) (*[]model.AppealTypes, error) {
	var u []model.AppealTypes
	stmt := SELECT(AppealTypes.AllColumns).
		FROM(AppealTypes).
		WHERE(AppealTypes.AppealCategoryID.EQ(Int32(*id))).
		ORDER_BY(AppealTypes.Title.ASC())
	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
