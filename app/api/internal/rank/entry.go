package rank

type Group struct{}

func (g *Group) Rank() *RankApi {
	return &insRank
}
