package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type TagRouter struct{}

func (r *TagRouter) InitTagRouter(router *gin.RouterGroup) gin.IRouter {
	tagRouter := router.Group("/tag")
	tagApi := api.Tag()
	{
		tagRouter.GET("/list", tagApi.Tag().GetTagList)
	}
	return tagRouter
}
