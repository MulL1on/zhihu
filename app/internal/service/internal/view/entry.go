package view

type Group struct{}

func (g *Group) View() *SView {
	return &insView
}
