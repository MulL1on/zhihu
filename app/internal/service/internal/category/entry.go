package category

type Group struct{}

func (g *Group) Info() *SInfo {
	return &insSInfo
}
