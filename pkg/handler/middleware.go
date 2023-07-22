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

	c.Set(userCtx, userId)
}