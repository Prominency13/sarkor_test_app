package model

type Phone struct {
	Id          int    `json:"id"`
	Phone       string `json:"phone" binding:"required"`
	Description string `json:"decription" binding:"required"`
	Is_fax      bool   `json:"is_fax" binding:"required"`
	User_id     User   `json:"user_id" binding:"required"`
}
