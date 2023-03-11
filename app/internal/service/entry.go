package service

import (
	"juejin/app/internal/service/internal/article"
	"juejin/app/internal/service/internal/category"
	"juejin/app/internal/service/internal/collection"
	"juejin/app/internal/service/internal/comment"
	"juejin/app/internal/service/internal/digg"
	"juejin/app/internal/service/internal/draft"
	"juejin/app/internal/service/internal/follower"
	"juejin/app/internal/service/internal/oidc"
	"juejin/app/internal/service/internal/rank"
	"juejin/app/internal/service/internal/tag"
	"juejin/app/internal/service/internal/user"
	"juejin/app/internal/service/internal/view"
)

var (
	insUser       = user.Group{}
	insDraft      = draft.Group{}
	insArticle    = article.Group{}
	insTag        = tag.Group{}
	insCollection = collection.Group{}
	insFollow     = follower.Group{}
	insComment    = comment.Group{}
	insView       = view.Group{}
	insDigg       = digg.Group{}
	insCategory   = category.Group{}
	insRank       = rank.Group{}
	insOidc       = oidc.Group{}
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

func Collection() *collection.Group {
	return &insCollection
}

func Follower() *follower.Group {
	return &insFollow
}

func Comment() *comment.Group {
	return &insComment
}

func View() *view.Group {
	return &insView
}

func Digg() *digg.Group {
	return &insDigg
}

func Category() *category.Group {
	return &insCategory
}

func Rank() *rank.Group {
	return &insRank
}

func Oidc() *oidc.Group {
	return &insOidc
}
