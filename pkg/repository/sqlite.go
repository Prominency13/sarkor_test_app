package repository

import (
	"github.com/sirupsen/logrus"

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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user(id INTEGER PRIMARY KEY, login TEXT UNIQUE, password TEXT, name TEXT, age TEXT);")
    if err != nil {
		logrus.Fatalf("Error occurred while processing SQL query %s", err.Error())
    }
	
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS phone(id INTEGER PRIMARY KEY, phone VARCHAR(12), description TEXT, is_fax BOOL, user_id INTEGER);")
    if err != nil {
		logrus.Fatalf("Error occurred while processing SQL query %s", err.Error())
    }

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}