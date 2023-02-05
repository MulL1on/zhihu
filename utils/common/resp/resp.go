package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseFail(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  message,
		"ok":   false,
	})
}

func ResponseSuccess(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  message,
		"ok":   true,
	})
}

func OkWithData(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    message,
		"data":   data,
		"ok":     true,
	})
}

func UploadOk(c *gin.Context, code int, url string) {
	c.JSON(code, gin.H{
		"code": code,
		"url":  url,
		"ok":   true,
	})
}

func UploadFail(c *gin.Context, code int, url string) {
	c.JSON(code, gin.H{
		"code": code,
		"url":  url,
		"ok":   false,
	})
}
