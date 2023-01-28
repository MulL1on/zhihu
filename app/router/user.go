package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) gin.IRouter {
	UserRouter := router.Group("/user")
	userApi := api.User()
	{
		UserRouter.POST("/login", userApi.Auth().Login)
		UserRouter.POST("/register", userApi.Auth().Register)
		UserRouter.POST("/code", userApi.Auth().SendCode)
	}
	return UserRouter
}

func (r *UserRouter) InitPrivateUserRouter(router *gin.RouterGroup) gin.IRouter {
	UserPrivateRouter := router.Group("/user")
	userApi := api.User()
	{
		UserPrivateRouter.PUT("/info", userApi.Info().UpdateUserInfo)
		UserPrivateRouter.GET("/info", userApi.Info().GetUserInfo)
		UserPrivateRouter.GET("/logout", userApi.Auth().Logout)
	}
	return UserPrivateRouter
}
