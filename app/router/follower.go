package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type FollowerRouter struct{}

func (r *FollowerRouter) InitFollowerRouter(router *gin.RouterGroup) gin.IRouter {
	followerRouter := router.Group("/follower")
	followerApi := api.Follower()
	{
		followerRouter.POST("/followee", followerApi.Follow().DoFollow)
		followerRouter.DELETE("/followee", followerApi.Follow().UndoFollow)
		followerRouter.GET("/followerList", followerApi.Follow().GetFollowerList)
		followerRouter.GET("/followeeList", followerApi.Follow().GetFolloweeList)

	}
	return followerRouter
}
