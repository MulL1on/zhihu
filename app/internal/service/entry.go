package service

import "juejin/app/internal/service/internal/user"

var (
	insUser = user.Group{}
)

func User() *user.Group {
	return &insUser
}
