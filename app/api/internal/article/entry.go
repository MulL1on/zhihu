package article

type Group struct{}

func (g *Group) Edit() *PublishApi {
	return &insPublish
}

func (g *Group) Info() *InfoApi {
	return &insInfo
}
