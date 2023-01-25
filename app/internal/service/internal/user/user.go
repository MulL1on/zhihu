package user

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	g "juejin/app/global"
	"juejin/app/internal/model/user"
	"juejin/utils/jwt"
	"time"
)

type SUser struct{}

var insUser = SUser{}

func (s *SUser) CheckUserIsExist(username string) error {
	var name string
	sqlStr := "select username from user_auth where username=?"
	err := g.MysqlDB.QueryRow(sqlStr, username).Scan(&name)
	if err != nil {
		if err != sql.ErrNoRows {
			g.Logger.Error("query mysql record fail", zap.Error(err))
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("username is already exist")
}

func (s *SUser) CheckMailIsExist(mail string) error {
	user := &user.Auth{}
	sqlStr := "select * from user_auth where mail=?"
	err := g.MysqlDB.QueryRow(sqlStr, mail).Scan(&user)
	if err != nil {
		if err != sql.ErrNoRows {
			g.Logger.Error("query mysql record fail", zap.Error(err))
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("username is already exist")
}

func (s *SUser) EncryptPassword(password string) (string, error) {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptPassword), nil
}

func (s *SUser) CreateUser(userSubject *user.Auth) error {
	sqlStr := "insert into user_auth (username,password,mail,phone,create_time) values (?,?,?,?,?)"
	_, err := g.MysqlDB.Exec(sqlStr, userSubject.Username, userSubject.Password, userSubject.Mail, userSubject.Phone, time.Now())
	if err != nil {
		g.Logger.Error("create mysql record failed", zap.Error(err))
		return err
	}
	return nil
}
func (s *SUser) GetEncryptPassword(username string) (string, error) {
	var pwd string
	sqlStr := "select password from user_auth where username = ?"
	err := g.MysqlDB.QueryRow(sqlStr, username).Scan(&pwd)
	if err != nil {
		g.Logger.Error("get encrypt password failed", zap.Error(err))
		return "", err
	}
	return pwd, nil
}
func (s *SUser) CheckPassword(password, encryptPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPwd), []byte(password))
	return err == nil
}
func (s *SUser) GenerateToken(ctx context.Context, user *user.Auth) (string, error) {
	config := g.Config.Middleware.Jwt
	j := jwt.NewJWT(&jwt.Config{
		SecretKey:  config.SecretKey,
		ExpireTime: config.ExpiresTime,
		BufferTime: config.BufferTime,
		Issuer:     config.Issuer})
	claims := j.CreateClaims(&jwt.BaseClaims{
		Id:         user.Id,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	})
	tokenString, err := j.GenerateToken(&claims)
	if err != nil {
		g.Logger.Error("generate token failed.", zap.Error(err))
		return "", fmt.Errorf("internal err")
	}
	err = g.Rdb.Set(ctx, fmt.Sprintf("jwt:%d", user.Id), tokenString, time.Duration(config.ExpiresTime)*time.Second).Err()
	if err != nil {
		g.Logger.Error("set redis cache failed.",
			zap.Error(err), zap.String("key", "jwt:[id]"),
			zap.Int64("id", user.Id),
		)
		return "", fmt.Errorf("internal err")
	}
	return tokenString, nil

}
