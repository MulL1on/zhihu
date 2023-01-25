package api

import "juejin/app/api/user"

var (
	insUser = user.Group{}
)

func User() *user.Group {
	return &insUser
}
