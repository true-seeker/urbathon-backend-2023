package repository

import (
	"database/sql"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/storage"
)

type OrganizationRepository struct {
	db *sql.DB
}

func NewOrganizationRepository(s storage.Sql) *OrganizationRepository {
	return &OrganizationRepository{db: s.GetDb()}
}

func (a *OrganizationRepository) Register(organization *model.Organizations, organizationInputCategories *[]int32) (*model.Organizations, error) {
	var u model.Organizations
	stmt := Organizations.
		INSERT(Organizations.AllColumns.Except(Organizations.ID)).
		MODEL(organization).
		RETURNING(Organizations.AllColumns)

	if err := stmt.Query(a.db, &u); err != nil {
		return nil, err
	}

	stmt = OrganizationAppealCategory.
		INSERT(OrganizationAppealCategory.OrganizationID, OrganizationAppealCategory.AppealCategoryID)
	for _, e := range *organizationInputCategories {
		stmt = stmt.VALUES(u.ID, Int32(e))
	}
	if _, err := stmt.Exec(a.db); err != nil {
		return nil, err
	}

	return &u, nil
}
func (a *OrganizationRepository) AddUser(organizationId int32, userId int32) error {
	stmt := Users.
		UPDATE(Users.OrganizationID, Users.Role).
		SET(Int32(organizationId), Int32(2)).
		WHERE(Users.ID.EQ(Int32(userId)))
	fmt.Println(stmt.Sql())
	if _, err := stmt.Exec(a.db); err != nil {
		return err
	}
	return nil
}
