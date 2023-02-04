package article

import (
	"github.com/gin-gonic/gin"
	"juejin/app/internal/model/article"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
)

type InfoApi struct{}

var insInfo InfoApi

func (a *InfoApi) GetArticleDetail(c *gin.Context) {
	articleId := c.Query("article_id")
	authorId := c.Query("author_id")
	err := service.View().View().CountView(c, authorId, articleId, 2)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "view counter internal error")
		return
	}
	var articleSubject = &article.Article{}
	err = service.Article().Info().GetArticleMajor(articleId, articleSubject)
	if err != nil {
		if err.Error() == "no such article" {
			resp.ResponseFail(c, http.StatusBadRequest, "no such article")
			return
		}
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	resp.OkWithData(c, "get article detail successfully", articleSubject)

}
