package collection

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/collection"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strconv"
)

type EditApi struct{}

var insEdit EditApi

type CollectArticle struct {
	ArticleId    string `json:"article_id" form:"article_id"`
	CollectionId int64  `json:"collection_id" form:"collection_id"`
}

func (a *EditApi) CreatCollectionSet(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var cs = &collection.Set{}

	err := c.BindJSON(&cs)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "collection set pattern incorrect")
		return
	}

	if cs.CollectionName == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "collection name can not be null")
		return
	}

	cs.CollectionId = service.Collection().Edit().GenerateUid()

	err = service.Collection().Edit().CreateCollectionSet(userId, cs)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = service.Collection().Edit().GetCollectionInfo(cs.CollectionId, cs)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.OkWithData(c, "create collection set successfully", cs)
}

func (a *EditApi) CollectArticle(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var ca = &CollectArticle{}
	err := c.BindJSON(ca)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "request pattern error")
		return
	}

	//验证编辑权限
	err = service.Collection().Edit().CheckEditAuth(ca.CollectionId, userId)
	if err != nil {
		if err.Error() == "no such collectionSet" {
			resp.ResponseFail(c, http.StatusBadRequest, "no such collection set")
			return
		}
		if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	//查看是否已经收藏
	err = service.Collection().Edit().CheckArticleIsExist(ca.CollectionId, ca.ArticleId)
	if err != nil {
		if err.Error() == "article is already exist" {
			resp.ResponseFail(c, http.StatusBadRequest, "article is already exist")
			return
		} else {
			resp.ResponseFail(c, http.StatusInternalServerError, "check article's existence error")
			return
		}
	}

	err = service.Collection().Edit().AddSelectArticle(ca.CollectionId, ca.ArticleId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "add article to collection set successfully")
}

func (a *EditApi) ModifyCollectionSet(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var cs = &collection.Set{}
	err := c.BindJSON(cs)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "collection set pattern incorrect")
		return
	}

	//检查编辑权限
	err = service.Collection().Edit().CheckEditAuth(cs.CollectionId, userId)
	if err != nil {
		if err.Error() == "no such collectionSet" {
			resp.ResponseFail(c, http.StatusOK, "no such collection set")
			return
		}
		if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = service.Collection().Edit().ModifyCollectionSet(cs.CollectionId, cs)

	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "modify collection set successfully")

}

func (a *EditApi) DeleteCollectionSet(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	collectionId, err := strconv.ParseInt(c.PostForm("collection_id"), 10, 64)

	//检查编辑权限
	err = service.Collection().Edit().CheckEditAuth(collectionId, userId)
	if err != nil {
		if err.Error() == "no such collectionSet" {
			resp.ResponseFail(c, http.StatusOK, "no such collection set")
			return
		}
		if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	err = service.Collection().Edit().DeleteCollectionSet(collectionId, userId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.ResponseSuccess(c, http.StatusOK, "delete collection set successfully")

}

func (a *EditApi) RemoveSelectedArticle(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}
	var ca = &CollectArticle{}
	err := c.BindJSON(ca)
	if err != nil {
		resp.ResponseFail(c, http.StatusBadRequest, "request pattern error")
		return
	}

	//验证编辑权限
	err = service.Collection().Edit().CheckEditAuth(ca.CollectionId, userId)
	if err != nil {
		if err.Error() == "no such collectionSet" {
			resp.ResponseFail(c, http.StatusBadRequest, "no such collection set")
			return
		}
		if err.Error() == "unauthorized" {
			resp.ResponseFail(c, http.StatusUnauthorized, "unauthorized")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	//查看是否已经收藏
	err = service.Collection().Edit().CheckArticleIsExist(ca.CollectionId, ca.ArticleId)
	if err != nil {
		if err.Error() != "article is already exist" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
		err = service.Collection().Edit().RemoveSelectedArticle(ca.ArticleId, ca.CollectionId)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
		resp.ResponseSuccess(c, http.StatusOK, "undo collect successfully")
		return
	}
	resp.ResponseFail(c, http.StatusInternalServerError, "article is not collected yet")
	return

}
