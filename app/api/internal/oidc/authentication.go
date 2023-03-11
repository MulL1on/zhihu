package oidc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	g "juejin/app/global"
	"juejin/app/internal/model/oidc"
	"juejin/app/internal/service"
	"juejin/utils/common/resp"
	"net/http"
	"strings"
)

type AuthenticationApi struct{}

var insAuthentication AuthenticationApi

func (a *AuthenticationApi) GetEUAuth(c *gin.Context) {
	//EU login & give authorization
	clientId := c.Query("client_id")
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "username cannot be null")
		return
	}
	if password == "" {
		resp.ResponseFail(c, http.StatusBadRequest, "password cannot be null")
		return
	}

	err := service.User().Auth().CheckUserIsExist(username)
	if err != nil {
		if err.Error() != "username is already exist" {
			resp.ResponseFail(c, http.StatusInternalServerError, "check username's existence error")
			return
		}
	} else {
		resp.ResponseFail(c, http.StatusBadRequest, "user doesn't exist")
		return
	}
	encryptPwd, err := service.User().Auth().GetEncryptPassword(username)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "get encrypt password failed")
		return
	}
	if !service.User().Auth().CheckPassword(password, encryptPwd) {
		resp.ResponseFail(c, http.StatusBadRequest, "invalid password or username")
		return
	}

	//get user id
	userId, err := service.Oidc().Oidc().GetUserId(username)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	//get redirect uri
	redirectUri, err := service.Oidc().Oidc().GetRedirectUri(clientId)

	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}

	//generate code & set code in redis
	code, err := service.Oidc().Oidc().GenerateCode(c, userId, clientId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	//redirect to uri selected
	c.Redirect(http.StatusMovedPermanently, redirectUri+"?code="+code)
}

func (a *AuthenticationApi) DistributeToken(c *gin.Context) {
	code := c.PostForm("code")
	//handle the code
	err := service.Oidc().Oidc().CheckCode(c, code)
	if err != nil {
		if err.Error() != "code expired or invalid" {
			resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
			return
		}
		resp.ResponseFail(c, http.StatusBadRequest, "code expired or invalid")
		return
	}

	//generate id token
	str, err := g.Rdb.Get(c, fmt.Sprintf("code:%s", code)).Result()
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	parts := strings.SplitN(str, ":", 2)
	userId := parts[0]
	clientId := parts[1]
	token, err := service.Oidc().Oidc().GenerateIdToken(userId, clientId)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	var res = &oidc.TokenResponse{
		TokenType: "Bearer",
		ExpiresIn: g.Config.Oidc.ExpireTime,
		IdToken:   token,
	}
	resp.OkWithData(c, "successfully", res)
}

func (a *AuthenticationApi) GetUserInfoByAC(c *gin.Context) {
	//handle the access token

	//return user info
}

func (a *AuthenticationApi) DistributeClientId(c *gin.Context) {
	//must login first to creat a oidc application
	appName := c.PostForm("app_name")
	redirectUri := c.PostForm("redirect_uri")
	//generate client id & set in database (client id | app_name)
	clientId, err := service.Oidc().Oidc().RegisterODICClient(redirectUri, appName)
	if err != nil {
		resp.ResponseFail(c, http.StatusInternalServerError, "internal error")
		return
	}
	//generate public key

	//return public key &return client idd
	type data struct {
		ClientId  string `json:"client_id"`
		PublicKey string `json:"public_key"`
	}
	var d = &data{
		ClientId:  clientId,
		PublicKey: "",
	}
	resp.ResponseSuccess(c, http.StatusOK)
}
