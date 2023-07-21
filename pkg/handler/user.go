package handler

import (
	"fmt"
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

func (uh *UserHandler) getUserName(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (uh *UserHandler) addUserPhone(c *gin.Context) {
	fmt.Println("It worked somehow")
}

func (uh *UserHandler) getUserPhone(c *gin.Context) {

}

func (uh *UserHandler) deleteUserPhone(c *gin.Context) {

}
