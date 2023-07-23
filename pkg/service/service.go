package service

import (
	"sarkor/test/pkg/model"
	"sarkor/test/pkg/repository"
)

type UserApi interface {
	RegisterUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
	FindUserByName(name string) (model.User, error)
	AddUserPhone(phone model.Phone, userId int) (int, error)
	FindUsersByPhone(phone string) ([]model.Phone, error)
	DeletePhoneByPhoneId(phoneId int) error
}

type Service struct {
	UserApi
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserApi: NewUserService(repos.UserApi),
	}
}
