package repository

type UserApi interface{}

type Repository struct{
	UserApi
}

func NewRepository() *Repository{
	return &Repository{}
}