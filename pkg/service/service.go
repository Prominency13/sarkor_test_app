package service

import "sarkor/test/pkg/repository"

type UserApi interface{}

type Service struct{
	UserApi
}

func NewService(repos *repository.Repository) *Service{
	return &Service{}
}