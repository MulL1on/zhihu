package oidc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	g "juejin/app/global"
	"juejin/utils/jwt"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type SOidc struct{}

var insOidc SOidc

func (s *SOidc) GenerateCode(ctx context.Context, userId int64, clientId string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 7)
	for i := range code {
		code[i] = letters[rand.Int63()&63]
	}
	str := fmt.Sprintf("oidc:%s", string(code))
	err := g.Rdb.Set(ctx, str, fmt.Sprintf("%d:%s", userId, clientId), time.Duration(g.Config.Oidc.ExpireTime)*time.Second).Err()
	if err != nil {
		g.Logger.Error("set oidc code error", zap.Error(err))
		return "", err
	}
	return string(code), nil
}

func (s *SOidc) GetRedirectUri(clientId string) (string, error) {
	sql := "select redirect_uri from oidc_client where client_id=?"
	var redirectUri string
	err := g.MysqlDB.QueryRow(sql, clientId).Scan(&redirectUri)
	if err != nil {
		g.Logger.Error("get redirect uri error", zap.Error(err))
		return "", err
	}
	return redirectUri, nil
}

func (s *SOidc) CheckCode(ctx context.Context, code string) error {
	if err := g.Rdb.Get(ctx, fmt.Sprintf("oidc:%s", code)).Err(); err != nil {
		if err != redis.Nil {
			g.Logger.Error("get oidc from redis error", zap.Error(err))
			return err
		}
		return fmt.Errorf("code expired or invalid")
	}

	return nil

}

func (s *SOidc) GetUserId(username string) (int64, error) {
	sql := "select id from user_auth where username=?"
	var userId int64
	err := g.MysqlDB.QueryRow(sql, username).Scan(&userId)
	if err != nil {
		g.Logger.Error("get user id error")
		return 0, err
	}
	return userId, nil
}

func (s *SOidc) GenerateIdToken(userId, clientId string) (string, error) {
	config := g.Config.Middleware.Jwt
	var cId = make([]string, 1)
	cId[0] = clientId
	j := jwt.NewJWT(&jwt.Config{
		SecretKey:  config.SecretKey,
		ExpireTime: config.ExpiresTime,
		BufferTime: config.BufferTime,
		Issuer:     config.Issuer})
	claims := j.CreateOIDCClaims(userId, cId)
	tokenString, err := j.GenerateIdToken(&claims)
	if err != nil {
		g.Logger.Error("generate token failed.", zap.Error(err))
		return "", err
	}

	return tokenString, nil
}

func (s *SOidc) RegisterODICClient(redirectUri string, appName string) (string, error) {
	sql := "insert into oidc_client (client_id, redirect_uri, app_name) values (?,?,?)"
	//generate client id
	rand.Seed(time.Now().UnixNano())
	clientId := make([]byte, 8)
	for i := range clientId {
		clientId[i] = letters[rand.Int63()&63]
	}
	_, err := g.MysqlDB.Exec(sql, clientId, redirectUri, appName)
	if err != nil {
		g.Logger.Error("set odic client error", zap.Error(err))
		return "", err
	}
	return string(clientId), nil
}
