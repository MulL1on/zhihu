package router

import (
	"github.com/gin-gonic/gin"
	g "juejin/app/global"
	"juejin/app/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS(), middleware.RateLimit(), middleware.ZapLogger(g.Logger), middleware.ZapRecovery(g.Logger, true))
	routerGroup := new(Group)

	publicGroup := r.Group("/api")
	{
		routerGroup.InitArticleRouter(publicGroup)
		routerGroup.InitTagRouter(publicGroup)
		routerGroup.InitUserRouter(publicGroup)
		routerGroup.InitRankRouter(publicGroup)
	}

	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.JWTAuthMiddleware())
	{
		routerGroup.InitPrivateArticleRouter(privateGroup)
		routerGroup.InitDraftRouter(privateGroup)
		routerGroup.InitPrivateUserRouter(privateGroup)
		routerGroup.InitCollectionPrivateRouter(privateGroup)
		routerGroup.InitFollowerRouter(privateGroup)
		routerGroup.InitCommentRouter(privateGroup)
		routerGroup.InitDiggRouter(privateGroup)
		routerGroup.InitUploadRouter(privateGroup)
	}
	g.Logger.Info("initialize routers successfully")
	return r
}
