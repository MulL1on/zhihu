package api

import (
	"juejin/app/api/internal/article"
	"juejin/app/api/internal/draft"
	"juejin/app/api/internal/user"
)

var (
	insUser    = user.Group{}
	insArticle = article.Group{}
	insDraft   = draft.Group{}
)

func User() *user.Group {
	return &insUser
}

func Article() *article.Group {
	return &insArticle
}

func Draft() *draft.Group {
	return &insDraft
}
