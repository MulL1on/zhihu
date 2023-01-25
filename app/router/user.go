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
	}
	return UserRouter
}
