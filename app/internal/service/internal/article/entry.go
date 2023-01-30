package article

type Group struct{}

func (g *Group) Publish() *SPublish {
	return &insPublish
}

func (g *Group) Info() *SInfo {
	return &insInfo
}
