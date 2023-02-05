package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type DiggRouter struct{}

func (r *DiggRouter) InitDiggRouter(router *gin.RouterGroup) gin.IRouter {
	diggRouter := router.Group("/digg")
	diggApi := api.Digg()
	{
		diggRouter.POST("/", diggApi.Like().DoDigg)

		diggRouter.DELETE("/", diggApi.Like().UndoDigg)
	}
	return diggRouter
}
