package repository

import (
	"database/sql"
	"errors"
	"urbathon-backend-2023/internal/app/model/entity"
	"urbathon-backend-2023/internal/app/model/input"
	"urbathon-backend-2023/internal/app/storage"
)

type UserRepository struct {
	db                 *sql.DB
	incidentRepository *IncidentRepository
}

func NewUserRepository(s storage.Sql) *UserRepository {
	return &UserRepository{db: s.GetDb()}
}

func (a *UserRepository) GetByCreds(loginInput *input.Login) (*entity.User, error) {
	row := a.db.QueryRow("select id, name, email from users WHERE email = $1 AND password = $2", loginInput.Email, loginInput.Password)

	var u entity.User

	switch err := row.Scan(&u.Id, &u.Name, &u.Email); {
	case errors.Is(err, sql.ErrNoRows):
		return nil, err
	}
	return &u, nil
}
func (a *UserRepository) Get(id *int) (*entity.User, error) {
	row := a.db.QueryRow("select id, name, email from users WHERE id = $1 ", id)

	var u entity.User

	switch err := row.Scan(&u.Id, &u.Name, &u.Email); {
	case errors.Is(err, sql.ErrNoRows):
		return nil, err
	}
	return &u, nil
}

func (a *UserRepository) Create(userInput *input.User) (*entity.User, error) {
	var id int
	_ = a.db.QueryRow("insert into users(email, password, name) values ($1, $2, $3) RETURNING id",
		userInput.Email, userInput.Password, userInput.Name).
		Scan(&id)

	e, err := a.Get(&id)
	if err != nil {
		return nil, err
	}

	return e, nil
}
