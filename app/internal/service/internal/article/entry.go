package article

type Group struct{}

func (g *Group) Publish() *SPublish {
	return &insPublish
}
