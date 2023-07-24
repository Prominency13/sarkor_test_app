package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const(
	userCtx = "userId"
	cookieToken = "SESSTOKEN"
)

func(uh *UserHandler) getUserId(c *gin.Context) (int, error){
	cookie, err := c.Request.Cookie(cookieToken)

	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "No cookie set")
	}

	tokenString := cookie.Value
	userId, err := uh.services.ParseToken(tokenString)

	return userId, err
}

func(uh *UserHandler) cookieIdentity(c *gin.Context) {
	cookie, err := c.Request.Cookie(cookieToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "No cookie set")
		return
	}

	tokenString := cookie.Value
	userId, err := uh.services.ParseToken(tokenString)

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, &userId)
}
