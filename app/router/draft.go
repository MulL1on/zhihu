package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type DraftRouter struct{}

func (r *DraftRouter) InitDraftRouter(router *gin.RouterGroup) gin.IRouter {
	draftRouter := router.Group("/content/draft")
	draftApi := api.Draft()
	{
		draftRouter.POST("/", draftApi.Audit().CreateDraft)
		draftRouter.PUT("/", draftApi.Audit().UpdateDraft)
		draftRouter.GET("/detail", draftApi.Audit().GetDraftDetail)
		draftRouter.DELETE("/", draftApi.Audit().DeleteDraft)
	}
	return draftRouter
}
