package router

type Group struct {
	UserRouter
	DraftRouter
	ArticleRouter
	CollectionRouter
	FollowerRouter
	CommentRouter
	DiggRouter
	UploadRouter
	TagRouter
	RankRouter
}
