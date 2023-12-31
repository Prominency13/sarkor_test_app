package repository

import (
	"sarkor/test/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UserApi interface{
	RegisterUser(user model.User) (int, error)
	GetUser(login, password string) (model.User, error)
	GetUserByName(name string) (model.User, error)
	AddUserPhone(phone model.Phone, userId int) (int, error)
	FindPhoneDuplicate(phone string, userId int) (string, error)
	GetUsersByPhone(phone string) ([]model.Phone, error)
	UpdatePhone(userId int, phone model.UpdatePhoneInput) error
	DeletePhoneByPhoneId(phoneId int) error
}

type Repository struct{
	UserApi
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		UserApi: NewUserSql(db),
	}
}