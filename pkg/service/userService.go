package service

import (
	"sarkor/test/pkg/model"
	"sarkor/test/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserApi
}

const(
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Login  string `json:"login"`
}

func NewUserService(repo repository.UserApi) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user model.User) (int, error) {
	user.Password = s.generatePwdHash(user.Password)
	return s.repo.RegisterUser(user)
}

func (s *UserService) generatePwdHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func (s *UserService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, s.generatePwdHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
		user.Login,
	})

	return token.Signature , nil

}

