package oidc

type Group struct{}

func (g *Group) Oidc() *SOidc {
	return &insOidc
}
