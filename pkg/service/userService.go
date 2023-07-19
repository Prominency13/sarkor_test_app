package service

import (
	"crypto/sha1"
	"fmt"
	"sarkor/test/pkg/model"
	"sarkor/test/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserApi
}

const(
	salt       = "qpug24fpih3wch1poh"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
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
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }

	// return string(hash)
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

	return token.SignedString([]byte(signingKey))

}

