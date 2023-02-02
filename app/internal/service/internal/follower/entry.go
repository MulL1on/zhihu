package follower

type Group struct{}

func (g *Group) Follow() *SFollow {
	return &insFollow
}
