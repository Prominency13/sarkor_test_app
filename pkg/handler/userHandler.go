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
		userApi.GET("/:name", uh.userIdentity, uh.getUserByName)
		userApi.POST("/phone", uh.userIdentity,uh.addUserPhone)
		userApi.GET("/phone", uh.userIdentity, uh.getUserPhone)
		userApi.PUT("/phone", uh.userIdentity, uh.editUserPhone)
		userApi.DELETE("/phone/:phone_id", uh.userIdentity, uh.deleteUserPhone)
	}
	return userRouter
}