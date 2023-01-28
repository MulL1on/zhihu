package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/utils/common/resp"
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
			resp.ResponseFail(c, http.StatusUnauthorized, "not logged in.")
			c.Abort()
			return
		}

		//check black list
		if err := g.Rdb.Get(c, fmt.Sprintf("black_list:%s", token)).Err(); err != nil {
			if err != redis.Nil {
				resp.ResponseFail(c, http.StatusInternalServerError, "get token black list fail")
				c.Abort()
				return
			}
		} else {
			resp.ResponseFail(c, http.StatusUnauthorized, "token is invalid")
			c.Abort()
			return
		}

		//parse token
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
			//将旧的token放进黑名单
			err = g.Rdb.Set(c, fmt.Sprintf("black_list:%s", token), "", time.Duration(mc.ExpiresAt.Unix()-time.Now().Unix())*time.Second).Err()
			if err != nil {
				g.Logger.Error("set redis key failed", zap.Error(err))
				resp.ResponseFail(c, http.StatusInternalServerError, "set token  to blacklist failed")
				c.Abort()
				return
			}
			mc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Middleware.Jwt.ExpiresTime) * time.Second))
			newToken, _ := j.GenerateToken(mc)
			cookieWriter.Set("token", newToken)
		}
		c.Set("id", mc.BaseClaims.Id)
		c.Next()
	}
}
