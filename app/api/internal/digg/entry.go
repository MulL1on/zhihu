package digg

type Group struct{}

func (g *Group) Like() *LikeApi {
	return &insLike
}
