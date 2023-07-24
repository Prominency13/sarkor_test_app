package handler

import (
	"sarkor/test/pkg/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	services *service.Service
}

func NewHandler(services *service.Service) *UserHandler{
	return &UserHandler{services: services}
}

func(uh *UserHandler) InitRoutes() *gin.Engine{
	userRouter := gin.New()

	userApi:= userRouter.Group("/user")
	{
		userApi.POST("/register", uh.register)
		userApi.POST("/auth", uh.auth)
		userApi.GET("/:name", uh.cookieIdentity, uh.getUserByName)
		userApi.POST("/phone", uh.cookieIdentity,uh.addUserPhone)
		userApi.GET("/phone", uh.cookieIdentity, uh.getUserPhone)
		userApi.PUT("/phone", uh.cookieIdentity, uh.editUserPhone)
		userApi.DELETE("/phone/:phone_id", uh.cookieIdentity, uh.deleteUserPhone)
	}
	return userRouter
}