package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/storage"
)

type UserRepository struct {
	db *sql.DB
}

func NewAuthRepository(s storage.Sql) *UserRepository {
	return &UserRepository{db: s.GetDb()}
}

func (a *UserRepository) GetByEmail(email *string) (*model.Users, error) {
	var u model.Users
	stmt := SELECT(Users.ID, Users.FirstName, Users.LastName, Users.Patronymic, Users.Email, Users.Password, Users.Salt).
		FROM(Users).
		WHERE(Users.Email.EQ(String(*email)))

	err := stmt.Query(a.db, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (a *UserRepository) Get(id *int32) (*model.Users, error) {
	var u model.Users
	stmt := SELECT(Users.ID, Users.FirstName, Users.LastName, Users.Patronymic, Users.Email).
		FROM(Users).
		WHERE(Users.ID.EQ(Int32(*id)))

	err := stmt.Query(a.db, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (a *UserRepository) Create(user *model.Users) (*model.Users, error) {
	var u model.Users
	stmt := Users.INSERT(Users.Email, Users.Password, Users.FirstName, Users.LastName, Users.Patronymic, Users.Salt).
		MODEL(user).
		RETURNING(Users.AllColumns)

	err := stmt.Query(a.db, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
