package test

type User struct{
	Id int `json:"id"`
	Login string `json:"login"`
	Password string
	Name string `json:"name"`
	Age	int16 `json:"age"`
}