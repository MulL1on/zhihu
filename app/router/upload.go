package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type UploadRouter struct{}

func (r *UploadRouter) InitUploadRouter(router *gin.RouterGroup) gin.IRouter {
	uploadRouter := router.Group("/upload")
	uploadApi := api.Upload()
	{
		uploadRouter.POST("/user/user-avatar", uploadApi.Upload().UserAvatarUpload)
	}
	return uploadRouter
}
