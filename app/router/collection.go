package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type CollectionRouter struct{}

//func (r *CollectionRouter) InitCollectionRouter(router *gin.RouterGroup) gin.IRouter {
//	collectionRouter := router.Group("/collection")
//
//	collectionApi := api.Collection()
//
//	{
//	}
//	return collectionRouter
//}

func (r *CollectionRouter) InitCollectionPrivateRouter(router *gin.RouterGroup) gin.IRouter {
	collectionRouter := router.Group("/collection")
	collectionApi := api.Collection()

	{
		collectionRouter.POST("/set", collectionApi.Edit().CreatCollectionSet)
		collectionRouter.POST("/article", collectionApi.Edit().CollectArticle)
		collectionRouter.GET("/detail", collectionApi.Info().GetCollectedArticle)
		collectionRouter.DELETE("/set", collectionApi.Edit().DeleteCollectionSet)
		collectionRouter.PUT("/set", collectionApi.Edit().ModifyCollectionSet)
		collectionRouter.DELETE("/article", collectionApi.Edit().RemoveSelectedArticle)
	}
	return collectionRouter
}
