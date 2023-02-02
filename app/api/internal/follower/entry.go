package follower

type Group struct{}

func (g *Group) Follow() *FollowApi {
	return &insFollow
}
