package tag

type Group struct{}

func (g *Group) Tag() *TagApi {
	return &insTag
}
