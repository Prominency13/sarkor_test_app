package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const(
	authHeader = "Authorization"
	userCtx = "userId"
)

func(uh *UserHandler) userIdentity(c *gin.Context){
	header := c.GetHeader(authHeader)
	if header == ""{
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Token contains an invalid number of segments")
		return
	}

	userId, err := uh.services.UserApi.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func(uh *UserHandler) checkCookie(c* gin.Context) (int, error){
	cookie, err := c.Request.Cookie("SESSTOKEN")

	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	tokenString := cookie.Value

	userId, err := uh.services.ParseToken(tokenString)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return userId, err
}

func(uh *UserHandler) getUserId(c *gin.Context) (int, error){
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User id not found")
		return 0, errors.New("User id not found")
	}

	// Приведение id к типу int
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User id type is invalid")
		return 0, errors.New("User id type is invalid")
	}

	return idInt, nil
}