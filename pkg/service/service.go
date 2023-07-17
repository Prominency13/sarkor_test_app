package service

import (
	"sarkor/test/pkg/model"
	"sarkor/test/pkg/repository"
)

type UserApi interface {
	RegisterUser(user model.User) (int, error)
	// GenerateToken(login, password string) (int, error)
	// GetUser(username, password string) (model.User, error)
}

type Service struct {
	UserApi
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserApi: NewUserService(repos.UserApi),
	}
}
