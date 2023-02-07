package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type RankRouter struct{}

func (r *RankRouter) InitRankRouter(router *gin.RouterGroup) gin.IRouter {
	rankRouter := router.Group("/rank")
	rankApi := api.Rank()
	{
		rankRouter.GET("/author", rankApi.Rank().GetAuthorRankings)
	}
	return rankRouter
}
