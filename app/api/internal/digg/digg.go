package digg

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/digg"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
)

type LikeApi struct{}

var insLike LikeApi

func (a *LikeApi) DoDigg(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var doDigg = &digg.Like{}
	err := c.BindJSON(doDigg)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "json patter incorrect")
		return
	}
	err = service.Digg().Like().DoDigg(c, userId, doDigg)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "do digg successfully")
}

func (a *LikeApi) UndoDigg(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var doDigg = &digg.Like{}
	err := c.BindJSON(doDigg)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "json patter incorrect")
		return
	}
	err = service.Digg().Like().UndoDigg(c, userId, doDigg)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "undo digg successfully")
}
