package user

type Group struct{}

func (g *Group) Auth() *AuthApi {
	return &insAuth
}

func (g *Group) Info() *InfoApi {
	return &insInfo
}
