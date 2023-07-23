package service

import (
	"crypto/sha1"
	"errors"
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

func (s *UserService) ParseToken(accessToken string) (int, error){
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Token claims aren't type of *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *UserService) FindUserByName(name string) (model.User, error){
	return s.repo.GetUserByName(name)
}

func (s *UserService) AddUserPhone(phone model.Phone, userId int) (int, error){
	duplicatePhone, err := s.repo.FindPhoneDuplicate(phone.Phone, userId)
	if duplicatePhone != "" {
		return 0, err
	}

	return s.repo.AddUserPhone(phone, userId)
}

func (s *UserService) FindUsersByPhone(phone string) ([]model.Phone, error){
	return s.repo.GetUsersByPhone(phone)
}

func(s *UserService) UpdatePhone(userId int, phone model.UpdatePhoneInput) error{
	return s.repo.UpdatePhone(userId, phone)
}

func(s *UserService) DeletePhoneByPhoneId(phoneId int) error{
	return s.repo.DeletePhoneByPhoneId(phoneId)

}
