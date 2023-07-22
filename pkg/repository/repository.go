package repository

import (
	"sarkor/test/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UserApi interface{
	RegisterUser(user model.User) (int, error)
	GetUser(login, password string) (model.User, error)
	GetUserByName(name string) (model.User, error)
}

type Repository struct{
	UserApi
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		UserApi: NewUserSql(db),
	}
}