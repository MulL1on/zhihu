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
		articleRouter.GET("/list", articleApi.Info().GetArticleList)
		articleRouter.GET("/detail", articleApi.Info().GetArticleDetail)
	}
	return articleRouter
}

func (r *ArticleRouter) InitPrivateArticleRouter(router *gin.RouterGroup) gin.IRouter {
	articleRouter := router.Group("/content/article")
	articleApi := api.Article()
	{
		articleRouter.GET("/listByUser", articleApi.Info().GetArticleList)
		articleRouter.GET("/detailByUser", articleApi.Info().GetArticleDetail)
		articleRouter.POST("/", articleApi.Edit().Publish)

	}
	return articleRouter
}
