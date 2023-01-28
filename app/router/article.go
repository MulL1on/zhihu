package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type ArticleRouter struct{}

func (r *ArticleRouter) InitArticleRouter(router *gin.RouterGroup) gin.IRouter {
	ArticleRouter := router.Group("/content/article")
	articleApi := api.Article()
	{
		ArticleRouter.POST("/", articleApi.Edit().Publish)
	}
	return ArticleRouter
}
