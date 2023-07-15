package handler

import (
	"net/http"
	"sarkor/test/pkg/model"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) register(c *gin.Context) {
	var input model.User

	// id := c.Query("id")
	// page := c.DefaultQuery("page", "0")
	// name := c.PostForm("name")
	// message := c.PostForm("message")
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

func (uh *UserHandler) auth(c *gin.Context) {

}

func (uh *UserHandler) getUserName(c *gin.Context) {

}

func (uh *UserHandler) addUserPhone(c *gin.Context) {

}

func (uh *UserHandler) getUserPhone(c *gin.Context) {

}

func (uh *UserHandler) deleteUserPhone(c *gin.Context) {

}
