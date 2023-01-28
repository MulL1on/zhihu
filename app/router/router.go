package router

import (
	"github.com/gin-gonic/gin"
	g "juejin/app/global"
	"juejin/app/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS(), middleware.ZapLogger(g.Logger), middleware.ZapRecovery(g.Logger, true))
	routerGroup := new(Group)

	publicGroup := r.Group("/api")
	{
		routerGroup.InitUserRouter(publicGroup)
	}

	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.JWTAuthMiddleware())
	{
		routerGroup.InitDraftRouter(privateGroup)
		routerGroup.InitPrivateUserRouter(privateGroup)
		routerGroup.InitArticleRouter(privateGroup)
	}
	g.Logger.Info("initialize routers successfully")
	return r
}
