package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/user")
	userApi := api.User()
	{
		userRouter.POST("/login", userApi.Auth().Login)
		userRouter.POST("/register", userApi.Auth().Register)
		userRouter.POST("/code", userApi.Auth().SendCode)
	}
	return userRouter
}

func (r *UserRouter) InitPrivateUserRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/user")
	userApi := api.User()
	{
		userRouter.PUT("/info", userApi.Info().UpdateUserInfo)
		userRouter.GET("/info", userApi.Info().GetUserInfo)
		userRouter.GET("/logout", userApi.Auth().Logout)
	}
	return userRouter
}
