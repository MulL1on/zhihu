package user

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/user"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
)

type InfoApi struct{}

var insInfo = InfoApi{}

func (a *InfoApi) GetUserInfo(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}

	var userBasic = &user.Basic{}
	var userCounter = &user.Counter{}
	err := service.User().Info().GetUserInfo(userBasic, userCounter, id)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	var infoPack = user.InfoPack{
		Counter: *userCounter,
		Basic:   *userBasic,
	}
	resp.OkWithData(c, "get user info successfully", infoPack)
}

func (a *InfoApi) UpdateUserInfo(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var userBasic = &user.Basic{}
	err := c.BindJSON(userBasic)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "bind json error")
		return
	}
	err = service.User().Info().UpdateUserInfo(userBasic, id)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "update user info successfully")
}
