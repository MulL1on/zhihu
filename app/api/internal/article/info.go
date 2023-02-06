package article

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/article"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strconv"
)

type InfoApi struct{}

var insInfo InfoApi

func (a *InfoApi) GetArticleDetail(c *gin.Context) {
	articleId := c.Query("article_id")
	userId, _ := c.Get("id")
	var targetArticle = &article.List{}
	err := service.Article().Info().GetArticleDetail(articleId, targetArticle)
	if err != nil {
		if err.Error() == "no such article" {
			resp.ResponseFail(c, http.StatusBadRequest, "no such article")
			return
		}
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
	if userId != nil {
		targetArticle.ArticleInfo.IsDigg, err = service.Digg().Like().CheckIsDigg(c, userId, targetArticle.ArticleInfo.ArticleId, 2)
		if err != nil {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
	}
	targetArticle.ArticleInfo.IsDigg = false

	//浏览数增加
	err = service.View().View().CountView(c, targetArticle.ArticleInfo.UserId, targetArticle.ArticleInfo.ArticleId, 2)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	resp.OkWithData(c, "get article detail successfully", targetArticle)

}

func (a *InfoApi) GetArticleList(c *gin.Context) {
	categoryId := c.Query("category")
	tagId := c.Query("tag")
	sortBy, _ := strconv.Atoi(c.Query("sort_by"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	pageNo, _ := strconv.Atoi(c.Query("page_no"))
	userId, _ := c.Get("id")
	var articleList *[]string
	var err error
	switch sortBy {
	case 1:
		articleList, err = service.Article().Info().GetArticleListByDigg(limit, pageNo, categoryId, tagId)
	case 2:
		articleList, err = service.Article().Info().GetArticleListByTime(limit, pageNo, categoryId, tagId)
	}
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	if len(*articleList) == 0 {
		resp.ResponseSuccess(c, http.StatusOK, "no record")
		return
	}

	var data = make([]article.List, 0)
	for _, v := range *articleList {
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
		if userId != nil {
			targetArticle.ArticleInfo.IsDigg, err = service.Digg().Like().CheckIsDigg(c, userId, targetArticle.ArticleInfo.ArticleId, 2)
			if err != nil {
				resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
				return
			}
		}
		targetArticle.ArticleInfo.IsDigg = false

		data = append(data, *targetArticle)
	}
	resp.OkWithData(c, "get article list successfully", data)
}
