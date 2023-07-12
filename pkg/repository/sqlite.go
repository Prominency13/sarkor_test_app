package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "user"
	phoneTable = "phone"
)

type Config struct{
	Path string
}

func NewSqliteDB(cfg Config) (*sqlx.DB, error){
	db,err := sqlx.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}