package router

import (
	"github.com/gin-gonic/gin"
	"juejin/app/api"
)

type DraftRouter struct{}

func (r *DraftRouter) InitDraftRouter(router *gin.RouterGroup) gin.IRouter {
	DraftRouter := router.Group("/content/draft")
	draftApi := api.Draft()
	{
		DraftRouter.POST("/", draftApi.Audit().CreateDraft)
		DraftRouter.PUT("/", draftApi.Audit().UpdateDraft)
		DraftRouter.GET("/", draftApi.Audit().GetDraftDetail)
		DraftRouter.DELETE("/", draftApi.Audit().DeleteDraft)
	}
	return DraftRouter
}
