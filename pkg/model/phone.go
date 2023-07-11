package test

type Phone struct{
	Id int `json:"id"`
	Phone string `json:"phone"`
	Description string `json:"decription"`
	Is_fax bool `json:"is_fax"`
	User_id User `json:"user_id"`
}