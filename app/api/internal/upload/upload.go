package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"juejin/utils/common/resp"
	"juejin/utils/common/upload"
	"net/http"
)

type UploadApi struct{}

var insUpload UploadApi

func (a *UploadApi) UserAvatarUpload(c *gin.Context) {
	f, _ := c.FormFile("avatar")
	path := "user/user-avatar/" + getUuid() + ".png"
	code, url := upload.ToQiniu(f, path)
	if code != 0 {
		resp.UploadFail(c, code, url)
		return
	}
	resp.UploadOk(c, http.StatusOK, url)
}

func getUuid() string {
	u := uuid.New()
	return u.String()
}

func (a *UploadApi) ArticleCoverUpload(c *gin.Context) {
	f, _ := c.FormFile("cover")
	path := "article/cover/" + getUuid() + ".png"
	code, url := upload.ToQiniu(f, path)
	if code != 0 {
		resp.UploadFail(c, code, url)
		return
	}
	resp.UploadOk(c, http.StatusOK, url)
}
