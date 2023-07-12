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
