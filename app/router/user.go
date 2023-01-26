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
		UserRouter.POST("/login", userApi.Sign().Login)
		UserRouter.POST("/register", userApi.Sign().Register)
		UserRouter.POST("/code", userApi.Sign().SendCode)
	}
	return UserRouter
}

func (r *UserRouter) InitPrivateUserRouter(router *gin.RouterGroup) gin.IRouter {
	UserPrivateRouter := router.Group("/user")
	userApi := api.User()
	{
		UserPrivateRouter.GET("/logout", userApi.Sign().Logout)
	}
	return UserPrivateRouter
}
