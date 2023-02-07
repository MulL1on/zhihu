package rank

type Group struct{}

func (g *Group) Rank() *SRank {
	return &insSRank
}
