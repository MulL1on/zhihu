package draft

type Group struct{}

func (g *Group) Audit() *SAudit {
	return &insAudit
}
