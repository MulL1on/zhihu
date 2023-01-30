package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type ArticleRouter struct{}

func (r *ArticleRouter) InitArticleRouter(router *gin.RouterGroup) gin.IRouter {
	articleRouter := router.Group("/content/article")
	articleApi := api.Article()
	{
		articleRouter.GET("/detail", articleApi.Info().GetArticleDetail)
		articleRouter.POST("/", articleApi.Edit().Publish)
	}
	return articleRouter
}
