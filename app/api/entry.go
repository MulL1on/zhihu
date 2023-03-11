package api

import (
	"juejin/app/api/internal/article"
	"juejin/app/api/internal/collection"
	"juejin/app/api/internal/comment"
	"juejin/app/api/internal/digg"
	"juejin/app/api/internal/draft"
	"juejin/app/api/internal/follower"
	"juejin/app/api/internal/oidc"
	"juejin/app/api/internal/rank"
	"juejin/app/api/internal/tag"
	"juejin/app/api/internal/upload"
	"juejin/app/api/internal/user"
)

var (
	insUser       = user.Group{}
	insArticle    = article.Group{}
	insDraft      = draft.Group{}
	insCollection = collection.Group{}
	insFollower   = follower.Group{}
	insComment    = comment.Group{}
	insDigg       = digg.Group{}
	insUpload     = upload.Group{}
	insTag        = tag.Group{}
	insRank       = rank.Group{}
	insOidc       = oidc.Group{}
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

func Collection() *collection.Group {
	return &insCollection
}

func Follower() *follower.Group {
	return &insFollower
}
func Comment() *comment.Group {
	return &insComment
}
func Digg() *digg.Group {
	return &insDigg
}

func Upload() *upload.Group {
	return &insUpload
}

func Tag() *tag.Group {
	return &insTag
}

func Rank() *rank.Group {
	return &insRank
}

func Oidc() *oidc.Group {
	return &insOidc
}
