package repository

import (
	"sarkor/test/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UserApi interface{
	RegisterUser(user model.User) (int, error)
}

type Repository struct{
	UserApi
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		UserApi: NewUserSql(db),
	}
}