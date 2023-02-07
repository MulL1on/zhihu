package user

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/user"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strconv"
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
	err := service.User().Info().GetUserInfo(c, userBasic, userCounter, id)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	userBasic.IsFollow = false
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

func (a *InfoApi) GetOthersInfo(c *gin.Context) {
	userId, _ := c.Get("id")
	othersId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	var userBasic = &user.Basic{}
	var userCounter = &user.Counter{}

	err := service.User().Info().GetUserInfo(c, userBasic, userCounter, othersId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	if userId != nil {
		err = service.Follower().Follow().CheckIsFollowed(othersId, userId)
		if err != nil {
			if err.Error() != "user is already followed" {
				resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
				return
			}
			userBasic.IsFollow = true
		}
		userBasic.IsFollow = false
	} else {
		userBasic.IsFollow = false
	}

	var infoPack = user.InfoPack{
		Counter: *userCounter,
		Basic:   *userBasic,
	}
	resp.OkWithData(c, "get user info successfully", infoPack)
}
