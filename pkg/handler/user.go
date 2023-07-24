package handler

import (
	"net/http"
	"sarkor/test/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) register(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	name := c.PostForm("name")
	ageStr := c.PostForm("age")

	age, err := strconv.ParseInt(ageStr, 10, 16)
    if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Can not convert age to integer")
    }

	user := model.User{
		Login: login,
		Password: password,
		Name: name,
		Age: int16(age),
	}

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := uh.services.UserApi.RegisterUser(user)
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

	c.SetCookie("SESSTOKEN", token, 3600, "/", "localhost", false, false)

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

	userId, err := uh.getUserId(c)
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

	users, err := uh.services.FindUsersByPhone(phone)
	if  err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) editUserPhone(c *gin.Context){
	var phone model.UpdatePhoneInput

	if err := c.BindJSON(&phone); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(*phone.Phone) > 12{
		newErrorResponse(c, http.StatusBadRequest, "Phone length can't be longer than 12 characters")
		return
	}

	// Get user id from JWT
	userId, err := uh.getUserId(c)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "Couldn't get user id")
		return
	}

	err = uh.services.UpdatePhone(userId, phone)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK!",
	})
}

func (uh *UserHandler) deleteUserPhone(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("phone_id"))
	if  err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = uh.services.DeletePhoneByPhoneId(id)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK!",
	})
}
