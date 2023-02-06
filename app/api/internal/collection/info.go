package collection

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/article"
	"juejin/app/internal/model/collection"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strconv"
)

type InfoApi struct{}

var insInfo InfoApi

func (a *InfoApi) GetCollectedArticle(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		resp.ResponseFail(c, http.StatusUnauthorized, "not log in")
		return
	}

	collectionId, _ := strconv.ParseInt(c.Query("collection_id"), 10, 64)
	limit, _ := strconv.Atoi(c.Query("limit"))
	pageNo, _ := strconv.Atoi(c.Query("page_no"))

	err := service.Collection().Edit().CheckViewAuth(collectionId, userId)
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

	idList, err := service.Collection().Edit().GetSelectArticleId(collectionId, limit, pageNo)

	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	var articles = make([]article.List, 0)
	for _, v := range *idList {
		var targetArticle = &article.List{}
		err = service.Article().Info().GetArticle(v, targetArticle)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}

		//获取分类信息
		err = service.Category().Info().GetCategoryInfo(targetArticle.ArticleInfo.CategoryId, &targetArticle.Category)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		}

		//获取标签信息
		tagList, err := service.Tag().Edit().GetTagListByItem(targetArticle.ArticleInfo.ArticleId, 2)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
		targetArticle.Tags = *tagList

		//检查是否点赞
		if userId != "" {
			targetArticle.ArticleInfo.IsDigg, err = service.Digg().Like().CheckIsDigg(c, userId, targetArticle.ArticleInfo.ArticleId, 2)
			if err != nil {
				resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
				return
			}
		}
		targetArticle.ArticleInfo.IsDigg = false

		articles = append(articles, *targetArticle)
	}

	//获取收藏夹信息
	var cs = &collection.Set{}
	err = service.Collection().Edit().GetCollectionInfo(collectionId, cs)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
	}

	type data struct {
		Articles       []article.List  `json:"articles"`
		CollectionInfo *collection.Set `json:"collection_info"`
	}

	var d = &data{
		Articles:       articles,
		CollectionInfo: cs,
	}

	resp.OkWithData(c, "get collected article successfully", d)

}
