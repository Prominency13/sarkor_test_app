package model

type User struct {
	Id       int    `json:"id" db:"id"`
	Login    string `json:"login" binding:"required"`
	Password string	`json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Age      int16  `json:"age" binding:"required"`
}
