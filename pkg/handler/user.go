package handler

import (
	"net/http"
	"sarkor/test/pkg/model"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) register(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := uh.services.UserApi.RegisterUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type authInput struct{
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uh *UserHandler) auth(c *gin.Context) {
	var input authInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := uh.services.UserApi.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// expirationTime := 60 * time.Minute
	c.SetCookie("SESSTOKEN", token, 3600, "/", "localhost", false, false)
	// cookie := http.Cookie{
	// 	Name:       "SESSTOKEN",
	// 	Value:      token,
	// 	Expires:    time.Now().Add(expirationTime),
	// }

	// http.SetCookie(rw, &cookie)
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (uh *UserHandler) getUserByName(c *gin.Context) {
	var user model.User
	name := c.Param("name")
	if err := c.BindJSON(&user); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := uh.services.UserApi.FindUserByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": user.Id,
		"name": user.Name,
		"age": user.Age,
	})
}

func (uh *UserHandler) addUserPhone(c *gin.Context) {
	var phone model.Phone

	if err := c.BindJSON(&phone); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(phone.Phone) > 12{
		newErrorResponse(c, http.StatusBadRequest, "Phone length can't be longer than 12 characters")
		return
	}

	cookie, err := c.Request.Cookie("SESSTOKEN")

	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tokenString := cookie.Value

	userId, err := uh.services.ParseToken(tokenString)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := uh.services.AddUserPhone(phone, userId)
	if  err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"phoneId": id})
}

func (uh *UserHandler) getUserPhone(c *gin.Context) {
	phone := c.Query("number")

	if err := c.BindJSON(&phone); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	users, err := uh.services.FindUsersByPhone(phone)
	if  err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"user_id": users[len(users)-1].User_id, "phone": users[len(users)-1].Phone, "is_fax": users[len(users)-1].Is_fax})
}

func (uh *UserHandler) editUserPhone(c *gin.Context){

}

func (uh *UserHandler) deleteUserPhone(c *gin.Context) {

}
