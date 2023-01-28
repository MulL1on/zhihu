package tag

type Group struct{}

func (g *Group) Edit() *SEdit {
	return &insEdit
}
