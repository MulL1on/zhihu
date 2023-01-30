package collection

type Group struct{}

func (g *Group) Edit() *SEdit {
	return &insEdit
}
