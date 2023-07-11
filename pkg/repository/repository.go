package repository

import "github.com/jmoiron/sqlx"

type UserApi interface{}

type Repository struct{
	UserApi
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{}
}