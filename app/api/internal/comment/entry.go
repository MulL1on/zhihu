package comment

type Group struct{}

func (g *Group) Review() *ReviewApi {
	return &insReview
}

func (g *Group) Reply() *ReplyApi {
	return &insReply
}
