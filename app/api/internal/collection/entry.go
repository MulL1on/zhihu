package collection

type Group struct{}

func (g *Group) Edit() *EditApi {
	return &insEdit
}

func (g *Group) Info() *InfoApi {
	return &insInfo
}
