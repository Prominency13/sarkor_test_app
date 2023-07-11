package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct{

}

func(uh *UserHandler) InitRoutes() *gin.Engine{
	userRouter := gin.New()

	userApi:= userRouter.Group("/user")
	{
		userApi.POST("/register", uh.register)
		userApi.GET("/auth", uh.auth)
		userApi.GET("/:name", uh.getUserName)
		userApi.POST("/phone", uh.addUserPhone)
		userApi.GET("/phone?q=", uh.getUserPhone)
		userApi.DELETE("/phone/:phone_id", uh.deleteUserPhone)
	}
	return userRouter
}