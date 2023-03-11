package oidc

type Group struct{}

func (g *Group) Authentication() *AuthenticationApi {
	return &insAuthentication
}
