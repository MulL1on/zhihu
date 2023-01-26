package user

type Group struct{}

func (g *Group) Auth() *SAuth {
	return &insAuth
}

func (g *Group) Info() *SInfo {
	return &insInfo
}
