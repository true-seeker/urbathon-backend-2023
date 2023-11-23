package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/storage"
)

type AppealStatusRepository struct {
	db *sql.DB
}

func NewAppealStatusRepository(s storage.Sql) *AppealStatusRepository {
	return &AppealStatusRepository{db: s.GetDb()}
}

func (a *AppealStatusRepository) Get(id *int32) (*model.AppealStatus, error) {
	var u model.AppealStatus
	stmt := SELECT(AppealStatus.AllColumns).
		FROM(AppealStatus).
		WHERE(AppealStatus.ID.EQ(Int32(*id)))

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *AppealStatusRepository) GetAll() (*[]model.AppealStatus, error) {
	var u []model.AppealStatus
	stmt := SELECT(AppealStatus.AllColumns).
		FROM(AppealStatus).
		ORDER_BY(AppealStatus.ID.ASC())
	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
