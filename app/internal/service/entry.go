package service

import (
	"juejin/app/internal/service/internal/article"
	"juejin/app/internal/service/internal/draft"
	"juejin/app/internal/service/internal/tag"
	"juejin/app/internal/service/internal/user"
)

var (
	insUser    = user.Group{}
	insDraft   = draft.Group{}
	insArticle = article.Group{}
	insTag     = tag.Group{}
)

func User() *user.Group {
	return &insUser
}

func Draft() *draft.Group {
	return &insDraft
}

func Article() *article.Group {
	return &insArticle
}

func Tag() *tag.Group {
	return &insTag
}
