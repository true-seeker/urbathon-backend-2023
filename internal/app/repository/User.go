package repository

import (
	"database/sql"
	. "github.com/go-jet/jet/v2/postgres"
	"urbathon-backend-2023/.gen/urbathon/public/model"
	. "urbathon-backend-2023/.gen/urbathon/public/table"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/storage"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(s storage.Sql) *UserRepository {
	return &UserRepository{db: s.GetDb()}
}

func (a *UserRepository) GetByEmail(loginInput *input.Login) (*model.Users, error) {
	var u model.Users
	stmt := SELECT(Users.ID, Users.Name, Users.Email, Users.Password, Users.Salt).
		FROM(Users).
		WHERE(Users.Email.EQ(String(*loginInput.Email)))

	err := stmt.Query(a.db, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (a *UserRepository) Get(id *int32) (*model.Users, error) {
	var u model.Users
	stmt := SELECT(Users.ID, Users.Name, Users.Email).
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
	stmt := Users.INSERT(Users.Email, Users.Password, Users.Name, Users.Salt).
		MODEL(user).
		RETURNING(Users.AllColumns)

	err := stmt.Query(a.db, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
