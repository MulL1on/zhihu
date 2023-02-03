package comment

type Group struct{}

func (g *Group) Review() *SReview {
	return &insReview
}

func (g *Group) Reply() *SReply {
	return &insReply
}
