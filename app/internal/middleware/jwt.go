package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/app/internal/model/resp"
	"juejin/utils/cookie"
	myjwt "juejin/utils/jwt"
	"net/http"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		cookieConfig := g.Config.App.Cookie
		cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
			cookie.Option{
				Config: cookieConfig.Cookie,
				Ctx:    c,
			})
		ok := cookieWriter.Get("x-token", &token)
		if token == "" || !ok {
			resp.ResponseFail(c, http.StatusUnauthorized, "not login")
			c.Abort()
			return
		}
		jwtConfig := g.Config.Middleware.Jwt
		j := myjwt.NewJWT(&myjwt.Config{SecretKey: jwtConfig.SecretKey})
		mc, err := j.ParseToken(token)
		if err != nil {
			resp.ResponseFail(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		//refresh token
		if mc.ExpiresAt.Unix()-time.Now().Unix() < mc.BufferTime && mc.ExpiresAt.Unix()-time.Now().Unix() > 0 {
			mc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Middleware.Jwt.ExpiresTime) * time.Second))
			newToken, _ := j.GenerateToken(mc)
			newClaims, _ := j.ParseToken(newToken)
			cookieWriter.Set("x-token", newToken)
			err = g.Rdb.Set(c, fmt.Sprintf("jwt:%d", newClaims.BaseClaims.Id), newToken, time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
			if err != nil {
				g.Logger.Error("set redis key failed", zap.Error(err), zap.String("key", "jwt[:id]"), zap.Int64("id", newClaims.BaseClaims.Id))
				resp.ResponseFail(c, http.StatusInternalServerError, "set token failed")
				c.Abort()
				return
			}
		}
		c.Set("id", mc.BaseClaims.Id)
		c.Next()
	}
}
