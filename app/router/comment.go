package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type CommentRouter struct{}

func (r *CommentRouter) InitCommentRouter(router *gin.RouterGroup) gin.IRouter {
	commentRouter := router.Group("/comment")
	commentApi := api.Comment()
	{
		commentRouter.POST("/", commentApi.Review().PostComment)
		commentRouter.DELETE("/", commentApi.Review().DeleteComment)
		commentRouter.POST("/reply", commentApi.Reply().PostReply)
		commentRouter.DELETE("/reply", commentApi.Reply().DeleteReply)
		commentRouter.GET("/list", commentApi.Review().GetCommentList)

	}
	return commentRouter
}
