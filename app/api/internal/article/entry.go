package article

type Group struct{}

func (g *Group) Edit() *PublishApi {
	return &insPublish
}
