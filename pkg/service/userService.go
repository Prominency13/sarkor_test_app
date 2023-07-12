package service

import (
	"sarkor/test/pkg/model"
	"sarkor/test/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{
	repo repository.UserApi
}

func NewUserService(repo repository.UserApi) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user model.User) (int, error){
	user.Password = s.generatePwdHash(user.Password)
	return s.repo.RegisterUser(user)
}

func (s *UserService) generatePwdHash(password string) string{
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
