package follower

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/user"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strconv"
)

type FollowApi struct{}

var insFollow FollowApi

func (a *FollowApi) DoFollow(c *gin.Context) {
	followerId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	followeeId, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	err := service.Follower().Follow().CheckIsFollowed(followeeId, followerId)
	if err != nil {
		if err.Error() == "user is already followed" {
			resp.ResponseFail(c, http.StatusBadRequest, "user is already followed")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	err = service.Follower().Follow().DoFollow(followerId, followeeId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "do follow successfully")

}

func (a *FollowApi) UndoFollow(c *gin.Context) {
	followerId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	followeeId, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	err := service.Follower().Follow().CheckIsFollowed(followeeId, followerId)
	if err != nil {
		if err.Error() != "user is already followed" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
	} else {
		resp.ResponseFail(c, http.StatusBadRequest, "not follow this user yet")
		return
	}
	err = service.Follower().Follow().UndoFollow(followerId, followeeId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "undo follow successfully")
}

func (a *FollowApi) GetFollowerList(c *gin.Context) {
	userId, ok := c.Get("id")
	limit, _ := strconv.Atoi(c.Query("limit"))
	pageNo, _ := strconv.Atoi(c.Query("page_no"))
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}

	//获取关注你的用户的id
	list, err := service.Follower().Follow().GetFollowerList(userId, limit, pageNo)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
		return
	}

	//获取用户详情
	var followers = make([]*user.InfoPack, len(*list))
	if len(*list) != 0 {
		for k, v := range *list {
			var uInfo = &user.InfoPack{}
			err = service.User().Info().GetUserInfo(&uInfo.Basic, &uInfo.Counter, v)
			if err != nil {
				resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
				return
			}
			followers[k] = uInfo
		}
	}
	resp.OkWithData(c, "get follower list successfully", followers)

}

func (a *FollowApi) GetFolloweeList(c *gin.Context) {
	userId, ok := c.Get("id")
	limit, _ := strconv.Atoi(c.Query("limit"))
	pageNo, _ := strconv.Atoi(c.Query("page_no"))
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	//获取你关注的用户的id
	list, err := service.Follower().Follow().GetFolloweeList(userId, limit, pageNo)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal err")
		return
	}

	//获取用户详情
	var followee = make([]*user.InfoPack, len(*list))
	if len(*list) != 0 {
		for k, v := range *list {
			var uInfo = &user.InfoPack{}
			err = service.User().Info().GetUserInfo(&uInfo.Basic, &uInfo.Counter, v)
			if err != nil {
				resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
				return
			}
			followee[k] = uInfo
		}
	}
	resp.OkWithData(c, "get follower list successfully", followee)
}
