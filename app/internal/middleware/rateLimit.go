package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/utils/common/resp"
	"net/http"
)

func RateLimit() gin.HandlerFunc {
	capacity := g.Config.Middleware.RateLimit.Capacity
	fillInterval := g.Config.Middleware.RateLimit.GetFillInterval(g.Config.Middleware.RateLimit.FillInterval)
	quantum := g.Config.Middleware.RateLimit.Quantum
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum)
	g.Logger.Info("new bucket successfully", zap.Int64("capacity", bucket.Capacity()), zap.Float64("rate", bucket.Rate()))
	return func(c *gin.Context) {

		if bucket.TakeAvailable(1) < 1 {
			resp.ResponseFail(c, http.StatusOK, "rate limit")
			c.Abort()
			return
		}
		c.Next()
	}

}
