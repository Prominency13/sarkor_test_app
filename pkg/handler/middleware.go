package handler

import (
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

	bearerToken := strings.Split(header, " ")[1]

	cookie, err := c.Request.Cookie("SESSTOKEN")

	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "No cookie set")
		return
	}

	tokenString := cookie.Value

	if bearerToken != tokenString{
		newErrorResponse(c, http.StatusForbidden, "Cookie does not match auth token")
		return
	}

	c.Set(userCtx, userId)
}

func(uh *UserHandler) getUserId(c *gin.Context) (int, error){
	cookie, err := c.Request.Cookie("SESSTOKEN")

	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "No cookie set")
	}

	tokenString := cookie.Value
	userId, err := uh.services.ParseToken(tokenString)

	return userId, err
}
