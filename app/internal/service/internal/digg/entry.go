package digg

type Group struct{}

func (g *Group) Like() *SLike {
	return &insLike
}
