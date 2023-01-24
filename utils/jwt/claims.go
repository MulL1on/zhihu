package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"juejin/utils/cookie"
	"time"
)

type CustomClaims struct {
	BufferTime int64
	jwt.RegisteredClaims
	BaseClaims
}

type BaseClaims struct {
	Id         int64
	Username   string
	CreateTime time.Time
	UpdateTime time.Time
}

func GetClaims(secret string, cookie *cookie.Cookie) (*CustomClaims, error) {
	var token string
	ok := cookie.Get("x-token", &token)
	if !ok {
		err := errors.New("get token by cookie failed")
		return nil, err
	}
	j := NewJWT(&Config{SecretKey: secret})
	claims, err := j.ParseToken(token)
	if err != nil {
		err := errors.New("parse token failed")
		return nil, err
	}
	return claims, nil
}

func GetUserInfo(secret string, cookie *cookie.Cookie, err error) (*BaseClaims, error) {
	if cl, err := GetClaims(secret, cookie); err != nil {
		return nil, err
	} else {
		return &cl.BaseClaims, nil
	}
}

func GetUserId(secret string, cookie *cookie.Cookie) (int64, error) {
	if cl, err := GetClaims(secret, cookie); err != nil {
		return -1, err
	} else {
		return cl.BaseClaims.Id, nil
	}
}
