package model

type Phone struct {
	Id          int    `json:"id"`
	Phone       string `json:"phone" binding:"required"`
	Description string `json:"description" binding:"required"`
	Is_fax      bool   `json:"is_fax"`
	User_id     int    `json:"user_id"`
}
