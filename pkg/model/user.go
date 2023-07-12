package model

type User struct {
	Id       int    `json:"id" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string
	Name     string `json:"name" binding:"required"`
	Age      int16  `json:"age" binding:"required"`
}
