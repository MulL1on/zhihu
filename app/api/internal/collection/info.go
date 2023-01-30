package collection

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/article"
	"juejin/app/internal/model/collection"
	"juejin/app/internal/model/user"
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
	var articleBriefList = make([]*article.Brief, len(*idList))

	//获取收藏的文章
	if len(*idList) != 0 {
		for k, v := range *idList {
			var articleSubject = &article.Article{}
			var uBasic = &user.Basic{}
			var articleBrief = &article.Brief{}
			err = service.Article().Info().GetArticleMajor(v, articleSubject)
			if err != nil {
				resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
				return
			}
			err = service.User().Info().GetUserBasic(uBasic, articleSubject.UserId)
			if err != nil {
				resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
				return
			}
			articleBrief.ArticleInfo = articleSubject
			articleBrief.ArticleId = articleSubject.ArticleId
			articleBrief.AuthorInfo = uBasic
			articleBriefList[k] = articleBrief
		}
	}

	//获取收藏夹信息
	var cs = &collection.Set{}
	err = service.Collection().Edit().GetCollectionInfo(collectionId, cs)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
	}

	type data struct {
		Articles       []*article.Brief `json:"articles"`
		CollectionInfo *collection.Set  `json:"collection_info"`
	}

	var d = &data{
		Articles:       articleBriefList,
		CollectionInfo: cs,
	}

	resp.OkWithData(c, "get collected article successfully", d)

}
